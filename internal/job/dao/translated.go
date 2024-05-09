package dao

import "backend-go/internal/app/model"

func (d *Dao) DeleteQuestTranslated(tokenId string) (err error) {
	err = d.db.Model(&model.QuestTranslated{}).
		Where("token_id = ?", tokenId).
		Delete(&model.QuestTranslated{}).Error
	return
}
