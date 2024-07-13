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
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func (s *Service) GetUserInfo(address string, loginAddress string) (res interface{}, err error) {
	var user model.Users
	if user, err = s.dao.GetUser(address); err != nil {
		return
	}
	// 是否管理员
	isAdmin, err := s.dao.IsAdmin(loginAddress)
	if err != nil {
		return
	}
	// 管理员显示
	if isAdmin {
		var showStr string
		showStr = fmt.Sprintf("%s...%s", address[:6], address[len(address)-4:])
		if user.NickName != nil && *user.NickName != "" {
			showStr = *user.NickName
		}
		if user.Name != nil && *user.Name != "" {
			showStr += "-" + *user.Name
		}
		// 显示标签
		tags, err := s.dao.GetUserTags(user.ID)
		if err == nil {
			for i := 0; i < len(tags); i++ {
				showStr += "，" + tags[i]
			}
		}
		user.NickName = &showStr
		return user, err
	}
	// default nickname
	if user.NickName == nil || *user.NickName == "" {
		if len(user.Address) > 10 {
			nickName := fmt.Sprintf("%s...%s", address[:6], address[len(address)-4:])
			user.NickName = &nickName
		}
	}
	return user, err
}

func (s *Service) UpdateUserInfo(address string, user request.UpdateUserInfo) (err error) {
	err = s.dao.UpdateUser(address, model.Users{Avatar: user.Avatar, Description: user.Description, NickName: user.NickName})
	if err != nil {
		log.Errorv("UpdateUser error", zap.Error(err))
	}
	return
}

func (s *Service) GetDiscordInfo(address string) (res interface{}, err error) {
	var socials string
	if socials, err = s.dao.GetSocialsInfo(&model.Users{Address: address}); err != nil {
		return
	}
	discord := gjson.Get(socials, "discord").Value()
	if discord == "" {
		return nil, errors.New("NoBindingDetected")
	}
	return discord, err
}

func (s *Service) GetTwitterInfo(address string) (res interface{}, err error) {
	var socials string
	if socials, err = s.dao.GetSocialsInfo(&model.Users{Address: address}); err != nil {
		return
	}
	id := gjson.Get(socials, "twitter.id").String()
	username := gjson.Get(socials, "twitter.username").String()
	if id == "" {
		return nil, errors.New("NoBindingDetected")
	}
	res = map[string]string{"id": id,
		"username": username,
	}
	return res, err
}

func (s *Service) UpdateAvatar(address string, header *multipart.FileHeader) (p string, err error) {
	ext := strings.ToLower(path.Ext(header.Filename))
	// 文件名
	key := uuid.NewV4().String()
	filename := key + ext
	director := s.c.Local.Path + "/"
	p = director + filename
	// 创建路径
	err = os.MkdirAll(director, os.ModePerm)
	if err != nil {
		log.Errorv("MkdirAll error", zap.Error(err))
		return
	}
	out, err := os.Create(p)
	if err != nil {
		log.Errorv("os.Create error", zap.Error(err))
		return
	}
	f, err := header.Open() // 读取文件
	if err != nil {
		log.Errorv("header.Open error", zap.Error(err))
		return
	}
	defer f.Close()
	_, err = io.Copy(out, f)
	if err != nil {
		log.Errorv("io.Copy error", zap.Error(err))
		return
	}
	//if err = s.dao.UpdateAvatar(address, filename); err != nil {
	//	return err
	//}
	if err = s.dao.UploadFile(model.Upload{Address: address, Name: filename, Key: key}); err != nil {
		return
	}
	return "/" + p, err
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
		return token, errors.New("SignatureVerificationFailed")
	}
	// 获取Nonce
	indexNonce := strings.LastIndex(req.Message, "Nonce:")
	if indexNonce == -1 {
		return token, errors.New("SignatureExpired")
	}
	nonce := req.Message[indexNonce+7:]
	// 获取Address
	indexAddress := strings.LastIndex(req.Message, "Wallet address:")
	if indexAddress == -1 {
		return token, errors.New("AddressError")
	}
	address := req.Message[indexAddress+16 : indexNonce]
	// 校验address
	if strings.TrimSpace(address) != req.Address {
		return token, errors.New("AddressError")
	}
	// 校验Nonce
	hasNonce, err := s.dao.HasNonce(context.Background(), nonce)
	if err != nil {
		log.Errorv("HasNonce error", zap.String("nonce", nonce))
		return token, errors.New("SignatureExpired")
	}
	if !hasNonce {
		return token, errors.New("SignatureExpired")
	}
	// 删除Nonce
	if err = s.dao.DelNonce(context.Background(), nonce); err != nil {
		log.Errorv("DelNonce error", zap.String("nonce", nonce)) // not important and continue
	}
	// 校验签名信息
	if req.Message[:indexAddress] != s.c.BlockChain.Signature {
		return token, errors.New("SignatureVerificationFailed")
	}
	// 保存用户信息
	user, err := s.createUser(model.Users{Address: req.Address, ParticleUserinfo: req.ParticleUserinfo})
	if err != nil {
		log.Errorv("createUser error", zap.Any("address", req.Address), zap.Error(err))
		return token, errors.New("UnexpectedError")
	}
	// 更新用户社交账户
	if len(req.ParticleUserinfo) != 0 {
		if err = s.dao.ParticleUpdateSocialsInfo(user.Address, req.ParticleUserinfo); err != nil {
			log.Errorv("UpdateSocialsInfo error", zap.Error(err))
		}
	}
	// 验证成功返回JWT
	claims := midAuth.CreateClaims(auth.BaseClaims{
		UserID:  user.ID,
		Address: req.Address,
	})
	token, err = midAuth.CreateToken(claims)
	if err != nil {
		log.Error("CreateToken error (%+v)", err)
		return token, errors.New("UnexpectedError")
	}
	return token, nil
}

