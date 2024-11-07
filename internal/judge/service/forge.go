package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"backend-go/pkg/log"
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

func (s *Service) TestSolidity(req request.TestReq, spjCode string) (res response.TestRes, err error) {
	if strings.TrimSpace(req.Address) == "" {
		req.Address = common.HexToAddress("0").String()
	}
	foundryPath := path.Join(s.c.Judge.SolidityWorkPath, req.Address, "foundry")
	// 获取合约名称
	re := regexp.MustCompile(`(?m)^[^\n//]*contract\s+(\w+)`)
	result := re.FindStringSubmatch(req.Code)
	if len(result) < 1 {
		return res, errors.New("TestSolidity contract name no match")
	}
	// 保存代码
	fileName := time.Now().Format("20060102150405.000") + ".sol"
	relativeFilePath := path.Join("src", fileName)
	p := path.Join(foundryPath, relativeFilePath)
	if err = os.WriteFile(p, []byte(req.Code+"\n"+spjCode), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}
	contract := "--match-path=" + relativeFilePath
	//args := []string{"test", contract, "--offline", "--json"}
	//execRes, err := execCommand(foundryPath, "forge", args...)
	command := fmt.Sprintf("cd /foundry && forge test --evm-version cancun %s --json", contract)
	args := []string{"exec", "-i", req.Address, "bash", "-c", command}
	execRes, err := execCommand("", "docker", args...)
	//if err := os.Rename(p, p+".bak"); err != nil {
	//	log.Errorv("os.Rename error", zap.Error(err))
	//}
	if err := os.Remove(p); err != nil {
		log.Errorv("os.Remove error", zap.Error(err))
	}
	if err != nil {
		return
	}
	execResList := strings.Split(execRes, "\n")

	if len(execResList) == 0 || !gjson.Valid(execResList[len(execResList)-2]) {
		res.Output = strings.Replace(execRes, req.Address+"/"+fileName, result[1]+".sol", -1)
		res.Status = 1
		return
	}
	resultArray := gjson.Get(execResList[len(execResList)-2], "*.test_results").Array()
	//fmt.Println("resultArray", resultArray)
	for _, v := range resultArray {
		gjson.Parse(v.String()).ForEach(func(key, value gjson.Result) bool {
			res.TotalTestcases += 1
			return true
		})
		gjson.Parse(v.String()).ForEach(func(key, value gjson.Result) bool {
			status := gjson.Get(value.String(), "status").String()
			reason := gjson.Get(value.String(), "reason").String()
			decodedLogs := gjson.Get(value.String(), "decoded_logs").String()
			fmt.Println("status", status)
			fmt.Println("reason", reason)
			fmt.Println("decodedLogs", decodedLogs)
			if status != "Success" {
				if reason != "" {
					res.Output = fmt.Sprintf("%s: %s", key, reason)
				} else {
					res.Output = fmt.Sprintf("%s: %s", key, decodedLogs)
				}
				res.Status = 3
				return false
			}
			res.TotalCorrect++
			return true // keep iterating
		})
	}
	res.Status = 3 // 运行成功状态
	return res, nil
}
