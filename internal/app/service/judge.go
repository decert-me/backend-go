package service

import (
	"backend-go/internal/judge/model/request"
	"backend-go/pkg/log"
	"github.com/imroc/req/v3"
	"go.uber.org/zap"
	"time"
)

func (s *Service) TryRun(body request.TryRunReq) (result string, err error) {
	body.Quest, err = s.dao.GetQuestByTokenID(body.TokenID)
	if err != nil {
		return "", err
	}
	client := req.C().SetTimeout(180 * time.Second)
	url := s.W.Next().Item + "/run/tryRun"
	res, err := client.R().SetBody(body).Post(url)
	if err != nil {
		log.Errorv("Post error", zap.Error(err))
	}
	return res.String(), nil
}
