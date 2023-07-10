package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/auth/model/request"
	"backend-go/internal/auth/utils"
	"backend-go/pkg/log"
	"errors"
	"fmt"
	twitterClient "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"strconv"
)

// TwitterAuthorizationURL 推特获取登陆链接
func (s *Service) TwitterAuthorizationURL(req request.TwitterAuthorizationReq) (res string, err error) {
	config := oauth1.Config{
		ConsumerKey:    s.c.Auth.Twitter.ConsumerKey,
		ConsumerSecret: s.c.Auth.Twitter.ConsumerSecret,
		CallbackURL:    req.Callback,
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
		CallbackURL:    req.Callback,
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
	// 查找是否已经绑定过
	binding, err := s.dao.TwitterIsBinding(user.ID)
	if err != nil {
		return result, errors.New("UnexpectedError")
	}
	if binding {
		return result, errors.New("AlreadyBinding")
	}
	// 绑定到用户
	if err = s.dao.TwitterBinding(address, user.ID, user.Name, accessToken, accessSecret); err != nil {
		return result, err
	}
	return
}

// TwitterUserTweet 查找关于Decert的推特
func (s *Service) TwitterUserTweet(claimReq request.TwitterClaimReq) (err error) {
	// 查找用户推特信息
	twitterData, err := s.dao.TwitterQueryByAddress(claimReq.Address)
	if err != nil || twitterData == "" {
		return errors.New("RecordNotFound")
	}
	// 查询用户推特列表
	accessToken := gjson.Get(twitterData, "accessToken").String()
	accessSecret := gjson.Get(twitterData, "accessSecret").String()
	token := oauth1.NewToken(accessToken, accessSecret)
	oaConfig := oauth1.NewConfig(s.c.Auth.Twitter.ConsumerKey, s.c.Auth.Twitter.ConsumerSecret)
	httpClient := oaConfig.Client(oauth1.NoContext, token)
	client := twitterClient.NewClient(httpClient)
	userID := gjson.Get(twitterData, "id").Int()
	list, _, err := client.Timelines.UserTimeline(&twitterClient.UserTimelineParams{UserID: userID, Count: 50})
	if err != nil {
		log.Errorv("List error", zap.Error(err))
		return
	}
	// 匹配推文
	for _, v := range list {
		if utils.CheckIfMatchClaimTweet(strconv.Itoa(int(claimReq.TokenId)), v.Text) {
			// 写入待空投列表
			s.dao.TwitterCreateTweetClaim(&model.ClaimBadgeTweet{
				Address: claimReq.Address,
				TokenId: claimReq.TokenId,
				Score:   claimReq.Score,
				TweetId: strconv.Itoa(int(v.ID)),
			})
		}
	}
	return
}
