package service

import (
	"backend-go/internal/app/utils"
	"backend-go/internal/judge/model/request"
	"backend-go/pkg/log"
	"github.com/imroc/req/v3"
	"go.uber.org/zap"
	"strings"
	"time"
)

func (s *Service) TryRun(body request.TryRunReq) (result string, err error) {
	if utils.IsUUID(body.TokenID) {
		body.Quest, err = s.dao.GetQuestByUUID(body.TokenID)
	} else {
		body.Quest, err = s.dao.GetQuestByTokenID(body.TokenID)
	}
	if err != nil {
		return "", err
	}
	client := req.C().SetTimeout(180 * time.Second)
	i := 0
	var item string
	// 存活检测
	for {
		if i > 2 {
			break
		}
		w := s.W.Next()
		item = w.Item
		res, err := req.C().SetTimeout(5 * time.Second).R().SetBody(body).Get(strings.Replace(item, "v1", "", 1) + "health")
		if err == nil && res.String() == "\"ok\"" {
			w.OnInvokeSuccess()
			break
		} else {
			w.OnInvokeFault()
		}
		i++
	}
	url := item + "/run/tryRun"
	res, err := client.R().SetBody(body).Post(url)
	if err != nil {
		log.Errorv("Post error", zap.Error(err))
		return "", err
	}
	return res.String(), nil
}

func (s *Service) TryTestRun(body request.TryTestRunReq) (result string, err error) {
	client := req.C().SetTimeout(180 * time.Second)
	i := 0
	var item string
	// 存活检测
	for {
		if i > 2 {
			break
		}
		w := s.W.Next()
		item = w.Item
		res, err := req.C().SetTimeout(5 * time.Second).R().SetBody(body).Get(strings.Replace(item, "v1", "", 1) + "health")
		if err == nil && res.String() == "\"ok\"" {
			w.OnInvokeSuccess()
			break
		} else {
			w.OnInvokeFault()
		}
		i++
	}
	url := item + "/run/tryTestRun"
	res, err := client.R().SetBody(body).Post(url)
	if err != nil {
		log.Errorv("Post error", zap.Error(err))
		return "", err
	}
	return res.String(), nil
}
