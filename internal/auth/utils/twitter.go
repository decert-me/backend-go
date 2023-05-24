package utils

import (
	"backend-go/internal/job/config"
	"regexp"
	"strconv"
	"strings"
)

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
