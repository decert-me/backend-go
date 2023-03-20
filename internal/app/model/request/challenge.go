package request

type GetChallengeListRequest struct {
	PageInfo
	Address    string
	ReqAddress string
	Type       uint8 `form:"type"`
}
