package service

import (
	"backend-go/internal/app/blockchain"
	"backend-go/internal/app/global"
	"backend-go/internal/app/model"
)

func HashSubmit(address string, hash string) (err error) {
	transHash := model.Transaction{SendAddr: address, Hash: hash}
	// save
	if err = global.DB.Model(&model.Transaction{}).Create(&transHash).Error; err != nil {
		return
	}
	blockchain.TransactionCh <- transHash
	return nil

}
