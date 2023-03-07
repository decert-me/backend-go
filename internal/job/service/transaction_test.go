package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/job/dao"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestBlockChain_StartTransaction(t *testing.T) {
	hash := "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa"
	deleteQuest()
	deleteTransaction()
	time.Sleep(time.Second * 1)
	err := s.dao.CreateTransaction(&model.Transaction{
		Hash: hash,
	})
	assert.Nil(t, err)
	s.handleTransactionReceipt(taskTx{task: &model.Transaction{Hash: hash}, txMap: new(sync.Map), countMap: new(sync.Map)})
	//b.TaskChain <- model.Transaction{Hash: hash}
	//waitForQuestCreated(10003)
	var transaction model.Transaction
	err = s.dao.DB().Where("hash", hash).First(&transaction).Error
	assert.Nil(t, err)
	assert.Equal(t, uint8(1), transaction.Status)
	deleteQuest()
	// fail for transaction
	hashFail := "0x25ec0b145615fb231fe5ce3546ee5ac84cfd2af154001b672b409c76bb493ce2"
	hashWait := "0x3e47e241a2b7a5bcaecacc89c563b1eb70231222b561369c82d9f951d39b75f1"
	err = s.dao.CreateTransaction(&model.Transaction{
		Hash: hashFail,
	})
	assert.Nil(t, err)
	err = s.dao.CreateTransaction(&model.Transaction{
		Hash: hashWait,
	})
	//assert.Nil(t, err)
	s.handleTransactionReceipt(taskTx{task: &model.Transaction{Hash: hashFail}, txMap: new(sync.Map), countMap: new(sync.Map)})
	s.TaskChain <- taskTx{task: &model.Transaction{Hash: hashWait}, txMap: new(sync.Map), countMap: new(sync.Map)}
	var transactionFail model.Transaction
	err = s.dao.DB().Where("hash", hashFail).First(&transactionFail).Error
	assert.Nil(t, err)
	assert.Equal(t, uint8(2), transactionFail.Status)
	// wait for transaction
	assert.Nil(t, err)
	var transactionWait model.Transaction
	err = s.dao.DB().Where("hash", hashWait).First(&transactionWait).Error
	assert.Nil(t, err)
	assert.Equal(t, uint8(0), transactionWait.Status)
	deleteTransaction()
	deleteQuest()
}

func TestTransactionServiceCrash(t *testing.T) {
	s.dao.Close() // Service Crash
	// Start testing
	s.StartTransaction()

	// restart
	s.dao = dao.New(c)
}
