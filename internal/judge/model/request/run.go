package request

type TryRunReq struct {
	Lang       string `json:"lang"`
	TokenID    int64  `json:"token_id"`
	QuestIndex uint8  `json:"quest_index"`
	Code       string `json:"code"`
	Input      string `json:"input"`
	Type       string `json:"type"`
}

type TryTestRunReq struct {
	Lang          string   `json:"lang"`
	CodeSnippet   string   `json:"code_snippet"`
	Code          string   `json:"code"`
	ExampleCode   string   `json:"example_code"`
	SpjCode       string   `json:"spj_code"`
	Input         string   `json:"input"`
	ExampleInput  []string `json:"example_input"`
	ExampleOutput []string `json:"example_output"`
	QuestIndex    uint8
}
