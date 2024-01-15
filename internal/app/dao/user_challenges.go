package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/clause"
)

func (d *Dao) CreateChallenges(challenges *model.UserChallenges) (err error) {
	err = d.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "token_id"}},
		UpdateAll: true,
	}).Create(&challenges).Error
	return
}

func (d *Dao) CreateChallengesList(tokenId string, receivers []common.Address) (err error) {
	var challenge []model.UserChallenges
	for _, v := range receivers {
		challenge = append(challenge, model.UserChallenges{
			Address: v.String(),
			TokenId: tokenId,
			Claimed: true,
			Status:  2,
		})
	}
	err = d.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "token_id"}},
		UpdateAll: true,
	}).Create(&challenge).Error
	return
}

func (d *Dao) CreateChallengesOne(tokenId string, receiver string, uerScore int64, nftAddress string, badgeTokenID string, chainID int64) (err error) {
	challenge := model.UserChallenges{
		Address:      receiver,
		TokenId:      tokenId,
		Claimed:      true,
		Status:       2,
		UserScore:    uerScore,
		NFTAddress:   nftAddress,
		BadgeTokenID: badgeTokenID,
		ChainID:      chainID,
	}
	err = d.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "token_id"}},
		UpdateAll: true,
	}).Create(&challenge).Error
	return
}

// GetOwnerChallengeList 挑战列表包含可领取
func (d *Dao) GetOwnerChallengeList(req *request.GetChallengeListRequest) (res []response.GetChallengeListRes, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	countSQL := `
		WITH ranked_logs AS (
		  SELECT *,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY id DESC) as rn
		  FROM user_challenge_log
		  WHERE address = ? AND deleted_at IS NULL
		),ranked_logs2 AS (SELECT *,ROW_NUMBER () OVER (PARTITION BY token_id ORDER BY id DESC) AS rn 
		FROM user_open_quest
		WHERE address= ? AND deleted_at IS NULL)
		SELECT COUNT(1) FROM quest A
		LEFT JOIN ranked_logs b ON a.token_id = b.token_id 
		LEFT JOIN ranked_logs2 c ON a.token_id = c.token_id
		LEFT JOIN user_challenges d ON a.token_id=d.token_id AND d.address = b.address
		WHERE b.rn = 1 AND (C.rn IS NULL or C.rn = 1 ) AND ((d.claimed=true OR b.pass=true) OR b.is_open_quest=true) AND (c.pass IS NULL OR c.pass=true OR (c.pass=false AND c.open_quest_review_status=1 AND A.status = 1))
	`
	// 查询记录条数
	if err = d.db.Raw(countSQL, req.Address, req.Address).Scan(&total).Error; err != nil {
		fmt.Println(err)
		return res, total, err
	}
	dataSQL := `
		WITH ranked_logs AS (
		  SELECT *,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY id DESC) as rn
		  FROM user_challenge_log
		  WHERE address = ? AND deleted_at IS NULL
		),ranked_logs2 AS (SELECT *,ROW_NUMBER () OVER (PARTITION BY token_id ORDER BY id DESC) AS rn 
		FROM user_open_quest
		WHERE address= ? AND deleted_at IS NULL)
		SELECT a.*,COALESCE(tr.title,a.title) as title,COALESCE(tr.description,a.description) as description,COALESCE(c.pass,b.pass) as claimable,COALESCE(c.open_quest_review_status,0) as open_quest_review_status,b.is_open_quest,COALESCE(d.nft_address,'') as nft_address,NOT (d.claimed = false AND zc.quest_id IS NULL) as claimed,d.badge_token_id,d.chain_id as badge_chain_id FROM quest A
		LEFT JOIN ranked_logs b ON a.token_id = b.token_id 
		LEFT JOIN ranked_logs2 c ON a.token_id = c.token_id
		LEFT JOIN user_challenges d ON a.token_id=d.token_id AND d.address = b.address
		LEFT JOIN (WITH zcloak_status AS (SELECT quest_id,ROW_NUMBER() OVER (PARTITION BY quest_id ORDER BY id DESC) as rn FROM zcloak_card WHERE address=? AND deleted_at IS NULL) SELECT quest_id FROM zcloak_status WHERE rn = 1) zc ON a.id = zc.quest_id
		LEFT JOIN quest_translated tr ON a.token_id = tr.token_id AND tr.language = ?
		WHERE b.rn = 1 AND (C.rn IS NULL or C.rn = 1 ) AND ((d.claimed=true OR b.pass=true) OR b.is_open_quest=true) AND (c.pass IS NULL OR c.pass=true OR (c.pass=false AND c.open_quest_review_status=1 AND A.status = 1))
		ORDER BY b.ID DESC LIMIT ? OFFSET ?
	`
	// 查询数据
	db := d.db.Raw(dataSQL, req.Address, req.Address, req.Address, req.Language, limit, offset)
	if err = db.Scan(&res).Error; err != nil {
		return res, total, err
	}
	// 查询领取状态
	for i, v := range res {
		if v.Claimed {
			// 查询证书领取状态
			if err := d.db.Model(&model.ZcloakCard{}).Where("address = ? AND quest_id = ?", req.Address, v.ID).First(&model.ZcloakCard{}).Error; err != nil {
				res[i].ClaimStatus = 1
				continue
			}
			// 查询NFT领取状态
			if err := d.db.Model(&model.UserChallenges{}).Where("address = ? AND token_id = ?", req.Address, v.TokenId).First(&model.UserChallenges{}).Error; err != nil {
				res[i].ClaimStatus = 2
				continue
			}
			res[i].ClaimStatus = 3
		}
	}
	return res, total, nil
}

