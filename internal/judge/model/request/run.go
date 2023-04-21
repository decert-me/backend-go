package request

type TryRunReq struct {
	Lang       string `json:"lang"`
	TokenID    int64  `json:"token_id"`
	QuestIndex uint8  `json:"quest_index"`
	Code       string `json:"code"`
	Input      string `json:"input"`
}
