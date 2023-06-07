package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/judge/model/judge"
	"backend-go/internal/judge/model/request"
	"backend-go/internal/judge/model/response"
	"backend-go/pkg/log"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"strings"
	"time"
)

func (s *Service) PythonTryRun(runReq request.TryRunReq) (tryRunRes response.TryRunRes, err error) {
	quest, err := s.dao.GetQuest(&model.Quest{TokenId: runReq.TokenID})
	if err != nil {
		return
	}
	questType := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.type", runReq.QuestIndex)).String()
	if questType != "coding" {
		return tryRunRes, errors.New("不是编程题目")
	}
	inputArray := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.input", runReq.QuestIndex)).Array()
	outputArray := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.output", runReq.QuestIndex)).Array()
	codeSnippetRaw := gjson.Get(string(quest.QuestData), fmt.Sprintf("questions.%d.code_snippets.#(lang=Python).code", runReq.QuestIndex)).String()
	var inputs, outputs []string
	for _, v := range inputArray {
		inputs = append(inputs, v.String())
	}
	for _, v := range outputArray {
		outputs = append(outputs, v.String())
	}
	// 请求体
	javaScriptRun, err := s.PythonScriptRun(runReq.Code, codeSnippetRaw, inputs)
	if err != nil {
		log.Errorv("JavaScriptRun error", zap.Error(err))
		return tryRunRes, err
	}
	// 发送请求到判题沙箱
	client := req.C().SetTimeout(30 * time.Second)
	url := s.c.Judge.SandBoxService + "run"
	// 发送请求
	res, err := client.R().SetBodyJsonMarshal(javaScriptRun).Post(url)
	if err != nil {
		return tryRunRes, err
	}
	// 获取结果
	var runResult []judge.RunResult
	err = json.Unmarshal(res.Bytes(), &runResult)
	if err != nil {
		return tryRunRes, err
	}
	// 判断是否正确
	if len(runResult) != len(outputs) {
		return tryRunRes, errors.New("返回数量不相等")
	}
	for i, input := range inputs {
		if runResult[i].Status != "Accepted" {
			tryRunRes.Status = 1
			if runResult[i].Files.Stderr != "" {
				tryRunRes.Msg = runResult[i].Files.Stderr
			} else {
				tryRunRes.Msg = runResult[i].Error
			}
			return
		}
		runOutput := strings.Replace(strings.TrimRight(runResult[i].Files.Stdout, "\n"), ", ", ",", -1)
		if runOutput == outputs[i] {
			tryRunRes.TotalCorrect = tryRunRes.TotalCorrect + 1
			continue
		}
		if runResult[i].Files.Stderr != "" {
			tryRunRes.Status = 2
			tryRunRes.Msg = runResult[i].Files.Stderr
			return
		}
		tryRunRes.Status = 3
		tryRunRes.LastInput = input
		tryRunRes.LastOutput = runOutput
		tryRunRes.LastExpect = outputs[i]
		return tryRunRes, nil
	}
	tryRunRes.Correct = true
	tryRunRes.Status = 3
	return
}

func (s *Service) PythonTryTestRun(runReq request.TryTestRunReq) (tryRunRes response.TryRunRes, err error) {
	// 请求体
	javaScriptRun, err := s.PythonScriptRun(runReq.ExampleCode, runReq.CodeSnippet, runReq.ExampleInput)
	if err != nil {
		log.Errorv("JavaScriptRun error", zap.Error(err))
		return tryRunRes, err
	}
	// 发送请求到判题沙箱
	client := req.C().SetTimeout(30 * time.Second)
	url := s.c.Judge.SandBoxService + "run"
	// 输入输出数量不相等
	if len(runReq.ExampleInput) != len(runReq.ExampleOutput) {
		return tryRunRes, errors.New("输入输出数量不相等")
	}
	tryRunRes.TotalTestcases = len(runReq.ExampleInput)
	// 发送请求
	res, err := client.R().SetBodyJsonMarshal(javaScriptRun).Post(url)
	if err != nil {
		return tryRunRes, err
	}
	//fmt.Println(javaScriptRun)
	fmt.Println(res.String())
	// 获取结果
	var runResult []judge.RunResult
	err = json.Unmarshal(res.Bytes(), &runResult)
	if err != nil {
		return tryRunRes, err
	}
	// 判断是否正确
	if len(runResult) != len(runReq.ExampleInput) {
		return tryRunRes, errors.New("返回数量不相等")
	}
	for i, input := range runReq.ExampleInput {
		if runResult[i].Status != "Accepted" {
			tryRunRes.Status = 1
			if runResult[i].Files.Stderr != "" {
				tryRunRes.Msg = runResult[i].Files.Stderr
			} else {
				tryRunRes.Msg = runResult[i].Error
			}
			return
		}
		runOutput := strings.Replace(strings.TrimRight(runResult[i].Files.Stdout, "\n"), ", ", ",", -1)
		if runOutput == runReq.ExampleOutput[i] {
			tryRunRes.TotalCorrect = tryRunRes.TotalCorrect + 1
			continue
		}
		if runResult[i].Files.Stderr != "" {
			tryRunRes.Status = 2
			tryRunRes.Msg = runResult[i].Files.Stderr
			return
		}
		tryRunRes.Status = 3
		tryRunRes.LastInput = input
		tryRunRes.LastOutput = runOutput
		tryRunRes.LastExpect = runReq.ExampleOutput[i]
		return tryRunRes, nil
	}
	tryRunRes.Correct = true
	tryRunRes.Status = 3
	return
}
