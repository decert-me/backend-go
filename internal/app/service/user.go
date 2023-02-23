package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"backend-go/pkg/auth"
	"backend-go/pkg/log"
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"strings"
)

func (s *Service) GetDiscordInfo(address string) (res interface{}, err error) {
	var socials string
	if socials, err = s.dao.GetSocialsInfo(&model.Users{Address: address}); err != nil {
		return
	}
	return gjson.Get(socials, "discord").Value(), err
}

// GetLoginMessage
// @description: 获取登录签名消息
// @param: address string
// @return: loginMessage string, err error
func (s *Service) GetLoginMessage(address string) (loginMessage string, err error) {
	loginMessage = fmt.Sprintf(s.c.BlockChain.Signature+"Wallet address:\n%s\n\n", address)
	UUID := uuid.NewV4() // 生成UUID
	// 存到Local Cache里
	if err = s.dao.SetNonce(context.Background(), UUID.String()); err != nil {
		log.Errorv("set nonce error: ", zap.Error(err))
		return loginMessage, err
	}
	return fmt.Sprintf(loginMessage+"Nonce:\n%s", UUID), nil
}

// AuthLoginSignRequest
// @description: 校验签名并返回Token
// @param: c *gin.Context, req request.AuthLoginSignRequest
// @return: token string, err error
func (s *Service) AuthLoginSignRequest(req request.AuthLoginSignRequest) (token string, err error) {
	midAuth := auth.New(s.c.Auth)
	if !utils.VerifySignature(req.Address, req.Signature, []byte(req.Message)) {
		return token, errors.New("签名校验失败")
	}
	// 获取Nonce
	index := strings.LastIndex(req.Message, "Nonce:")
	if index == -1 {
		return token, errors.New("nonce获取失败")
	}
	nonce := req.Message[index+7:]
	// 校验Nonce
	hasNonce, err := s.dao.HasNonce(context.Background(), nonce)
	if err != nil {
		log.Errorv("HasNonce error", zap.String("nonce", nonce))
	}
	if !hasNonce {
		return token, errors.New("签名已失效")
	}
	// 删除Nonce
	if err = s.dao.DelNonce(context.Background(), nonce); err != nil {
		log.Errorv("DelNonce error", zap.String("nonce", nonce)) // not important and continue
	}
	// 保存用户信息
	user := model.Users{Address: req.Address}
	if err = s.dao.SaveUser(&user); err != nil {
		log.Errorv("SaveUser error", zap.Any("user", user))
	}
	// 验证成功返回JWT
	claims := midAuth.CreateClaims(auth.BaseClaims{
		UserID:  user.ID,
		Address: req.Address,
	})
	token, err = midAuth.CreateToken(claims)
	if err != nil {
		log.Error("CreateToken error (%v)", err)
		return token, errors.New("获取token失败")
	}
	return token, nil
}
