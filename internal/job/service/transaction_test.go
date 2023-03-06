package service

import (
	"backend-go/internal/app/dao"
	"backend-go/internal/app/model"
	"backend-go/internal/job/service/blockchain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockChain_StartTransaction(t *testing.T) {
	hash := "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa"
	blockchain.deleteQuest()
	blockchain.deleteTransaction()
	err := blockchain.b.dao.CreateTransaction(&model.Transaction{
		Hash: hash,
	})
	assert.Nil(t, err)
	blockchain.b.handleTransactionReceipt(blockchain.b.client, model.Transaction{Hash: hash})
	//b.TaskChain <- model.Transaction{Hash: hash}
	//waitForQuestCreated(10003)
	var transaction model.Transaction
	err = blockchain.b.dao.DB().Where("hash", hash).First(&transaction).Error
	assert.Nil(t, err)
	assert.Equal(t, uint8(1), transaction.Status)
	blockchain.deleteQuest()
	// fail for transaction
	hashFail := "0x25ec0b145615fb231fe5ce3546ee5ac84cfd2af154001b672b409c76bb493ce2"
	hashWait := "0x3e47e241a2b7a5bcaecacc89c563b1eb70231222b561369c82d9f951d39b75f1"
	err = blockchain.b.dao.CreateTransaction(&model.Transaction{
		Hash: hashFail,
	})
	assert.Nil(t, err)
	err = blockchain.b.dao.CreateTransaction(&model.Transaction{
		Hash: hashWait,
	})
	//assert.Nil(t, err)
	blockchain.b.handleTransactionReceipt(blockchain.b.client, model.Transaction{Hash: hashFail})
	blockchain.b.TaskChain <- model.Transaction{Hash: hashWait}
	var transactionFail model.Transaction
	err = blockchain.b.dao.DB().Where("hash", hashFail).First(&transactionFail).Error
	assert.Nil(t, err)
	assert.Equal(t, uint8(2), transactionFail.Status)
	// wait for transaction
	assert.Nil(t, err)
	var transactionWait model.Transaction
	err = blockchain.b.dao.DB().Where("hash", hashWait).First(&transactionWait).Error
	assert.Nil(t, err)
	assert.Equal(t, uint8(0), transactionWait.Status)
	blockchain.deleteTransaction()
	blockchain.deleteQuest()
}

func TestTransactionServiceCrash(t *testing.T) {
	blockchain.b.dao.Close() // Service Crash
	// Start testing
	blockchain.b.StartTransaction()

	// restart
	blockchain.b.dao = dao.New(blockchain.c)
}

func TestBlockChain_handleTransaction(t *testing.T) {
	assert.Equal(t, true, blockchain.b.traversed.Load())
	blockchain.b.handleTransaction()
	assert.Equal(t, true, blockchain.b.traversed.Load())
}
