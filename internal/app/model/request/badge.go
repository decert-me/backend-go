package request

type PermitClaimBadgeReq struct {
	TokenId int64  `json:"tokenId" binding:"required"`
	Score   int64  `json:"score"`
	Answer  string `json:"answer" binding:"required"`
}
type SubmitClaimTweetReq struct {
	TokenId  int64  `json:"tokenId"`
	TweetUrl string `json:"tweetUrl"`
	Score    int64  `json:"score"  binding:"required"`
	Answer   string `json:"answer" binding:"required"`
}
