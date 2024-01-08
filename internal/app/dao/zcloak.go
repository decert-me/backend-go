package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/pkg/log"
	"errors"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func (d *Dao) SaveSignAndDid(address, nonce string, request request.SaveSignAndDidRequest) (err error) {
	if address == "" {
		return errors.New("ParameterError")
	}
	// 查询 Did 是否已经被绑定
	didAddress, err := d.GetDidAddress(request.DidAddress)
	if err != nil {
		log.Errorv("查询 Did 是否已经被绑定失败", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	if didAddress != "" {
		return errors.New("DidAlreadyBound")
	}
	// 查询地址是否绑定过Did
	did, err := d.GetAddressDid(address)
	if err != nil {
		log.Errorv("查询地址是否绑定过Did失败", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	if did != "" {
		return errors.New("AddressAlreadyBoundDid")
	}
	zcloakDid := model.ZcloakDid{
		Nonce:      nonce,
		Signature:  request.Sign,
		KeyFile:    request.KeyFile,
		Address:    address,
		DidAddress: request.DidAddress,
	}
	err = d.db.Model(&model.ZcloakDid{}).Create(&zcloakDid).Error
	if err != nil {
		log.Errorv("保存签名和DID账号失败", zap.Error(err))
		return errors.New("UnexpectedError")
	}
	return
}

func (d *Dao) GetAddressDid(address string) (did string, err error) {
	err = d.db.Model(&model.ZcloakDid{}).Select("did_address").Where("address", address).First(&did).Error
	if err == gorm.ErrRecordNotFound {
		return "", nil
	}
	return
}

func (d *Dao) GetDidAddress(did string) (address string, err error) {
	err = d.db.Model(&model.ZcloakDid{}).Select("address").Where("did_address", did).First(&address).Error
	if err == gorm.ErrRecordNotFound {
		return "", nil
	}
	return
}

func (d *Dao) GetVcInfo(address, questID string) (vc interface{}, err error) {
	err = d.db.Model(&model.ZcloakCard{}).
		Select("vc").
		Joins("LEFT JOIN zcloak_did ON zcloak_card.holder = zcloak_did.did_address").
		Where("zcloak_did.address = ? AND quest_id = ?", address, questID).First(&vc).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return
}

func (d *Dao) GetDidHighestScore(did string, questID uint) (highestScore int64, err error) {
	err = d.db.Model(&model.ZcloakCard{}).
		Select("score").
		Where("did = ? AND quest_id = ?", did, questID).
		Order("score desc").
		First(&highestScore).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	return
}

func (d *Dao) SaveZcloakCard(zcloakCard model.ZcloakCard) (err error) {
	err = d.db.Model(&model.ZcloakCard{}).Create(&zcloakCard).Error
	return nil
}

func (d *Dao) AddressHasCard(address string) (hasCard bool, err error) {
	err = d.db.Model(&model.ZcloakCard{}).Where("address = ?", address).First(&model.ZcloakCard{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// GetKeyFileWithSignature 获取签名
func (d *Dao) GetKeyFileWithSignature(address string) (signature string, keyFile datatypes.JSON, nonce string, err error) {
	type Data struct {
		Signature string
		KeyFile   datatypes.JSON
		Nonce     string
	}
	var data Data
	err = d.db.Model(&model.ZcloakDid{}).Where("address = ?", address).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return signature, keyFile, nonce, nil
		}
		return signature, keyFile, nonce, err
	}
	return data.Signature, data.KeyFile, data.Nonce, nil
}
