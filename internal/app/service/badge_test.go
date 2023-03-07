package service

import (
	"backend-go/internal/app/dao"
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_PermitClaimBadge(t *testing.T) {
	// delete exist
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	deleteQuest()
	deleteTransaction()
	deleteChallenges()
	// Start testing
	s.HashSubmit("", "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa")
	waitForQuestCreated(10003)
	res, err := s.PermitClaimBadge("0x7d32D1DE76acd73d58fc76542212e86ea63817d8", request.PermitClaimBadgeReq{
		TokenId: 10003,
		Score:   100,
		Answer:  "[0,[0,1],\"true\"]",
	})
	assert.Nil(t, err)
	assert.Equal(t, "0x07de3bb6c6e6c4889c25d6f1aab4f282ae41296460e97f699af5d36527ab28d7645169d4bc6aca9482623db91f42e9044335127092172d252ee895cc492ba1141b", res, "sign should equal")
	// tokenID invalid
	_, err = s.PermitClaimBadge(address, request.PermitClaimBadgeReq{
		TokenId: 10,
		Score:   100,
		Answer:  "[0,[0,1],\"true\"]",
	})
	assert.EqualErrorf(t, err, "题目不存在", "")
	// clear
	deleteQuest()
	deleteTransaction()
	deleteChallenges()
}
func TestService_PermitClaimBadge2(t *testing.T) {
	// delete exist
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	deleteQuest()
	deleteTransaction()
	s.HashSubmit("", "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa")
	waitForQuestCreated(10003)
	// answer error
	_, err := s.PermitClaimBadge(address, request.PermitClaimBadgeReq{
		TokenId: 10003,
		Score:   100,
		Answer:  "[1,[0,1],\"trued\"]",
	})
	assert.EqualErrorf(t, err, "答案错误", "")
	// answer length error
	_, err = s.PermitClaimBadge(address, request.PermitClaimBadgeReq{
		TokenId: 10003,
		Score:   100,
		Answer:  "[[0,1],\"true\"]",
	})
	assert.EqualErrorf(t, err, "出现错误", "")

	// clear
	deleteQuest()
	deleteTransaction()
}

func TestService_SubmitClaimTweet(t *testing.T) {
	// delete exist
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	// Start testing
	deleteTransaction()
	deleteQuest()
	deleteBadgeTweet()
	s.HashSubmit("", "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa")
	waitForQuestCreated(10003)
	err := s.SubmitClaimTweet(address, request.SubmitClaimTweetReq{
		TokenId:  10003,
		TweetUrl: "https://twitter.com/liangjies/status/1633028821715927041?s=20",
		Score:    100,
		Answer:   "[0,[0,1],\"true\"]",
	})
	assert.Nil(t, err)
	var claimTweet model.ClaimBadgeTweet
	err = s.dao.DB().Where("token_id", 10003).Where("address", address).First(&claimTweet).Error
	assert.Nil(t, err)
	assert.NotZero(t, claimTweet.AddTs)
	claimTweetExcept := model.ClaimBadgeTweet{
		ID:         claimTweet.ID,
		Address:    address,
		TokenId:    10003,
		Url:        "https://twitter.com/liangjies/status/1633028821715927041?s=20",
		TweetId:    "1633028821715927041",
		AddTs:      claimTweet.AddTs,
		Airdropped: false,
	}
	assert.Equal(t, claimTweetExcept, claimTweet)
	// repeated tweet
	err = s.SubmitClaimTweet(address, request.SubmitClaimTweetReq{
		TokenId:  10003,
		TweetUrl: "https://twitter.com/liangjies/status/1633028821715927041?s=20",
		Score:    100,
		Answer:   "[0,[0,1],\"true\"]",
	})
	assert.EqualErrorf(t, err, "推文重复使用", "")

	deleteBadgeTweet()
	// invalid quest
	err = s.SubmitClaimTweet(address, request.SubmitClaimTweetReq{
		TokenId:  10,
		TweetUrl: "https://twitter.com/liangjies/status/1633028821715927041?s=20",
		Score:    100,
		Answer:   "[0,[0,1],\"true\"]",
	})
	assert.EqualErrorf(t, err, "题目不存在", "")
	// answer error
	err = s.SubmitClaimTweet(address, request.SubmitClaimTweetReq{
		TokenId:  10003,
		TweetUrl: "https://twitter.com/liangjies/status/1633028821715927041?s=20",
		Score:    100,
		Answer:   "[1,[0,1],\"false\"]",
	})
	assert.EqualErrorf(t, err, "答案错误", "")

	// clear
	deleteQuest()
	deleteTransaction()
	deleteBadgeTweet()

}

func TestService_SubmitClaimTweet2(t *testing.T) {
	// delete exist
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	// Start testing
	deleteTransaction()
	deleteQuest()
	deleteBadgeTweet()
	s.HashSubmit("", "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa")
	waitForQuestCreated(10003)
	// answer length error
	err := s.SubmitClaimTweet(address, request.SubmitClaimTweetReq{
		TokenId:  10003,
		TweetUrl: "https://twitter.com/liangjies/status/1633028821715927041?s=20",
		Score:    100,
		Answer:   "[[0,1],\"true\"]",
	})
	assert.EqualErrorf(t, err, "出现错误", "")
	// cannot find tweet id
	err = s.SubmitClaimTweet(address, request.SubmitClaimTweetReq{
		TokenId:  10003,
		TweetUrl: "https://twitter.com/liangjies/544t3tq/",
		Score:    100,
		Answer:   "[0,[0,1],\"true\"]",
	})
	assert.EqualErrorf(t, err, "链接错误", "")
	// tweet cannot match
	err = s.SubmitClaimTweet(address, request.SubmitClaimTweetReq{
		TokenId:  10003,
		TweetUrl: "https://twitter.com/taylorswift13/status/1587420273325838336",
		Score:    100,
		Answer:   "[0,[0,1],\"true\"]",
	})
	assert.EqualErrorf(t, err, "推文不匹配", "")

	// clear
	deleteQuest()
	deleteTransaction()
	deleteBadgeTweet()
}

func TestBadgeServiceCrash(t *testing.T) {
	s.dao.Close() // Service Crash
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	// Start testing
	err := s.SubmitClaimTweet(address, request.SubmitClaimTweetReq{
		TokenId:  10003,
		TweetUrl: "https://twitter.com/liangjies/status/1630110919815733248",
		Score:    100,
		Answer:   "[0,[0,1],\"true\"]",
	})
	assert.EqualErrorf(t, err, "题目不存在", "")
	// restart
	d = dao.New(c)
	s = New(c)
}
