package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"fmt"
	"github.com/tidwall/gjson"
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
	// Quest
	questSQL := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		tx = tx.Where("quest.status = 1 AND quest.disabled = false AND collection_status=1")
		tx = tx.Where(&req.Quest)
		if req.SearchKey != "" {
			tx = tx.Where("quest.title ILIKE ? OR quest.description ILIKE ?", "%"+req.SearchKey+"%", "%"+req.SearchKey+"%")
		}
		if req.Address != "" {
			tx = tx.Select("quest.id,quest.uuid,quest.title,quest.label,quest.disabled,quest.description,quest.dependencies,quest.is_draft,quest.add_ts,quest.token_id,quest.type,quest.difficulty,quest.estimate_time,quest.creator,quest.meta_data,quest.quest_data,quest.extra_data,quest.uri,quest.pass_score,quest.total_score,quest.recommend,quest.status,quest.style,quest.cover,quest.author,quest.sort,quest.collection_status,c.claimed,COALESCE(o.open_quest_review_status,0) as open_quest_review_status,COALESCE(o.pass,p.pass,false) as claimable")
			tx = tx.Joins("LEFT JOIN user_challenges c ON quest.token_id = c.token_id AND c.address = ?", req.Address)
			tx = tx.Joins("LEFT JOIN (WITH ranked_statuses AS (SELECT token_id, open_quest_review_status,pass,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY id DESC) as rn FROM user_open_quest WHERE address=? AND deleted_at IS NULL) SELECT open_quest_review_status,token_id,pass FROM ranked_statuses WHERE rn = 1) o ON quest.token_id = o.token_id", req.Address)
			tx = tx.Joins("LEFT JOIN (WITH ranked_log AS (SELECT token_id,pass,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY id DESC) as rn FROM user_challenge_log WHERE address = ? AND deleted_at IS NULL) SELECT token_id,pass FROM ranked_log WHERE rn = 1) p ON quest.token_id = p.token_id", req.Address)
		} else {
			tx = tx.Select("quest.id,quest.uuid,quest.title,quest.label,quest.disabled,quest.description,quest.dependencies,quest.is_draft,quest.add_ts,quest.token_id,quest.type,quest.difficulty,quest.estimate_time,quest.creator,quest.meta_data,quest.quest_data,quest.extra_data,quest.uri,quest.pass_score,quest.total_score,quest.recommend,quest.status,quest.style,quest.cover,quest.author,quest.sort,quest.collection_status,FALSE as claimed,0 as open_quest_review_status,false as claimable")
		}
		return tx.Find(&[]response.GetQuestListRes{})
	})
	// Collection
	collectionSQL := d.db.Model(&model.Collection{}).ToSQL(func(tx *gorm.DB) *gorm.DB {
		tx = tx.Select("id,uuid,title,label,disabled,description,dependencies,is_draft,add_ts,token_id,type,difficulty,estimate_time,creator,meta_data,quest_data,extra_data,uri,pass_score,total_score,recommend,status,style,cover,author,sort,collection_status,FALSE as claimed,0 as open_quest_review_status,false as claimable")
		return tx.Where("status = 1").Find(&[]response.GetQuestListRes{})
	})
	//fmt.Println("questSQL", questSQL)
	//fmt.Println("collectionSQL", collectionSQL)
	// 执行 UNION 查询
	unionSQL := fmt.Sprintf("SELECT * FROM((%s) UNION (%s)) as t ORDER BY sort desc,add_ts desc LIMIT %d OFFSET %d", questSQL, collectionSQL, limit, offset)
	unionCountSQL := fmt.Sprintf("SELECT count(1) FROM((%s) UNION (%s)) as t", questSQL, collectionSQL)
	db.Raw(unionCountSQL).Scan(&total)
	db.Raw(unionSQL).Scan(&questList)
	// 查询集合数量
	for i := 0; i < len(questList); i++ {
		if questList[i].Style == 2 {
			if questList[i].ID == 0 {
				continue
			}
			// 合辑作者
			if questList[i].Author != "" {
				d.db.Model(&model.Users{}).Where("address = ?", questList[i].Author).First(&questList[i].AuthorInfo)
			}

			// 数量
			var collectionQuestList []model.Quest
			err = d.db.Model(&model.Quest{}).Where("collection_id = ?", questList[i].ID).Find(&collectionQuestList).Error
			if err != nil {
				return questList, total, err
			}
			questList[i].CollectionCount = int64(len(collectionQuestList))
			// 预估时间
			var estimateTimeTotal int64
			for _, quest := range collectionQuestList {
				estimateTimeTotal += gjson.Get(string(quest.QuestData), "estimateTime").Int()
				fmt.Println("estimateTimeTotal", estimateTimeTotal)
			}

			fmt.Println("estimateTimeTotal", estimateTimeTotal)
			if estimateTimeTotal != 0 {
				questList[i].EstimateTime = estimateTimeTotal
			}
		}
	}
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
		Select("quest.*,b.claimed,b.user_score,b.nft_address,COALESCE(o.open_quest_review_status,0) as open_quest_review_status,COALESCE(o.answer,l.answer) as answer").
		Joins("left join user_challenges b ON quest.token_id=b.token_id AND b.address= ?", address).
		Joins("left join user_challenge_log l ON quest.token_id=l.token_id AND l.address= ? AND l.deleted_at IS NULL", address).
		Joins("left join user_open_quest o ON quest.token_id=o.token_id AND o.address= ? AND o.deleted_at IS NULL", address).
		Where("quest.token_id", id).
		Order("l.add_ts desc,o.id desc").
		First(&quest).Error
	return
}

func (d *Dao) GetQuestWithClaimStatusByUUID(uuid string, address string) (quest response.GetQuestRes, err error) {
	err = d.db.Model(&model.Quest{}).
		Select("quest.*,b.claimed,b.user_score,b.nft_address,COALESCE(o.open_quest_review_status,0) as open_quest_review_status,COALESCE(o.answer,l.answer) as answer").
		Joins("left join user_challenges b ON quest.token_id=b.token_id AND b.address= ?", address).
		Joins("left join user_challenge_log l ON quest.token_id=l.token_id AND l.address= ? AND l.deleted_at IS NULL", address).
		Joins("left join user_open_quest o ON quest.token_id=o.token_id AND o.address= ? AND o.deleted_at IS NULL", address).
		Where("quest.uuid", uuid).
		Order("l.add_ts desc,o.id desc").
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
	db.Where("quest.status = 1")
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
