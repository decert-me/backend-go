package initialize

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/model"
	"gorm.io/gorm"
)

// NewPgSQL new pgsql db
func NewPgSQL(c *config.Pgsql) *gorm.DB {
	db := GormPgSql(c)
	if c.AutoMigrate {
		RegisterTables(db) // 初始化表
		//initMultiChainV2(db)
	}
	return db
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Users{},
		model.ClaimBadgeTweet{},
		model.Quest{},
		model.Collection{},
		model.UserChallenges{},
		model.Transaction{},
		model.Upload{},
		model.UserChallengeLog{},
		model.ReadProgress{},
		model.UserChallengeClaim{},
		model.UserOpenQuest{},
		model.UserMessage{},
		model.OpenQuestPerm{},
		model.QuestTranslated{},
		model.CollectionTranslated{},
		model.ZcloakDid{},
		model.ZcloakCard{},
		model.CollectionRelate{},
	)
	if err != nil {
		panic("register table failed")
	}
}

func initMultiChainV2(db *gorm.DB) {
	// 查询是否存在链ID
	var chainID int64
	err := db.Model(&model.Collection{}).Select("max(chain_id)").Scan(&chainID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		panic("initMultiChainV2 failed")
	}
	if chainID != 0 {
		return
	}
	// 初始化
	err = db.Model(&model.Collection{}).Where("token_id != '0' AND uuid !='' ").Update("chain_id", 137).Error
	if err != nil {
		panic("initMultiChainV2 failed")
	}
	// 初始化
	err = db.Model(&model.Collection{}).Where("token_id = '0'").Update("token_id", "").Error
	if err != nil {
		panic("initMultiChainV2 failed")
	}
}
