package service

import (
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"go.uber.org/zap"
)

func (s *Service) HashSubmit(address string, hash string) (err error) {
	transHash := model.Transaction{SendAddr: address, Hash: hash}
	// save
	if err = s.dao.CreateTransaction(&transHash); err != nil {
		log.Errorv("CreateTransaction error", zap.Error(err))
		return
	}
	s.blockchain.TaskChain <- transHash
	return nil

}
