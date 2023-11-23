package service

import "backend-go/internal/app/model"

// GetUnreadMessage 获取未读消息
func (s *Service) GetUnreadMessage(address string) (list []model.UserMessage, err error) {
	return s.dao.GetUnreadMessage(address)
}

// ReadMessage 阅读消息
func (s *Service) ReadMessage(address string, id int64) (err error) {
	return s.dao.ReadMessage(address, id)
}
