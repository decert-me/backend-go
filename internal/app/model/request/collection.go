package request

type GetCollectionChallengeUser struct {
	CollectionID string `json:"collection_id" form:"collection_id" binding:"required"`
	PageInfo
}

type CollectionClaimRequest struct {
	TokenID int64 `json:"token_id" form:"token_id" binding:"required"`
}

type CheckQuestInCollectionRequest struct {
	TokenID int64 `json:"token_id" form:"token_id" binding:"required"`
}

type GetCollectionHolderRankRequest struct {
	PageInfo
}
