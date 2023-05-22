package service

import (
	"backend-go/internal/auth/model/request"
	"fmt"
	twitterClient "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter"
)

func (s *Service) TwitterAuthorizationURL() (res string, err error) {
	config := oauth1.Config{
		ConsumerKey:    s.c.Auth.Github.ConsumerKey,
		ConsumerSecret: s.c.Auth.Github.ConsumerSecret,
		CallbackURL:    s.c.Auth.Github.CallbackURL,
		Endpoint:       twitter.AuthorizeEndpoint,
	}
	requestToken, _, err := config.RequestToken()
	if err != nil {
		return "", err
	}
	authorizationURL, err := config.AuthorizationURL(requestToken)
	if err != nil {
		return "", err
	}
	return authorizationURL.String(), err
}

func (s *Service) TwitterCallback(req request.TwitterCallbackReq) {
	config := oauth1.Config{
		ConsumerKey:    s.c.Auth.Github.ConsumerKey,
		ConsumerSecret: s.c.Auth.Github.ConsumerSecret,
		CallbackURL:    s.c.Auth.Github.CallbackURL,
		Endpoint:       twitter.AuthorizeEndpoint,
	}
	accessToken, accessSecret, err := config.AccessToken(req.RequestToken, "secret does not matter", req.Verifiers)
	if err != nil {
		return
	}

	oaConfig := oauth1.NewConfig(s.c.Auth.Github.ConsumerKey, s.c.Auth.Github.ConsumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := oaConfig.Client(oauth1.NoContext, token)
	client := twitterClient.NewClient(httpClient)
	user, _, err := client.Accounts.VerifyCredentials(nil)
	if err != nil {
		fmt.Println(err)
	}
	user.ID
}
