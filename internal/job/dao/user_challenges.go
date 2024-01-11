package dao

import (
	"backend-go/internal/app/model"
	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/clause"
	"math/big"
)

func (d *Dao) CreateChallenges(challenges *model.UserChallenges) (err error) {
	err = d.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "token_id"}},
		UpdateAll: true,
	}).Create(&challenges).Error
	return
}

func (d *Dao) CreateChallengesList(tokenIds []*big.Int, receivers []common.Address, scores []*big.Int) (err error) {
	var challenge []model.UserChallenges
	for i, _ := range receivers {
		challenge = append(challenge, model.UserChallenges{
			Address:   receivers[i].String(),
			TokenId:   tokenIds[i].Int64(),
			UserScore: scores[i].Int64(),
			Claimed:   true,
			Status:    2,
		})
	}
	err = d.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}, {Name: "token_id"}},
		UpdateAll: true,
	}).Create(&challenge).Error
	return
}

func (d *Dao) DeleteUserChallengeLogByTokenId(tokenId string) (err error) {
	err = d.db.Where("token_id = ?", tokenId).Delete(&model.UserChallengeLog{}).Error
	if err != nil {
		return err
	}
	err = d.db.Where("token_id = ?", tokenId).Delete(&model.UserOpenQuest{}).Error
	if err != nil {
		return err
	}
	return
}
