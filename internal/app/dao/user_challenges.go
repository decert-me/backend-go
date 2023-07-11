package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"math/big"
	"sync"
)

var lock sync.Mutex

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
	lock.Lock()
	defer lock.Unlock()
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := d.db
	// 临时数据
	var claimable []request.Claimable
	json.Unmarshal([]byte(req.Claimable), &claimable)
	randomBytes := make([]byte, 8)
	_, err = rand.Read(randomBytes)
	if err != nil {
		log.Errorv("rand.Read error", zap.Error(err))
		return
	}
	randomInt, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		log.Errorv("rand.Int error", zap.Error(err))
		return
	}
	tempDB := "temp_table_" + randomInt.String()
	if err = db.Exec(fmt.Sprintf("CREATE TEMPORARY TABLE %s (token_id int8, add_ts int8)", tempDB)).Error; err != nil {
		return res, total, err
	}
	defer func() {
		db.Exec(fmt.Sprintf("DROP TABLE %s", tempDB))
	}()

	// 跳过已存在的
	var tokenList []int64
	for _, claim := range claimable {
		tokenList = append(tokenList, claim.TokenId)
	}
	var existTokenList []int64
	db.Raw("SELECT token_id FROM claim_badge_tweet WHERE token_id in ? AND address= ? AND status=0 UNION "+
		"SELECT token_id FROM user_challenges WHERE token_id in ? AND address = ?", tokenList, req.Address, tokenList, req.Address).Scan(&existTokenList)
	for _, v := range claimable {
		if utils.SliceIsExist(existTokenList, v.TokenId) {
			continue
		}
		if err = db.Exec(fmt.Sprintf("INSERT INTO %s (token_id,add_ts) VALUES (?, ?)", tempDB), v.TokenId, v.AddTs).Error; err != nil {
			return res, total, err
		}
	}

	err = db.Raw("SELECT count(1) FROM (SELECT a.claimed,a.add_ts as complete_ts,b.* FROM user_challenges a JOIN quest b ON a.token_id=b.token_id WHERE address = ? "+
		" UNION "+
		"SELECT 'f' as claimed,a.add_ts as complete_ts,b.* as claimed FROM claim_badge_tweet a JOIN quest b ON a.token_id=b.token_id WHERE a.address = ? AND a.status=0"+
		" UNION "+
		fmt.Sprintf("SELECT 'f' as claimed,a.add_ts as complete_ts,b.* as claimed FROM %s a JOIN quest b ON a.token_id=b.token_id", tempDB)+
		") a1", req.Address, req.Address).Scan(&total).Error
	if err != nil {
		return res, total, err
	}
	err = db.Raw("SELECT * FROM ((SELECT a.claimed,a.add_ts as complete_ts,b.* FROM user_challenges a JOIN quest b ON a.token_id=b.token_id  WHERE address = ? ORDER BY a.add_ts DESC"+
		") UNION ("+
		"SELECT 'f' as claimed,a.add_ts as complete_ts,b.* as claimed FROM claim_badge_tweet a JOIN quest b ON a.token_id=b.token_id WHERE a.address = ? AND a.status=0 ORDER BY a.add_ts DESC"+
		") UNION ("+
		fmt.Sprintf("SELECT 'f' as claimed,a.add_ts as complete_ts,b.* as claimed FROM %s a JOIN quest b ON a.token_id=b.token_id", tempDB)+
		")) a1 ORDER BY complete_ts DESC LIMIT ? OFFSET ? ",
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

	err = db.Raw("SELECT COUNT(1) FROM user_challenges a JOIN quest b ON a.token_id=b.token_id WHERE address = ?", req.Address).Scan(&total).Error
	if err != nil {
		return res, total, err
	}
	err = db.Raw("SELECT a.claimed,a.add_ts as complete_ts,b.* FROM user_challenges a JOIN quest b ON a.token_id=b.token_id WHERE address = ? ORDER BY a.add_ts DESC LIMIT ? OFFSET ?",
		req.Address, limit, offset).Scan(&res).Error
	if err != nil {
		return res, total, err
	}
	return res, total, nil
}

// GetChallengeNotClaimList 未领取挑战列表
func (d *Dao) GetChallengeNotClaimList(req *request.GetChallengeListRequest) (res []response.GetChallengeListRes, total int64, err error) {
	lock.Lock()
	defer lock.Unlock()
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := d.db
	// 临时数据
	var claimable []request.Claimable
	json.Unmarshal([]byte(req.Claimable), &claimable)
	// 跳过已存在的
	var tokenList []int64
	for _, claim := range claimable {
		tokenList = append(tokenList, claim.TokenId)
	}
	var existTokenList []int64
	db.Raw("SELECT token_id FROM claim_badge_tweet WHERE token_id in ? AND address= ? AND status=0 UNION "+
		"SELECT token_id FROM user_challenges WHERE token_id in ? AND address = ?", tokenList, req.Address, tokenList, req.Address).Scan(&existTokenList)

	randomBytes := make([]byte, 8)
	_, err = rand.Read(randomBytes)
	if err != nil {
		log.Errorv("rand.Read error", zap.Error(err))
		return
	}
	randomInt, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		log.Errorv("rand.Int error", zap.Error(err))
		return
	}
	tempDB := "temp_table_" + randomInt.String()
	if err = db.Exec(fmt.Sprintf("CREATE TEMPORARY TABLE %s (token_id int8, add_ts int8)", tempDB)).Error; err != nil {
		return res, total, err
	}
	defer func() {
		db.Exec(fmt.Sprintf("DROP TABLE %s", tempDB))
	}()
	for _, v := range claimable {
		if utils.SliceIsExist(existTokenList, v.TokenId) {
			continue
		}
		if err = db.Exec(fmt.Sprintf("INSERT INTO %s (token_id,add_ts) VALUES (?, ?)", tempDB), v.TokenId, v.AddTs).Error; err != nil {
			return res, total, err
		}
	}
	err = db.Raw("SELECT COUNT(1) FROM claim_badge_tweet a LEFT JOIN quest b ON a.token_id=b.token_id  WHERE address = ? AND a.status=0", req.Address).Scan(&total).Error
	if err != nil {
		return res, total, err
	}
	total = total + int64(len(claimable))
	err = db.Raw("SELECT * FROM ((SELECT 'f' as claimed,a.add_ts as complete_ts,b.* FROM claim_badge_tweet a LEFT JOIN quest b ON a.token_id=b.token_id WHERE a.address = ? AND a.status=0 "+
		") UNION ("+
		fmt.Sprintf("SELECT 'f' as claimed,a.add_ts as complete_ts,b.* as claimed FROM %s a JOIN quest b ON a.token_id=b.token_id", tempDB)+
		")) a1 ORDER BY complete_ts DESC LIMIT ? OFFSET ?",
		req.Address, limit, offset).Scan(&res).Error
	if err != nil {
		return res, total, err
	}
	return res, total, nil
}
