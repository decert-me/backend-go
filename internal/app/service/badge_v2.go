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

// GenerateMintSignature - 生成用户自主 mint NFT 的签名
// 用户通过前端调用 BadgeMinter.claim() 时使用此签名
func (s *Service) GenerateMintSignature(address string, req request.GenerateMintSignatureReq, lang string) (res map[string]interface{}, err error) {
	log.Infov("GenerateMintSignature called",
		zap.String("address", address),
		zap.String("tokenId", req.TokenId),
		zap.Int64("score", req.Score),
		zap.Int64("chainId", req.ChainID))

	// 1. 校验是否绑定社交账号
	data, err := s.dao.HasBindSocialAccount(address)
	var isBinding bool
	for _, v := range data {
		if v == true {
			isBinding = true
			break
		}
	}
	if !isBinding {
		log.Errorv("Social account not bound", zap.String("address", address))
		return res, errors.New("NoBindingDetected")
	}

	// 2. 校验是否已经领取
	if s.dao.HasAirdrop(address, req.TokenId) {
		return res, errors.New("AlreadyAirdrop")
	}

	// 3. 校验分数正确性
	quest, err := s.dao.GetQuestByTokenIDWithLang(lang, req.TokenId)
	if err != nil {
		return res, errors.New("TokenIDInvalid")
	}

	// 4. 校验题目是否更新 (与原有逻辑保持一致)
	if req.Uri != "" && req.Uri != quest.Uri {
		log.Warnv("Quest URI mismatch - quest may have been updated",
			zap.String("address", address),
			zap.String("tokenId", req.TokenId),
			zap.String("receivedUri", req.Uri),
			zap.String("expectedUri", quest.Uri))
		return res, errors.New("QuestUpdate")
	}

	// 5. 校验答案
	_, _, _, pass, err := s.AnswerCheck(s.c.Quest.EncryptKey, req.Answer, address, req.Score, &quest.Quest, true)
	if err != nil {
		log.Errorv("AnswerCheck error", zap.Error(err))
		return res, errors.New("UnexpectedError")
	}
	if !pass {
		return res, errors.New("AnswerIncorrect")
	}

	// 6. 获取私钥
	privateKey, err := crypto.HexToECDSA(s.c.BlockChain.SignPrivateKey)
	if err != nil {
		log.Errorv("HexToECDSA error", zap.Error(err))
		return res, errors.New("UnexpectedError")
	}

	// 7. 获取 BadgeMinter 合约地址
	badgeMinterAddress := ""
	switch req.ChainID {
	case 11155420: // OP Sepolia
		badgeMinterAddress = "0xEdC46868f04d482f04A8c29E915aBED72C03cD35"
	case 10: // Optimism
		badgeMinterAddress = "0x0aa319263401eEcecd5Fa2C34636b1057A8B2BFB"
	case 137: // Polygon
		badgeMinterAddress = "0x0aa319263401eEcecd5Fa2C34636b1057A8B2BFB"
	case 42161: // Arbitrum
		badgeMinterAddress = "0x0aa319263401eEcecd5Fa2C34636b1057A8B2BFB"
	default:
		return res, errors.New("UnsupportedChain")
	}

	// 8. 构造签名数据
	// 签名格式: keccak256(chainId, to, questId, uri, contractAddress, callerAddress)
	tokenId, set := big.NewInt(0).SetString(req.TokenId, 10)
	if !set {
		return res, errors.New("TokenIDInvalid")
	}

	hash := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "uint256", "string", "address", "address"},
		[]interface{}{
			big.NewInt(req.ChainID), // chainId
			address,                 // to (接收者)
			tokenId,                 // questId
			req.ImageUri,            // uri (NFT的图片/元数据URI)
			badgeMinterAddress,      // BadgeMinter 合约地址
			address,                 // caller (调用者，用户自己)
		},
	)

	// 9. 添加以太坊前缀并签名
	prefixedHash := solsha3.SoliditySHA3WithPrefix(hash)
	signature, err := crypto.Sign(prefixedHash, privateKey)
	if err != nil {
		log.Errorv("Sign error", zap.Error(err))
		return res, errors.New("UnexpectedError")
	}
	signature[64] += 27

	log.Infov("Signature generated successfully",
		zap.String("address", address),
		zap.String("tokenId", req.TokenId),
		zap.String("badgeMinterAddress", badgeMinterAddress),
		zap.String("signature", hexutil.Encode(signature)))

	// 10. 保存用户 claim 记录（标记为待 claim 状态）
	// 当用户真正在链上 claim 成功后，前端会调用 ConfirmUserClaim API 更新状态
	if err = s.dao.CreateUserChallengeClaim(&model.UserChallengeClaim{
		Address: address,
		TokenId: req.TokenId,
		ChainID: req.ChainID,
	}); err != nil {
		log.Errorv("CreateUserChallengeClaim error", zap.Error(err))
		// 不影响签名返回，继续执行
	}

	// 11. 返回签名数据
	res = map[string]interface{}{
		"to":        address,
		"questId":   req.TokenId,
		"uri":       req.ImageUri, // 返回 NFT 的图片/元数据 URI
		"signature": hexutil.Encode(signature),
		"chainId":   req.ChainID,
	}

	return res, nil
}

// ConfirmUserMint - 确认用户自主 mint 成功
// 前端在用户成功调用 BadgeMinter.claim() 后调用此接口，更新数据库状态
func (s *Service) ConfirmUserMint(address string, tokenID string, txHash string) error {
	log.Infov("ConfirmUserMint called",
		zap.String("address", address),
		zap.String("tokenId", tokenID),
		zap.String("txHash", txHash))

	// 更新 UserChallengeClaim 状态为成功 (status = 2)
	err := s.dao.UpdateUserChallengeClaimStatus(address, tokenID, 2)
	if err != nil {
		log.Errorv("UpdateUserChallengeClaimStatus error", zap.Error(err))
		return errors.New("UpdateStatusFailed")
	}

	log.Infov("User mint confirmed successfully",
		zap.String("address", address),
		zap.String("tokenId", tokenID))

	return nil
}
