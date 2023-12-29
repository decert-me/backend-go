package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"errors"
	"fmt"
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
		err = d.db.Model(&model.Collection{}).Select("collection.*,COALESCE(tr.title,collection.title) as title,COALESCE(tr.description,collection.description) as description").
			Joins("LEFT JOIN collection_translated as tr ON collection.id = tr.collection_id AND tr.language = ?", r.Language).
			Where("collection.id", collectionID).First(&collection).Error
		if err != nil {
			return questList, collection, err
		}
	} else {
		// 查询合辑信息
		err = d.db.Model(&model.Collection{}).Select("collection.*,COALESCE(tr.title,collection.title) as title,COALESCE(tr.description,collection.description) as description").
			Joins("LEFT JOIN collection_translated as tr ON collection.id = tr.collection_id AND tr.language = ?", r.Language).
			Where("collection.uuid", r.ID).First(&collection).Error
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
	db := d.db.Model(&model.CollectionRelate{}).
		Joins("left join quest ON collection_relate.quest_id=quest.id").
		Joins("left join quest_translated as tr ON quest.token_id = tr.token_id AND tr.language = ?", r.Language).
		Where("collection_relate.collection_id = ? AND quest.status=1", collection.ID)
	if r.Address != "" {
		db.Select("quest.*,c.claimed,COALESCE(tr.title,quest.title) as title,COALESCE(tr.description,quest.description) as description")
		db.Joins("LEFT JOIN user_challenges c ON quest.token_id = c.token_id AND c.address = ?", r.Address)
	} else {
		db.Select("quest.*,COALESCE(tr.title,quest.title) as title,COALESCE(tr.description,quest.description) as description")
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

// GetCollectionFlashRank 获取合辑闪电榜
func (d *Dao) GetCollectionFlashRank(address, collectionID string) (res response.GetCollectionFlashRankRes, err error) {
	// 查询合辑信息
	var collection model.Collection
	err = d.db.Model(&model.Collection{}).Where("id", collectionID).First(&collection).Error
	if err != nil {
		return res, err
	}
	// 合辑未完结直接返回
	if collection.TokenId == 0 {
		return res, nil
	}
	// 查询合辑内挑战列表
	var questList []model.Quest
	err = d.db.Model(&model.CollectionRelate{}).
		Select("quest.token_id,quest.quest_data").
		Joins("left join quest ON collection_relate.token_id=quest.token_id").
		Where("collection_relate.collection_id", collection.ID).
		Where("collection_relate.status = 1").Find(&questList).Error
	if err != nil {
		return res, err
	}
	// 区分开放题
	var allTokenIDList []int64
	for _, quest := range questList {
		allTokenIDList = append(allTokenIDList, quest.TokenId)
	}
	allTokenIDCount := len(allTokenIDList)
	var havingSQL string
	rawSQL := `SELECT address
			FROM user_challenge_log
			WHERE token_id IN ? AND address !='' AND pass=true AND deleted_at IS NULL
			GROUP BY address
			HAVING COUNT(DISTINCT token_id) = ?`
	havingSQL = d.db.Model(&model.Collection{}).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Raw(rawSQL, allTokenIDList, allTokenIDCount)
	})
	// 获取合辑闪电榜
	rankListSQL := `
		WITH total AS(
			SELECT address,token_id, created_at
			FROM user_challenge_log
			WHERE token_id IN ? AND user_challenge_log.address !='' AND pass=true AND is_open_quest=false AND deleted_at IS NULL
			UNION
			SELECT address,token_id, created_at
			FROM user_open_quest
			WHERE token_id IN ? AND pass=true AND deleted_at IS NULL
		),ranked AS(
		SELECT total.address,token_id, created_at,ROW_NUMBER() OVER (PARTITION BY total.address,total.token_id ORDER BY created_at ASC) as rn
		FROM total
		INNER JOIN ( 
	`
	rankListSQL = rankListSQL + havingSQL + ` ) AS a ON total.address = a.address
		),ranked_list AS(
		SELECT address,token_id, created_at,ROW_NUMBER() OVER (PARTITION BY address ORDER BY created_at DESC) as rn
		FROM ranked WHERE rn=1 
		)
		SELECT ROW_NUMBER() OVER (ORDER BY ranked_list.created_at ASC) as rank,ranked_list.address,ranked_list.created_at as finish_time,users.avatar
		FROM ranked_list
		LEFT JOIN users ON ranked_list.address=users.address
		WHERE rn=1 ORDER BY created_at asc LIMIT 10;
	`
	err = d.db.Raw(rankListSQL, allTokenIDList, allTokenIDList).Scan(&res.RankList).Error
	if err != nil {
		return res, err
	}
	// 查询用户排名
	userRankSQL := `
		WITH total AS(
			SELECT address,token_id, created_at
			FROM user_challenge_log
			WHERE token_id IN ? AND user_challenge_log.address !='' AND pass=true AND is_open_quest=false  AND deleted_at IS NULL
			UNION
			SELECT address,token_id, created_at
			FROM user_open_quest
			WHERE token_id IN ? AND pass=true AND deleted_at IS NULL
		),ranked AS(
		SELECT total.address,token_id, created_at,ROW_NUMBER() OVER (PARTITION BY total.address,total.token_id ORDER BY created_at ASC) as rn
		FROM total
		INNER JOIN (
	`
	userRankSQL = userRankSQL + havingSQL + ` ) AS a ON total.address = a.address
		),ranked_list AS(
			SELECT address,token_id, created_at,ROW_NUMBER() OVER (PARTITION BY address ORDER BY created_at DESC) as rn
			FROM ranked WHERE rn=1 
		),ranked_score_user AS(
			SELECT ROW_NUMBER() OVER (ORDER BY ranked_list.created_at ASC) as rank,ranked_list.address,ranked_list.created_at as finish_time
			FROM ranked_list
			WHERE rn=1
		)
		SELECT rank,ranked_score_user.address,ranked_score_user.finish_time,users.avatar
		FROM ranked_score_user
		LEFT JOIN users ON ranked_score_user.address=users.address
		WHERE ranked_score_user.address = ? LIMIT 1;
	`
	err = d.db.Raw(userRankSQL, allTokenIDList, allTokenIDList, address).Scan(&res).Error
	return
}

