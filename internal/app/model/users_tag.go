package model

type UsersTag struct {
	ID     uint `gorm:"primarykey" json:"id"`                                                  // 主键ID
	UserID uint `gorm:"column:user_id;comment:用户ID;index:user_id_tag_id,unique" json:"userID"` // 用户ID
	TagID  uint `gorm:"column:tag_id;comment:标签ID;index:user_id_tag_id,unique" json:"tagID"`   // 标签ID
}

func (UsersTag) TableName() string {
	return "users_tag"
}
