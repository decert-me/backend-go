package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"encoding/json"
	"errors"
	"fmt"
	reqV3 "github.com/imroc/req/v3"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"strings"
	"time"
)

// SaveSignAndDid 保存签名和DID账号
func (s *Service) SaveSignAndDid(address string, req request.SaveSignAndDidRequest) (err error) {
	if !utils.VerifySignature(address, req.SignHash, []byte(req.Sign)) {
		return errors.New("SignatureVerificationFailed")
	}
	// 校验Address
	if req.Sign[0:42] != address {
		return errors.New("AddressError")
	}
	// 校验did账号
	indexAddress := strings.LastIndex(req.Sign, "did:zk:")
	if indexAddress == -1 {
		return errors.New("AddressError")
	}
	did := req.Sign[indexAddress : indexAddress+49]
	if strings.TrimSpace(did) != req.DidAddress {
		return errors.New("AddressError")
	}
	err = s.dao.SaveSignAndDid(address, req)
	if err != nil {
		return err
	}
	return
}

// GenerateCardInfoTask TODO: task
func (s *Service) GenerateCardInfoTask(address, didAddress string) (err error) {
	// 是否生成过证书，跳过已经生成过
	hasCard, err := s.dao.AddressHasCard(address)
	if err != nil {
		return errors.New("UnexpectedError")
	}
	if hasCard {
		return nil
	}
	// 查询所有历史挑战最高分
	res, err := s.dao.GetAddressHighScore(address)
	if err != nil {
		log.Errorv("查询所有历史挑战最高分失败", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	for _, v := range res {
		// 生成证书
		err = s.GenerateCardInfo(address, v.Score, request.GenerateCardInfoRequest{
			TokenId: v.TokenId,
			Answer:  string(v.Answer),
		})
		if err != nil {
			log.Errorv("生成证书失败", zap.Error(err))
			return errors.New("UnexpectedError")
		}
	}
	return
}

// GetAddressDid 查询地址绑定的Did
func (s *Service) GetAddressDid(address string) (did string, err error) {
	return s.dao.GetAddressDid(address)
}

// GetVcInfo 获取 VC 信息
func (s *Service) GetVcInfo(address, tokenID string) (vc interface{}, err error) {
	return s.dao.GetVcInfo(address, tokenID)
}

type DidCardRequest struct {
	Receiver string               `json:"receiver"`
	Params   DidCardParamsRequest `json:"params"`
}

type DidCardParamsRequest struct {
	Title       string `json:"Title"`
	ChallengeID string `json:"ChallengeID"`
	Pass        bool   `json:"Pass"`
	Score       int64  `json:"Score"`
	Content     string `json:"Content"`
}

// GenerateCardInfo 生成 card 信息
func (s *Service) GenerateCardInfo(address string, score int64, req request.GenerateCardInfoRequest) (err error) {
	// 获取did 账号
	did, err := s.dao.GetAddressDid(address)
	if err != nil {
		return errors.New("DIDNotFound")
	}
	if did == "" {
		return errors.New("DIDNotFound")
	}
	// 校验分数正确性
	quest, err := s.dao.GetQuestByTokenID(req.TokenId)
	if err != nil {
		return errors.New("TokenIDInvalid")
	}
	// 校验题目
	if req.Uri != "" && req.Uri != quest.Uri {
		return errors.New("QuestUpdate")
	}
	pass := true
	if score == 0 {
		score, pass, err = s.AnswerCheck(s.c.Quest.EncryptKey, req.Answer, address, 0, &quest, true)
		if err != nil {
			log.Errorv("AnswerCheck error", zap.Error(err))
			return errors.New("UnexpectedError")
		}
	}
	// 未通过跳过
	if !pass {
		return nil
	}
	// 将用户答案写入metadata
	metadata, err := sjson.Set(string(quest.MetaData), "attributes.user_answer", gjson.Parse(req.Answer).Value())
	if err != nil {
		log.Errorv("sjson set error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	// 将metadata上传到IPFS
	err, hash := s.IPFSUploadJSON(gjson.Parse(metadata).Value())
	if err != nil || hash == "" {
		log.Errorv("IPFSUploadJSON error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	// 构造请求
	data := DidCardRequest{
		Receiver: did,
		Params: DidCardParamsRequest{
			Title:       quest.Title,
			ChallengeID: quest.UUID,
			Pass:        pass,
			Score:       score,
			Content:     "ipfs://" + hash,
		},
	}
	// 发送请求获取vc
	res, err := reqV3.C().R().SetBodyJsonMarshal(data).Post(s.c.ZCloak.Url + "/vc/issue")
	if err != nil {
		log.Errorv("get VC error", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	if gjson.Get(res.String(), "code").Int() != 0 {
		log.Errorv("get VC error", zap.Error(err), zap.String("res", res.String()))
		return errors.New("UnexpectedError")
	}
	// 保存 VC
	zcloakCard := model.ZcloakCard{
		Address: address,
		Did:     did,
		QuestID: quest.ID,
		Score:   score,
		VC:      []byte(gjson.Get(res.String(), "data.vc").String()),
	}
	err = s.dao.SaveZcloakCard(zcloakCard)
	if err != nil {
		return err
	}
	err = s.SaveToNFTCollection(SaveCardInfoRequest{
		Chain:           "polygon",
		AccountAddress:  strings.ToLower(address),
		ContractAddress: strings.ToLower("0xc8e9cd4921e54c4163870092ca8d9660e967b53d"),
		TokenID:         cast.ToString(req.TokenId),
		ImageURI:        strings.TrimPrefix(gjson.Get(string(quest.MetaData), "image").String(), "ipfs://"),
		ErcType:         "erc1155",
		Name:            gjson.Get(string(quest.MetaData), "name").String(),
		DidAddress:      did,
		MetadataJson:    string(quest.MetaData),
	})
	if err != nil {
		return err
	}
	return
}

// GetDidSignMessage 获取DID签名信息
func (s *Service) GetDidSignMessage(did, ethAddress string) (message string, err error) {
	message = fmt.Sprintf("%s\n\n%s\n\n%d", ethAddress, did, time.Now().UnixNano()/1000000)
	return message, nil
}

type SaveCardInfoRequest struct {
	Chain           string `json:"chain" form:"chain" binding:"required"`
	AccountAddress  string `json:"account_address" form:"account_address" binding:"required"`
	ContractAddress string `json:"contract_address" form:"contract_address" binding:"required"`
	TokenID         string `json:"token_id" form:"token_id" binding:"required"`
	ImageURI        string `json:"image_uri" form:"image_uri" binding:"required"`
	ErcType         string `json:"erc_type" form:"erc_type" binding:"required"`
	Name            string `json:"name" form:"name" binding:"required"`
	DidAddress      string `json:"did_address" form:"did_address" binding:"required"`
	MetadataJson    string `json:"metadata_json" form:"metadata_json"`
}

// SaveToNFTCollection 保存到NFT
func (s *Service) SaveToNFTCollection(saveCardInfo SaveCardInfoRequest) (err error) {
	if s.c.NFT.API == "" {
		return
	}
	// 发送请求
	client := reqV3.C().SetCommonHeader("x-api-key", s.c.NFT.APIKey)
	fmt.Println(s.c.NFT.API + "/zcloak/saveCardInfo")
	data, _ := json.Marshal(saveCardInfo)
	fmt.Println(string(data))
	r, err := client.R().SetBodyJsonMarshal(saveCardInfo).Post(s.c.NFT.API + "/zcloak/saveCardInfo")
	if err != nil {
		log.Errorv("SaveToNFT error", zap.Error(err), zap.String("res", r.String()))
		return err
	}
	if r.StatusCode != 200 {
		log.Errorv("SaveToNFT error", zap.Error(err), zap.String("res", r.String()))
		return errors.New("UnexpectedError")
	}
	if gjson.Get(r.String(), "status").Int() != 0 {
		log.Errorv("SaveToNFT error", zap.Error(err), zap.String("res", r.String()))
		return errors.New("UnexpectedError")
	}
	return nil
}

// GetKeyFileWithSignature 获取KeyFiles
func (s *Service) GetKeyFileWithSignature(address string) (keyFile datatypes.JSON, err error) {
	return s.dao.GetKeyFileWithSignature(address)
}

// GenerateCard 生成证书
func (s *Service) GenerateCard(address string, tokenID string) (err error) {
	quest, err := s.dao.GetQuestByTokenID(tokenID)
	if err != nil {
		return errors.New("TokenIDInvalid")
	}
	var userScore int64
	var answer string
	// 判断是否是开放题
	if IsOpenQuest(gjson.Get(string(quest.QuestData), "questions").String()) {
		// 获取用户最新回答
		answer, userScore, err = s.dao.GetLatestOpenQuestPassAnswer(address, tokenID)
		if err != nil {
			log.Errorv("GetLatestOpenQuestPassAnswer error", zap.Error(err))
			return errors.New("UnexpectedError")
		}
	} else {
		answer, userScore, err = s.dao.GetLatestQuestPassAnswer(address, tokenID)
		if err != nil {
			log.Errorv("GetLatestQuestPassAnswer error", zap.Error(err))
			return errors.New("UnexpectedError")
		}
		userScore = userScore / 100
	}
	fmt.Println("answer", answer)
	s.GenerateCardInfo(address, userScore, request.GenerateCardInfoRequest{
		TokenId: tokenID,
		Answer:  answer,
	})
	return
}

// GetDidCardInfo 获取 DID Card 信息
func (s *Service) GetDidCardInfo(req request.GetDidCardInfoRequest) (didCardInfo datatypes.JSON, err error) {
	return s.dao.GetDidCardInfo(req.Address, req.TokenId)
}
