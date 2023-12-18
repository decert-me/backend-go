package service

import (
	"backend-go/pkg/log"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
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
	r, err := client.R().SetBodyJsonMarshal(discordCallback).Post(s.c.Social.Wechat.CallURL + "/v1/callback/discord")
	if err != nil {
		return errors.New("FailedObtainDiscordInfo")
	}
	fmt.Println(r.String())
	if gjson.Get(r.String(), "status").Int() != 0 {
		return errors.New("FailedObtainDiscordInfo")
	}
	discordID := gjson.Get(r.String(), "data.id").String()
	username := gjson.Get(r.String(), "data.username").String()
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
