package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_HashSubmit(t *testing.T) {
	// delete exist
	err := s.dao.DB().Where("hash", QuestCreatedHash).Delete(&model.Transaction{}).Error
	assert.Nil(t, err)
	// start testing
	err = s.HashSubmit(ADDRESS, request.HashSubmitRequest{Hash: QuestCreatedHash})
	assert.Nil(t, err)
	// repeat should error
	err = s.HashSubmit(ADDRESS, request.HashSubmitRequest{Hash: QuestCreatedHash})
	assert.Error(t, err)
}
