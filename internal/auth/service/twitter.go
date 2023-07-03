package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/auth/model/receive"
	"backend-go/internal/auth/model/request"
	"backend-go/internal/auth/utils"
	"backend-go/pkg/log"
	"encoding/json"
	"errors"
	"fmt"
	twitterClient "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/imroc/req/v3"
	"go.uber.org/zap"
	"time"
)

// TwitterAuthorizationURL 推特获取登陆链接
func (s *Service) TwitterAuthorizationURL() (res string, err error) {
	config := oauth1.Config{
		ConsumerKey:    s.c.Auth.Twitter.ConsumerKey,
		ConsumerSecret: s.c.Auth.Twitter.ConsumerSecret,
		CallbackURL:    s.c.Auth.Twitter.CallbackURL,
		Endpoint:       twitter.AuthorizeEndpoint,
	}
	requestToken, requestSecret, err := config.RequestToken()
	if err != nil {
		log.Errorv("RequestToken error", zap.Error(err))
		return "", err
	}
	fmt.Println(requestToken, requestSecret)
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		log.Errorv("AuthorizationURL error", zap.Error(err))
		return "", err
	}
	return authorizationURL.String(), err
}

// TwitterCallback 推特回调登陆
func (s *Service) TwitterCallback(address string, req request.TwitterCallbackReq) (result interface{}, err error) {
	config := oauth1.Config{
		ConsumerKey:    s.c.Auth.Twitter.ConsumerKey,
		ConsumerSecret: s.c.Auth.Twitter.ConsumerSecret,
		CallbackURL:    s.c.Auth.Twitter.CallbackURL,
		Endpoint:       twitter.AuthorizeEndpoint,
	}
	accessToken, accessSecret, err := config.AccessToken(req.RequestToken, "secret does not matter", req.Verifier)
	if err != nil {
		log.Errorv("AccessToken error", zap.Error(err))
		return
	}

	oaConfig := oauth1.NewConfig(s.c.Auth.Twitter.ConsumerKey, s.c.Auth.Twitter.ConsumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := oaConfig.Client(oauth1.NoContext, token)
	client := twitterClient.NewClient(httpClient)

	user, _, err := client.Accounts.VerifyCredentials(nil)
	if err != nil {
		log.Errorv("VerifyCredentials error", zap.Error(err))
		return
	}
	list, _, err := client.Timelines.UserTimeline(&twitterClient.UserTimelineParams{UserID: 2648158200})
	if err != nil {
		log.Errorv("List error", zap.Error(err))
		return
	}
	fmt.Println(list)
	return
	// 查找是否已经绑定过
	binding, err := s.dao.TwitterIsBinding(user.ID)
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	if binding {
		return result, errors.New("AlreadyBinding")
	}
	// 绑定到用户
	if err = s.dao.TwitterBinding(address, user.ID, user.Name); err != nil {
		return result, err
	}
	return
}

// TwitterUserTweet 查找关于Decert的推特
func (s *Service) TwitterUserTweet(claimReq request.TwitterClaimReq) (err error) {
	// 查找用户推特信息
	twitterID, err := s.dao.TwitterQueryIdByAddress(claimReq.Address)
	if err != nil || twitterID == "" {
		return errors.New("RecordNotFound")
	}
	// 查询用户推特列表
	client := req.C().
		SetTimeout(30 * time.Second).
		SetCommonRetryCount(2).
		SetUserAgent("v2TweetLookupJS").
		SetCommonBearerAuthToken(s.c.Auth.Twitter.BearerToken)
	url := fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets", twitterID)
	req, err := client.R().Get(url)
	fmt.Println(req.String())
	var tweets receive.TweetsGenerated
	if err = json.Unmarshal(req.Bytes(), &tweets); err != nil {
		log.Errorv("json.Unmarshal error", zap.Error(err))
		return err
	}
	fmt.Println(tweets.Data)
	return
	// 匹配推文
	for _, v := range tweets.Data {
		if utils.CheckIfMatchClaimTweet(twitterID, v.Text) {
			// 写入待空投列表
			s.dao.TwitterCreateTweetClaim(&model.ClaimBadgeTweet{
				Address: claimReq.Address,
				TokenId: claimReq.TokenId,
				Score:   claimReq.Score,
				TweetId: v.ID,
			})
		}
	}
	return
}
