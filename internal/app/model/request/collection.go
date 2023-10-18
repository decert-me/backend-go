package request

type GetCollectionChallengeUser struct {
	CollectionID uint `json:"collection_id" form:"collection_id" binding:"required"`
	PageInfo
}
