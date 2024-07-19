package service

import (
	"backend-go/pkg/log"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"strings"
	"time"
)

// WechatService 微信服务
func (s *Service) WechatService(c *gin.Context) (err error) {
	req := c.Request
	rw := c.Writer
	// 创建wechat对象
	wc := wechat.NewWechat()
	// 本地内存保存access_token
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:          s.c.Auth.Wechat.AppID,
		AppSecret:      s.c.Auth.Wechat.AppSecret,
		Token:          s.c.Auth.Wechat.Token,
		EncodingAESKey: s.c.Auth.Wechat.EncodingAESKey,
		Cache:          memory,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	// 传入request和responseWriter
	server := officialAccount.GetServer(req, rw)
	server.SetMessageHandler(s.messageHandler)
	//处理消息接收以及回复
	err = server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
	return
}

// messageHandler 消息处理
func (s *Service) messageHandler(msg *message.MixMessage) *message.Reply {
	switch msg.MsgType {
	case message.MsgTypeEvent:
		return s.handleEvent(msg)
	case message.MsgTypeText:
		return s.handleText(msg)
	}
	return nil
}

// handleText 处理文本消息
func (s *Service) handleText(msg *message.MixMessage) *message.Reply {
	if msg.Content == "确认" || msg.Content == "确定" {
		return s.handleConfirmation(msg)
	}
	return nil
}

// handleEvent 处理事件消息
func (s *Service) handleEvent(msg *message.MixMessage) *message.Reply {
	if msg.Event == message.EventSubscribe || msg.Event == message.EventScan {
		return s.handleSubscription(msg)
	}
	if msg.Event == message.EventClick && msg.EventKey == "V1001_GOOD" {
		return &message.Reply{MsgType: message.MsgTypeImage, MsgData: message.NewImage("hE1FXKcBCLuXNojNQvWrBrvVQgQfoMgW2Eqv6hePAkLC4MZPj7ZjQr0wvHtuIjbB")}
	}
	return nil
}

// handleSubscription 处理订阅事件
func (s *Service) handleSubscription(msg *message.MixMessage) *message.Reply {
	err := s.cache.Set("wechat::"+string(msg.FromUserName), []byte(msg.EventKey))
	if err != nil {
		log.Errorv("缓存失败", zap.Error(err))
		return nil
	}
	msgData, err := s.WechatBindAddress(msg.EventKey, string(msg.FromUserName), false)
	if err != nil && msgData != "" {
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("钱包地址绑定失败：" + msgData)}
	}
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("钱包地址绑定成功")}
}

// handleConfirmation 处理确认绑定消息
func (s *Service) handleConfirmation(msg *message.MixMessage) *message.Reply {
	eventKey, err := s.cache.Get("wechat::" + string(msg.FromUserName))
	if err != nil {
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("绑定信息已过期，请重新绑定")}
	}
	msgData, err := s.WechatBindAddress(string(eventKey), string(msg.FromUserName), true)
	if err != nil && msgData != "" {
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("钱包地址绑定失败：" + msgData)}
	}
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("钱包地址绑定成功")}
}

// GetWechatQrcode 获取关注二维码
func (s *Service) GetWechatQrcode(c *gin.Context, app, address string) (data string, err error) {
	// 项目配置
	wechatConfig := s.c.Project[app]
	// 校验key
	if c.GetHeader("x-api-key") != wechatConfig.APIKey {
		log.Errorv("非法请求", zap.String("x-api-key", c.GetHeader("x-api-key")))
		return "", errors.New("非法请求")
	}
	tq := basic.NewTmpQrRequest(time.Second*120, fmt.Sprintf("%s::bind::%s", app, address))
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:          s.c.Auth.Wechat.AppID,
		AppSecret:      s.c.Auth.Wechat.AppSecret,
		Token:          s.c.Auth.Wechat.Token,
		EncodingAESKey: s.c.Auth.Wechat.EncodingAESKey,
		Cache:          memory,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	ticket, err := officialAccount.GetBasic().GetQRTicket(tq)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	return basic.ShowQRCode(ticket), nil
}

// WechatBindAddress 处理地址绑定
func (s *Service) WechatBindAddress(eventKey, fromUserName string, replace bool) (msg string, err error) {
	// 判断是否为绑定事件
	if !strings.Contains(strings.Split(eventKey, "::")[1], "bind") {
		log.Errorv("非绑定事件", zap.String("eventKey", eventKey))
		return "", errors.New("非绑定事件")
	}
	// 清除qrscene_开头
	eventKey = strings.TrimPrefix(eventKey, "qrscene_")
	// 地址
	address := strings.Split(eventKey, "::")[2]
	// 项目配置
	project := strings.Split(eventKey, "::")[0]
	wechatConfig := s.c.Project[project]
	// 发送请求
	client := req.C().SetCommonHeader("x-api-key", wechatConfig.APIKey)
	type WechatBind struct {
		Address string `json:"address" form:"address" binding:"required"`
		Code    string `json:"code" form:"code" binding:"required"`
		Replace bool   `json:"replace" form:"replace"`
	}
	wechatBind := WechatBind{
		Address: address,
		Code:    fromUserName,
		Replace: replace,
	}
	r, err := client.R().SetBodyJsonMarshal(wechatBind).Post(wechatConfig.CallBackURL + "/v1/social/wechatBindAddress")
	if err != nil || r.StatusCode != 200 {
		return "绑定失败", errors.New("绑定失败")
	}
	//fmt.Println(r.String())
	// 绑定失败
	if gjson.Get(r.String(), "status").Int() != 0 {
		return gjson.Get(r.String(), "message").String(), errors.New("绑定失败")
	}
	// 绑定成功
	return gjson.Get(r.String(), "data").String(), nil
}
