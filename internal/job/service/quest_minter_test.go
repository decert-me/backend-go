package service

import (
	"backend-go/internal/app/dao"
	"backend-go/internal/app/model"
	"backend-go/internal/job/service/blockchain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandleClaimed(t *testing.T) {
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	blockchain.deleteQuest()
	blockchain.deleteChallenges()
	blockchain.deleteTransaction()
	// no such tokenId in quest
	blockchain.b.handleTransactionReceipt(blockchain.b.client, model.Transaction{Hash: "0xd4a9528e8600cab85835c4ac6282771e66d5cab6c62f9e34b0f955917f6f1511"})
	assert.Error(t, blockchain.b.dao.DB().Where("token_id", 10003).Where("address", address).First(&model.UserChallenges{}).Error)
	// normal
	blockchain.deleteQuest()
	blockchain.deleteChallenges()
	blockchain.deleteTransaction()

	blockchain.b.handleTransactionReceipt(blockchain.b.client, model.Transaction{Hash: "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa"})
	blockchain.waitForQuestCreated(10003)
	blockchain.b.handleTransactionReceipt(blockchain.b.client, model.Transaction{Hash: "0xd4a9528e8600cab85835c4ac6282771e66d5cab6c62f9e34b0f955917f6f1511"})
	blockchain.waitForClaimed(10003, address)
	var challenge model.UserChallenges
	err := blockchain.d.DB().Where("token_id", 10003).Where("address", address).First(&challenge).Error
	assert.Nil(t, err)
	assert.NotZero(t, challenge.AddTs)
	assert.NotZero(t, challenge.UpdateTs)
	assert.NotZero(t, challenge.ClaimTs)
	challengeExpected := model.UserChallenges{
		ID:       challenge.ID,
		Address:  address,
		TokenId:  10003,
		Status:   2,
		Claimed:  true,
		ClaimTs:  challenge.ClaimTs,
		AddTs:    challenge.AddTs,
		UpdateTs: challenge.UpdateTs,
	}
	assert.Equal(t, challengeExpected, challenge)
	//
	err = blockchain.b.handleClaimed("", &types.Log{Data: []byte("test")})
	assert.Error(t, err, "should return error when error Log")
	// clear
	blockchain.deleteQuest()
	blockchain.deleteChallenges()
	blockchain.deleteTransaction()
}

func TestBlockChain_receiverNotClaimList(t *testing.T) {
	client, err := ethclient.Dial(blockchain.b.c.BlockChain.Provider)
	assert.Nil(t, err)
	receivers := []string{"0x7d32D1DE76acd73d58fc76542212e86ea63817d8", "0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2", "0xbe866fe4bafc11ae886238772afbd24570f9b530"}
	receiversNotClaim := blockchain.b.receiverNotClaimList(client, 10003, receivers)
	receiversNotClaimExpected := []common.Address{common.HexToAddress("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"), common.HexToAddress("0xbe866fe4bafc11ae886238772afbd24570f9b530")}
	assert.Equal(t, receiversNotClaimExpected, receiversNotClaim, "should filter claimed")

	receiversNotClaim = blockchain.b.receiverNotClaimList(client, 9999, receivers)
	receiversNotClaimExpected = []common.Address{common.HexToAddress("0x7d32D1DE76acd73d58fc76542212e86ea63817d8"), common.HexToAddress("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"), common.HexToAddress("0xbe866fe4bafc11ae886238772afbd24570f9b530")}
	assert.Equal(t, receiversNotClaimExpected, receiversNotClaim, "should not filter")

	receiversNotClaim = blockchain.b.receiverNotClaimList(client, 10003, []string{"0x7d32D1DE76acd73d58fc76542212e86ea638173232grerg43523", "0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"})
	receiversNotClaimExpected = []common.Address{common.HexToAddress("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2")}
	assert.Equal(t, receiversNotClaimExpected, receiversNotClaim, "should filter wrong addresses")

}

func TestBlockChain_AirdropBadge(t *testing.T) {
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	// ethclient dial error
	temp := blockchain.b.c.BlockChain.Provider
	blockchain.b.c.BlockChain.Provider = "httest://12312"
	assert.EqualErrorf(t, blockchain.b.AirdropBadge(), "ethclient dial error", "")
	blockchain.b.c.BlockChain.Provider = temp
	// no Airdrop list will nil error
	err := blockchain.b.AirdropBadge()
	assert.Nil(t, err)
	//
	blockchain.deleteBadgeTweet()
	blockchain.b.dao.CreateClaimBadgeTweet(&model.ClaimBadgeTweet{
		Address: address,
		TokenId: 10003,
	})
	err = blockchain.b.AirdropBadge()
	assert.Nil(t, err)
	blockchain.deleteBadgeTweet()
}

func TestQuestMinterServiceCrash(t *testing.T) {
	blockchain.b.dao.Close() // Service Crash
	// Start testing
	assert.EqualErrorf(t, blockchain.b.AirdropBadge(), "sql: database is closed", "")
	blockchain.b.handleTransactionReceipt(blockchain.b.client, model.Transaction{Hash: "0xd4a9528e8600cab85835c4ac6282771e66d5cab6c62f9e34b0f955917f6f1511"})
	blockchain.b.receiverNotClaimList(blockchain.b.client, 10003, []string{"0x7d32D1DE76acd73d58fc76542212e86ea638173232grerg43523", "0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"})
	// restart
	blockchain.b.dao = dao.New(blockchain.c)
}
