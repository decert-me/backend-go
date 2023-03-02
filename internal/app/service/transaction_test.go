package service

import (
	"backend-go/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_HashSubmit(t *testing.T) {
	hash := "0x60b66b2e0627aaadb42981d7edeacd7150cc7632801a11aba1e01e895105fcfa"
	adress := "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
	// delete exist
	err := s.dao.DB().Where("hash", hash).Delete(&model.Transaction{}).Error
	assert.Nil(t, err)
	// start testing
	err = s.HashSubmit(adress, hash)
	assert.Nil(t, err)
	// repeat should error
	err = s.HashSubmit(adress, hash)
	assert.Error(t, err)
}
