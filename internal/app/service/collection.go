package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
)

func (s *Service) GetCollectionChallengeUser(r request.GetCollectionChallengeUser) (data response.GetCollectionChallengeUserRes, total int64, err error) {
	return s.dao.GetCollectionChallengeUserByID(r)
}

// GetCollectionQuest
func (s *Service) GetCollectionQuest(r request.GetCollectionQuestRequest) (res []response.GetQuestListRes, collection model.Collection, err error) {
	return s.dao.GetCollectionQuest(r)
}
