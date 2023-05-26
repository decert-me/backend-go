package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/utils"
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"regexp"
	"strings"
	"time"
)

func (s *Service) TryRun(req request.TryRunReq) (tryRunRes response.TryRunRes, err error) {
	quest, err := s.dao.GetQuest(&model.Quest{TokenId: req.TokenID})
	if err != nil {
		return
	}
	questType := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.type", req.QuestIndex)).String()

	if questType != "coding" {
		return tryRunRes, errors.New("不是编程题目")
	}
	// 普通编程题目
	input := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.input", req.QuestIndex)).Array()
	if len(input) != 0 {
		tryRunRes, err = s.RunNormalSolidity(req, quest)
		// 错误提前返回
		if err != nil || tryRunRes.Status != 3 {
			return
		}
	}
	var tryRunResTemp response.TryRunRes
	// 特殊编程题目
	spjCode := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.spj_code", req.QuestIndex)).Array()
	if len(spjCode) != 0 {
		tryRunResTemp, err = s.RunNormalSpecialSolidity(req, quest)
		if err != nil {
			return tryRunResTemp, err
		}
	}
	if tryRunRes.Status == 0 {
		return tryRunResTemp, nil
	}
	return
}

func (s *Service) TryTestRun(req request.TryTestRunReq) (tryRunRes response.TryRunRes, err error) {
	// 普通编程题目
	if len(req.ExampleInput) != 0 {
		tryRunRes, err = s.RunTestSolidity(req)
		// 错误提前返回
		if err != nil {
			return
		}
	}
	// 特殊编程题目
	if len(req.SpjCode) != 0 {
		tryRunResTemp, err := s.RunTestSpecialSolidity(req)
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

func (s *Service) RunTestSolidity(req request.TryTestRunReq) (tryRunRes response.TryRunRes, err error) {
	// 获取运行函数
	var functionName string
	re := regexp.MustCompile(`function\s+(\w+)\s*\(`)
	matches := re.FindStringSubmatch(req.CodeSnippet)
	if len(matches) > 1 {
		functionName = matches[1]
	}
	fmt.Println(req.ExampleInput)
	fmt.Println(req.ExampleOutput)
	runReq := runSolidityReq{
		InputArray:    req.ExampleInput,
		OutputArray:   req.ExampleOutput,
		CorrectAnswer: req.ExampleCode,
		Input:         req.Input,
		Code:          req.Code,
		FunctionName:  functionName,
	}
	return s.RunSolidity(runReq)
}

func (s *Service) RunNormalSolidity(req request.TryRunReq, quest model.Quest) (tryRunRes response.TryRunRes, err error) {
	// 获取运行函数
	var functionName string
	codeSnippetRaw := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.code_snippets.#(lang=Solidity).code", req.QuestIndex)).String()
	re := regexp.MustCompile(`function\s+(\w+)\s*\(`)
	matches := re.FindStringSubmatch(codeSnippetRaw)
	if len(matches) > 1 {
		functionName = matches[1]
	}
	inputArray := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.input", req.QuestIndex)).Array()
	outputArray := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.output", req.QuestIndex)).Array()
	correctAnswerRaw := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.code_snippets.#(lang=Solidity).correctAnswer", req.QuestIndex)).String()
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
	runReq := runSolidityReq{
		InputArray:    inputArrayString,
		OutputArray:   outputArrayString,
		CorrectAnswer: correctAnswer,
		Input:         req.Input,
		Code:          req.Code,
		FunctionName:  functionName,
	}
	return s.RunSolidity(runReq)
}

type runSolidityReq struct {
	InputArray    []string
	OutputArray   []string
	CorrectAnswer string
	Input         string
	Code          string
	FunctionName  string
	SpjCode       string
}

func (s *Service) RunSolidity(req runSolidityReq) (tryRunRes response.TryRunRes, err error) {
	private := GetPrivate()
	// 参数
	inputArray := req.InputArray
	outputArray := req.OutputArray
	correctAnswer := req.CorrectAnswer
	tryRunRes.Input = req.Input
	functionName := req.FunctionName
	tryRunRes.TotalTestcases = len(outputArray)
	// 编译
	var contract response.BuildRes
	if req.Code != "" {
		contract, err = s.BuildSolidity(private, request.BuildReq{Code: req.Code})
		if err != nil || contract.Status == 1 {
			tryRunRes.Status = 1
			tryRunRes.Msg = contract.Output
			return
		}
	}
	// 编译标准答案
	correctContract, err := s.BuildSolidity(private, request.BuildReq{Code: correctAnswer})
	if err != nil || correctContract.Status == 1 {
		tryRunRes.Status = 1
		tryRunRes.Msg = correctContract.Output
		return
	}
	fmt.Println("err", err)
	// quest.QuestData
	// 使用gjson解析JSON字符串
	var jsonParsed gjson.Result
	if req.Code != "" {
		jsonParsed = gjson.Parse(contract.ABI)
	} else {
		// 标准答案
		jsonParsed = gjson.Parse(correctContract.ABI)
	}

	// 获取函数index
	var index int
	for i, v := range jsonParsed.Get("#.name").Array() {
		if v.String() == functionName {
			index = i
			break
		}
	}
	// 获取add(uint256,uint256)(uint256)
	inputsType := jsonParsed.Get(fmt.Sprintf("%d.inputs.#.type", index)).Array()
	var inputsTypes []string
	for _, v := range inputsType {
		inputsTypes = append(inputsTypes, v.String())
	}

	outputType := jsonParsed.Get(fmt.Sprintf("%d.outputs.#.type", index)).Array()
	var outputsTypes []string
	for _, v := range outputType {
		outputsTypes = append(outputsTypes, v.String())
	}

	inputTypes := strings.Join(inputsTypes, ",")
	outputTypes := strings.Join(outputsTypes, ",")

	name := jsonParsed.Get(fmt.Sprintf("%d.name", index)).String()
	method := fmt.Sprintf("%v(%v)(%v)", name, inputTypes, outputTypes)
	fmt.Println(method)
	var outPut, exceptOutPut strings.Builder
	for _, v := range strings.Split(req.Input, "\n") {
		if v == "" {
			continue
		}
		v = strings.Replace(v, "\n", "", -1)
		v = strings.TrimSpace(v)
		// 运行
		startTime := time.Now()
		res, err := s.CastCall(request.CastCallReq{
			To:     contract.ContractAddress,
			Method: method,
			Data:   v,
		})
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		fmt.Println("程序运行时间：", elapsedTime)
		if err != nil {
			tryRunRes.Status = 2
			tryRunRes.Msg = err.Error()
			return tryRunRes, err
		}
		if res.Status == 1 {
			tryRunRes.Status = 2
			tryRunRes.Msg = res.Msg
			return tryRunRes, err
		}
		outPut.WriteString(res.Data)
		outPut.WriteString("\n")
		// 标准输出
		startTime = time.Now()
		exceptRes, err := s.CastCall(request.CastCallReq{
			To:     correctContract.ContractAddress,
			Method: method,
			Data:   v,
		})
		endTime = time.Now()
		elapsedTime = endTime.Sub(startTime)
		fmt.Println("程序运行时间：", elapsedTime)
		exceptOutPut.WriteString(exceptRes.Data)
		exceptOutPut.WriteString("\n")
	}
	tryRunRes.ExceptOutput = strings.TrimRight(exceptOutPut.String(), "\n")
	tryRunRes.Output = strings.TrimRight(outPut.String(), "\n")
	tryRunRes.Status = 3
	// 检查是否通过
	for i, v := range inputArray {
		v = strings.TrimSpace(v)
		var res response.CastCallRes
		if req.Code == "" {
			res, err = s.CastCall(request.CastCallReq{
				To:     correctContract.ContractAddress,
				Method: method,
				Data:   strings.Replace(v, "\n", "", -1),
			})
			if err != nil {
				tryRunRes.Status = 2
				tryRunRes.Msg = err.Error()
				return tryRunRes, err
			}
		} else {
			res, err = s.CastCall(request.CastCallReq{
				To:     contract.ContractAddress,
				Method: method,
				Data:   strings.Replace(v, "\n", "", -1),
			})
			if err != nil {
				tryRunRes.Status = 2
				tryRunRes.Msg = err.Error()
				return tryRunRes, err
			}
		}
		if res.Data != strings.TrimSpace(outputArray[i]) {
			tryRunRes.Correct = false
			tryRunRes.LastInput = v
			tryRunRes.LastOutput = res.Data
			tryRunRes.LastExpect = outputArray[i]
			return tryRunRes, err
		}
		tryRunRes.TotalCorrect++
	}
	tryRunRes.Correct = true

	return
}

func (s *Service) RunSpecialSolidity(req runSolidityReq) (tryRunRes response.TryRunRes, err error) {
	private := GetPrivate()
	spjCode := req.SpjCode
	// 清除
	var spjCodeNew strings.Builder
	for _, v := range strings.Split(spjCode, "\n") {
		if strings.Contains(v, "SPDX-License-Identifier") {
			continue
		}
		spjCodeNew.WriteString(v)
	}
	// 编译
	contract, err := s.BuildSolidity(private, request.BuildReq{Code: req.Code})
	if err != nil || contract.Status == 1 {
		tryRunRes.Status = 1
		tryRunRes.Msg = contract.Output
		return
	}
	// 测试
	res, err := s.TestSolidity(request.ForgeTestReq{
		Code:    req.Code,
		Address: "",
	}, spjCodeNew.String())
	tryRunRes.TotalCorrect = res.TotalCorrect
	tryRunRes.TotalTestcases = res.TotalTestcases
	if err != nil || res.Status == 1 {
		tryRunRes.Status = 2
		tryRunRes.Msg = res.Output
		return
	}
	tryRunRes.Status = 3
	tryRunRes.Correct = true
	return
}

func (s *Service) RunSpecialHardhatSolidity(req runSolidityReq) (tryRunRes response.TryRunRes, err error) {
	// 测试
	res, err := s.HardhatTestSolidity(request.ForgeTestReq{
		Code:    req.Code,
		Address: "",
	}, req.SpjCode)

	tryRunRes.TotalCorrect = res.TotalCorrect
	tryRunRes.TotalTestcases = res.TotalTestcases
	tryRunRes.Msg = res.Output
	tryRunRes.Status = res.Status
	tryRunRes.Correct = true
	return
}

func (s *Service) RunNormalSpecialSolidity(req request.TryRunReq, quest model.Quest) (tryRunRes response.TryRunRes, err error) {
	spjCodeList := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.spj_code", req.QuestIndex)).Array()
	if len(spjCodeList) == 0 {
		return tryRunRes, errors.New("no spj code found")
	}
	for _, v := range spjCodeList {
		frame := gjson.Get(v.String(), "frame").String()
		spjCode := gjson.Get(v.String(), "code").String()
		if frame == "Foundry" {
			runReq := runSolidityReq{
				SpjCode: spjCode,
				Code:    req.Code,
			}
			tryRunRes, err = s.RunSpecialSolidity(runReq)
			// 错误提前终止
			if err != nil || tryRunRes.Status != 3 {
				return
			}
		} else if frame == "Hardhat" {
			runReq := runSolidityReq{
				SpjCode: spjCode,
				Code:    req.Code,
			}
			tryRunRes, err = s.RunSpecialHardhatSolidity(runReq)
			// 错误提前终止
			if err != nil || tryRunRes.Status != 3 {
				return
			}
		}
	}
	return
}

func (s *Service) RunTestSpecialSolidity(req request.TryTestRunReq) (tryRunRes response.TryRunRes, err error) {
	for _, v := range req.SpjCode {
		if v.Frame == "Foundry" {
			runReq := runSolidityReq{
				SpjCode: v.Code,
				Code:    req.ExampleCode,
			}
			tryRunRes, err = s.RunSpecialSolidity(runReq)
			// 错误提前终止
			if err != nil || tryRunRes.Status != 3 {
				return
			}
		} else if v.Frame == "Hardhat" {
			runReq := runSolidityReq{
				SpjCode: v.Code,
				Code:    req.ExampleCode,
			}
			tryRunRes, err = s.RunSpecialHardhatSolidity(runReq)
			return tryRunRes, err
			// 错误提前终止
			if err != nil || tryRunRes.Status != 3 {
				return
			}
		}
	}
	return
}