// createUser 创建用户
func (s *Service) createUser(userInfo model.Users) (user model.Users, err error) {
	user, err = s.dao.GetUser(userInfo.Address)
	if err == nil {
		if err = s.dao.UpdateUser(userInfo.Address, userInfo); err != nil {
			return
		}
		return
	}
	// create user
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user = userInfo
		if err = s.dao.CreateUser(&user); err != nil {
			return
		}
	}
	return
}

// AuthLoginSignRequestSolana
// @description: 校验签名并返回Token
// @param: c *gin.Context, req request.AuthLoginSignRequest
// @return: token string, err error
func (s *Service) AuthLoginSignRequestSolana(req request.AuthLoginSignRequest) (token string, err error) {
	midAuth := auth.New(s.c.Auth)
	if !utils.VerifySignatureSolana(req.Address, req.Signature, []byte(req.Message)) {
		return token, errors.New("SignatureVerificationFailed")
	}
	// 获取Nonce
	indexNonce := strings.LastIndex(req.Message, "Nonce:")
	if indexNonce == -1 {
		return token, errors.New("SignatureExpired")
	}
	nonce := req.Message[indexNonce+7:]
	// 获取Address
	indexAddress := strings.LastIndex(req.Message, "Wallet address:")
	if indexAddress == -1 {
		return token, errors.New("AddressError")
	}
	address := req.Message[indexAddress+16 : indexNonce]
	// 校验address
	if strings.TrimSpace(address) != req.Address {
		return token, errors.New("AddressError")
	}
	// 校验Nonce
	hasNonce, err := s.dao.HasNonce(context.Background(), nonce)
	if err != nil {
		log.Errorv("HasNonce error", zap.String("nonce", nonce))
		return token, errors.New("SignatureExpired")
	}
	if !hasNonce {
		return token, errors.New("SignatureExpired")
	}
	// 删除Nonce
	if err = s.dao.DelNonce(context.Background(), nonce); err != nil {
		log.Errorv("DelNonce error", zap.String("nonce", nonce)) // not important and continue
	}
	// 校验签名信息
	if req.Message[:indexAddress] != s.c.BlockChain.Signature {
		return token, errors.New("SignatureVerificationFailed")
	}
	// 保存用户信息
	user, err := s.createUser(model.Users{Address: req.Address, ParticleUserinfo: req.ParticleUserinfo})
	if err != nil {
		log.Errorv("createUser error", zap.Any("address", req.Address), zap.Error(err))
		return token, errors.New("UnexpectedError")
	}
	// 验证成功返回JWT
	claims := midAuth.CreateClaims(auth.BaseClaims{
		UserID:  user.ID,
		Address: req.Address,
	})
	token, err = midAuth.CreateToken(claims)
	if err != nil {
		log.Error("CreateToken error (%+v)", err)
		return token, errors.New("UnexpectedError")
	}
	return token, nil
}

// HasOpenQuestPerm 获取用户是否有创建开放题权限
func (s *Service) HasOpenQuestPerm(address string) (perm bool, beta bool, err error) {
	return s.dao.HasOpenQuestPerm(address)
}

// HasBindSocialAccount 获取用户是否绑定社交账号
func (s *Service) HasBindSocialAccount(address string) (data map[string]bool, err error) {
	return s.dao.HasBindSocialAccount(address)
}
