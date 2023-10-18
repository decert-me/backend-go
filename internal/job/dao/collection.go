package dao

import "backend-go/internal/app/model"

func (d *Dao) UpdateCollection(collectionID int64, quest model.Quest) error {
	collection := model.Collection{
		UUID:      quest.UUID,
		TokenId:   quest.TokenId,
		Uri:       quest.Uri,
		MetaData:  quest.MetaData,
		Recommend: quest.Recommend,
	}
	return d.db.Model(&model.Collection{}).Where("id = ?", collectionID).Updates(collection).Error
}
