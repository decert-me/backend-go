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
	deleteQuest()
	deleteTransaction()
	time.Sleep(time.Second * 1)
	err := s.dao.CreateTransaction(&model.Transaction{
		Hash: QuestCreatedHash,
	})
	assert.Nil(t, err)
	s.handleTransactionReceipt(taskTx{task: &model.Transaction{Hash: QuestCreatedHash}, txMap: new(sync.Map), countMap: new(sync.Map)})
	waitForQuestCreated(TOKENID)
	var transaction model.Transaction
	err = s.dao.DB().Where("hash", QuestCreatedHash).First(&transaction).Error
	assert.Nil(t, err)
	assert.Equal(t, uint8(1), transaction.Status)
	deleteQuest()
	// fail for transaction
	err = s.dao.CreateTransaction(&model.Transaction{
		Hash: FailHash,
	})
	assert.Nil(t, err)
	err = s.dao.CreateTransaction(&model.Transaction{
		Hash: WaitHash,
	})
	//assert.Nil(t, err)
	s.handleTransactionReceipt(taskTx{task: &model.Transaction{Hash: FailHash}, txMap: new(sync.Map), countMap: new(sync.Map)})
	s.TaskChain <- taskTx{task: &model.Transaction{Hash: WaitHash}, txMap: new(sync.Map), countMap: new(sync.Map)}
	var transactionFail model.Transaction
	err = s.dao.DB().Where("hash", FailHash).First(&transactionFail).Error
	assert.Nil(t, err)
	assert.Equal(t, uint8(2), transactionFail.Status)
	// wait for transaction
	assert.Nil(t, err)
	var transactionWait model.Transaction
	err = s.dao.DB().Where("hash", WaitHash).First(&transactionWait).Error
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