// GetCollectionHighRank 获取合辑高分榜
func (d *Dao) GetCollectionHighRank(address, collectionID string) (res response.GetCollectionHighRankRes, err error) {
	// 查询合辑信息
	var collection model.Collection
	err = d.db.Model(&model.Collection{}).Where("id", collectionID).First(&collection).Error
	if err != nil {
		return res, err
	}
	// 合辑未完结直接返回
	if collection.TokenId == 0 {
		return res, nil
	}
	// 查询合辑内挑战列表
	var questList []model.Quest
	err = d.db.Model(&model.CollectionRelate{}).
		Select("quest.token_id,quest.quest_data").
		Joins("left join quest ON collection_relate.token_id=quest.token_id").
		Where("collection_relate.collection_id", collection.ID).
		Where("collection_relate.status = 1").Find(&questList).Error
	if err != nil {
		return res, err
	}
	// 区分开放题
	var allTokenIDList []int64
	for _, quest := range questList {
		allTokenIDList = append(allTokenIDList, quest.TokenId)
	}
	allTokenIDCount := len(allTokenIDList)
	var havingSQL string
	rawSQL := `SELECT address
			FROM user_challenge_log
			WHERE token_id IN ? AND address !='' AND pass=true AND deleted_at IS NULL
			GROUP BY address
			HAVING COUNT(DISTINCT token_id) = ?`
	havingSQL = d.db.Model(&model.Collection{}).ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Raw(rawSQL, allTokenIDList, allTokenIDCount)
	})

	// 获取合辑高分榜
	rankListSQL := `
		 WITH total AS(
		SELECT address,token_id, created_at,user_score/100 as score
		FROM user_challenge_log
		WHERE token_id IN ? AND user_challenge_log.address !='' AND pass=true AND is_open_quest=false AND deleted_at IS NULL
		UNION
		SELECT address,token_id, created_at,open_quest_score as score
		FROM user_open_quest
		WHERE token_id IN ? AND pass=true AND deleted_at IS NULL
		),ranked AS(
		SELECT total.address,token_id, created_at,score,ROW_NUMBER() OVER (PARTITION BY total.address,total.token_id ORDER BY score DESC) as rn
		FROM total
		INNER JOIN ( 
	`
	rankListSQL = rankListSQL + havingSQL + `) AS a ON total.address = a.address
		),ranked_score AS(
		SELECT ranked.address, SUM(score) as total_score,max(created_at) as created_at
		FROM ranked WHERE rn=1
		GROUP BY ranked.address
		)
		SELECT ROW_NUMBER() OVER (ORDER BY ranked_score.total_score DESC,ranked_score.created_at ASC) as rank,total_score as score,ranked_score.address,ranked_score.created_at as finish_time,users.avatar
		FROM ranked_score
		LEFT JOIN users ON ranked_score.address=users.address
		ORDER BY total_score DESC,ranked_score.created_at ASC LIMIT 10;
	`
	err = d.db.Raw(rankListSQL, allTokenIDList, allTokenIDList).Scan(&res.RankList).Error
	if err != nil {
		return res, err
	}
	// 查询用户排名
	userRankSQL := `
		 WITH total AS(
		SELECT address,token_id, created_at,user_score/100 as score
		FROM user_challenge_log
		WHERE token_id IN ? AND user_challenge_log.address !='' AND pass=true AND is_open_quest=false AND deleted_at IS NULL
		UNION
		SELECT address,token_id, created_at,open_quest_score as score
		FROM user_open_quest
		WHERE token_id IN ? AND pass=true AND deleted_at IS NULL
		),ranked AS(
		SELECT total.address,token_id, created_at,score,ROW_NUMBER() OVER (PARTITION BY total.address,total.token_id ORDER BY score DESC) as rn
		FROM total
		INNER JOIN ( 
	`
	userRankSQL = userRankSQL + havingSQL + ` ) AS a ON total.address = a.address
		),ranked_score AS(
		SELECT ranked.address, SUM(score) as total_score,max(created_at) as created_at
		FROM ranked WHERE rn=1
		GROUP BY ranked.address
		),ranked_score_user AS(
			SELECT ROW_NUMBER() OVER (ORDER BY ranked_score.total_score DESC,ranked_score.created_at ASC) as rank,total_score as score,ranked_score.address,ranked_score.created_at as finish_time
			FROM ranked_score
		)
		SELECT rank,score,ranked_score_user.address,ranked_score_user.finish_time,users.avatar
		FROM ranked_score_user
		LEFT JOIN users ON ranked_score_user.address=users.address
		WHERE ranked_score_user.address = ? LIMIT 1;
	`
	err = d.db.Raw(userRankSQL, allTokenIDList, allTokenIDList, address).Scan(&res).Error
	return
}

// GetCollectionHolderRank 获取合辑 Holder 榜单
func (d *Dao) GetCollectionHolderRank(address string, id int64, page, pageSize int) (res []response.GetCollectionHolderListRes, total int64, err error) {
	// 分页参数
	limit := pageSize
	offset := pageSize * (page - 1)
	// 查询合辑的tokenID
	var tokenId int64
	err = d.db.Model(&model.Collection{}).Select("token_id").Where("id", id).Scan(&tokenId).Error
	if err != nil {
		return res, total, err
	}
	err = d.db.Model(&model.UserChallenges{}).Where("token_id", tokenId).Count(&total).Error
	if err != nil {
		return res, total, err
	}
	err = d.db.Model(&model.UserChallenges{}).
		Select("ROW_NUMBER() OVER (ORDER BY user_challenges.add_ts ASC,user_challenges.id ASC) as rank,users.*,to_timestamp(user_challenges.add_ts) as claim_time").
		Joins("LEFT JOIN users ON user_challenges.address=users.address").
		Where("user_challenges.token_id", tokenId).
		Order("user_challenges.add_ts ASC,user_challenges.id ASC").
		Limit(limit).Offset(offset).
		Find(&res).Error
	fmt.Println(res)
	return res, total, err
}
