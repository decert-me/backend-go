package request

import "backend-go/internal/app/model"

type TryRunReq struct {
	Lang       string      `json:"lang"`
	TokenID    string      `json:"token_id"`
	QuestIndex uint8       `json:"quest_index"`
	Code       string      `json:"code"`
	Input      string      `json:"input"`
	Type       string      `json:"type"`
	Address    string      `json:"address"`
	Quest      model.Quest `json:"quest"`
}

type SpjCodeList struct {
	Frame string `json:"frame"`
	Code  string `json:"code"`
}
type TryTestRunReq struct {
	Lang          string        `json:"lang"`
	CodeSnippet   string        `json:"code_snippet"`
	Code          string        `json:"code"`
	ExampleCode   string        `json:"example_code"`
	SpjCode       []SpjCodeList `json:"spj_code"`
	Input         string        `json:"input"`
	ExampleInput  []string      `json:"example_input"`
	ExampleOutput []string      `json:"example_output"`
	QuestIndex    uint8
	Address       string `json:"-"`
}
