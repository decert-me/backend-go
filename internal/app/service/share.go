package service

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/utils"
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
	body := response.ShareCallbackResponse{
		ShareCode: shareCode,
		Params:    params,
	}
	url := s.c.Share.Callback + "/v1/url/saveAirdrop"
	// 生成校验hash和时间戳
	timestamp, hashValue := utils.HashData(body, s.c.Share.VerifyKey)
	headers := map[string]string{
		"verify":    hashValue,
		"timestamp": strconv.Itoa(int(timestamp)),
	}
	res, err := client.R().SetHeaders(headers).SetBodyJsonMarshal(body).Post(url)
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
	body := response.ClickCallbackResponse{
		App:       "decert",
		ShareCode: shareCode,
		IP:        clientIP,
		UserAgent: userAgent,
	}
	url := s.c.Share.Callback + "/v1/url/saveAccess"
	// 生成校验hash和时间戳
	timestamp, hashValue := utils.HashData(body, s.c.Share.VerifyKey)
	headers := map[string]string{
		"verify":    hashValue,
		"timestamp": strconv.Itoa(int(timestamp)),
	}
	res, err := client.R().SetHeaders(headers).SetBodyJsonMarshal(body).Post(url)
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
func (s *Service) AirdropCallback(c *gin.Context, req request.AirdropCallbackRequest) (err error) {
	// 校验
	verify := c.Request.Header.Get("verify")
	timestamp := c.Request.Header.Get("timestamp")
	if !utils.VerifyData(req, s.c.Share.VerifyKey, verify, timestamp) {
		return errors.New("校验失败")
	}

	if req.Status == 2 {
		s.dao.AirdropFailNotice(req.Receiver, req.TokenId, req.Msg)
		return
	}
	tokenId, err := strconv.Atoi(req.TokenId)
	if err != nil {
		log.Errorv("strconv.Atoi error", zap.Error(err))
		return
	}
	if err = s.dao.UpdateAirdroppedOne(int64(tokenId), req.Receiver, req.Hash); err != nil {
		log.Errorv("UpdateAirdroppedOne error", zap.Any("error", err))
	}
	if err = s.dao.CreateChallengesOne(int64(tokenId), req.Receiver); err != nil {
		log.Errorv("CreateChallengesOne error ", zap.Any("error", err))
	}
	s.dao.AirdropSuccessNotice(req.Receiver, req.TokenId)
	return err
}
