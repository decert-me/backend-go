package service

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
)

func (s *Service) GetCollectionChallengeUser(r request.GetCollectionChallengeUser) (data response.GetCollectionChallengeUserRes, total int64, err error) {
	return s.dao.GetCollectionChallengeUserByID(r)
}
