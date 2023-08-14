package utils

import (
	"github.com/imroc/req/v3"
	"regexp"
	"strings"
	"time"
)

// CheckIfMatchClaimTweet
// @description: 检查推特帖子内容相符
// @param: tokenId uint64, tweet string
// @return: bool
func CheckIfMatchClaimTweet(tokenId string, tweet string) bool {
	const twitter = "@decertme"
	const twitterLink = "https://decert.me/quests/"
	// 推文包含有「@DecertMe」
	if !strings.Contains(strings.ToLower(tweet), twitter) {
		return false
	}
	// 包含挑战链接
	pattern := regexp.MustCompile(`((https?|http)://[^\s/$.?#].[^\s]*)`)
	matches := strings.Replace(pattern.FindString(tweet), "\\n", "", -1)
	expectURL := strings.TrimSpace(twitterLink) + tokenId
	client := req.C().SetTimeout(30 * time.Second).
		SetCommonRetryCount(2).SetRedirectPolicy(req.NoRedirectPolicy())
	res, err := client.R().Get(strings.TrimSpace(matches))
	if err != nil {
		return false
	}
	if len(res.Header["Location"]) == 0 || res.Header["Location"][0] != expectURL {
		return false
	}
	return true
}