// GetChallengeList 挑战列表不包含可领取
func (d *Dao) GetChallengeList(req *request.GetChallengeListRequest) (res []response.GetChallengeListRes, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	countSQL := `
		WITH all_quest AS(
					SELECT nft_address,token_id,add_ts
					FROM user_challenges
					WHERE  address = ?
					UNION
					SELECT '' as nft_address,token_id,zc.add_ts
					FROM zcloak_card zc
					LEFT JOIN quest ON quest.id = zc.quest_id
					WHERE address = ?
		),ranked_quest AS (
		 SELECT nft_address,token_id,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY add_ts ASC) as rn 
				 FROM all_quest
		)
		SELECT count(1) FROM ranked_quest a
		JOIN quest b ON a.token_id=b.token_id 
		LEFT JOIN user_challenges d ON a.token_id=d.token_id AND d.address = ?
		LEFT JOIN quest_translated tr ON b.token_id = tr.token_id AND tr.language = ? 
		WHERE rn=1
	`
	// 查询记录条数
	if err = d.db.Raw(countSQL, req.Address, req.Address, req.Address, req.Language).Scan(&total).Error; err != nil {
		fmt.Println(err)
		return res, total, err
	}
	dataSQL := `
		WITH all_quest AS(
					SELECT nft_address,token_id,add_ts
					FROM user_challenges
					WHERE  address = ?
					UNION
					SELECT '' as nft_address,token_id,zc.add_ts
					FROM zcloak_card zc
					LEFT JOIN quest ON quest.id = zc.quest_id
					WHERE address = ?
		),ranked_quest AS (
		 SELECT nft_address,add_ts,token_id,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY add_ts ASC) as rn 
				 FROM all_quest
		)
		SELECT a.nft_address,true as claimed,a.add_ts as complete_ts,b.*,COALESCE(tr.title,b.title) as title,COALESCE(tr.description,b.description) as description
		FROM ranked_quest a
		JOIN quest b ON a.token_id=b.token_id 
		LEFT JOIN user_challenges d ON a.token_id=d.token_id AND d.address = ?
		LEFT JOIN quest_translated tr ON b.token_id = tr.token_id AND tr.language = ? 
		WHERE rn=1
		ORDER BY a.add_ts DESC LIMIT ? OFFSET ?
	`
	// 查询数据
	db := d.db.Raw(dataSQL, req.Address, req.Address, req.Address, req.Language, limit, offset)
	if err = db.Scan(&res).Error; err != nil {
		return res, total, err
	}
	// 查询领取状态
	for i, v := range res {
		// 查询证书领取状态
		if err := d.db.Model(&model.ZcloakCard{}).Where("address = ? AND quest_id = ?", req.Address, v.ID).First(&model.ZcloakCard{}).Error; err != nil {
			res[i].ClaimStatus = 1
			continue
		}
		// 查询NFT领取状态
		if err := d.db.Model(&model.UserChallenges{}).Where("address = ? AND token_id = ?", req.Address, v.TokenId).First(&model.UserChallenges{}).Error; err != nil {
			res[i].ClaimStatus = 2
			continue
		}
		res[i].ClaimStatus = 3
	}
	return res, total, nil
}

