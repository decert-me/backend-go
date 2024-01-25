package request

type PermitClaimBadgeReq struct {
	TokenId string `json:"tokenId" binding:"required"`
	Score   int64  `json:"score"`
	Answer  string `json:"answer" binding:"required"`
	Uri     string `json:"uri"`
}

type SubmitClaimTweetReq struct {
	TokenId  string `json:"tokenId"`
	TweetUrl string `json:"tweetUrl"`
	Score    int64  `json:"score"  binding:"required"`
	Answer   string `json:"answer" binding:"required"`
	Uri      string `json:"uri"`
}

type UpdateBadgeURIRequest struct {
	TokenId string `json:"token_id" binding:"required"`
	Uri     string `json:"uri" binding:"required"`
	ChainID int64  `json:"chain_id"`
}

type SubmitClaimShareReq struct {
	TokenId string `json:"tokenId"`
	Score   int64  `json:"score"  binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Uri     string `json:"uri"`
}

type SubmitAirdropReq struct {
	TokenId string `json:"tokenId"`
	Score   int64  `json:"score"  binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	Uri     string `json:"uri"`
}

type SubmitClaimShareV2Req struct {
	TokenId  string `json:"tokenId"`
	Score    int64  `json:"score"  binding:"required"`
	Answer   string `json:"answer" binding:"required"`
	Uri      string `json:"uri"`
	ChainID  int64  `json:"chain_id"`
	ImageUri string `json:"image_uri"`
}
