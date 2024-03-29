package dao

import (
	"backend-go/internal/app/model"
	"github.com/shopspring/decimal"
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
func (d *Dao) GetUserOpenQuestReviewed(address string, tokenID string) (userOpenQuest model.UserOpenQuest, err error) {
	err = d.db.Model(&model.UserOpenQuest{}).
		Where("address", address).
		Where("token_id", tokenID).
		Where("open_quest_review_status = 2").
		Order("id desc").First(&userOpenQuest).Error
	return
}

// GetLatestQuestPassAnswer 获取用户通过的最新回答
func (d *Dao) GetLatestQuestPassAnswer(address string, tokenID string) (answer string, score int64, err error) {
	type UserChallengeLog struct {
		Answer    string `gorm:"column:answer"`
		UserScore int64  `gorm:"column:user_score"`
	}
	var userChallengeLog UserChallengeLog
	err = d.db.Model(&model.UserChallengeLog{}).
		Where("address", address).
		Where("token_id", tokenID).
		Where("pass=true").
		Order("id desc").First(&userChallengeLog).Error
	return userChallengeLog.Answer, userChallengeLog.UserScore, err
}

// GetUserChallengeLastedScore 获取用户最新挑战分数
func (d *Dao) GetUserChallengeLastedScore(address string, tokenID string) (score float64, err error) {
	var userChallengeLog model.UserChallengeLog
	err = d.db.Model(&model.UserChallengeLog{}).
		Where("address", address).
		Where("token_id", tokenID).
		Order("id desc").First(&userChallengeLog).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, nil
		}
		return 0, err
	}
	userScore := userChallengeLog.UserScore
	if userChallengeLog.IsOpenQuest {
		// 查询开放题最新分数
		var userOpenQuest model.UserOpenQuest
		err = d.db.Model(&model.UserOpenQuest{}).
			Where("address", address).
			Where("token_id", tokenID).
			Where("open_quest_review_status = 2").
			Order("id desc").First(&userOpenQuest).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return 0, nil
			}
			return 0, err
		}
		userScore = userOpenQuest.UserScore
	}
	_userScore := decimal.NewFromInt(userScore)
	userScoreRes, _ := _userScore.Div(decimal.NewFromInt(100)).Round(2).Float64()
	return userScoreRes, nil
}
