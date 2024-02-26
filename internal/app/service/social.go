package service

import (
	"backend-go/internal/app/assets"
	"backend-go/pkg/log"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/jordan-wright/email"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"html/template"
	"net/smtp"
	"strings"
)

// GetWechatQrcode 获取关注二维码
func (s *Service) GetWechatQrcode(address string) (data string, err error) {
	// 发送请求
	client := req.C().SetCommonHeader("x-api-key", s.c.Social.Wechat.APIKey)
	r, err := client.R().SetQueryParams(map[string]string{"address": address, "app": "decert"}).Get(s.c.Social.Wechat.CallURL + "/wechat/getWechatQrcode")
	if err != nil {
		return
	}
	if r.StatusCode != 200 {
		return "", errors.New("UnexpectedError")
	}
	if gjson.Get(r.String(), "status").Int() != 0 {
		return "", errors.New(gjson.Get(r.String(), "message").String())
	}
	return gjson.Get(r.String(), "data").String(), err
}

// WechatBindAddress 处理地址绑定
func (s *Service) WechatBindAddress(c *gin.Context, address, fromUserName string) (err error) {
	// 校验key
	if c.GetHeader("x-api-key") != s.c.Social.Wechat.APIKey {
		log.Errorv("非法请求", zap.String("x-api-key", c.GetHeader("x-api-key")))
		return errors.New("非法请求")
	}
	// 判断是否已经绑定过
	wechatData, err := s.dao.WechatQueryByAddress(address)
	if err != nil {
		return errors.New("服务器内部错误")
	}
	if wechatData != "{}" {
		return errors.New("钱包地址已绑定，请勿重复操作")
	}
	// 判断微信是否被别的地址绑定过
	isBinding, err := s.dao.WechatIsBinding(fromUserName)
	if err != nil {
		return errors.New("服务器内部错误")
	}
	if isBinding {
		return errors.New("微信已经绑定过地址")
	}
	// 绑定
	return s.dao.WechatBindAddress(address, fromUserName)
}

// DiscordAuthorizationURL 获取 Discord 授权链接
func (s *Service) DiscordAuthorizationURL(callback string) (data string, err error) {
	// 发送请求
	client := req.C().SetCommonHeader("x-api-key", s.c.Social.Discord.APIKey)
	r, err := client.R().SetQueryParam("callback", callback).Get(s.c.Social.Discord.CallURL + "/v1/authorization/discord")
	if err != nil {
		return "", errors.New("FailedObtainDiscordInfo")
	}
	if r.StatusCode != 200 {
		return "", errors.New("FailedObtainDiscordInfo")
	}
	if gjson.Get(r.String(), "status").Int() != 0 {
		return "", errors.New("FailedObtainDiscordInfo")
	}
	data = gjson.Get(r.String(), "data").String()
	if data == "" {
		return "", errors.New("FailedObtainDiscordInfo")
	}
	return data, nil
}

// DiscordCallback Discord 回调绑定
func (s *Service) DiscordCallback(address string, discordCallback interface{}) (err error) {
	// 跳过已绑定地址
	discordData, _ := s.dao.DiscordQueryByAddress(address)
	if string(discordData) != "{}" {
		return errors.New("AddressAlreadyLinkedDiscord")
	}
	// 发送请求获取 Discord 用户信息
	client := req.C().SetCommonHeader("x-api-key", s.c.Social.Wechat.APIKey)
	r, err := client.R().SetBodyJsonMarshal(discordCallback).Post(s.c.Social.Discord.CallURL + "/v1/callback/discord")
	if err != nil {
		return errors.New("FailedObtainDiscordInfo")
	}
	fmt.Println(r.String())
	if gjson.Get(r.String(), "status").Int() != 0 {
		return errors.New("FailedObtainDiscordInfo")
	}
	discordID := gjson.Get(r.String(), "data.id").String()
	username := gjson.Get(r.String(), "data.username").String()
	if discordID == "" || username == "" {
		return errors.New("FailedObtainDiscordInfo")
	}
	// 跳过已绑定 Discord
	Binding, err := s.dao.DiscordIsBinding(discordID)
	if err != nil {
		return errors.New("UnexpectedError")
	}
	if Binding {
		return errors.New("DiscordAlreadyLinked")
	}

	return s.dao.DiscordBindAddress(discordID, username, address)
}

