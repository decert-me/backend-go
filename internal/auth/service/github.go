package service

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// GithubAuthorizationURL Github获取登陆链接
func (s *Service) GithubAuthorizationURL() (res string, err error) {
	config := &oauth2.Config{
		ClientID:     "your-client-id",
		ClientSecret: "your-client-secret",
		Endpoint:     github.Endpoint,
	}
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return url, nil
}
