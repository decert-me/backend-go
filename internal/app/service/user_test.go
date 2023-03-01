package service

import (
	"backend-go/internal/app/model"
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
	//token, err := s.AuthLoginSignRequest(
	//	request.AuthLoginSignRequest{
	//		Address:   "0x7d32D1DE76acd73d58fc76542212e86ea63817d8",
	//		Message:   "",
	//		Signature: "",
	//	},
	//)
	//assert.Nil(t, err)
}
