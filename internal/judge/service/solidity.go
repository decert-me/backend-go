package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func (s *Service) RunSolidity(req request.TryRunReq) (tryRunRes response.TryRunRes, err error) {
	quest, err := s.dao.GetQuest(&model.Quest{TokenId: req.TokenID})
	if err != nil {
		return
	}
	if gjson.Get(string(quest.MetaData), fmt.Sprintf("questions.%d.type", req.QuestIndex)).Int() != 3 {
		return tryRunRes, errors.New("不是编程题目")
	}
	//input := "[\"[2,7,11,15],9\",\"[3,2,4],6\",\"[3,3],6\"]"
	//output := "[\"[0,1]\",\"[1,2]\",\"[0,1]\"]"
	inputArray := gjson.Get(string(quest.MetaData), fmt.Sprintf("questions.%d.input", req.QuestIndex)).Array()
	outputArray := gjson.Get(string(quest.MetaData), fmt.Sprintf("questions.%d.output", req.QuestIndex)).Array()
	tryRunRes.Input = req.Input
	tryRunRes.TotalTestcases = len(outputArray)
	// 编译
	contract, err := s.BuildSolidity(request.BuildReq{Code: req.Code})
	if err != nil || contract.Status == 1 {
		tryRunRes.Status = 1
		tryRunRes.Msg = contract.Output
		return
	}
	// 使用gjson解析JSON字符串
	jsonParsed := gjson.Parse(contract.ABI)

	// 获取add(uint256,uint256)(uint256)
	inputsType := jsonParsed.Get("0.inputs.#.type").Array()
	var inputsTypes []string
	for _, v := range inputsType {
		inputsTypes = append(inputsTypes, v.String())
	}

	outputType := jsonParsed.Get("0.outputs.#.type").Array()
	var outputsTypes []string
	for _, v := range outputType {
		outputsTypes = append(outputsTypes, v.String())
	}

	inputTypes := strings.Join(inputsTypes, ",")
	outputTypes := strings.Join(outputsTypes, ",")

	name := jsonParsed.Get("0.name").String()
	method := fmt.Sprintf("%v(%v)(%v)", name, inputTypes, outputTypes)
	fmt.Println(method)
	var outPut strings.Builder
	for _, v := range strings.Split(req.Input, "\n") {
		// 运行
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
		if res.Status == 1 {
			tryRunRes.Status = 2
			tryRunRes.Msg = res.Msg
			return tryRunRes, err
		}
		outPut.WriteString(res.Data)
		outPut.WriteString("\n")
	}
	tryRunRes.Output = strings.TrimRight(outPut.String(), "\n")
	tryRunRes.Status = 3
	// 检查是否通过
	for i, v := range inputArray {
		fmt.Println(v)
		res, err := s.CastCall(request.CastCallReq{
			To:     contract.ContractAddress,
			Method: method,
			Data:   v.String(),
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
	tryRunRes.JudgeID, err = s.dao.SaveJudgeResult(model.JudgeResult{
		TokenID:    req.TokenID,
		QuestIndex: req.QuestIndex,
		ScoreRaw:   gjson.Get(string(quest.MetaData), fmt.Sprintf("questions.%d.score", req.QuestIndex)).Int(),
		Pass:       true,
	})
	if err != nil {
		return
	}
	return
}

func (s *Service) RunSpecialSolidity(req request.TryRunReq) (tryRunRes response.TryRunRes, err error) {
	quest, err := s.dao.GetQuest(&model.Quest{TokenId: req.TokenID})
	if err != nil {
		return
	}
	spjCode := gjson.Get(string(quest.MetaData), fmt.Sprintf("questions.%d.spj_code", req.QuestIndex)).String()
	if spjCode == "" {
		return tryRunRes, errors.New("no spj code found")
	}
	// 编译
	res, err := s.TestSolidity(request.ForgeTestReq{
		Code:    req.Code,
		Address: "",
	}, spjCode)
	tryRunRes.TotalCorrect = res.TotalCorrect
	tryRunRes.TotalTestcases = res.TotalTestcases
	if err != nil || res.Status == 1 {
		tryRunRes.Status = 1
		tryRunRes.Msg = res.Output
		return
	}
	tryRunRes.Correct = true
	tryRunRes.JudgeID, err = s.dao.SaveJudgeResult(model.JudgeResult{
		TokenID:    req.TokenID,
		QuestIndex: req.QuestIndex,
		ScoreRaw:   gjson.Get(string(quest.MetaData), fmt.Sprintf("questions.%d.score", req.QuestIndex)).Int(),
		Pass:       true,
	})
	if err != nil {
		return
	}
	return
}
