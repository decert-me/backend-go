package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/pkg/log"
	"errors"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
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

func (s *Service) CreateChallengeLog(req request.SaveChallengeLogRequest) (err error) {
	// 校验分数正确性
	quest, err := s.dao.GetQuest(&model.Quest{TokenId: req.TokenId})
	if err != nil {
		return errors.New("TokenIDInvalid")
	}
	userScore, pass, err := s.AnswerScore(s.c.Quest.EncryptKey, req.Answer, quest.Uri, quest)
	if err != nil {
		log.Errorv("AnswerCheck error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	err = s.dao.CreateChallengeLog(&model.UserChallengeLog{
		Address:   req.Address,
		TokenId:   req.TokenId,
		Answer:    []byte(gjson.Parse(req.Answer).Raw),
		UserScore: userScore,
		Pass:      pass,
	})
	if err != nil {
		return errors.New("OperationFailed")
	}
	return nil
}
