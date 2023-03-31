package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/pkg/log"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

func (s *Service) HashSubmit(address string, req request.HashSubmitRequest) (err error) {
	transHash := model.Transaction{SendAddr: address, Hash: req.Hash, Params: []byte(gjson.Parse(req.Params).Raw)}
	// save
	if err = s.dao.CreateTransaction(&transHash); err != nil {
		log.Errorv("CreateTransaction error", zap.Error(err))
		return
	}
	return nil

}