// GetChallengeNotClaimList 未领取挑战列表
func (d *Dao) GetChallengeNotClaimList(req *request.GetChallengeListRequest) (res []response.GetChallengeListRes, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	countSQL := `
		WITH ranked_logs AS (
		  SELECT *,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY id DESC) as rn
		  FROM user_challenge_log
		  WHERE address = ?
		),ranked_logs2 AS (
		SELECT *,ROW_NUMBER () OVER (PARTITION BY token_id ORDER BY id DESC) AS rn 
		FROM user_open_quest
		WHERE address= ?)
		SELECT COUNT(1) FROM quest A
		LEFT JOIN ranked_logs b ON a.token_id = b.token_id
		LEFT JOIN ranked_logs2 c ON a.token_id = c.token_id
		LEFT JOIN user_challenges d ON a.token_id=d.token_id AND d.address = b.address
		LEFT JOIN zcloak_card zc ON a.id = zc.quest_id AND d.address = zc.address
		WHERE zc.id IS NULL AND b.rn = 1 AND (C.rn IS NULL or C.rn = 1) AND (b.pass=true OR b.is_open_quest=true) AND (c.pass IS NULL OR c.pass=true) AND d.claimed is null
	`
	// 查询记录条数
	if err = d.db.Raw(countSQL, req.Address, req.Address).Scan(&total).Error; err != nil {
		return res, total, err
	}
	dataSQL := `
		WITH ranked_logs AS (
		  SELECT *,ROW_NUMBER() OVER (PARTITION BY token_id ORDER BY id DESC) as rn
		  FROM user_challenge_log
		  WHERE address = ?
		),ranked_logs2 AS (
		SELECT *,ROW_NUMBER () OVER (PARTITION BY token_id ORDER BY id DESC) AS rn 
		FROM user_open_quest
		WHERE address= ?)
		SELECT a.*,COALESCE(tr.title,a.title) as title,COALESCE(tr.description,a.description) as description,COALESCE(c.pass,b.pass) as claimable,COALESCE(c.open_quest_review_status,0) as open_quest_review_status,b.is_open_quest,COALESCE(d.nft_address,'') as nft_address,COALESCE(d.claimed,false) as claimed FROM quest A
		LEFT JOIN ranked_logs b ON a.token_id = b.token_id
		LEFT JOIN ranked_logs2 c ON a.token_id = c.token_id
		LEFT JOIN user_challenges d ON a.token_id=d.token_id AND d.address = b.address
		LEFT JOIN quest_translated tr ON a.token_id = tr.token_id AND tr.language = ?
		LEFT JOIN zcloak_card zc ON a.id = zc.quest_id AND b.address = zc.address
		WHERE zc.id IS NULL AND b.rn = 1 AND (C.rn IS NULL or C.rn = 1 ) AND (b.pass=true OR b.is_open_quest=true) AND (c.pass IS NULL OR c.pass=true) AND d.claimed is null
		ORDER BY b.ID DESC LIMIT ? OFFSET ?
	`
	// 查询数据
	if err = d.db.Raw(dataSQL, req.Address, req.Address, req.Language, limit, offset).Scan(&res).Error; err != nil {
		return res, total, err
	}
	return res, total, nil
}

// GetChallengeWaitReviewList 待审核挑战列表
func (d *Dao) GetChallengeWaitReviewList(req *request.GetChallengeListRequest) (res []response.GetChallengeListRes, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := d.db

	err = db.Raw("SELECT COUNT(1) FROM user_open_quest a JOIN quest b ON a.token_id=b.token_id WHERE address = ? AND open_quest_review_status=1 AND a.deleted_at IS NULL", req.Address).Scan(&total).Error
	if err != nil {
		return res, total, err
	}
	err = db.Raw("SELECT a.open_quest_review_status,b.*,COALESCE(tr.title,b.title) as title,COALESCE(tr.description,b.description) as description FROM user_open_quest a JOIN quest b ON a.token_id=b.token_id LEFT JOIN quest_translated tr ON b.token_id = tr.token_id AND tr.language = ? WHERE address = ? AND open_quest_review_status=1 AND a.deleted_at IS NULL ORDER BY a.id DESC LIMIT ? OFFSET ?",
		req.Language, req.Address, limit, offset).Scan(&res).Error
	if err != nil {
		return res, total, err
	}
	return res, total, nil
}
