package service

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"strings"
	"time"
)

func GetDiscordInfo(address string) (res interface{}, err error) {
	var socials string
	err = global.DB.Model(&model.Users{}).Select("socials").
		Where("address = ?", address).
		First(&socials).Error
	return gjson.Get(socials, "discord").Value(), err
}

// GetLoginMessage
// @description: 获取登录签名消息
// @param: address string
// @return: err error, loginMessage string
func GetLoginMessage(address string) (err error, loginMessage string) {
	loginMessage = fmt.Sprintf(global.CONFIG.Contract.Signature+"Wallet address:\n%s\n\n", address)
	UUID := uuid.NewV4() // 生成UUID
	fmt.Println(UUID.String())
	// 存到Local Cache里
	if err = global.REDIS.Set(context.Background(), global.CONFIG.Redis.Prefix+UUID.String(), "", time.Hour*24).Err(); err != nil {
		return err, loginMessage
	}
	loginMessage = fmt.Sprintf(loginMessage+"Nonce:\n%s", UUID)
	return err, loginMessage
}

// AuthLoginSignRequest
// @description: 校验签名并返回Token
// @param: c *gin.Context, req request.AuthLoginSignRequest
// @return: token string, err error
func AuthLoginSignRequest(req request.AuthLoginSignRequest) (token string, err error) {
	if !utils.VerifySignature(req.Address, req.Signature, []byte(req.Message)) {
		return token, errors.New("签名校验失败")
	}
	// 获取Nonce
	index := strings.LastIndex(req.Message, "Nonce:")
	if index == -1 {
		return token, errors.New("nonce获取失败")
	}
	// 校验Nonce
	cacheErr := global.REDIS.Get(context.Background(), global.CONFIG.Redis.Prefix+req.Message[index+7:]).Err()
	if cacheErr != nil {
		return token, errors.New("签名已失效")
	}
	// 删除Nonce
	_ = global.REDIS.Del(context.Background(), global.CONFIG.Redis.Prefix+req.Message[index+7:])
	// 获取用户名--不存在则新增
	var user model.Users
	if errUser := global.DB.Model(&model.Users{}).Where("address = ?", req.Address).First(&user).Error; errUser != nil {
		if errUser == gorm.ErrRecordNotFound {
			user.Address = req.Address
			if err = global.DB.Model(&model.Users{}).Save(&user).Error; err != nil {
				return token, err
			}
		} else {
			return token, err
		}
	}
	// 验证成功返回JWT
	j := utils.NewJWT()
	claims := j.CreateClaims(utils.BaseClaims{
		UserID:  user.ID,
		Address: req.Address,
	})
	token, err = j.CreateToken(claims)
	if err != nil {
		return token, errors.New("获取token失败")
	}
	//  存入local cache
	//if err = global.TokenCache.Set(req.Address, []byte(token)); err != nil {
	//	return token, errors.New("保存token失败")
	//}
	return token, nil
}
