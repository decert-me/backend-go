package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"backend-go/internal/judge/utils"
	"backend-go/pkg/log"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (s *Service) MoveTryRun(req request.TryRunReq) (tryRunRes response.TryRunRes, err error) {
	// 默认零地址
	if strings.TrimSpace(req.Address) == "" {
		req.Address = common.HexToAddress("0").String()
	}
	// Docker启动
	lock, err := s.DockerInit(req.Address)
	defer lock.Unlock()
	if err != nil {
		return tryRunRes, errors.New("UnexpectedError")
	}
	questType := gjson.Get(string(req.Quest.QuestData), fmt.Sprintf("questions.%d.type", req.QuestIndex)).String()

	if questType != "coding" {
		return tryRunRes, errors.New("不是编程题目")
	}
	// 普通编程题目
	input := gjson.Get(string(req.Quest.QuestData), fmt.Sprintf("questions.%d.input", req.QuestIndex)).Array()
	if len(input) != 0 {
		tryRunRes, err = s.RunNormalMove(req, req.Quest)
		// 错误提前返回
		if err != nil || tryRunRes.Status != 3 {
			return
		}
	}
	var tryRunResTemp response.TryRunRes
	// 特殊编程题目
	spjCode := gjson.Get(string(req.Quest.QuestData), fmt.Sprintf("questions.%d.spj_code", req.QuestIndex)).Array()
	if len(spjCode) != 0 {
		tryRunResTemp, err = s.RunNormalSpecialMove(req, req.Quest)
		if err != nil {
			return tryRunResTemp, err
		}
	}
	if tryRunRes.Status == 0 {
		return tryRunResTemp, nil
	}
	return response.TryRunRes{}, err
}
func (s *Service) RunNormalMove(req request.TryRunReq, quest model.Quest) (tryRunRes response.TryRunRes, err error) {
	// 获取运行函数
	var functionName string
	codeSnippetRaw := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.code_snippets.#(lang=Move).code", req.QuestIndex)).String()
	re := regexp.MustCompile(`fun\s+(\w+)\((.*)\):\s*(\((.*)\)|\w+)\s*{`)
	matches := re.FindStringSubmatch(codeSnippetRaw)
	if len(matches) < 2 {
		return tryRunRes, errors.New("UnexpectedError")
	}
	functionName = matches[1]
	// 获取模块和地址
	re = regexp.MustCompile(`module\s+(\w+)::(\w+)\s*{`)
	matches = re.FindStringSubmatch(codeSnippetRaw)
	if len(matches) < 2 {
		return tryRunRes, errors.New("UnexpectedError")
	}
	contractAddress := matches[1]
	model := matches[2]

	inputArray := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.input", req.QuestIndex)).Array()
	outputArray := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.output", req.QuestIndex)).Array()
	correctAnswerRaw := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.code_snippets.#(lang=Move).correctAnswer", req.QuestIndex)).String()
	correctAnswer := utils.AnswerDecode(s.c.Quest.EncryptKey, correctAnswerRaw)
	correctAnswer = gjson.Parse(correctAnswer).String()
	var inputArrayString, outputArrayString []string
	// 类型转换
	for _, v := range inputArray {
		inputArrayString = append(inputArrayString, v.String())
	}
	for _, v := range outputArray {
		outputArrayString = append(outputArrayString, v.String())
	}
	runReq := runMoveReq{
		Address:         req.Address,
		ContractAddress: contractAddress,
		InputArray:      inputArrayString,
		OutputArray:     outputArrayString,
		CorrectAnswer:   correctAnswer,
		Input:           req.Input,
		Code:            req.Code,
		FunctionName:    functionName,
		Model:           model,
	}
	return s.RunMove(runReq)
}
func (s *Service) RunNormalSpecialMove(req request.TryRunReq, quest model.Quest) (tryRunRes response.TryRunRes, err error) {
	spjCodeList := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.spj_code", req.QuestIndex)).Array()
	if len(spjCodeList) == 0 {
		return tryRunRes, errors.New("no spj code found")
	}
	for _, v := range spjCodeList {
		frame := gjson.Get(v.String(), "frame").String()
		spjCode := gjson.Get(v.String(), "code").String()
		if frame == "Move" {
			runReq := runMoveReq{
				SpjCode: spjCode,
				Code:    req.Code,
				Address: req.Address,
			}
			tryRunRes, err = s.RunSpecialMove(frame, runReq)
			// 错误提前终止
			if err != nil || tryRunRes.Status != 3 {
				return
			}
		}
	}
	return
}

