package request

type GetCollectionChallengeUser struct {
	CollectionID string `json:"collection_id" form:"collection_id" binding:"required"`
	PageInfo
}
