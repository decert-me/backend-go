package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"backend-go/pkg/log"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

func (s *Service) BuildSolidity(private string, req request.BuildReq) (res response.BuildRes, err error) {
	if strings.TrimSpace(req.Address) == "" {
		req.Address = common.HexToAddress("0").String()
	}
	if private == "" {
		private = GetPrivate()
	}
	foundryPath := path.Join(s.c.Judge.SolidityWorkPath, req.Address, "foundry")
	// 获取合约名称
	re := regexp.MustCompile(`contract\s+(\w+)\s*{`)
	result := re.FindStringSubmatch(req.Code)
	if len(result) < 1 {
		return res, errors.New("contract name no match")
	}
	// 保存代码
	fileName := time.Now().Format("20060102150405.000") + ".sol"
	relativeFilePath := path.Join("src", fileName)
	p := path.Join(foundryPath, relativeFilePath)
	if err = os.WriteFile(p, []byte(req.Code), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}
	contract := relativeFilePath + ":" + result[1]
	command := fmt.Sprintf("cd /foundry && forge create %s %s --json", contract, fmt.Sprintf(" --private-key=%s", private))
	args := []string{"exec", "-i", req.Address, "bash", "-c", command}
	execRes, err := execCommand("", "docker", args...)
	if err != nil {
		return
	}
	//if err := os.Rename(p, p+".bak"); err != nil {
	//	log.Errorv("os.Rename error", zap.Error(err))
	//}
	if err := os.Remove(p); err != nil {
		log.Errorv("os.Remove error", zap.Error(err))
	}

	if !gjson.Valid(execRes) {
		fmt.Println(req.Address + "/" + fileName)
		res.Output = strings.Replace(execRes, req.Address+"/"+fileName, result[1]+".sol", -1)
		res.Status = 1
		return
	}
	res.Output = "Compiler successful"
	// 合约地址
	res.ContractAddress = gjson.Get(execRes, "deployedTo").String()
	// 读取ABI
	abiFilePath := "/foundry/out/" + fileName + "/" + result[1] + ".json"
	commandRead := fmt.Sprintf("cat %s", abiFilePath)
	argsRead := []string{"exec", "-i", req.Address, "bash", "-c", commandRead}
	execRead, err := execCommand("", "docker", argsRead...)
	if err != nil {
		return
	}
	//data, _ := os.ReadFile(abiFilePath)
	res.ABI = gjson.Get(execRead, "abi").String()
	// Gas消耗
	commandGas := fmt.Sprintf("tx %s --json", gjson.Get(execRes, "transactionHash").String())
	argsExec := []string{"exec", "-i", req.Address, "bash", "-c", commandGas}
	gasRes, err := execCommand("", "docker", argsExec...)
	if err != nil {
		return
	}
	res.Gas = gjson.Get(gasRes, "gas").String()
	return res, nil
}
