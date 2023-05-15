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
	Lang        string `json:"lang"`
	CodeSnippet string `json:"code_snippet"`
	Code        string `json:"code"`
	SpjCode     string `json:"spj_code"`
	Input       string `json:"input"`
	Output      string `json:"output"`
	QuestIndex  uint8  `json:"quest_index"`
}
