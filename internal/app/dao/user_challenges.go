package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
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

// GetOwnerChallengeList 挑战列表包含可领取
func (d *Dao) GetOwnerChallengeList(req *request.GetChallengeListRequest) (res []response.GetChallengeListRes, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := d.db

	err = db.Raw("SELECT count(1) FROM (SELECT a.claimed,a.add_ts as complete_ts,b.* FROM user_challenges a LEFT JOIN quest b ON a.token_id=b.token_id  WHERE address = ? "+
		" UNION "+
		"SELECT 'f' as claimed,a.add_ts as complete_ts,b.* as claimed FROM claim_badge_tweet a LEFT JOIN quest b ON a.token_id=b.token_id WHERE a.address = ? AND a.status=0) a1", req.Address, req.Address).Scan(&total).Error
	if err != nil {
		return res, total, err
	}
	err = db.Raw("SELECT * FROM ((SELECT a.claimed,a.add_ts as complete_ts,b.* FROM user_challenges a LEFT JOIN quest b ON a.token_id=b.token_id  WHERE address = ? ORDER BY a.add_ts DESC"+
		") UNION ("+
		"SELECT 'f' as claimed,a.add_ts as complete_ts,b.* as claimed FROM claim_badge_tweet a LEFT JOIN quest b ON a.token_id=b.token_id WHERE a.address = ? AND a.status=0 ORDER BY a.add_ts DESC)) a1 ORDER BY add_ts DESC LIMIT ? OFFSET ? ",
		req.Address, req.Address, limit, offset).Scan(&res).Error
	if err != nil {
		return res, total, err
	}
	return res, total, nil
}

// GetChallengeList 挑战列表不包含可领取
func (d *Dao) GetChallengeList(req *request.GetChallengeListRequest) (res []response.GetChallengeListRes, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := d.db

	err = db.Raw("SELECT COUNT(1) FROM user_challenges a LEFT JOIN quest b ON a.token_id=b.token_id  WHERE address = ?", req.Address).Scan(&total).Error
	if err != nil {
		return res, total, err
	}
	err = db.Raw("SELECT a.claimed,a.add_ts as complete_ts,b.* FROM user_challenges a LEFT JOIN quest b ON a.token_id=b.token_id  WHERE address = ? ORDER BY a.add_ts DESC LIMIT ? OFFSET ?",
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
	db := d.db

	err = db.Raw("SELECT COUNT(1) FROM claim_badge_tweet a LEFT JOIN quest b ON a.token_id=b.token_id  WHERE address = ? AND a.status=0", req.Address).Scan(&total).Error
	if err != nil {
		return res, total, err
	}
	err = db.Raw("SELECT 'f' as claimed,a.add_ts as complete_ts,b.* FROM claim_badge_tweet a LEFT JOIN quest b ON a.token_id=b.token_id WHERE a.address = ? AND a.status=0 ORDER BY a.add_ts DESC LIMIT ? OFFSET ?",
		req.Address, limit, offset).Scan(&res).Error
	if err != nil {
		return res, total, err
	}
	return res, total, nil
}
