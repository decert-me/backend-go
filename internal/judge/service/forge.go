package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"backend-go/pkg/log"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"os"
	"regexp"
	"strings"
	"time"
)

func (s *Service) TestSolidity(req request.ForgeTestReq, spjCode string) (res response.ForgeTestRes, err error) {
	if req.Address == "" {
		req.Address = common.HexToAddress("0").String()
	}
	foundryPath := s.c.Foundry.WorkPath
	// 获取合约名称
	re := regexp.MustCompile(`contract\s+(\w+)\s*{`)

	result := re.FindStringSubmatch(req.Code)
	if len(result) < 1 {
		return res, errors.New("contract name no match")
	}
	// 保存代码
	fileName := time.Now().Format("20060102150405.000") + ".sol"
	relativePath := "src/" + req.Address + "/"
	relativeFilePath := relativePath + fileName
	if err = os.MkdirAll(foundryPath+"/"+relativePath, os.ModePerm); err != nil {
		log.Errorv("os.MkdirAll() Filed", zap.Error(err))
		return
	}
	p := foundryPath + "/" + relativeFilePath
	if err = os.WriteFile(p, []byte(req.Code+"\n"+spjCode), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}
	contract := "--match-path=" + relativeFilePath
	args := []string{"test", contract, "--offline", "--json"}
	execRes, err := execCommand(foundryPath, "forge", args...)
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
	res.TotalTestcases = len(resultArray)
	for _, v := range resultArray {
		if !gjson.Get(v.String(), "*.success").Bool() {
			if gjson.Get(v.String(), "*.reason").String() != "" {
				res.Output = gjson.Get(v.String(), "*.reason").String()
			} else {
				res.Output = gjson.Get(v.String(), "*.decoded_logs").String()
			}
			res.Status = 1
			return
		}
		res.TotalCorrect++
	}

	return res, nil
}
