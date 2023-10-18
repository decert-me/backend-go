package response

type GetCollectionChallengeUserRes struct {
	Users []ChallengeUsers `gorm:"users" json:"users"`
	Times int64
}

type GetCollectionQuestPageResult struct {
	List       interface{} `json:"list"`
	Collection interface{} `json:"collection"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
}

type GetCollectionChallengeUserPageDataResult struct {
	GetCollectionChallengeUserRes
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}
