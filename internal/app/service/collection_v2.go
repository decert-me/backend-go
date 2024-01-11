package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

// CollectionClaimV2 领取合辑奖励
func (s *Service) CollectionClaimV2(r request.CollectionClaimRequest, address string) error {
	// 查询合辑信息
	collection, err := s.dao.GetCollectionByTokenID(r.TokenID)
	if err != nil {
		return err
	}
	// 查询合辑包含的Quest列表
	questList, err := s.dao.GetQuestListByCollectionID(collection.ID)
	if len(questList) == 0 {
		return errors.New("UnexpectedError")
	}
	for _, quest := range questList {
		// 查询是否完成
		status, err := s.dao.HasClaimed(address, quest.TokenId)
		if err != nil {
			return err
		}
		if status == 0 {
			return errors.New("QuestNotComplete")
		}
	}
	// 校验是否已经空投
	if s.dao.HasAirdrop(address, collection.TokenId) {
		return errors.New("AlreadyAirdrop")
	}
	var app string
	// 判断地址
	if utils.IsValidAddress(address) {
		app = "decert_v2"
	} else {
		app = "decert_solana"
	}
	// 生成分享码
	paramsMap := map[string]interface{}{
		"app": app,
		"params": map[string]interface{}{
			"receiver":       address,
			"tokenId":        collection.TokenId,
			"challenge_type": "collection",
			"uri":            collection.Uri,
			"title":          collection.Title,
			"startTs":        gjson.Get(string(collection.ExtraData), "startTs").Int(),
			"endTs":          gjson.Get(string(collection.ExtraData), "endTs").Int(),
			"creator":        collection.Creator,
			"chain_id":       r.ChainID,
		},
	}
	// 将Map转换为JSON格式的字节数组
	paramsData, err := json.Marshal(paramsMap)
	if err != nil {
		log.Errorv("JSON encoding error:", zap.Error(err))
		return err
	}
	_, err = s.GenerateShare(request.GenerateShareRequest{Params: string(paramsData)})
	if err != nil {
		log.Errorv("GenerateShare error:", zap.Error(err))
		return err
	}
	// 保存记录
	if err = s.dao.CreateUserChallengeClaim(&model.UserChallengeClaim{
		Address: address,
		TokenId: collection.TokenId,
	}); err != nil {
		log.Errorv("CreateUserChallengeClaim error", zap.Error(err))
		return err
	}
	return nil
}
