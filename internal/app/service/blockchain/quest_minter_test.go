package blockchain

import (
	"backend-go/internal/app/dao"
	"backend-go/internal/app/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHandleClaimed(t *testing.T) {
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	time.Sleep(time.Second)
	deleteQuest()
	deleteChallenges()
	deleteTransaction()
	time.Sleep(time.Second)
	// no such tokenId in quest
	b.handleTransactionReceipt(b.client, model.Transaction{Hash: "0xd4a9528e8600cab85835c4ac6282771e66d5cab6c62f9e34b0f955917f6f1511"})
	assert.Error(t, b.dao.DB().Where("token_id", 10003).Where("address", address).First(&model.UserChallenges{}).Error)
	// normal
	deleteQuest()
	deleteChallenges()
	deleteTransaction()

	b.handleTransactionReceipt(b.client, model.Transaction{Hash: "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa"})
	waitForQuestCreated(10003)
	b.handleTransactionReceipt(b.client, model.Transaction{Hash: "0xd4a9528e8600cab85835c4ac6282771e66d5cab6c62f9e34b0f955917f6f1511"})
	waitForClaimed(10003, address)
	var challenge model.UserChallenges
	err := d.DB().Where("token_id", 10003).Where("address", address).First(&challenge).Error
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
	err = b.handleClaimed("", &types.Log{Data: []byte("test")})
	assert.Error(t, err, "should return error when error Log")
	// clear
	deleteQuest()
	deleteChallenges()
	deleteTransaction()
}

func TestBlockChain_receiverNotClaimList(t *testing.T) {
	client, err := ethclient.Dial(b.c.BlockChain.Provider)
	assert.Nil(t, err)
	receivers := []string{"0x7d32D1DE76acd73d58fc76542212e86ea63817d8", "0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2", "0xbe866fe4bafc11ae886238772afbd24570f9b530"}
	receiversNotClaim := b.receiverNotClaimList(client, 10003, receivers)
	receiversNotClaimExpected := []common.Address{common.HexToAddress("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"), common.HexToAddress("0xbe866fe4bafc11ae886238772afbd24570f9b530")}
	assert.Equal(t, receiversNotClaimExpected, receiversNotClaim, "should filter claimed")

	receiversNotClaim = b.receiverNotClaimList(client, 9999, receivers)
	receiversNotClaimExpected = []common.Address{common.HexToAddress("0x7d32D1DE76acd73d58fc76542212e86ea63817d8"), common.HexToAddress("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"), common.HexToAddress("0xbe866fe4bafc11ae886238772afbd24570f9b530")}
	assert.Equal(t, receiversNotClaimExpected, receiversNotClaim, "should not filter")

	receiversNotClaim = b.receiverNotClaimList(client, 10003, []string{"0x7d32D1DE76acd73d58fc76542212e86ea638173232grerg43523", "0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"})
	receiversNotClaimExpected = []common.Address{common.HexToAddress("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2")}
	assert.Equal(t, receiversNotClaimExpected, receiversNotClaim, "should filter wrong addresses")

}

func TestBlockChain_AirdropBadge(t *testing.T) {
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	// ethclient dial error
	temp := b.c.BlockChain.Provider
	b.c.BlockChain.Provider = "httest://12312"
	assert.EqualErrorf(t, b.AirdropBadge(), "ethclient dial error", "")
	b.c.BlockChain.Provider = temp
	// no Airdrop list will nil error
	err := b.AirdropBadge()
	assert.Nil(t, err)
	//
	deleteBadgeTweet()
	b.dao.CreateClaimBadgeTweet(&model.ClaimBadgeTweet{
		Address: address,
		TokenId: 10003,
	})
	err = b.AirdropBadge()
	assert.Nil(t, err)
	deleteBadgeTweet()
}

func TestQuestMinterServiceCrash(t *testing.T) {
	b.dao.Close() // Service Crash
	// Start testing
	assert.EqualErrorf(t, b.AirdropBadge(), "sql: database is closed", "")
	b.handleTransactionReceipt(b.client, model.Transaction{Hash: "0xd4a9528e8600cab85835c4ac6282771e66d5cab6c62f9e34b0f955917f6f1511"})
	b.receiverNotClaimList(b.client, 10003, []string{"0x7d32D1DE76acd73d58fc76542212e86ea638173232grerg43523", "0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"})
	// restart
	b.dao = dao.New(c)
}
