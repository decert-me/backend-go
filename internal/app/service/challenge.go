package service

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
)

func (s *Service) GetUserChallengeList(req request.GetChallengeListRequest) (res []response.GetChallengeListRes, total int64, err error) {
	if req.Type == 1 {
		res, total, err = s.dao.GetChallengeNotClaimList(&req)
		return
	}
	if req.Type == 2 {
		res, total, err = s.dao.GetChallengeList(&req)
		return
	}
	if req.Address == req.ReqAddress {
		res, total, err = s.dao.GetOwnerChallengeList(&req)
	} else {
		res, total, err = s.dao.GetChallengeList(&req)
	}
	return
}
