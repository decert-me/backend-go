package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"errors"
)

func (d *Dao) GetCollectionChallengeUserByID(r request.GetCollectionChallengeUser) (res response.GetCollectionChallengeUserRes, total int64, err error) {
	// 分页
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	// 查询合辑列表
	var tokenIDList []uint
	err = d.db.Model(&model.CollectionRelate{}).Where("collection_id", r.CollectionID).Pluck("token_id", &tokenIDList).Error
	if err != nil {
		return res, total, err
	}
	if len(tokenIDList) == 0 {
		return res, total, errors.New("FetchFailed")
	}
	// 查询挑战用户
	err = d.db.Model(&model.UserChallenges{}).
		Select("DISTINCT ON (users.id) users.*").
		Joins("LEFT JOIN users ON user_challenges.address=users.address").
		Where("user_challenges.token_id in ?", tokenIDList).
		Order("users.id,user_challenges.add_ts desc").
		Limit(limit).Offset(offset).
		Find(&res.Users).Error
	// 查询挑战次数
	err = d.db.Model(&model.UserChallenges{}).Select("COUNT(DISTINCT user_challenges.address)").Where("token_id", tokenIDList).Count(&res.Times).Error
	if err != nil {
		return res, total, err
	}
	total = res.Times
	return res, total, err
}

func (d *Dao) GetCollectionQuest(r request.GetCollectionQuestRequest) (questList []response.GetQuestListRes, collection model.Collection, err error) {
	// 查询合辑信息
	err = d.db.Model(&model.Collection{}).Where("id", r.ID).First(&collection).Error
	if err != nil {
		return questList, collection, err
	}
	// 查询合辑内挑战
	db := d.db.Model(&model.CollectionRelate{}).Joins("left join quest ON collection_relate.quest_id=quest.id").
		Where("collection_relate.collection_id = ? AND quest.status=1", r.ID)
	if r.Address != "" {
		db.Select("quest.*,c.claimed")
		db.Joins("LEFT JOIN user_challenges c ON quest.token_id = c.token_id AND c.address = ?", r.Address)
	} else {
		db.Select("*")
	}
	err = db.Order("collection_relate.sort desc").Find(&questList).Error
	return
}

func (d *Dao) GetCollectionByID(id int) (collection model.Collection, err error) {
	err = d.db.Model(&model.Collection{}).Where("id", id).First(&collection).Error
	return
}
