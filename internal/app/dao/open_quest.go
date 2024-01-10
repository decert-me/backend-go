package dao

import "backend-go/internal/app/model"

// GetLatestOpenQuestPassAnswer 获取用户通过的最新回答
func (d *Dao) GetLatestOpenQuestPassAnswer(address string, tokenID int64) (answer string, score int64, err error) {
	type UserOpenQuest struct {
		Answer         string `gorm:"column:answer"`
		OpenQuestScore int64  `gorm:"column:open_quest_score"`
	}
	var userOpenQuest UserOpenQuest
	err = d.db.Model(&model.UserOpenQuest{}).
		Where("address", address).
		Where("token_id", tokenID).
		Where("open_quest_review_status = 2 AND pass=true").
		Order("id desc").First(&userOpenQuest).Error

	return userOpenQuest.Answer, userOpenQuest.OpenQuestScore, err
}
