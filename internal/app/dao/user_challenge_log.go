package dao

import (
	"backend-go/internal/app/model"
	"gorm.io/gorm"
)

func (d *Dao) CreateChallengeLog(log *model.UserChallengeLog) (err error) {
	return d.db.Create(&log).Error
}

func (d *Dao) CreateUserOpenQuest(userOpenQuest *model.UserOpenQuest) (err error) {
	// 获取最新的审核记录
	var userOpenQuestReviewed model.UserOpenQuest
	if err := d.db.Model(&model.UserOpenQuest{}).
		Where("address", userOpenQuest.Address).
		Where("token_id", userOpenQuest.TokenId).
		Where("open_quest_review_status = 2").
		Order("id desc").First(&userOpenQuestReviewed).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 查询是否有未审核的记录
			var userOpenQuestNotReviewed model.UserOpenQuest
			if err := d.db.Model(&model.UserOpenQuest{}).
				Where("address", userOpenQuest.Address).
				Where("token_id", userOpenQuest.TokenId).
				Where("open_quest_review_status = 1").
				Order("id desc").First(&userOpenQuestNotReviewed).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					// 创建新纪录
					return d.db.Model(&model.UserOpenQuest{}).Create(&userOpenQuest).Error
				}
				return err
			} else {
				return d.db.Model(&model.UserOpenQuest{}).Where("id = ?", userOpenQuestNotReviewed.ID).Updates(&userOpenQuest).Error
			}
		}
		return err
	}
	// 判断是否有新的提交
	var userOpenQuestNext model.UserOpenQuest
	if err := d.db.Model(&model.UserOpenQuest{}).
		Where("address", userOpenQuest.Address).
		Where("token_id", userOpenQuest.TokenId).
		Where("id > ?", userOpenQuestReviewed.ID).
		Order("id desc").First(&userOpenQuestNext).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 创建新纪录
			return d.db.Model(&model.UserOpenQuest{}).Create(&userOpenQuest).Error
		}
		return err
	}
	// 否则更新数据
	return d.db.Model(&model.UserOpenQuest{}).Where("id = ?", userOpenQuestNext.ID).Updates(&userOpenQuest).Error
}

// GetUserOpenQuestReviewed 获取用户最新已审核开放题
func (d *Dao) GetUserOpenQuestReviewed(address string, tokenID int64) (userOpenQuest model.UserOpenQuest, err error) {
	err = d.db.Model(&model.UserOpenQuest{}).
		Where("address", address).
		Where("token_id", tokenID).
		Where("open_quest_review_status = 2").
		Order("id desc").First(&userOpenQuest).Error
	return
}
