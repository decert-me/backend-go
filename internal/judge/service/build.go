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

func (s *Service) BuildSolidity(req request.BuildReq) (res response.BuildRes, err error) {
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
	if err = os.WriteFile(p, []byte(req.Code), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}
	contract := relativeFilePath + ":" + result[1]
	args := []string{"create", contract, "--private-key=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", "--offline", "--json"}
	execRes, err := execCommand(foundryPath, "forge", args...)
	if err := os.Rename(p, p+".bak"); err != nil {
		log.Errorv("os.Remove error", zap.Error(err))
	}
	if err != nil {
		return
	}

	if !gjson.Valid(execRes) {
		res.Output = strings.Replace(execRes, req.Address+"/"+fileName, result[1]+".sol", -1)
		return
	}
	res.Output = "Compiler successful"
	// 合约地址
	res.ContractAddress = gjson.Get(execRes, "deployedTo").String()
	// 读取ABI
	abiFilePath := foundryPath + "/out/" + fileName + "/" + result[1] + ".json"
	data, _ := os.ReadFile(abiFilePath)
	res.ABI = gjson.Get(string(data), "abi").String()
	// Gas消耗
	argsTx := []string{"tx", gjson.Get(execRes, "transactionHash").String(), "--json"}
	gasRes, err := execCommand(foundryPath, "cast", argsTx...)
	if err != nil {
		return
	}
	res.Gas = gjson.Get(gasRes, "gas").String()
	return res, nil
}
