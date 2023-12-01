package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func (d *Dao) GetCollectionChallengeUserByID(r request.GetCollectionChallengeUser) (res response.GetCollectionChallengeUserRes, total int64, err error) {
	// 兼容UUID
	collectionID, idErr := cast.ToUintE(r.CollectionID)
	if idErr != nil {
		// 查询合辑信息
		err = d.db.Model(&model.Collection{}).Select("id").Where("uuid", r.CollectionID).First(&collectionID).Error
		if err != nil {
			return res, total, err
		}
	}
	// 分页
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	// 查询合辑列表
	var tokenIDList []uint
	err = d.db.Model(&model.CollectionRelate{}).Where("collection_id", collectionID).Pluck("token_id", &tokenIDList).Error
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

func (d *Dao) GetCollectionQuest(r request.GetCollectionQuestRequest) (questList []response.GetQuestListRes, collection response.GetCollectionRes, err error) {
	collectionID, idErr := cast.ToUintE(r.ID)
	if idErr == nil {
		// 查询合辑信息
		err = d.db.Model(&model.Collection{}).
			Select("collection.*,c.claimed,c.badge_token_id").
			Joins("LEFT JOIN user_challenges c ON collection.token_id = c.token_id AND c.address = ?", r.Address).
			Where("collection.id", collectionID).First(&collection).Error
		if err != nil {
			return questList, collection, err
		}
	} else {
		// 查询合辑信息
		err = d.db.Model(&model.Collection{}).Where("uuid", r.ID).First(&collection).Error
		if err != nil {
			return questList, collection, err
		}
	}
	// 查询是否领取
	if r.Address != "" {
		var claimed int
		err = d.db.Model(&model.UserChallenges{}).
			Select("COUNT(1)").
			Where("token_id = ?", collection.TokenId).
			Where("address = ?", r.Address).
			Scan(&claimed).Error
		if err != nil {
			return questList, collection, err
		}
		if claimed > 0 {
			collection.Claimed = true
		}
	}
	// 查询合辑内挑战
	db := d.db.Model(&model.CollectionRelate{}).Joins("left join quest ON collection_relate.quest_id=quest.id").
		Where("collection_relate.collection_id = ? AND quest.status=1", collection.ID)
	if r.Address != "" {
		db.Select("quest.id,quest.version,quest.uuid,quest.title,quest.label,quest.disabled,quest.description,quest.dependencies,quest.is_draft,quest.add_ts,quest.token_id,quest.type,quest.difficulty,quest.estimate_time,quest.creator,quest.meta_data,quest.quest_data,quest.extra_data,quest.uri,quest.pass_score,quest.total_score,quest.recommend,quest.status,quest.style,quest.cover,quest.author,quest.sort,quest.collection_status,c.claimed,COALESCE(o.open_quest_review_status,0) as open_quest_review_status,COALESCE(o.pass,p.pass,false) as claimable")
		db.Joins("LEFT JOIN user_challenges c ON quest.token_id = c.token_id AND c.address = ?", r.Address)
		db.Joins("LEFT JOIN (WITH ranked_statuses AS (SELECT token_id, open_quest_review_status,pass,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY id DESC) as rn FROM user_open_quest WHERE address=? AND deleted_at IS NULL) SELECT open_quest_review_status,token_id,pass FROM ranked_statuses WHERE rn = 1) o ON quest.token_id = o.token_id", r.Address)
		db.Joins("LEFT JOIN (WITH ranked_log AS (SELECT token_id,pass,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY id DESC) as rn FROM user_challenge_log WHERE address = ? AND deleted_at IS NULL) SELECT token_id,pass FROM ranked_log WHERE rn = 1) p ON quest.token_id = p.token_id", r.Address)
		//db.Select("quest.*,c.claimed,COALESCE(c.badge_token_id,quest.token_id) as badge_token_id")
		//db.Joins("LEFT JOIN user_challenges c ON quest.token_id = c.token_id AND c.address = ?", r.Address)
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

func (d *Dao) GetCollectionByTokenID(tokenID int64) (collection model.Collection, err error) {
	err = d.db.Model(&model.Collection{}).Where("token_id", tokenID).First(&collection).Error
	return
}

func (d *Dao) GetQuestListByCollectionID(collectionID uint) (questList []model.Quest, err error) {
	err = d.db.Model(&model.CollectionRelate{}).
		Joins("left join quest ON collection_relate.quest_id=quest.id").
		Where("collection_relate.collection_id = ? AND quest.status=1", collectionID).
		Order("collection_relate.sort desc").Find(&questList).Error
	return
}

// CheckQuestInCollection 查询挑战是否在合辑内
func (d *Dao) CheckQuestInCollection(r request.CheckQuestInCollectionRequest) (res response.CheckQuestInCollectionRes, err error) {
	err = d.db.Model(&model.CollectionRelate{}).
		Select("collection_id").
		Where("token_id = ?", r.TokenID).
		First(&res.CollectionID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return res, err
	}
	if res.CollectionID != 0 {
		res.IsInCollection = true
	}
	return
}
