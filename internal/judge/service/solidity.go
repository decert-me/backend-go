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
	if questType == "coding" {
		return s.RunSolidity(req, quest)
	} else if questType == "special_judge_coding" {
		return s.RunSpecialSolidity(req, quest)
	} else {
		return tryRunRes, errors.New("不是编程题目")
	}
}

func (s *Service) TryTestRun(req request.TryTestRunReq) (tryRunRes response.TryTestRunRes, err error) {
	if req.SpjCode == "" {
		return s.RunTestSolidity(req)
	} else {
		return s.RunTestSpecialSolidity(req)
	}
}

func (s *Service) RunTestSolidity(req request.TryTestRunReq) (tryRunRes response.TryTestRunRes, err error) {
	private := GetPrivate()
	// 获取运行函数
	var functionName string
	re := regexp.MustCompile(`function\s+(\w+)\s*\(`)
	matches := re.FindStringSubmatch(req.CodeSnippet)
	if len(matches) > 1 {
		functionName = matches[1]
	}
	// 编译
	contract, err := s.BuildSolidity(private, request.BuildReq{Code: req.Code})
	if err != nil || contract.Status == 1 {
		tryRunRes.Status = 1
		tryRunRes.Msg = contract.Output
		return
	}
	jsonParsed := gjson.Parse(contract.ABI)
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
	//fmt.Println(method)
	tryRunRes.Status = 3
	// 检查是否通过
	tryRunRes.TotalTestcases = len(req.Input)
	for i, v := range req.Input {
		//fmt.Println(v)
		res, err := s.CastCall(request.CastCallReq{
			To:     contract.ContractAddress,
			Method: method,
			Data:   v,
		})
		if err != nil {
			tryRunRes.Status = 2
			tryRunRes.Msg = err.Error()
			return tryRunRes, err
		}
		if res.Data != req.Output[i] {
			tryRunRes.Correct = false
			tryRunRes.LastInput = v
			tryRunRes.LastOutput = res.Data
			return tryRunRes, err
		}
		tryRunRes.TotalCorrect++
	}
	tryRunRes.Correct = true
	return
}

func (s *Service) RunSolidity(req request.TryRunReq, quest model.Quest) (tryRunRes response.TryRunRes, err error) {
	private := GetPrivate()
	//input := "[\"[2,7,11,15],9\",\"[3,2,4],6\",\"[3,3],6\"]"
	//output := "[\"[0,1]\",\"[1,2]\",\"[0,1]\"]"
	inputArray := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.input", req.QuestIndex)).Array()
	outputArray := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.output", req.QuestIndex)).Array()
	correctAnswerRaw := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.code_snippets.#(lang=Solidity).correctAnswer", req.QuestIndex)).String()
	correctAnswer := utils.AnswerDecode(s.c.Quest.EncryptKey, correctAnswerRaw)
	correctAnswer = gjson.Parse(correctAnswer).String()
	tryRunRes.Input = req.Input
	tryRunRes.TotalTestcases = len(outputArray)
	// 获取运行函数
	var functionName string
	codeSnippetRaw := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.code_snippets.#(lang=Solidity).code", req.QuestIndex)).String()
	re := regexp.MustCompile(`function\s+(\w+)\s*\(`)
	matches := re.FindStringSubmatch(codeSnippetRaw)
	if len(matches) > 1 {
		functionName = matches[1]
	}
	// 编译
	contract, err := s.BuildSolidity(private, request.BuildReq{Code: req.Code})
	if err != nil || contract.Status == 1 {
		tryRunRes.Status = 1
		tryRunRes.Msg = contract.Output
		return
	}
	// 编译标准答案
	correctContract, err := s.BuildSolidity(private, request.BuildReq{Code: correctAnswer})
	if err != nil || contract.Status == 1 {
		tryRunRes.Status = 1
		tryRunRes.Msg = contract.Output
		return
	}
	// quest.QuestData
	// 使用gjson解析JSON字符串
	jsonParsed := gjson.Parse(contract.ABI)
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
	//fmt.Println(method)
	var outPut, exceptOutPut strings.Builder
	for _, v := range strings.Split(req.Input, "\n") {
		if v == "" {
			continue
		}
		v = strings.Replace(v, "\n", "", -1)
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
		//fmt.Println(v)
		res, err := s.CastCall(request.CastCallReq{
			To:     contract.ContractAddress,
			Method: method,
			Data:   strings.Replace(v.String(), "\n", "", -1),
		})

		if err != nil {
			tryRunRes.Status = 2
			tryRunRes.Msg = err.Error()
			return tryRunRes, err
		}
		if res.Data != outputArray[i].String() {
			tryRunRes.Correct = false
			tryRunRes.LastInput = v.String()
			tryRunRes.LastOutput = res.Data
			tryRunRes.LastExpect = outputArray[i].String()
			return tryRunRes, err
		}
		tryRunRes.TotalCorrect++
	}
	tryRunRes.Correct = true

	return
}

func (s *Service) RunSpecialSolidity(req request.TryRunReq, quest model.Quest) (tryRunRes response.TryRunRes, err error) {
	private := GetPrivate()
	spjCode := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.spj_code", req.QuestIndex)).String()
	if spjCode == "" {
		return tryRunRes, errors.New("no spj code found")
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
	}, spjCode)
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

func (s *Service) RunTestSpecialSolidity(req request.TryTestRunReq) (tryRunRes response.TryTestRunRes, err error) {
	private := GetPrivate()
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
	}, req.SpjCode)
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
