package service

import (
	"backend-go/internal/app/model/request"
	"backend-go/pkg/log"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"strconv"
)

// GenerateShare 生成分享码
func (s *Service) GenerateShare(req request.GenerateShareRequest) (res string, err error) {
	// 1. 创建 MD5 散列算法实例
	hasher := md5.New()
	// 2. 将字符串转换为字节数组并计算 MD5 散列
	hasher.Write([]byte(req.Params))
	// 3. 获取生成的 MD5 散列值的字节数组
	hashBytes := hasher.Sum(nil)
	// 4. 将字节数组转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)
	// 5. 返回前 15 个字符
	// Callback
	return hashString[:15], s.ShareCallback(hashString[:15], req.Params)
}

// ShareCallback 分享回调
func (s *Service) ShareCallback(shareCode, params string) (err error) {
	client := req.C()
	body := map[string]string{
		"share_code": shareCode,
		"params":     params,
	}
	res, err := client.R().SetBodyJsonMarshal(body).Post(s.c.Share.ShareCallback)
	if err != nil {
		log.Errorv("req post error", zap.Error(err))
		return
	}
	if res.StatusCode != 200 {
		return errors.New("error")
	}
	if gjson.Get(res.String(), "status").Int() != 0 {
		return errors.New(gjson.Get(res.String(), "message").String())
	}
	return nil
}

// ClickShare 点击分享链接
func (s *Service) ClickShare(c *gin.Context, req request.ClickShareRequest) (err error) {
	// 获取请求IP
	clientIP := c.ClientIP()
	// 获取User-Agent信息
	userAgent := c.Request.UserAgent()
	return s.ClickCallback(req.ShareCode, clientIP, userAgent)
}

// ClickCallback 点击回调
func (s *Service) ClickCallback(shareCode, clientIP, userAgent string) (err error) {
	client := req.C()
	body := map[string]string{
		"share_code": shareCode,
		"ip":         clientIP,
		"user_agent": userAgent,
	}
	res, err := client.R().SetBodyJsonMarshal(body).Post(s.c.Share.ClickCallback)
	if err != nil {
		log.Errorv("req post error", zap.Error(err))
		return
	}
	if gjson.Get(res.String(), "status").Int() != 0 {
		return errors.New(gjson.Get(res.String(), "message").String())
	}
	return nil
}

// AirdropCallback 空投回调处理
func (s *Service) AirdropCallback(req request.AirdropCallbackRequest) (err error) {
	tokenId, err := strconv.Atoi(req.TokenId)
	if err != nil {
		return
	}
	if err = s.dao.UpdateAirdroppedOne(int64(tokenId), req.Receiver, req.Hash); err != nil {
		log.Errorv("updateAirdropStatus", zap.Any("error", err))
	}
	if err = s.dao.CreateChallengesOne(int64(tokenId), req.Receiver); err != nil {
		log.Errorv("updateAirdropStatus", zap.Any("error", err))
	}
	return err
}
