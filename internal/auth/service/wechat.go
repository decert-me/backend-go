package service

import (
	"backend-go/pkg/log"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount/basic"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
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
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		if msg.MsgType == message.MsgTypeEvent {
			if msg.Event == message.EventSubscribe || msg.Event == message.EventScan {
				if msgData, err := s.WechatBindAddress(msg.EventKey, string(msg.FromUserName)); err != nil {
					if msgData != "" {
						return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("钱包地址绑定失败：" + msgData)}
					} else {
						return nil
					}
				}
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText("钱包地址绑定成功")}
			}
		}
		return nil
	})
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

// GetWechatQrcode 获取关注二维码
func (s *Service) GetWechatQrcode(address string) (data string, err error) {
	tq := basic.NewTmpQrRequest(time.Second*120, fmt.Sprintf("bind::%s", address))
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
func (s *Service) WechatBindAddress(eventKey, fromUserName string) (msg string, err error) {
	// 判断是否为绑定事件
	if !strings.Contains(strings.Split(eventKey, "::")[0], "bind") {
		log.Errorv("非绑定事件", zap.String("eventKey", eventKey))
		return "", errors.New("非绑定事件")
	}
	// 地址
	address := strings.Split(eventKey, "::")[1]
	// 判断是否已经绑定过
	wechatData, err := s.dao.WechatQueryByAddress(address)
	if err != nil {
		return "服务器内部错误", err
	}
	if wechatData != "{}" {
		return "钱包地址已绑定，请勿重复操作", errors.New("钱包地址已绑定，请勿重复操作")
	}
	// 判断微信是否被别的地址绑定过
	isBinding, err := s.dao.WechatIsBinding(fromUserName)
	if err != nil {
		return "服务器内部错误", err
	}
	if isBinding {
		return "微信账号已绑定其他钱包地址", errors.New("微信已经绑定过地址")
	}
	// 绑定
	return s.dao.WechatBindAddress(address, fromUserName)
}
