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
	deleteQuest()
	deleteTransaction()
	deleteChallenges()
	// Start testing
	s.HashSubmit("", request.HashSubmitRequest{Hash: QuestCreatedHash})
	waitForQuestCreated(TOKENID)
	res, err := s.PermitClaimBadge(ADDRESS, request.PermitClaimBadgeReq{
		TokenId: TOKENID,
		Score:   SCORE,
		Answer:  ANSWER,
	})
	assert.Nil(t, err)
	assert.Equal(t, "0xeddb0acc916fdcc5dff6d3f75818f146dce95974239dbaf18370af72ac37f0ca0449ec58ac2c57c1998463609770d85dc09c3489cf86c06d9c5edbce04e264391c", res, "sign should equal")
	// tokenID invalid
	_, err = s.PermitClaimBadge(ADDRESS, request.PermitClaimBadgeReq{
		TokenId: 10,
		Score:   SCORE,
		Answer:  ANSWER,
	})
	assert.EqualErrorf(t, err, "TokenIDInvalid", "")
	// clear
	deleteQuest()
	deleteTransaction()
	deleteChallenges()
}
func TestService_PermitClaimBadge2(t *testing.T) {
	// delete exist
	deleteQuest()
	deleteTransaction()
	s.HashSubmit("", request.HashSubmitRequest{Hash: QuestCreatedHash})
	waitForQuestCreated(TOKENID)
	// answer error
	_, err := s.PermitClaimBadge(ADDRESS, request.PermitClaimBadgeReq{
		TokenId: TOKENID,
		Score:   SCORE,
		Answer:  "[1,[0,1],\"trued\"]",
	})
	assert.EqualErrorf(t, err, "AnswerIncorrect", "")
	// answer length error
	_, err = s.PermitClaimBadge(ADDRESS, request.PermitClaimBadgeReq{
		TokenId: TOKENID,
		Score:   SCORE,
		Answer:  "[[0,1],\"true\"]",
	})
	assert.EqualErrorf(t, err, "UnexpectedError", "")

	// clear
	deleteQuest()
	deleteTransaction()
}

func TestService_SubmitClaimTweet(t *testing.T) {
	// delete exist
	deleteTransaction()
	deleteQuest()
	deleteBadgeTweet()
	// Start testing
	s.HashSubmit("", request.HashSubmitRequest{Hash: QuestCreatedHash})
	waitForQuestCreated(TOKENID)
	err := s.SubmitClaimTweet(ADDRESS, request.SubmitClaimTweetReq{
		TokenId:  TOKENID,
		TweetUrl: TWEETURL,
		Score:    SCORE,
		Answer:   ANSWER,
	})
	assert.Nil(t, err)

	var claimTweet model.ClaimBadgeTweet
	err = s.dao.DB().Where("token_id", TOKENID).Where("address", ADDRESS).First(&claimTweet).Error
	assert.Nil(t, err)
	assert.NotZero(t, claimTweet.AddTs)
	claimTweetExcept := model.ClaimBadgeTweet{
		ID:      claimTweet.ID,
		Address: ADDRESS,
		TokenId: TOKENID,
		Url:     TWEETURL,
		TweetId: TWEETID,
		Score:   SCORE,
		AddTs:   claimTweet.AddTs,
	}
	assert.Equal(t, claimTweetExcept, claimTweet)
	// repeated tweet
	err = s.SubmitClaimTweet(ADDRESS, request.SubmitClaimTweetReq{
		TokenId:  TOKENID,
		TweetUrl: TWEETURL,
		Score:    SCORE,
		Answer:   ANSWER,
	})
	assert.EqualErrorf(t, err, "TweetRepeated", "")

	deleteBadgeTweet()
	// invalid quest
	err = s.SubmitClaimTweet(ADDRESS, request.SubmitClaimTweetReq{
		TokenId:  10,
		TweetUrl: TWEETURL,
		Score:    SCORE,
		Answer:   ANSWER,
	})
	assert.EqualErrorf(t, err, "TokenIDInvalid", "")
	// answer error
	err = s.SubmitClaimTweet(ADDRESS, request.SubmitClaimTweetReq{
		TokenId:  TOKENID,
		TweetUrl: TWEETURL,
		Score:    SCORE,
		Answer:   "[1,[0,1],\"false\"]",
	})
	assert.EqualErrorf(t, err, "AnswerIncorrect", "")

	// clear
	deleteQuest()
	deleteTransaction()
	deleteBadgeTweet()

}

func TestService_SubmitClaimTweet2(t *testing.T) {
	// Start testing
	deleteTransaction()
	deleteQuest()
	deleteBadgeTweet()
	s.HashSubmit("", request.HashSubmitRequest{Hash: QuestCreatedHash})
	waitForQuestCreated(TOKENID)
	// answer length error
	err := s.SubmitClaimTweet(ADDRESS, request.SubmitClaimTweetReq{
		TokenId:  TOKENID,
		TweetUrl: TWEETURL,
		Score:    SCORE,
		Answer:   "[[0,1],\"true\"]",
	})
	assert.EqualErrorf(t, err, "UnexpectedError", "")
	// cannot find tweet id
	err = s.SubmitClaimTweet(ADDRESS, request.SubmitClaimTweetReq{
		TokenId:  TOKENID,
		TweetUrl: "https://twitter.com/liangjies/544t3tq/",
		Score:    SCORE,
		Answer:   ANSWER,
	})
	assert.EqualErrorf(t, err, "BrokenLink", "")
	// clear
	deleteQuest()
	deleteTransaction()
	deleteBadgeTweet()
}

func TestBadgeServiceCrash(t *testing.T) {
	s.dao.Close() // Service Crash
	// Start testing
	err := s.SubmitClaimTweet(ADDRESS, request.SubmitClaimTweetReq{
		TokenId:  TOKENID,
		TweetUrl: TWEETURL,
		Score:    SCORE,
		Answer:   ANSWER,
	})
	assert.EqualErrorf(t, err, "TokenIDInvalid", "")
	// restart
	d = dao.New(c)
	s = New(c)
}
