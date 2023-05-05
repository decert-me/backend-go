package request

type PermitClaimBadgeReq struct {
	Score   int64  `json:"score"  binding:"required"`
	ChainID int    `json:"chain_id" binding:"required"`
	To      string `json:"to" binding:"required"`
	TokenId int64  `json:"tokenId" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
}
type SubmitClaimTweetReq struct {
	TokenId  int64  `json:"tokenId"`
	TweetUrl string `json:"tweetUrl"`
	Score    int64  `json:"score"  binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}

type UpdateBadgeURIRequest struct {
	TokenId int64  `json:"token_id" binding:"required"`
	Uri     string `json:"uri" binding:"required"`
}
