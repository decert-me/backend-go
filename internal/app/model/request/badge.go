package request

type PermitClaimBadgeReq struct {
	TokenId int64 `json:"tokenId" binding:"required"`
	Score   int64 `json:"score,omitempty"`
}
type SubmitClaimTweetReq struct {
	TokenId  int64  `json:"tokenId"`
	TweetUrl string `json:"tweetUrl"`
}
