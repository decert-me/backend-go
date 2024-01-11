package request

type PermitClaimBadgeReq struct {
	TokenId int64  `json:"tokenId" binding:"required"`
	Score   int64  `json:"score"`
	Answer  string `json:"answer" binding:"required"`
	Uri     string `json:"uri"`
}

type SubmitClaimTweetReq struct {
	TokenId  int64  `json:"tokenId"`
	TweetUrl string `json:"tweetUrl"`
	Score    int64  `json:"score"  binding:"required"`
	Answer   string `json:"answer" binding:"required"`
	Uri      string `json:"uri"`
}

type UpdateBadgeURIRequest struct {
	TokenId int64  `json:"token_id" binding:"required"`
	Uri     string `json:"uri" binding:"required"`
	ChainID int64  `json:"chainId"`
}

type SubmitClaimShareReq struct {
	TokenId int64  `json:"tokenId"`
	Score   int64  `json:"score"  binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Uri     string `json:"uri"`
}

type SubmitAirdropReq struct {
	TokenId int64  `json:"tokenId"`
	Score   int64  `json:"score"  binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Uri     string `json:"uri"`
}

type SubmitClaimShareV2Req struct {
	TokenId int64  `json:"tokenId"`
	Score   int64  `json:"score"  binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Uri     string `json:"uri"`
	ChainID int64  `json:"chainId"`
}
