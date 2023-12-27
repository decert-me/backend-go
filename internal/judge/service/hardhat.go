package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"backend-go/pkg/log"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func SaveCode(path string, fileName string, code string) (err error) {
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		log.Errorv("os.MkdirAll() Filed", zap.Error(err))
		return
	}
	p := path + "/" + fileName
	if err = os.WriteFile(p, []byte(code), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}
	return
}

func (s *Service) HardhatTestSolidity(req request.TestReq, spjCode string) (res response.TestRes, err error) {
	// 未登陆默认0地址
	if strings.TrimSpace(req.Address) == "" {
		req.Address = common.HexToAddress("0").String()
	}

	hardhatPath := path.Join(s.c.Judge.SolidityWorkPath, req.Address, "hardhat")

	// 保存代码
	fileName := time.Now().Format("20060102150405.000") + ".sol"
	if err = SaveCode(path.Join(hardhatPath, "contracts"), fileName, req.Code); err != nil {
		return
	}
	// 保存测试代码
	fileNameTest := time.Now().Format("20060102150405.000") + ".js"
	if err = SaveCode(path.Join(hardhatPath, "test"), fileNameTest, spjCode); err != nil {
		return
	}
	relativePath := hardhatPath
	// docker执行
	command := fmt.Sprintf("cd /hardhat && npx hardhat test")
	args := []string{"exec", "-i", req.Address, "bash", "-c", command}
	execRes, err := execCommand(relativePath, "docker", args...)
	if err != nil {
		log.Errorv("execCommand error", zap.Any("args", args))
		return res, errors.New("UnexpectedError")
	}
	//fmt.Println(execRes)
	cleaned := strings.Replace(execRes, "\n\n", "\n", -1)
	// 清空文件
	if err := os.Remove(path.Join(hardhatPath, "test", fileNameTest)); err != nil {
		log.Errorv("os.Remove error", zap.Error(err))
	}
	if err := os.Remove(path.Join(hardhatPath, "contracts", fileName)); err != nil {
		log.Errorv("os.Remove error", zap.Error(err))
	}
	if strings.Contains(cleaned, "Error HH600: Compilation failed") {
		res.Output = cleaned
		res.Status = 1
		return
	}
	// 正则匹配结果
	// 匹配通过数量
	rePass := regexp.MustCompile(`(\d+)\s+passing`)
	matchPass := rePass.FindAllStringSubmatch(cleaned, 1)
	// 匹配失败数量
	reFail := regexp.MustCompile(`(\d+)\s+failing`)
	matchFail := reFail.FindAllStringSubmatch(cleaned, 1)
	if len(matchFail) != 0 {
		res.Output = cleaned
		res.Status = 2
		return
	}
	res.Output = cleaned
	res.Status = 3
	var passTotal, failTotal int
	if len(matchPass) > 0 && len(matchPass[0]) > 1 {
		passTotal, _ = strconv.Atoi(matchPass[0][1])
	}
	if len(matchFail) > 0 && len(matchFail[0]) > 1 {
		failTotal, _ = strconv.Atoi(matchFail[0][1])
	}
	res.TotalTestcases = passTotal + failTotal
	res.TotalCorrect = passTotal

	return res, nil
}
