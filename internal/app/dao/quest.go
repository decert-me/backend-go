package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"fmt"
	"gorm.io/gorm"
)

func (d *Dao) HasTokenId(tokenId int64) (has bool, err error) {
	var count int64
	err = d.db.Model(&model.Quest{}).Where("token_id", tokenId).Count(&count).Error
	if count > 0 {
		has = true
	}
	return
}

func (d *Dao) ValidTokenId(tokenId int64) (valid bool, err error) {
	var quest model.Quest
	err = d.db.
		Where("token_id", tokenId).Where("disabled", false).Where("is_draft", false).
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
	return d.db.Create(&req).Error
}

func (d *Dao) GetQuestList(req *request.GetQuestListRequest) (questList []response.GetQuestListRes, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	db := d.db.Model(&model.Quest{})

	db.Where(&req.Quest)
	err = db.Count(&total).Error
	if err != nil {
		return questList, total, err
	}
	if req.OrderKey == "token_id" {
		fmt.Println(req.OrderKey)
		fmt.Println(req.Desc)
		if req.Desc {
			db.Order("token_id desc")
		} else {
			db.Order("token_id asc")
		}
	} else {
		db.Order("token_id desc")
	}
	if req.SearchKey != "" {
		db.Where("quest.title ILIKE ? OR quest.description ILIKE ?", "%"+req.SearchKey+"%", "%"+req.SearchKey+"%")
	}
	if req.Address != "" {
		db.Select("quest.*,c.claimed")
		db.Joins("LEFT JOIN user_challenges c ON quest.token_id = c.token_id AND c.address = ?", req.Address)
	} else {
		db.Select("*")
	}
	err = db.Limit(limit).Offset(offset).Find(&questList).Error

	return questList, total, err
}

func (d *Dao) GetQuestByTokenID(id int64) (quest model.Quest, err error) {
	err = d.db.Model(&model.Quest{}).Where("token_id", id).First(&quest).Error
	return
}

func (d *Dao) GetQuestByUUID(uuid string) (quest model.Quest, err error) {
	err = d.db.Model(&model.Quest{}).Where("uuid", uuid).First(&quest).Error
	return
}

func (d *Dao) GetQuestWithClaimStatusByTokenID(id int64, address string) (quest response.GetQuestRes, err error) {
	err = d.db.Model(&model.Quest{}).
		Select("quest.*,b.claimed").
		Joins("left join user_challenges b ON quest.token_id=b.token_id AND b.address= ?", address).
		Where("quest.token_id", id).
		First(&quest).Error
	return
}

func (d *Dao) GetQuestWithClaimStatusByUUID(uuid string, address string) (quest response.GetQuestRes, err error) {
	err = d.db.Model(&model.Quest{}).
		Select("quest.*,b.claimed").
		Joins("left join user_challenges b ON quest.token_id=b.token_id AND b.address= ?", address).
		Where("quest.uuid", uuid).
		First(&quest).Error
	return
}

func (d *Dao) GetUserQuestList(req *request.GetUserQuestListRequest) (questList []response.GetUserQuestListRes, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := d.db.Model(&model.Quest{})
	db.Where(&req.Quest)
	err = db.Count(&total).Error
	if err != nil {
		return questList, total, err
	}
	err = db.Limit(limit).Offset(offset).Order("add_ts desc").Find(&questList).Error

	return questList, total, err
}

func (d *Dao) GetUserQuestListWithClaimed(req *request.GetUserQuestListRequest) (questList []response.QuestWithClaimed, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := d.db.Model(&response.QuestWithClaimed{})
	db.Select("quest.*,EXISTS (SELECT 1 FROM user_challenges WHERE quest.token_id = user_challenges.token_id) AS has_claim")
	db.Where(&req.Quest)
	err = db.Count(&total).Error
	if err != nil {
		return questList, total, err
	}
	err = db.Limit(limit).Offset(offset).Order("add_ts desc").Find(&questList).Error

	return questList, total, err
}

func (d *Dao) GetQuestChallengeUserByTokenID(tokenId int64) (res response.GetQuestChallengeUserRes, err error) {
	err = d.db.Model(&model.UserChallenges{}).Where("token_id", tokenId).Count(&res.Times).Error
	if err != nil {
		return res, err
	}
	err = d.db.Model(&model.UserChallenges{}).
		Select("users.*").
		Joins("LEFT JOIN users ON user_challenges.address=users.address").
		Where("user_challenges.token_id", tokenId).
		Order("user_challenges.add_ts desc").
		Limit(12).
		Find(&res.Users).Error
	return res, err
}

func (d *Dao) GetQuestChallengeUserByUUID(uuid string) (res response.GetQuestChallengeUserRes, err error) {
	var quest model.Quest
	err = d.db.Model(&model.Quest{}).Select("token_id").Where("uuid", uuid).First(&quest).Error
	if err != nil {
		return res, err
	}

	err = d.db.Model(&model.UserChallenges{}).Where("token_id", quest.TokenId).Count(&res.Times).Error
	if err != nil {
		return res, err
	}
	err = d.db.Model(&model.UserChallenges{}).
		Select("users.*").
		Joins("LEFT JOIN users ON user_challenges.address=users.address").
		Where("user_challenges.token_id", quest.TokenId).
		Order("user_challenges.add_ts desc").
		Limit(12).
		Find(&res.Users).Error
	return res, err
}

func (d *Dao) UpdateQuest(req *model.Quest) (err error) {
	return d.db.Where("token_id", req.TokenId).Updates(&req).Error
}
