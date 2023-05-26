package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"backend-go/pkg/log"
	"errors"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"os"
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

func (s *Service) HardhatTestSolidity(req request.ForgeTestReq, spjCode string) (res response.ForgeTestRes, err error) {
	uuidKey := uuid.New().String()
	dirPath := uuidKey
	hardhatPath := s.c.Hardhat.Path + "/"

	// 保存代码
	fileName := time.Now().Format("20060102150405.000") + ".sol"
	if err = SaveCode(hardhatPath+dirPath+"/contracts", fileName, req.Code); err != nil {
		return
	}
	// 保存测试代码
	fileNameTest := time.Now().Format("20060102150405.000") + ".js"
	if err = SaveCode(hardhatPath+dirPath+"/test", fileNameTest, spjCode); err != nil {
		return
	}
	relativePath := hardhatPath + dirPath
	//resourcePath, resourceName := filepath.Split(s.c.Hardhat.ResourcePath)
	// 解压依赖包
	//UnArgs := []string{"xzf", resourceName, "-C", hardhatPath + dirPath}
	//_, err = execCommand(resourcePath, "tar", UnArgs...)
	//if err != nil {
	//	log.Errorv("execCommand error", zap.Any("args", UnArgs))
	//	return res, errors.New("UnexpectedError")
	//}
	// solc目录
	solcCachePath := s.c.Hardhat.SolcCachePath
	solcCacheMap := solcCachePath + ":/root/.cache"
	// docker执行
	contractDir := relativePath + "/contracts" + ":/hardhat/contracts"
	testDir := relativePath + "/test" + ":/hardhat/test"
	// node:18.16-alpine
	args := []string{"run", "-v", contractDir, "-v", testDir, "-v", solcCacheMap, "--name", uuidKey, "myimage:1.0", "sh",
		"-c", "cd /hardhat && npm install > /dev/null 2>&1 && npx hardhat test"}
	execRes, err := execCommand(relativePath, "docker", args...)
	if err != nil {
		log.Errorv("execCommand error", zap.Any("args", args))
		return res, errors.New("UnexpectedError")
	}
	//fmt.Println(execRes)
	// 运行成功删除docker
	_, err = execCommand(relativePath, "docker", "rm", uuidKey)
	if err != nil {
		log.Errorv("execCommand error", zap.Any("rm", uuidKey))
		return res, errors.New("UnexpectedError")
	}
	cleaned := strings.Replace(execRes, "\n\n", "\n", -1)
	// 清空文件
	if err := os.RemoveAll(relativePath); err != nil {
		log.Errorv("os.Remove error", zap.Error(err))
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
