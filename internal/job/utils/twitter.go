package utils

import (
	"backend-go/internal/job/config"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"regexp"
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
	pattern := regexp.MustCompile(`((https?|http)://[^\s/$.?#].[^\s]*)`)
	matches := strings.Replace(pattern.FindString(tweet), "\\n", "", -1)
	tweetContent := strings.Replace(strings.Replace(tweet, matches, "", -1), " ", "", -1)
	var contentMatch bool
	var matchesConfig string
	for _, v := range c.Twitter.ClaimContent {
		matchesConfig = strings.Replace(pattern.FindString(v), "\\n", "", -1)
		configContent := strings.Replace(strings.Replace(v, matchesConfig, "", -1), " ", "", -1)
		if configContent == tweetContent {
			contentMatch = true
			break
		}
	}
	if !contentMatch {
		return false
	}
	expectURL := strings.TrimSpace(matchesConfig) + strconv.FormatInt(tokenId, 10)
	client := GetTwitterClient(c)
	res, err := client.clientReq.R().Get(strings.TrimSpace(matches))
	if err != nil {
		return false
	}
	if len(res.Header["Location"]) == 0 || res.Header["Location"][0] != expectURL {
		return false
	}

	return true
}
