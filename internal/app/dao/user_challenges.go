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

func (d *Dao) CreateChallengesList(tokenId int64, receivers []common.Address) (err error) {
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

func (d *Dao) CreateChallengesOne(tokenId int64, receiver string, uerScore int64, nftAddress string) (err error) {
	challenge := model.UserChallenges{
		Address:    receiver,
		TokenId:    tokenId,
		Claimed:    true,
		Status:     2,
		UserScore:  uerScore,
		NFTAddress: nftAddress,
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
		  WHERE address = ?
		),ranked_logs2 AS (SELECT *,ROW_NUMBER () OVER (PARTITION BY token_id ORDER BY id DESC) AS rn 
		FROM user_open_quest
		WHERE address= ?)
		SELECT COUNT(1) FROM quest A
		LEFT JOIN ranked_logs b ON a.token_id = b.token_id 
		LEFT JOIN ranked_logs2 c ON a.token_id = c.token_id
		LEFT JOIN user_challenges d ON a.token_id=d.token_id AND d.address = b.address
		WHERE b.rn = 1 AND (C.rn IS NULL or C.rn = 1 ) AND ((d.claimed=true OR b.pass=true) OR b.is_open_quest=true) AND (c.pass IS NULL OR c.pass=true OR c.pass=false AND c.open_quest_review_status=1)
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
		  WHERE address = ?
		),ranked_logs2 AS (SELECT *,ROW_NUMBER () OVER (PARTITION BY token_id ORDER BY id DESC) AS rn 
		FROM user_open_quest
		WHERE address= ?)
		SELECT a.*,COALESCE(c.pass,b.pass) as claimable,COALESCE(c.open_quest_review_status,0) as open_quest_review_status,b.is_open_quest,COALESCE(d.nft_address,'') as nft_address,COALESCE(d.claimed,false) as claimed FROM quest A
		LEFT JOIN ranked_logs b ON a.token_id = b.token_id 
		LEFT JOIN ranked_logs2 c ON a.token_id = c.token_id
		LEFT JOIN user_challenges d ON a.token_id=d.token_id AND d.address = b.address
		WHERE b.rn = 1 AND (C.rn IS NULL or C.rn = 1 ) AND ((d.claimed=true OR b.pass=true) OR b.is_open_quest=true) AND (c.pass IS NULL OR c.pass=true OR c.pass=false AND c.open_quest_review_status=1)
		ORDER BY b.ID DESC LIMIT ? OFFSET ?
	`
	// 查询数据
	db := d.db.Raw(dataSQL, req.Address, req.Address, limit, offset)
	if err = db.Scan(&res).Error; err != nil {
		return res, total, err
	}
	return res, total, nil
}

// GetChallengeList 挑战列表不包含可领取
func (d *Dao) GetChallengeList(req *request.GetChallengeListRequest) (res []response.GetChallengeListRes, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := d.db

	err = db.Raw("SELECT COUNT(1) FROM user_challenges a JOIN quest b ON a.token_id=b.token_id WHERE address = ?", req.Address).Scan(&total).Error
	if err != nil {
		return res, total, err
	}
	err = db.Raw("SELECT a.nft_address,a.claimed,a.add_ts as complete_ts,b.* FROM user_challenges a JOIN quest b ON a.token_id=b.token_id WHERE address = ? ORDER BY a.add_ts DESC LIMIT ? OFFSET ?",
		req.Address, limit, offset).Scan(&res).Error
	if err != nil {
		return res, total, err
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
		WHERE b.rn = 1 AND (C.rn IS NULL or C.rn = 1 ) AND (b.pass=true OR b.is_open_quest=true) AND (c.pass IS NULL OR c.pass=true) AND d.claimed is null
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
		SELECT a.*,COALESCE(c.pass,b.pass) as claimable,COALESCE(c.open_quest_review_status,0) as open_quest_review_status,b.is_open_quest,COALESCE(d.nft_address,'') as nft_address,COALESCE(d.claimed,false) as claimed FROM quest A
		LEFT JOIN ranked_logs b ON a.token_id = b.token_id
		LEFT JOIN ranked_logs2 c ON a.token_id = c.token_id
		LEFT JOIN user_challenges d ON a.token_id=d.token_id AND d.address = b.address
		WHERE b.rn = 1 AND (C.rn IS NULL or C.rn = 1 ) AND (b.pass=true OR b.is_open_quest=true) AND (c.pass IS NULL OR c.pass=true) AND d.claimed is null
		ORDER BY b.ID DESC LIMIT ? OFFSET ?
	`
	// 查询数据
	if err = d.db.Raw(dataSQL, req.Address, req.Address, limit, offset).Scan(&res).Error; err != nil {
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
	err = db.Raw("SELECT a.open_quest_review_status,b.* FROM user_open_quest a JOIN quest b ON a.token_id=b.token_id WHERE address = ? AND open_quest_review_status=1 AND a.deleted_at IS NULL ORDER BY a.id DESC LIMIT ? OFFSET ?",
		req.Address, limit, offset).Scan(&res).Error
	if err != nil {
		return res, total, err
	}
	return res, total, nil
}
