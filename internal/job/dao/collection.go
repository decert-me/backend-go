package dao

import "backend-go/internal/app/model"

func (d *Dao) UpdateCollectionOnce(collectionID int64, quest model.Quest) error {
	collection := model.Collection{
		UUID:        quest.UUID,
		TokenId:     quest.TokenId,
		Uri:         quest.Uri,
		MetaData:    quest.MetaData,
		Recommend:   quest.Recommend,
		Version:     quest.Version,
		Creator:     quest.Creator,
		ExtraData:   quest.ExtraData,
		QuestData:   quest.QuestData,
		Description: quest.Description,
	}
	return d.db.Model(&model.Collection{}).Where("id = ? AND token_id = 0", collectionID).Updates(collection).Error
}
