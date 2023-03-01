package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/pkg/auth"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"strings"
	"testing"
)

func TestService_GetDiscordInfo(t *testing.T) {
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	socials := "{\"discord\": {\"id\": \"1045278958163349535\", \"username\": \"liangjies\"}}"
	// mock discord robot
	err := s.dao.DB().Where("address", address).Delete(&model.Users{}).Error
	assert.Nil(t, err)
	// no data
	_, err = s.GetDiscordInfo(address)
	assert.Error(t, err)
	// normal
	err = s.dao.DB().Save(&model.Users{
		Address: address,
		Socials: []byte(gjson.Parse(socials).Raw),
	}).Error
	assert.Nil(t, err)
	// start testing
	res, err := s.GetDiscordInfo(address)
	assert.Nil(t, err)
	assert.Equal(t, gjson.Get(socials, "discord").Value(), res)
}

func TestService_GetLoginMessage(t *testing.T) {
	address := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	loginMessage, err := s.GetLoginMessage(address)
	assert.Nil(t, err)
	index := strings.LastIndex(loginMessage, "Nonce:")
	assert.NotEqual(t, -1, index, "nonce should exist")
	// login message should equal
	loginMessageExpect := fmt.Sprintf(s.c.BlockChain.Signature+"Wallet address:\n%s\n\n", address)
	assert.Equal(t, loginMessageExpect, loginMessage[:index])
	// nonce should be in redis
	nonce := loginMessage[index+7:]
	hasNonce, err := s.dao.HasNonce(context.Background(), nonce)
	assert.Nil(t, err)
	assert.True(t, hasNonce)
	// set nonce error

}

func TestService_AuthLoginSignRequest(t *testing.T) {
	// 签名已失效
	_, err := s.AuthLoginSignRequest(
		request.AuthLoginSignRequest{
			Address:   "0x7d32D1DE76acd73d58fc76542212e86ea63817d8",
			Message:   "Welcome to Decert!\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nYour authentication status will reset after 24 hours.\n\nWallet address:\n0x7d32D1DE76acd73d58fc76542212e86ea63817d8\n\nNonce:\nfee31011-c84d-4e06-8ff2-4e8c4dc29b31",
			Signature: "0x32b510e0bbb0a6e52d500631d550f47802001ab958f2e5893fed591cae59e92330f8de89e999f75ab0607ce1f100de78f1dcd2030714624adc6ccf5c870928c21c",
		},
	)
	assert.Equal(t, "签名已失效", err.Error())
	// 签名地址错误
	_, err = s.AuthLoginSignRequest(
		request.AuthLoginSignRequest{
			Address:   "0x7d32D1DE76acd73d58fc76542212e86ea63817d8",
			Message:   "Welcome to Decert!\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nYour authentication status will reset after 24 hours.\n\nWallet address:\n0x7d32D1DE76acd73d58fc76542212e86ea6381000\n\nNonce:\ncf152367-3739-4545-aa7b-ff8d22140364",
			Signature: "0x521f0d827ea82621971b74d40ca897565b1d80f3242ea1153544167716d8da6f4a8639aa6c8d889081e8d430471ae484c863cc9725399548637e7138b3b69b571b",
		},
	)
	assert.Equal(t, "签名地址错误", err.Error())
	// nonce获取失败
	_, err = s.AuthLoginSignRequest(
		request.AuthLoginSignRequest{
			Address:   "0x7d32D1DE76acd73d58fc76542212e86ea63817d8",
			Message:   "Welcome to Decert!\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nYour authentication status will reset after 24 hours.\n\nWallet address:\n0x7d32D1DE76acd73d58fc76542212e86ea63817d8\n\n",
			Signature: "0x400e034b9a53e5653d5a9e565a7d915a5793689cadac2942b6fbfe22afb881b1229b28e6ce4693b71335f697c67933a7d506de937f42b050206ed6b684ec17021c",
		},
	)
	assert.Equal(t, "nonce获取失败", err.Error())
	// address获取失败
	_, err = s.AuthLoginSignRequest(
		request.AuthLoginSignRequest{
			Address:   "0x7d32D1DE76acd73d58fc76542212e86ea63817d8",
			Message:   "Welcome to Decert!\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nYour authentication status will reset after 24 hours.\n\nNonce:\n29b05af3-aae0-4bed-9b58-77693f81044d",
			Signature: "0xb509ed2e9bf35579431f6da48051e5fe968a2af3202483754d91198d8cf7693c2b2edc5e359e37a83ca6969abbcb236f5f89483875cd1e27dddb28a2601af7cd1c",
		},
	)
	assert.Equal(t, "address获取失败", err.Error())
	// 签名信息错误
	err = s.dao.SetNonce(context.Background(), "0d2c1cb5-3cad-431a-a266-6e3262cb5fb7")
	assert.Nil(t, err)

	_, err = s.AuthLoginSignRequest(
		request.AuthLoginSignRequest{
			Address:   "0x7d32D1DE76acd73d58fc76542212e86ea63817d8",
			Message:   "Wallet address:\n0x7d32D1DE76acd73d58fc76542212e86ea63817d8\n\nNonce:\n0d2c1cb5-3cad-431a-a266-6e3262cb5fb7",
			Signature: "0xa37af8f9d6889ab2bdfd2ed0769d93514045d3d8a2e2b046167032567e67babe497512fe14c369f44bc1d39ed3d26b86c58d63dec7dd4a5c5ffa665a63ba70491c",
		},
	)
	assert.Equal(t, "签名信息错误", err.Error())

	// 登陆成功
	err = s.dao.SetNonce(context.Background(), "fee31011-c84d-4e06-8ff2-4e8c4dc29b31")
	assert.Nil(t, err)
	token, err := s.AuthLoginSignRequest(
		request.AuthLoginSignRequest{
			Address:   "0x7d32D1DE76acd73d58fc76542212e86ea63817d8",
			Message:   "Welcome to Decert!\n\nThis request will not trigger a blockchain transaction or cost any gas fees.\n\nYour authentication status will reset after 24 hours.\n\nWallet address:\n0x7d32D1DE76acd73d58fc76542212e86ea63817d8\n\nNonce:\nfee31011-c84d-4e06-8ff2-4e8c4dc29b31",
			Signature: "0x32b510e0bbb0a6e52d500631d550f47802001ab958f2e5893fed591cae59e92330f8de89e999f75ab0607ce1f100de78f1dcd2030714624adc6ccf5c870928c21c",
		},
	)
	assert.Nil(t, err)
	var midAuth *auth.Auth
	midAuth = auth.New(c.Auth)

	claims, err := midAuth.ParseToken(token)
	assert.Equal(t, "0x7d32D1DE76acd73d58fc76542212e86ea63817d8", claims.Address)
	assert.Equal(t, "Decert", claims.Issuer)
}