func (s *Service) MoveTryTestRun(req request.TryTestRunReq) (tryRunRes response.TryRunRes, err error) {
	// 默认零地址
	if strings.TrimSpace(req.Address) == "" {
		req.Address = common.HexToAddress("0").String()
	}
	// Docker启动
	//lock, err := s.DockerInit(address)
	//defer lock.Unlock()
	//if err != nil {
	//	return tryRunRes, errors.New("UnexpectedError")
	//}
	// 普通编程题目
	if len(req.ExampleInput) != 0 {
		tryRunRes, err = s.RunTestMove(req)
		// 错误提前返回
		if err != nil || tryRunRes.Status != 3 || tryRunRes.Correct != true {
			return
		}
	}
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

func (s *Service) RunTestMove(req request.TryTestRunReq) (tryRunRes response.TryRunRes, err error) {
	// 获取运行函数
	var functionName string
	re := regexp.MustCompile(`fun\s+(\w+)\((.*)\):\s*(\((.*)\)|\w+)\s*{`)
	matches := re.FindStringSubmatch(req.CodeSnippet)
	if len(matches) < 2 {
		return tryRunRes, errors.New("UnexpectedError")
	}
	functionName = matches[1]
	// 获取模块和地址
	re = regexp.MustCompile(`module\s+(\w+)::(\w+)\s*{`)
	matches = re.FindStringSubmatch(req.CodeSnippet)
	if len(matches) < 2 {
		return tryRunRes, errors.New("UnexpectedError")
	}
	contractAddress := matches[1]
	model := matches[2]

	runReq := runMoveReq{
		Address:         req.Address,
		ContractAddress: contractAddress,
		InputArray:      req.ExampleInput,
		OutputArray:     req.ExampleOutput,
		CorrectAnswer:   req.ExampleCode,
		Input:           req.Input,
		Code:            req.Code,
		FunctionName:    functionName,
		Model:           model,
	}
	return s.RunMove(runReq)
}

type runMoveReq struct {
	Address         string
	ContractAddress string
	InputArray      []string
	OutputArray     []string
	CorrectAnswer   string
	Input           string
	Code            string
	Model           string
	FunctionName    string
	SpjCode         string
}

func (s *Service) RunTestSpecialMove(req request.TryTestRunReq) (tryRunRes response.TryRunRes, err error) {
	var code string
	if req.Code != "" {
		code = req.Code
	} else {
		code = req.ExampleCode
	}
	for _, v := range req.SpjCode {
		if v.Frame == "Move" || v.Frame == "Aptos" || v.Frame == "Sui" {
			runReq := runMoveReq{
				SpjCode: v.Code,
				Code:    code,
				Address: req.Address,
			}
			tryRunRes, err = s.RunSpecialMove(v.Frame, runReq)
			// 错误提前终止
			if err != nil || tryRunRes.Status != 3 {
				return
			}
		}
	}
	return
}

func (s *Service) RunSpecialMove(frame string, req runMoveReq) (tryRunRes response.TryRunRes, err error) {
	spjCode := req.SpjCode
	// 测试
	var res response.TestRes
	if frame == "Move" {
		res, err = s.TestMove(request.TestReq{
			Code:    req.Code,
			Address: req.Address,
		}, spjCode)
	} else if frame == "Aptos" {
		res, err = s.TestAptos(request.TestReq{
			Code:    req.Code,
			Address: req.Address,
		}, spjCode)
	} else if frame == "Sui" {
		res, err = s.TestSui(request.TestReq{
			Code:    req.Code,
			Address: req.Address,
		}, spjCode)
	}

	tryRunRes.TotalCorrect = res.TotalCorrect
	tryRunRes.TotalTestcases = res.TotalTestcases
	tryRunRes.Status = res.Status
	tryRunRes.Msg = res.Output

	if res.TotalCorrect != 0 && res.TotalCorrect == res.TotalTestcases {
		tryRunRes.Correct = true
	}
	return
}

func (s *Service) TestMove(req request.TestReq, spjCode string) (res response.TestRes, err error) {
	if strings.TrimSpace(req.Address) == "" {
		req.Address = common.HexToAddress("0").String()
	}
	movePath := path.Join("/Users/mac/Code/move-test")
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
	p := path.Join(movePath, relativeFilePath)
	if err = os.WriteFile(p, []byte(code), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}
	//contract := "--match-path=" + relativeFilePath
	args := []string{"test", "--skip-fetch-latest-git-deps"}
	execRes, err := execCommand(movePath, "move", args...)
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

func (s *Service) TestSui(req request.TestReq, spjCode string) (res response.TestRes, err error) {
	if strings.TrimSpace(req.Address) == "" {
		req.Address = common.HexToAddress("0").String()
	}
	suiPath := path.Join("/Users/mac/Code/sui-test")
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
	p := path.Join(suiPath, relativeFilePath)
	if err = os.WriteFile(p, []byte(code), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}
	//contract := "--match-path=" + relativeFilePath
	args := []string{"move", "test", "--skip-fetch-latest-git-deps"}
	execRes, err := execCommand(suiPath, "sui", args...)
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

func (s *Service) TestAptos(req request.TestReq, spjCode string) (res response.TestRes, err error) {
	if strings.TrimSpace(req.Address) == "" {
		req.Address = common.HexToAddress("0").String()
	}
	suiPath := path.Join("/Users/mac/Code/sui-test")
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
	p := path.Join(suiPath, relativeFilePath)
	if err = os.WriteFile(p, []byte(code), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}
	//contract := "--match-path=" + relativeFilePath
	args := []string{"move", "test", "--skip-fetch-latest-git-deps"}
	execRes, err := execCommand(suiPath, "aptos", args...)
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

func (s *Service) RunMove(req runMoveReq) (tryRunRes response.TryRunRes, err error) {
	tryRunRes.TotalTestcases = len(req.OutputArray)
	tryRunRes.Status = 1
	use := fmt.Sprintf("use %s::%s;", req.Address, req.Model)
	var output, debug strings.Builder
	output.WriteString("(a1")
	debug.WriteString("debug::print(&a1);\n")
	for i := 2; i <= len(req.OutputArray); i++ {
		output.WriteString(fmt.Sprintf(",a%d", i))
		debug.WriteString(fmt.Sprintf("debug::print(&a%d);\n", i))
	}
	output.WriteString(")")
	input := strings.Join(req.InputArray, ",")
	let := fmt.Sprintf("let %s = %s::%s(%s)", output.String(), req.Model, req.FunctionName, input)
	debugFun := fmt.Sprintf("fun debug_script() {\n %s;\n %s}", let, debug.String())
	script := fmt.Sprintf("script {\n use std::debug;\n %s \n %s \n}", use, debugFun)

	// Code
	code := fmt.Sprintf("%s \n %s", req.Code, script)
	// 执行代码
	foundryPath := path.Join("/Users/mac/Code/move-test")
	// 保存代码
	fileName := time.Now().Format("20060102150405.000") + ".move"
	relativeFilePath := path.Join("sources", fileName)
	p := path.Join(foundryPath, relativeFilePath)
	if err = os.WriteFile(p, []byte(code), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}
	// 标准答案
	correctAnswer := fmt.Sprintf("%s \n %s", req.CorrectAnswer, script)
	fileNameCorrect := time.Now().Format("20060102150405.000") + ".move"
	relativeFilePathCorrect := path.Join("sources", fileNameCorrect)
	pCorrect := path.Join(foundryPath, relativeFilePathCorrect)
	if err = os.WriteFile(pCorrect, []byte(correctAnswer), 0664); err != nil {
		log.Errorv("os.WriteFile() Filed", zap.Error(err))
		return
	}

	args := []string{"sandbox", "publish", "--ignore-breaking-changes", "--skip-fetch-latest-git-deps"}
	_, err = execCommand(foundryPath, "move", args...)
	if err != nil {
		return
	}
	tryRunRes.Status = 3
	args = []string{"sandbox", "run", relativeFilePath, "--skip-fetch-latest-git-deps"}
	execRes, err := execCommand(foundryPath, "move", args...)
	//command := fmt.Sprintf("cd /foundry && forge test %s --json", contract)
	//args := []string{"exec", "-i", req.Address, "bash", "-c", command}
	//execRes, err := execCommand("", "docker", args...)
	if err := os.Remove(p); err != nil {
		log.Errorv("os.Remove error", zap.Error(err))
	}
	fmt.Println(execRes)
	// 定义正则表达式
	re := regexp.MustCompile(`\[debug\]\s+(\d+)`)
	// 匹配字符串
	matches := re.FindAllStringSubmatch(execRes, -1)
	// 判题
	for i, match := range matches {
		if match[1] != req.OutputArray[i] {
			tryRunRes.LastInput = req.InputArray[i]
			tryRunRes.LastOutput = match[1]
			tryRunRes.LastExpect = req.OutputArray[i]
			tryRunRes.Correct = false
			return
		}
	}

	args = []string{"sandbox", "run", relativeFilePathCorrect, "--skip-fetch-latest-git-deps"}
	execRes, err = execCommand(foundryPath, "move", args...)
	//command := fmt.Sprintf("cd /foundry && forge test %s --json", contract)
	//args := []string{"exec", "-i", req.Address, "bash", "-c", command}
	//execRes, err := execCommand("", "docker", args...)
	if err := os.Remove(pCorrect); err != nil {
		log.Errorv("os.Remove error", zap.Error(err))
	}
	//fmt.Println(execRes)
	// 匹配字符串
	matches = re.FindAllStringSubmatch(execRes, -1)
	// 判题
	for i, match := range matches {
		if match[1] != req.OutputArray[i] {
			tryRunRes.LastInput = req.InputArray[i]
			tryRunRes.LastOutput = match[1]
			tryRunRes.LastExpect = req.OutputArray[i]
			tryRunRes.Correct = false
			return
		}
		tryRunRes.TotalCorrect += 1
	}

	tryRunRes.Correct = true
	return
}
