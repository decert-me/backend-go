package utils

import (
	"backend-go/internal/job/config"
	"backend-go/pkg/log"
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
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

// GetSpyderTweetById
// @description: 获取推特帖子内容
// @param: tweetId string
// @return: string, error
func GetSpyderTweetById(c *config.Config, tweetId string) (string, error) {
	// 参数设置
	options := []chromedp.ExecAllocatorOption{
		//chromedp.UserDataDir(""),
		//chromedp.ExecPath("G:\\Program Files\\Chrome\\App\\Chrome.exe"),
		//启动chrome的时候不检查默认浏览器
		chromedp.Flag("no-default-browser-check", true),
		//启动chrome 不适用沙盒, 性能优先
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("headless", true),
		//chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	// 创建chrome示例
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 150*time.Hour)
	defer cancel()
	//var textBox bool
	//cookies := strings.Split(s.c.Spyder.Cookie, ";")
	var html string
	err := chromedp.Run(ctx,
		chromedp.Tasks{
			// 打开导航
			chromedp.Navigate(fmt.Sprintf("https://twitter.com/EmberCN/status/%s", tweetId)),
			// 等待元素加载完成
			chromedp.WaitVisible("body", chromedp.ByQuery),
			// 等待5秒
			chromedp.Sleep(15 * time.Second),
			//获取需要的数据
			chromedp.OuterHTML(`html`, &html, chromedp.ByQuery),
		},
	)
	if err != nil {
		log.Errorv("err", zap.Error(err))
		return "", err
	}
	// 将字符串转换为io.Reader
	reader := strings.NewReader(html)
	// 解析HTML文档
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Errorv("NewDocumentFromReader err", zap.Error(err))
		return "", err
	}

	// 匹配<meta>标签的内容
	meta := doc.Find("meta[property='og:description']")
	content, exists := meta.Attr("content")
	if exists {
		fmt.Println(content)
		return content, err
	}
	//.Each(func(index int, meta *goquery.Selection) {
	//	content, exists := meta.Attr("content")
	//	if exists {
	//		return content
	//	}
	//})
	return "", err
}

// CheckIfMatchClaimTweet
// @description: 检查推特帖子内容相符
// @param: tokenId uint64, tweet string
// @return: bool
func CheckIfMatchClaimTweet(c *config.Config, tokenId int64, tweet string) bool {
	const twitter = "@decertme"
	const twitterLink = "https://decert.me/quests/"
	// 推文包含有「@DecertMe」
	if !strings.Contains(strings.ToLower(tweet), twitter) {
		return false
	}
	// 包含挑战链接
	pattern := regexp.MustCompile(`((https?|http)://[^\s/$.?#].[^\s]*)`)
	matches := strings.Replace(pattern.FindString(tweet), "\\n", "", -1)
	expectURL := strings.TrimSpace(twitterLink) + strconv.FormatInt(tokenId, 10)
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
