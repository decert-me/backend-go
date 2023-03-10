package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/job/config"
	"backend-go/internal/job/dao"
	"backend-go/internal/job/initialize"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"math/big"
	"sync"
	"testing"
)

func TestHandleClaimed(t *testing.T) {
	deleteQuest()
	deleteChallenges()
	deleteTransaction()

	// no such tokenId in quest
	s.handleTransactionReceipt(taskTx{task: &model.Transaction{Hash: ClaimHash}, txMap: new(sync.Map), countMap: new(sync.Map)})
	assert.Error(t, s.dao.DB().Where("token_id", TOKENID).Where("address", ADDRESS).First(&model.UserChallenges{}).Error)
	// normal
	deleteQuest()
	deleteChallenges()
	deleteTransaction()

	s.handleTransactionReceipt(taskTx{task: &model.Transaction{Hash: QuestCreatedHash}, txMap: new(sync.Map), countMap: new(sync.Map)})
	waitForQuestCreated(TOKENID)
	s.handleTransactionReceipt(taskTx{task: &model.Transaction{Hash: ClaimHash}, txMap: new(sync.Map), countMap: new(sync.Map)})
	waitForClaimed(TOKENID, ADDRESS)
	var challenge model.UserChallenges
	err := d.DB().Where("token_id", TOKENID).Where("address", ADDRESS).First(&challenge).Error
	assert.Nil(t, err)
	assert.NotZero(t, challenge.AddTs)
	assert.NotZero(t, challenge.UpdateTs)
	assert.NotZero(t, challenge.ClaimTs)
	challengeExpected := model.UserChallenges{
		ID:       challenge.ID,
		Address:  ADDRESS,
		TokenId:  TOKENID,
		Status:   2,
		Claimed:  true,
		ClaimTs:  challenge.ClaimTs,
		AddTs:    challenge.AddTs,
		UpdateTs: challenge.UpdateTs,
	}
	assert.Equal(t, challengeExpected, challenge)
	//
	err = s.handleClaimed("", &types.Log{Data: []byte("test")})
	assert.Error(t, err, "should return error when error Log")
	// clear
	deleteQuest()
	deleteChallenges()
	deleteTransaction()
}

func TestBlockChain_receiverNotClaimList(t *testing.T) {
	client, err := ethclient.Dial(s.w.Next().Item)
	assert.Nil(t, err)
	tokenIds := []*big.Int{big.NewInt(10003), big.NewInt(10003), big.NewInt(10003)}
	receivers := []string{"0x7d32D1DE76acd73d58fc76542212e86ea63817d8", "0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2", "0xbe866fe4bafc11ae886238772afbd24570f9b530"}
	scores := []*big.Int{big.NewInt(100), big.NewInt(100), big.NewInt(100)}
	// should filter claimed
	_, receiversNotClaim, _ := s.receiverNotClaimList(client, tokenIds, receivers, scores)
	receiversNotClaimExpected := []common.Address{common.HexToAddress("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"), common.HexToAddress("0xbe866fe4bafc11ae886238772afbd24570f9b530")}
	assert.Equal(t, receiversNotClaimExpected, receiversNotClaim, "should filter claimed")
	// should not filter
	tokenIds = []*big.Int{big.NewInt(9999), big.NewInt(9999), big.NewInt(9999)}
	_, receiversNotClaim, _ = s.receiverNotClaimList(client, tokenIds, receivers, scores)
	receiversNotClaimExpected = []common.Address{common.HexToAddress("0x7d32D1DE76acd73d58fc76542212e86ea63817d8"), common.HexToAddress("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"), common.HexToAddress("0xbe866fe4bafc11ae886238772afbd24570f9b530")}
	assert.Equal(t, receiversNotClaimExpected, receiversNotClaim, "should not filter")
	// should filter wrong addresses
	tokenIds = []*big.Int{big.NewInt(10003), big.NewInt(10003)}
	_, receiversNotClaim, _ = s.receiverNotClaimList(client, tokenIds, []string{"0x7d32D1DE76acd73d58fc76542212e86ea638173232grerg43523", "0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2"}, scores)
	receiversNotClaimExpected = []common.Address{common.HexToAddress("0xBC5Ea980BdD0436a2798Bccf8fECc61bCb0010f2")}
	assert.Equal(t, receiversNotClaimExpected, receiversNotClaim, "should filter wrong addresses")

}

func TestBlockChain_AirdropBadge(t *testing.T) {
	// ethclient dial error
	temp := *s.c.BlockChain
	*s.c.BlockChain = config.BlockChain{Provider: []config.Provider{{"httest://12312", 5}}}
	s.w = initialize.InitProvider(s.c)
	assert.EqualErrorf(t, s.AirdropBadge(), "ethclient dial error", "")
	s.c.BlockChain = &temp
	s.w = initialize.InitProvider(s.c)
	// no Airdrop list will nil error
	err := s.AirdropBadge()
	assert.Nil(t, err)
	//
	deleteBadgeTweet()
	s.dao.CreateClaimBadgeTweet(&model.ClaimBadgeTweet{
		Address: ADDRESS,
		TokenId: TOKENID,
	})
	err = s.AirdropBadge()
	assert.Nil(t, err)
	deleteBadgeTweet()
}

func TestQuestMinterServiceCrash(t *testing.T) {
	s.dao.Close() // Service Crash
	// Start testing
	assert.EqualErrorf(t, s.AirdropBadge(), "sql: database is closed", "")
	s.handleTransactionReceipt(taskTx{task: &model.Transaction{Hash: ClaimHash}, txMap: new(sync.Map), countMap: new(sync.Map)})

	client, err := ethclient.Dial(s.w.Next().Item)
	assert.Nil(t, err)
	s.receiverNotClaimList(client, []*big.Int{big.NewInt(10003)}, []string{"0x7d32D1DE76acd73d58fc76542212e86ea638173232grerg43523"}, []*big.Int{big.NewInt(100)})
	// restart
	s.dao = dao.New(c)
}
