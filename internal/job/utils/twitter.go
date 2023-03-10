package utils

import (
	"backend-go/internal/job/config"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"strconv"
	"strings"
	"sync"
	"time"
)

var onceTwitter sync.Once

type configReqTwitter struct {
	clientTwitter *req.Client
	clientReq     *req.Client
}

var clientTwitter *configReqTwitter

func GetTwitterClient(c *config.Config) *configReqTwitter {
	onceTwitter.Do(func() {
		clientTwitter = new(configReqTwitter)
		clientTwitter.clientTwitter = req.C().
			SetTimeout(30 * time.Second).
			SetCommonRetryCount(2).
			SetUserAgent("v2TweetLookupJS").
			SetCommonBearerAuthToken(c.Twitter.BearToken)

		clientTwitter.clientReq = req.C().SetTimeout(30 * time.Second).
			SetCommonRetryCount(2).SetRedirectPolicy(req.NoRedirectPolicy())
	})
	return clientTwitter
}

// GetTweetById
// @description: 获取推特帖子内容
// @param: tweetId string
// @return: string, error
func GetTweetById(c *config.Config, tweetId string) (string, error) {
	client := GetTwitterClient(c)
	url := "https://api.twitter.com/2/tweets?ids=" + tweetId
	req, err := client.clientTwitter.R().Get(url)
	return gjson.Get(req.String(), "data.0.text").String(), err
}

// CheckIfMatchClaimTweet
// @description: 检查推特帖子内容相符
// @param: tokenId uint64, tweet string
// @return: bool
func CheckIfMatchClaimTweet(c *config.Config, tokenId int64, tweet string) bool {
	expect := strings.Split(c.Twitter.ClaimContent, "\n")
	actual := strings.Split(tweet, "\n")
	if len(actual) != len(expect) || strings.TrimSpace(actual[0]) != expect[0] || strings.TrimSpace(actual[2]) != expect[2] {
		return false
	}
	expectURL := strings.TrimSpace(expect[1]) + strconv.FormatInt(tokenId, 10)
	client := GetTwitterClient(c)
	res, err := client.clientReq.R().Get(strings.TrimSpace(actual[1]))
	if err != nil {
		return false
	}
	if len(res.Header["Location"]) == 0 || res.Header["Location"][0] != expectURL {
		return false
	}
	return true
}
