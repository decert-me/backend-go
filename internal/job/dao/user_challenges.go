package dao

import (
	"backend-go/internal/app/model"
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
