package utils

import (
	"backend-go/internal/app/config"
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

// GetTweetIdFromURL
// @description: 获取推特帖子ID
// @param: url string
// @return: tweetId string
func GetTweetIdFromURL(url string) (tweetId string) {
	url = strings.Split(url, "?")[0]
	reg := regexp.MustCompile(`https://twitter.com/.+\/status\/(\d+)`)
	result := reg.FindAllStringSubmatch(url, -1)
	if len(result) == 0 || len(result[0]) < 2 {
		return
	}
	return result[0][1]
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
