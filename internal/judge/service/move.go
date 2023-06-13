package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"backend-go/pkg/log"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (s *Service) MoveTryTestRun(req request.TryTestRunReq) (tryRunRes response.TryRunRes, err error) {
	// Docker启动
	//lock, err := s.DockerInit(address)
	//defer lock.Unlock()
	//if err != nil {
	//	return tryRunRes, errors.New("UnexpectedError")
	//}
	// 特殊编程题目
	if len(req.SpjCode) != 0 {
		tryRunResTemp, err := s.RunTestSpecialMove(req)
		if err != nil {
			return tryRunResTemp, err
		}
		tryRunRes.Correct = tryRunResTemp.Correct
		tryRunRes.Status = tryRunResTemp.Status
		tryRunRes.Msg = tryRunResTemp.Msg
		tryRunRes.TotalTestcases = tryRunResTemp.TotalTestcases
		tryRunRes.TotalCorrect = tryRunResTemp.TotalCorrect
	}
	return
}

type runMoveReq struct {
	Address       string
	InputArray    []string
	OutputArray   []string
	CorrectAnswer string
	Input         string
	Code          string
	FunctionName  string
	SpjCode       string
}

func (s *Service) RunTestSpecialMove(req request.TryTestRunReq) (tryRunRes response.TryRunRes, err error) {
	var code string
	if req.Code != "" {
		code = req.Code
	} else {
		code = req.ExampleCode
	}
	for _, v := range req.SpjCode {
		if v.Frame == "Move" {
			runReq := runMoveReq{
				SpjCode: v.Code,
				Code:    code,
				Address: req.Address,
			}
			tryRunRes, err = s.RunSpecialMove(runReq)
			// 错误提前终止
			if err != nil || tryRunRes.Status != 3 {
				return
			}
		}
	}
	return
}

func (s *Service) RunSpecialMove(req runMoveReq) (tryRunRes response.TryRunRes, err error) {
	spjCode := req.SpjCode
	// 测试
	res, err := s.TestMove(request.ForgeTestReq{
		Code:    req.Code,
		Address: req.Address,
	}, spjCode)
	tryRunRes.TotalCorrect = res.TotalCorrect
	tryRunRes.TotalTestcases = res.TotalTestcases
	tryRunRes.Status = res.Status
	tryRunRes.Msg = res.Output

	if res.TotalCorrect != 0 && res.TotalCorrect == res.TotalTestcases {
		tryRunRes.Correct = true
	}
	return
}

func (s *Service) TestMove(req request.ForgeTestReq, spjCode string) (res response.ForgeTestRes, err error) {
	if req.Address == "" {
		req.Address = common.HexToAddress("0").String()
	}
	foundryPath := path.Join("/Users/mac/Code/move-test")
	var spjTrimmedCode string
	startIndex := strings.Index(spjCode, "{")

	if startIndex != -1 {
		spjTrimmedCode = spjCode[startIndex+1:]
	}
	codeEndIndex := strings.LastIndex(req.Code, "}")
	code := req.Code[:codeEndIndex] + "\n" + spjTrimmedCode
	// 保存代码
	fileName := time.Now().Format("20060102150405.000") + ".move"
	relativeFilePath := path.Join("sources", fileName)
	p := path.Join(foundryPath, relativeFilePath)
	if err = os.WriteFile(p, []byte(code), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}
	//contract := "--match-path=" + relativeFilePath
	args := []string{"test", "--skip-fetch-latest-git-deps"}
	execRes, err := execCommand(foundryPath, "move", args...)
	//command := fmt.Sprintf("cd /foundry && forge test %s --json", contract)
	//args := []string{"exec", "-i", req.Address, "bash", "-c", command}
	//execRes, err := execCommand("", "docker", args...)
	if err := os.Remove(p); err != nil {
		log.Errorv("os.Remove error", zap.Error(err))
	}
	if err != nil {
		return
	}
	regExp := regexp.MustCompile(`Total tests: (\d+); passed: (\d+); failed: (\d+)`)
	resultIndex := strings.LastIndex(execRes, "Test result")
	if resultIndex == -1 {
		res.Output = execRes
		res.Status = 1
		return
	}
	match := regExp.FindAllStringSubmatch(execRes[resultIndex:], 1)
	if len(match) == 0 && len(match[0]) < 3 {
		res.Status = 2
		return
	}
	res.TotalTestcases, _ = strconv.Atoi(match[0][1])
	res.TotalCorrect, _ = strconv.Atoi(match[0][2])
	res.Status = 3
	return res, nil
}
