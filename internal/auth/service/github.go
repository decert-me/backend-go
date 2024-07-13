package service

import (
	"backend-go/pkg/log"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// GithubAuthorizationURL Github获取登陆链接
func (s *Service) GithubAuthorizationURL(callback string) (res string, err error) {
	config := &oauth2.Config{
		ClientID:     s.c.Auth.Github.ClientID,
		ClientSecret: s.c.Auth.Github.ClientSecret,
		Endpoint:     github.Endpoint,
		RedirectURL:  callback,
	}
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return url, nil
}

// GithubCallback Github回调
func (s *Service) GithubCallback(code, callback string) (id, username string, err error) {
	config := &oauth2.Config{
		ClientID:     s.c.Auth.Github.ClientID,
		ClientSecret: s.c.Auth.Github.ClientSecret,
		Endpoint:     github.Endpoint,
		RedirectURL:  callback,
	}
	token, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Errorv("Exchange error", zap.Error(err))
		return
	}
	client := req.C().SetCommonBearerAuthToken(token.AccessToken)
	res, err := client.R().Get("https://api.github.com/user")
	if err != nil {
		log.Errorv("Get user info error", zap.Error(err))
		return
	}
	id = gjson.Get(res.String(), "id").String()
	username = gjson.Get(res.String(), "name").String()
	return id, username, err
}