// GetEmailBindCode 获取邮箱绑定验证码
func (s *Service) GetEmailBindCode(address, emailAddress string) (err error) {
	// 发送内容
	type Code struct {
		Code string `json:"code"`
	}
	code, err := s.dao.EmailGetCode(address)
	if err != nil {
		log.Errorv("Get email code error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	test := Code{Code: code}
	t, err := template.ParseFS(assets.Assets, "email.html")
	if err != nil {
		log.Errorv("Parse email template error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, test)
	if err != nil {
		log.Errorv("Execute email template error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	// 发送邮件
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = fmt.Sprintf("%s <%s>", s.c.Email.Nickname, s.c.Email.From)

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = strings.Split(emailAddress, ";")
	// 设置主题
	em.Subject = "Decert邮箱验证码"

	em.HTML = tpl.Bytes()
	//设置服务器相关的配置
	addr := fmt.Sprintf("%s:%d", s.c.Email.Host, s.c.Email.Port)

	if !s.c.Email.IsSSL {
		plainAuth := smtp.PlainAuth("", s.c.Email.From, s.c.Email.Secret, s.c.Email.Host)
		err = em.Send(addr, plainAuth)
	} else {
		plainAuth := smtp.PlainAuth("", s.c.Email.From, s.c.Email.Secret, "")
		err = em.SendWithStartTLS(addr, plainAuth, &tls.Config{InsecureSkipVerify: true})
	}
	if err != nil {
		log.Errorv("Send email error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	return nil
}

// EmailBindAddress 处理邮箱绑定
func (s *Service) EmailBindAddress(address, emailAddress, code string) (err error) {
	// 校验验证码
	emailCode, err := s.dao.EmailQueryCode(address)
	if err != nil {
		return errors.New("UnexpectedError")
	}
	if emailCode != code {
		return errors.New("EmailCaptchaError")
	}
	// 判断是否已经绑定过
	emailData, err := s.dao.EmailQueryByAddress(address)
	if err != nil {
		return errors.New("EmailCaptchaError")
	}
	if emailData != "{}" {
		return errors.New("AddressAlreadyLinkedEmail")
	}
	// 判断邮箱是否被别的地址绑定过
	isBinding, err := s.dao.EmailIsBinding(emailAddress)
	if err != nil {
		return errors.New("UnexpectedError")
	}
	if isBinding {
		return errors.New("EmailAlreadyLinked")
	}
	// 绑定
	return s.dao.EmailBindAddress(address, emailAddress)
}

// GithubAuthorizationURL 获取 Github 授权链接
func (s *Service) GithubAuthorizationURL(callback string) (data string, err error) {
	// 发送请求
	client := req.C().SetCommonHeader("x-api-key", s.c.Social.Github.APIKey)
	r, err := client.R().SetQueryParam("callback", callback).Get(s.c.Social.Github.CallURL + "/v1/authorization/github")
	if err != nil {
		return "", errors.New("FailedObtainGithubInfo")
	}
	if r.StatusCode != 200 {
		return "", errors.New("FailedObtainGithubInfo")
	}
	if gjson.Get(r.String(), "status").Int() != 0 {
		return "", errors.New("FailedObtainGithubInfo")
	}
	data = gjson.Get(r.String(), "data").String()
	if data == "" {
		return "", errors.New("FailedObtainGithubInfo")
	}
	return data, nil
}

// GithubCallback Github 回调绑定
func (s *Service) GithubCallback(address string, githubCallback interface{}) (err error) {
	// 跳过已绑定地址
	githubData, _ := s.dao.GithubQueryByAddress(address)
	if string(githubData) != "{}" {
		return errors.New("AddressAlreadyLinkedGithub")
	}
	// 发送请求获取 Github 用户信息
	client := req.C().SetCommonHeader("x-api-key", s.c.Social.Wechat.APIKey)
	r, err := client.R().SetBodyJsonMarshal(githubCallback).Post(s.c.Social.Github.CallURL + "/v1/callback/github")
	if err != nil {
		return errors.New("FailedObtainGithubInfo")
	}
	if gjson.Get(r.String(), "status").Int() != 0 {
		return errors.New("FailedObtainGithubInfo")
	}
	githubID := gjson.Get(r.String(), "data.id").String()
	username := gjson.Get(r.String(), "data.username").String()
	if githubID == "" || username == "" {
		return errors.New("FailedObtainGithubInfo")
	}
	// 跳过已绑定 Github
	Binding, err := s.dao.GithubIsBinding(githubID)
	if err != nil {
		return errors.New("UnexpectedError")
	}
	if Binding {
		return errors.New("GithubAlreadyLinked")
	}

	return s.dao.GithubBindAddress(githubID, username, address)
}
