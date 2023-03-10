package utils

import (
	"regexp"
	"strings"
)

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
