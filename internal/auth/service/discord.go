package service

import (
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"golang.org/x/oauth2"
)

func (s *Service) DiscordAuthorizationURL(callback string) string {
	var discordOauthConfig = &oauth2.Config{
		ClientID:     s.c.Auth.Discord.ClientID,
		ClientSecret: s.c.Auth.Discord.ClientSecret,
		RedirectURL:  callback,
		Scopes:       []string{"identify"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
	}
	url := discordOauthConfig.AuthCodeURL("state", oauth2.AccessTypeOnline)
	return url
}

func (s *Service) DiscordCallback(code, callback string) (id, username string, err error) {
	var discordOauthConfig = &oauth2.Config{
		ClientID:     s.c.Auth.Discord.ClientID,
		ClientSecret: s.c.Auth.Discord.ClientSecret,
		RedirectURL:  callback,
		Scopes:       []string{"identify"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://discord.com/api/oauth2/authorize",
			TokenURL: "https://discord.com/api/oauth2/token",
		},
	}
	token, err := discordOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Failed to exchange token: ", err)
		return id, username, errors.New("钱包地址绑定失败：授权失败")
	}
	client := req.C().SetCommonBearerAuthToken(token.AccessToken)
	res, err := client.R().Get("https://discord.com/api/users/@me")
	if err != nil {
		return id, username, errors.New("钱包地址绑定失败：获取用户信息失败")
	}
	id = gjson.Get(res.String(), "id").String()
	username = gjson.Get(res.String(), "username").String()
	// 跳过已绑定地址
	return id, username, err
}
