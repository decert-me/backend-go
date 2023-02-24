package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"gorm.io/gorm"
)

func (d *Dao) HasTokenId(tokenId uint64) (id uint, err error) {
	var quest model.Quest
	err = d.db.Select("id").Where("tokenId", tokenId).First(&quest).Error
	return quest.ID, err
}

func (d *Dao) ValidTokenId(tokenId int64) (valid bool, err error) {
	var quest model.Quest
	err = d.db.
		Where("tokenId", tokenId).Where("disabled", false).Where("isDraft", false).
		First(&quest).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, err
}

func (d *Dao) CreateQuest(req *model.Quest) (err error) {
	return d.db.Model(&model.Quest{}).Create(&req).Error
}

func (d *Dao) GetQuestList(req *request.GetQuestListRequest) (questList []model.Quest, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	db := d.db.Model(&model.Quest{}).Where(&req.Quest)
	err = db.Count(&total).Error
	if err != nil {
		return questList, total, err
	}
	err = db.Limit(limit).Offset(offset).Order("add_ts desc").Find(&questList).Error

	return questList, total, err
}

func (d *Dao) GetQuest(req *model.Quest) (quest model.Quest, err error) {
	err = d.db.Model(&model.Quest{}).Where("tokenId", req.TokenId).First(&quest).Error
	return
}
