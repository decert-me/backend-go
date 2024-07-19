package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/liangjies/go-solidity-sha3"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"math/big"
)

func (s *Service) UpdateBadgeURIV2(address string, badgeURI request.UpdateBadgeURIRequest) (res string, err error) {
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.SignPrivateKey)
	if err != nil {
		return
	}
	tokenId, set := big.NewInt(0).SetString(badgeURI.TokenId, 10)
	if !set {
		return res, errors.New("TokenIDInvalid")
	}
	hash := solsha3.SoliditySHA3(
		// types
		[]string{"uint256", "string", "address", "address"},
		// values
		[]interface{}{
			tokenId, badgeURI.Uri, s.c.ContractV2[badgeURI.ChainID].Badge, address,
		},
	)
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	signature[64] += 27
	return hexutil.Encode(signature), err
}

func (s *Service) SubmitClaimShareV2(address string, req request.SubmitClaimShareV2Req, lang string) (res string, err error) {
	// 校验是否绑定社交账号
	data, err := s.dao.HasBindSocialAccount(address)
	var isBinding bool
	for _, v := range data {
		if v == true {
			isBinding = true
			break
		}
	}
	if !isBinding {
		return res, errors.New("NoBindingDetected")
	}
	// 校验是否已经空投
	if s.dao.HasAirdrop(address, req.TokenId) {
		return res, errors.New("AlreadyAirdrop")
	}
	// 校验分数正确性
	quest, err := s.dao.GetQuestByTokenIDWithLang(lang, req.TokenId)
	if err != nil {
		return res, errors.New("TokenIDInvalid")
	}
	// 校验题目
	if req.Uri != "" && req.Uri != quest.Uri {
		return res, errors.New("QuestUpdate")
	}
	_, _, _, pass, err := s.AnswerCheck(s.c.Quest.EncryptKey, req.Answer, address, req.Score, &quest.Quest, true)
	if err != nil {
		log.Errorv("AnswerCheck error", zap.Error(err))
		return res, errors.New("UnexpectedError")
	}
	if !pass {
		return res, errors.New("AnswerIncorrect")
	}
	var app string
	// 判断地址
	if utils.IsValidAddress(address) {
		app = "decert_v2"
	} else {
		app = "decert_solana"
	}
	// 解析用户答案
	answer, err := s.AnswerParse(s.c.Quest.EncryptKey, req.Answer, address, &quest.Quest)
	if err != nil {
		log.Errorv("AnswerParse error", zap.Error(err))
		return res, errors.New("UnexpectedError")
	}
	// 生成分享码
	paramsMap := map[string]interface{}{
		"app": app,
		"params": map[string]interface{}{
			"receiver":           address,
			"tokenId":            req.TokenId,
			"score":              req.Score,
			"uri":                quest.Uri,
			"answer":             answer,
			"title":              quest.Title,
			"startTs":            gjson.Get(string(quest.ExtraData), "startTs").Int(),
			"endTs":              gjson.Get(string(quest.ExtraData), "endTs").Int(),
			"creator":            quest.Creator,
			"chain_id":           req.ChainID,
			"image_uri":          req.ImageUri,
			"uuid":               quest.UUID,
			"challenge_ipfs_url": gjson.Get(string(quest.MetaData), "attributes.challenge_ipfs_url").String(),
			"description":        quest.Description,
		},
	}
	// 将Map转换为JSON格式的字节数组
	paramsData, err := json.Marshal(paramsMap)
	if err != nil {
		log.Errorv("JSON encoding error:", zap.Error(err))
		return
	}
	res, err = s.GenerateShare(request.GenerateShareRequest{Params: string(paramsData)})
	if err != nil {
		log.Errorv("GenerateShare error:", zap.Error(err))
		return
	}
	// 保存记录
	if err = s.dao.CreateUserChallengeClaim(&model.UserChallengeClaim{
		Address: address,
		TokenId: req.TokenId,
		ChainID: req.ChainID,
	}); err != nil {
		log.Errorv("CreateUserChallengeClaim error", zap.Error(err))
		return
	}
	return res, err
}
