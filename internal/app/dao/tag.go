package dao

import "backend-go/internal/app/model"

func (d *Dao) GetUserTags(userID uint) (tags []string, err error) {
	err = d.db.Model(&model.UsersTag{}).Select("tag.name").
		Joins("join tag on users_tag.tag_id = tag.id").
		Where("users_tag.user_id = ?", userID).
		Find(&tags).Error
	return
}

func (d *Dao) GetUserNameTagsByAddress(address string) (nickname, name string, tags []string, err error) {
	var user model.Users
	err = d.db.Model(&model.Users{}).
		Where("address = ?", address).
		First(&user).Error
	if err != nil {
		return
	}
	if user.Name != nil {
		name = *user.Name
	}
	if user.NickName != nil {
		nickname = *user.NickName
	}
	err = d.db.Model(&model.UsersTag{}).Select("tag.name").
		Joins("join tag on users_tag.tag_id = tag.id").
		Where("users_tag.user_id = ?", user.ID).
		Find(&tags).Error
	return
}
