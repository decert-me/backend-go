package dao

import "backend-go/internal/app/model"

// GetUnreadMessage 获取未读消息
func (d *Dao) GetUnreadMessage(address string) (list []model.UserMessage, err error) {
	err = d.db.Model(&model.UserMessage{}).Where("address = ? AND is_read = false", address).Order("id desc").Limit(3).Find(&list).Error
	return
}

// ReadMessage 阅读消息
func (d *Dao) ReadMessage(address string, id int64) (err error) {
	return d.db.Model(&model.UserMessage{}).Where("address = ? AND is_read = false AND id = ?", address, id).Update("is_read", true).Error
}
