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

// GenerateMintSignatureReq - 用于用户自主 mint NFT 的签名生成请求
type GenerateMintSignatureReq struct {
	TokenId  string `json:"tokenId" binding:"required"`
	Score    int64  `json:"score" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
	Uri      string `json:"uri"`      // Quest URI (用于验证挑战是否更新)
	ImageUri string `json:"image_uri"` // 图片 IPFS URI (用于合约 mint)
	ChainID  int64  `json:"chain_id" binding:"required"`
}

// ConfirmUserMintReq - 确认用户自主 mint 成功的请求
type ConfirmUserMintReq struct {
	TokenId string `json:"token_id" binding:"required"`
	TxHash  string `json:"tx_hash" binding:"required"`
}
