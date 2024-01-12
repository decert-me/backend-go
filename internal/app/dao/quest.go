package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"fmt"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
)

func (d *Dao) HasTokenId(tokenId string) (has bool, err error) {
	var count int64
	err = d.db.Model(&model.Quest{}).Where("token_id", tokenId).Count(&count).Error
	if count > 0 {
		has = true
	}
	return
}

func (d *Dao) ValidTokenId(tokenId string) (valid bool, err error) {
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
		tx = tx.Joins("LEFT JOIN quest_translated ON quest.token_id = quest_translated.token_id AND quest_translated.language = ?", req.Language)
		if req.SearchKey != "" {
			tx = tx.Where("quest.title ILIKE ? OR quest.description ILIKE ?", "%"+req.SearchKey+"%", "%"+req.SearchKey+"%")
		}
		if req.Address != "" {
			tx = tx.Select("quest.id,quest.uuid,COALESCE(quest_translated.title,quest.title) as title,quest.label,quest.disabled,COALESCE(quest_translated.description,quest.description) as description,quest.dependencies,quest.is_draft,quest.add_ts,quest.token_id,quest.type,quest.difficulty,quest.estimate_time,quest.creator,quest.meta_data,quest.quest_data,quest.extra_data,quest.uri,quest.pass_score,quest.total_score,quest.recommend,quest.status,quest.style,quest.cover,quest.author,quest.sort,quest.collection_status,NOT (c.claimed = false AND zc.quest_id IS NULL) as claimed,COALESCE(o.open_quest_review_status,0) as open_quest_review_status,COALESCE(o.pass,p.pass,false) as claimable")
			tx = tx.Joins("LEFT JOIN user_challenges c ON quest.token_id = c.token_id AND c.address = ?", req.Address)
			tx = tx.Joins("LEFT JOIN (WITH zcloak_status AS (SELECT quest_id,ROW_NUMBER() OVER (PARTITION BY quest_id ORDER BY id DESC) as rn FROM zcloak_card WHERE address=? AND deleted_at IS NULL) SELECT quest_id FROM zcloak_status WHERE rn = 1) zc ON quest.id = zc.quest_id", req.Address)
			tx = tx.Joins("LEFT JOIN (WITH ranked_statuses AS (SELECT token_id, open_quest_review_status,pass,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY id DESC) as rn FROM user_open_quest WHERE address=? AND deleted_at IS NULL) SELECT open_quest_review_status,token_id,pass FROM ranked_statuses WHERE rn = 1) o ON quest.token_id = o.token_id", req.Address)
			tx = tx.Joins("LEFT JOIN (WITH ranked_log AS (SELECT token_id,pass,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY id DESC) as rn FROM user_challenge_log WHERE address = ? AND deleted_at IS NULL) SELECT token_id,pass FROM ranked_log WHERE rn = 1) p ON quest.token_id = p.token_id", req.Address)
		} else {
			tx = tx.Select("quest.id,quest.uuid,COALESCE(quest_translated.title,quest.title) as title,quest.label,quest.disabled,COALESCE(quest_translated.description,quest.description) as description,quest.dependencies,quest.is_draft,quest.add_ts,quest.token_id,quest.type,quest.difficulty,quest.estimate_time,quest.creator,quest.meta_data,quest.quest_data,quest.extra_data,quest.uri,quest.pass_score,quest.total_score,quest.recommend,quest.status,quest.style,quest.cover,quest.author,quest.sort,quest.collection_status,FALSE as claimed,0 as open_quest_review_status,false as claimable")
		}
		return tx.Find(&[]response.GetQuestListRes{})
	})
	// Collection
	collectionSQL := d.db.Model(&model.Collection{}).ToSQL(func(tx *gorm.DB) *gorm.DB {
		tx = tx.Select("collection.id,uuid,COALESCE(tr.title,collection.title) as title,label,disabled,COALESCE(tr.description,collection.description) as description,dependencies,is_draft,add_ts,token_id,type,difficulty,estimate_time,creator,meta_data,quest_data,extra_data,uri,pass_score,total_score,recommend,status,style,cover,author,sort,collection_status,FALSE as claimed,0 as open_quest_review_status,false as claimable").
			Joins("LEFT JOIN collection_translated as tr ON collection.id = tr.collection_id AND tr.language = ?", req.Language)
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

func (d *Dao) GetQuestByTokenIDWithLang(language string, tokenID string) (quest response.GetQuestRes, err error) {
	err = d.db.Model(&model.Quest{}).Select("quest.*,COALESCE(tr.title,quest.title) as title,COALESCE(tr.description,quest.description) as description,"+
		"COALESCE(tr.meta_data,quest.meta_data) as meta_data,COALESCE(tr.quest_data,quest.quest_data) as quest_data").
		Joins("LEFT JOIN quest_translated tr ON quest.token_id = tr.token_id AND tr.language = ?", language).Where("quest.token_id", tokenID).First(&quest.Quest).Error
	if err != nil {
		return quest, err
	}
	// 获取所有答案
	err = d.db.Raw(`SELECT answer AS answers
		FROM (
		SELECT  quest_data->>'answers' AS answer FROM quest WHERE token_id = ?
		UNION
		SELECT answer FROM quest_translated WHERE token_id = ? AND answer IS NOT NULL) AS combined_data
		`, tokenID, tokenID).Scan(&quest.Answers).Error
	return
}

func (d *Dao) GetQuestByTokenID(tokenID string) (quest model.Quest, err error) {
	err = d.db.Model(&model.Quest{}).Where("token_id", tokenID).First(&quest).Error
	return
}

func (d *Dao) GetQuestByUUID(language, uuid string) (quest response.GetQuestRes, err error) {
	err = d.db.Model(&model.Quest{}).Select("quest.*,COALESCE(tr.title,quest.title) as title,COALESCE(tr.description,quest.description) as description,"+
		"COALESCE(tr.meta_data,quest.meta_data) as meta_data,COALESCE(tr.quest_data,quest.quest_data) as quest_data").
		Joins("LEFT JOIN quest_translated tr ON quest.token_id = tr.token_id AND tr.language = ?", language).
		Where("uuid", uuid).First(&quest.Quest).Error
	if err != nil {
		return quest, err
	}
	// 获取所有答案
	err = d.db.Raw(`SELECT answer AS answers
		FROM (
		SELECT  quest_data->>'answers' AS answer FROM quest WHERE token_id = ?
		UNION
		SELECT answer FROM quest_translated WHERE token_id = ? AND answer IS NOT NULL) AS combined_data
		`, quest.Quest.TokenId, quest.Quest.TokenId).Scan(&quest.Answers).Error
	return
}

func (d *Dao) GetQuestWithClaimStatusByTokenIDWithLang(language string, tokenID string, address string) (quest response.GetQuestRes, err error) {
	err = d.db.Model(&model.Quest{}).
		Select("quest.*,COALESCE(tr.title,quest.title) as title,COALESCE(tr.description,quest.description) as description,"+
			"COALESCE(tr.meta_data,quest.meta_data) as meta_data,COALESCE(tr.quest_data,quest.quest_data) as quest_data,"+
			"b.claimed,b.user_score,b.nft_address,b.badge_token_id,b.chain_id as badge_chain_id,COALESCE(o.open_quest_review_status,0) as open_quest_review_status,COALESCE(o.answer,l.answer) as answer").
		Joins("left join user_challenges b ON quest.token_id=b.token_id AND b.address= ?", address).
		Joins("left join user_challenge_log l ON quest.token_id=l.token_id AND l.address= ? AND l.deleted_at IS NULL", address).
		Joins("left join user_open_quest o ON quest.token_id=o.token_id AND o.address= ? AND o.deleted_at IS NULL", address).
		Joins("LEFT JOIN quest_translated tr ON quest.token_id = tr.token_id AND tr.language = ?", language).
		Where("quest.token_id", tokenID).
		Order("l.add_ts desc,o.id desc").
		First(&quest).Error
	if err != nil {
		return quest, err
	}
	// 获取所有答案
	err = d.db.Raw(`SELECT answer AS answers
		FROM (
		SELECT  quest_data->>'answers' AS answer FROM quest WHERE token_id = ?
		UNION
		SELECT answer FROM quest_translated WHERE token_id = ? AND answer IS NOT NULL) AS combined_data
		`, tokenID, tokenID).Scan(&quest.Answers).Error
	return
}

func (d *Dao) GetQuestWithClaimStatusByTokenID(tokenID string, address string) (quest response.GetQuestRes, err error) {
	err = d.db.Model(&model.Quest{}).
		Select("quest.*,b.claimed,b.user_score,b.nft_address,b.badge_token_id,b.chain_id as badge_chain_id,COALESCE(o.open_quest_review_status,0) as open_quest_review_status,COALESCE(o.answer,l.answer) as answer").
		Joins("left join user_challenges b ON quest.token_id=b.token_id AND b.address= ?", address).
		Joins("left join user_challenge_log l ON quest.token_id=l.token_id AND l.address= ? AND l.deleted_at IS NULL", address).
		Joins("left join user_open_quest o ON quest.token_id=o.token_id AND o.address= ? AND o.deleted_at IS NULL", address).
		Where("quest.token_id", tokenID).
		Order("l.add_ts desc,o.id desc").
		First(&quest).Error
	if err != nil {
		return quest, err
	}
	// 获取所有答案
	err = d.db.Raw(`SELECT answer AS answers
		FROM (
		SELECT  quest_data->>'answers' AS answer FROM quest WHERE token_id = ?
		UNION
		SELECT answer FROM quest_translated WHERE token_id = ? AND answer IS NOT NULL) AS combined_data
		`, tokenID, tokenID).Scan(&quest.Answers).Error
	return
}

func (d *Dao) GetQuestWithClaimStatusByUUID(language, uuid string, address string) (quest response.GetQuestRes, err error) {
	err = d.db.Model(&model.Quest{}).
		Select("quest.*,COALESCE(tr.title,quest.title) as title,COALESCE(tr.description,quest.description) as description,"+
			"COALESCE(tr.meta_data,quest.meta_data) as meta_data,COALESCE(tr.quest_data,quest.quest_data) as quest_data,"+
			"b.claimed,b.user_score,b.nft_address,b.badge_token_id,b.chain_id as badge_chain_id,COALESCE(o.open_quest_review_status,0) as open_quest_review_status,COALESCE(o.answer,l.answer) as answer").
		Joins("left join user_challenges b ON quest.token_id=b.token_id AND b.address= ?", address).
		Joins("left join user_challenge_log l ON quest.token_id=l.token_id AND l.address= ? AND l.deleted_at IS NULL", address).
		Joins("left join user_open_quest o ON quest.token_id=o.token_id AND o.address= ? AND o.deleted_at IS NULL", address).
		Joins("LEFT JOIN quest_translated tr ON quest.token_id = tr.token_id AND tr.language = ?", language).
		Where("quest.uuid", uuid).
		Order("l.add_ts desc,o.id desc").
		First(&quest).Error
	if err != nil {
		return quest, err
	}
	// 获取所有答案
	err = d.db.Raw(`SELECT answer AS answers
		FROM (
		SELECT  quest_data->>'answers' AS answer FROM quest WHERE token_id = ?
		UNION
		SELECT answer FROM quest_translated WHERE token_id = ? AND answer IS NOT NULL) AS combined_data
		`, quest.TokenId, quest.TokenId).Scan(&quest.Answers).Error
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
	if req.Creator == req.Address {
		db.Select("quest.*,EXISTS (SELECT 1 FROM user_challenges WHERE quest.token_id = user_challenges.token_id) AS has_claim")
	} else {
		db.Select("quest.*,COALESCE(tr.title,quest.title) as title,COALESCE(tr.description,quest.description) as description,EXISTS (SELECT 1 FROM user_challenges WHERE quest.token_id = user_challenges.token_id) AS has_claim").
			Joins("LEFT JOIN quest_translated tr ON quest.token_id = tr.token_id AND tr.language = ?", req.Language)
	}
	db.Where(&req.Quest)
	db.Where("quest.status = 1")
	err = db.Count(&total).Error
	if err != nil {
		return questList, total, err
	}
	err = db.Limit(limit).Offset(offset).Order("add_ts desc").Find(&questList).Error

	return questList, total, err
}

func (d *Dao) GetQuestChallengeUserByTokenID(tokenId string) (res response.GetQuestChallengeUserRes, err error) {
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

// GetQuestFlashRankByTokenID 获取闪电榜
func (d *Dao) GetQuestFlashRankByTokenID(address string, tokenId string) (res response.GetQuestFlashListRes, err error) {
	// 查询挑战
	var quest model.Quest
	err = d.db.Model(&model.Quest{}).Where("token_id", tokenId).First(&quest).Error
	if err != nil {
		return res, err
	}
	// 判断是否是开放题
	if !IsOpenQuest(gjson.Get(string(quest.QuestData), "questions").String()) {
		rankListSQL := `
		WITH ranked AS (
		 SELECT address,token_id, created_at,ROW_NUMBER() OVER (PARTITION BY address ORDER BY created_at ASC) as rn
	     FROM user_challenge_log
	     WHERE token_id = ? AND address !='' AND pass=true AND deleted_at IS NULL
		)
		SELECT ROW_NUMBER() OVER (ORDER BY created_at ASC) as rank,ranked.address,users.avatar,ranked.created_at as finish_time
		FROM ranked
		LEFT JOIN users ON ranked.address=users.address
		WHERE rn=1 ORDER BY created_at ASC LIMIT 10;
		`
		err = d.db.Raw(rankListSQL, tokenId).Scan(&res.RankList).Error
		if err != nil {
			return res, err
		}
		// 地址为空返回结果
		if address == "" {
			return res, err
		}
		userRankSQL := `
		WITH ranked AS (
		 SELECT address,token_id, created_at,ROW_NUMBER() OVER (PARTITION BY address ORDER BY created_at ASC) as rn 
		 FROM user_challenge_log 
		 WHERE token_id = ? AND address !='' AND pass=true AND deleted_at IS NULL
		),
		ranked_with_rank AS (
		 SELECT ROW_NUMBER() OVER (ORDER BY created_at ASC) as rank,address,created_at as finish_time 
		 FROM ranked 
		 WHERE rn=1 
		)
		SELECT ranked_with_rank.*,users.avatar
		FROM ranked_with_rank
		LEFT JOIN users ON ranked_with_rank.address=users.address
		WHERE ranked_with_rank.address = ?
		LIMIT 1;
		`
		err = d.db.Raw(userRankSQL, tokenId, address).Scan(&res).Error
		if err != nil {
			return res, err
		}
		// 直接返回
		return res, err
	}
	// 开放题
	rankListSQL := `
		WITH all_open_quest AS(
			SELECT address,created_at
			FROM user_open_quest
			WHERE token_id = ? AND pass=true
			UNION
			SELECT address,created_at
			FROM user_challenge_log
			WHERE token_id = ? AND pass=true AND is_open_quest=false
		),ranked_open_quest AS (
		 SELECT address,created_at,ROW_NUMBER() OVER (PARTITION BY address ORDER BY created_at ASC) as rn 
		 FROM all_open_quest
		 ),
		ranked_with_rank AS (
		 SELECT ROW_NUMBER() OVER (ORDER BY created_at ASC) as rank,address,created_at as finish_time 
		 FROM ranked_open_quest 
		 WHERE rn=1 
		)
		SELECT ranked_with_rank.*,users.avatar
		FROM ranked_with_rank
		LEFT JOIN users ON ranked_with_rank.address=users.address
		ORDER BY rank ASC 
		LIMIT 10;
	`
	err = d.db.Raw(rankListSQL, tokenId, tokenId).Scan(&res.RankList).Error
	if err != nil {
		return res, err
	}
	// 地址为空返回结果
	if address == "" {
		return res, err
	}
	userRankSQL := `
		WITH all_open_quest AS(
			SELECT address,created_at
			FROM user_open_quest
			WHERE token_id = ? AND pass=true
			UNION
			SELECT address,created_at
			FROM user_challenge_log
			WHERE token_id = ? AND pass=true AND is_open_quest=false
		),ranked_open_quest AS (
		 SELECT address,created_at,ROW_NUMBER() OVER (PARTITION BY address ORDER BY created_at DESC) as rn 
		 FROM all_open_quest
		 ),
		ranked_with_rank AS (
		 SELECT ROW_NUMBER() OVER (ORDER BY created_at ASC) as rank,address,created_at as finish_time 
		 FROM ranked_open_quest 
		 WHERE rn=1 
		)
		SELECT ranked_with_rank.*,users.avatar
		FROM ranked_with_rank
		LEFT JOIN users ON ranked_with_rank.address=users.address
		WHERE ranked_with_rank.address = ?
		LIMIT 1;
	`
	err = d.db.Raw(userRankSQL, tokenId, tokenId, address).Scan(&res).Error
	return res, err
}

// GetQuestFlashRankByUUID 获取闪电榜
func (d *Dao) GetQuestFlashRankByUUID(address string, uuid string) (res response.GetQuestFlashListRes, err error) {
	// 获取token_id
	var quest model.Quest
	err = d.db.Model(&model.Quest{}).Select("token_id").Where("uuid", uuid).First(&quest).Error
	if err != nil {
		return res, err
	}
	return d.GetQuestFlashRankByTokenID(address, quest.TokenId)
}

// IsOpenQuest 判断是否开放题
func IsOpenQuest(answerUser string) bool {
	answerU := gjson.Get(answerUser, "@this").Array()
	for _, v := range answerU {
		if v.Get("type").String() == "open_quest" {
			return true
		}
	}
	return false
}

// GetQuestHighRankByTokenID 获取高分榜
func (d *Dao) GetQuestHighRankByTokenID(address string, tokenId string) (res response.GetQuestHighScoreListRes, err error) {
	// 查询挑战
	var quest model.Quest
	err = d.db.Model(&model.Quest{}).Where("token_id", tokenId).First(&quest).Error
	if err != nil {
		return res, err
	}
	// 判断是否是开放题
	if !IsOpenQuest(gjson.Get(string(quest.QuestData), "questions").String()) {
		rankListSQL := `
			WITH ranked AS (
			 SELECT address,token_id,created_at, user_score,ROW_NUMBER() OVER (PARTITION BY address ORDER BY user_score DESC,created_at ASC) as rn
			 FROM user_challenge_log
			 WHERE token_id = ? AND address !='' AND pass=true AND deleted_at IS NULL
			)
			SELECT ROW_NUMBER() OVER (ORDER BY user_score DESC,created_at ASC) as rank,ranked.address,users.avatar,created_at as finish_time,ranked.user_score/100 as score
			FROM ranked
			LEFT JOIN users ON ranked.address=users.address
			WHERE rn=1 ORDER BY user_score DESC,created_at ASC LIMIT 10;
		`
		err = d.db.Raw(rankListSQL, tokenId).Scan(&res.RankList).Error
		if err != nil {
			return res, err
		}
		// 地址为空返回结果
		if address == "" {
			return res, err
		}
		userRankSQL := `
			WITH ranked AS (
			 SELECT address,token_id,created_at,user_score,ROW_NUMBER() OVER (PARTITION BY address ORDER BY user_score DESC,created_at ASC) as rn
			 FROM user_challenge_log
			 WHERE token_id = ? AND address !='' AND pass=true AND deleted_at IS NULL
			),ranked_with_rank AS (
			SELECT ROW_NUMBER() OVER (ORDER BY user_score DESC,created_at ASC) as rank,ranked.address,ranked.user_score/100 as score,created_at as finish_time
			FROM ranked WHERE rn=1
			)
			SELECT rank,ranked_with_rank.address,users.avatar,score,finish_time FROM ranked_with_rank
			LEFT JOIN users ON ranked_with_rank.address=users.address
			WHERE ranked_with_rank.address = ?
			LIMIT 1;
		`
		err = d.db.Raw(userRankSQL, tokenId, address).Scan(&res).Error
		if err != nil {
			return res, err
		}
		// 直接返回
		return res, err
	}
	// 开放题
	rankListSQL := `
		WITH all_open_quest AS(
			SELECT address,open_quest_score as score,created_at
			FROM user_open_quest
			WHERE token_id = ? AND pass=true
			UNION
			SELECT address,user_score/100 as score,created_at
			FROM user_challenge_log
			WHERE token_id = ? AND pass=true AND is_open_quest=false
		),ranked_open_quest AS (
		 SELECT address,score,created_at,ROW_NUMBER() OVER (PARTITION BY address ORDER BY score DESC,created_at ASC) as rn 
		 FROM all_open_quest
		 ),
		ranked_with_rank AS (
		 SELECT ROW_NUMBER() OVER (ORDER BY score DESC,created_at ASC) as rank,address,score,created_at as finish_time 
		 FROM ranked_open_quest 
		 WHERE rn=1
		)
		SELECT ranked_with_rank.*,users.avatar
		FROM ranked_with_rank
		LEFT JOIN users ON ranked_with_rank.address=users.address
		ORDER BY score DESC,ranked_with_rank.finish_time ASC
		LIMIT 10;
	`
	err = d.db.Raw(rankListSQL, tokenId, tokenId).Scan(&res.RankList).Error
	if err != nil {
		return res, err
	}
	// 地址为空返回结果
	if address == "" {
		return res, err
	}
	userRankSQL := `
		WITH all_open_quest AS(
			SELECT address,open_quest_score as score,created_at
			FROM user_open_quest
			WHERE token_id = ? AND pass=true
			UNION
			SELECT address,user_score/100 as score,created_at
			FROM user_challenge_log
			WHERE token_id = ? AND pass=true AND is_open_quest=false
		),ranked_open_quest AS (
		 SELECT address,score,created_at,ROW_NUMBER() OVER (PARTITION BY address ORDER BY score DESC,created_at ASC) as rn 
		 FROM all_open_quest
		 ),
		ranked_with_rank AS (
		 SELECT ROW_NUMBER() OVER (ORDER BY score DESC,created_at ASC) as rank,address,score,created_at as finish_time 
		 FROM ranked_open_quest 
		 WHERE rn=1
		)
		SELECT ranked_with_rank.* ,users.avatar
		FROM ranked_with_rank
		LEFT JOIN users ON ranked_with_rank.address=users.address
		WHERE ranked_with_rank.address = ?
		LIMIT 1;
	`
	err = d.db.Raw(userRankSQL, tokenId, tokenId, address).Scan(&res).Error
	return res, err
}

// GetQuestHighRankByUUID 获取高分榜
func (d *Dao) GetQuestHighRankByUUID(address string, uuid string) (res response.GetQuestHighScoreListRes, err error) {
	// 获取token_id
	var quest model.Quest
	err = d.db.Model(&model.Quest{}).Select("token_id").Where("uuid", uuid).First(&quest).Error
	if err != nil {
		return res, err
	}
	return d.GetQuestHighRankByTokenID(address, quest.TokenId)
}

// GetQuestHolderRankByTokenID 获取持有榜
func (d *Dao) GetQuestHolderRankByTokenID(address string, tokenId string, page, pageSize int) (res []response.GetQuestHolderListRes, total int64, err error) {
	// 分页参数
	limit := pageSize
	offset := pageSize * (page - 1)
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

// GetQuestHolderRankByUUID 获取持有榜
func (d *Dao) GetQuestHolderRankByUUID(address string, uuid string, page, pageSize int) (res []response.GetQuestHolderListRes, total int64, err error) {
	// 获取token_id
	var quest model.Quest
	err = d.db.Model(&model.Quest{}).Select("token_id").Where("uuid", uuid).First(&quest).Error
	if err != nil {
		return res, total, err
	}
	return d.GetQuestHolderRankByTokenID(address, quest.TokenId, page, pageSize)
}

// GetQuestAnswersByTokenId 获取挑战多语言答案
func (d *Dao) GetQuestAnswersByTokenId(tokenId string) (answers []string, err error) {
	err = d.db.Raw(`SELECT answer AS answers
		FROM (
		SELECT  quest_data->>'answers' AS answer FROM quest WHERE token_id = ?
		UNION
		SELECT answer FROM quest_translated WHERE token_id = ? AND answer IS NOT NULL) AS combined_data
		`, tokenId, tokenId).Scan(&answers).Error
	return
}

// GetAddressHighScore 获取地址最高分
func (d *Dao) GetAddressHighScore(address string) (res []response.GetAddressHighScore, err error) {
	userHighScoreSQL := `
		WITH all_quest AS(
				SELECT address,token_id,open_quest_score as score,answer
				FROM user_open_quest
				WHERE  address = ? AND pass=true
				UNION
				SELECT address,token_id,user_score/100 as score,answer
				FROM user_challenge_log
				WHERE address = ? AND pass=true AND is_open_quest=false
			),ranked_quest AS (
		SELECT address,token_id,score,answer,ROW_NUMBER() OVER (PARTITION BY address,token_id ORDER BY token_id DESC,score DESC) as rn 
		 FROM all_quest
			)
		 SELECT address,token_id,score,answer
		 FROM ranked_quest 
		 WHERE rn=1
	`
	err = d.db.Raw(userHighScoreSQL, address, address).Scan(&res).Error
	return
}
