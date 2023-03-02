package blockchain

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/dao"
	"backend-go/internal/app/initialize"
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

var c *config.Config
var d *dao.Dao
var b *BlockChain

func TestMain(m *testing.M) {
	c = initialize.Viper("../../../../config/config.yaml")
	// 初始化日志框架
	c.Log.Save = false
	c.Log.Level = "silent"
	c.Log.LogInConsole = true
	log.Init(c.Log)
	//c.Pgsql.LogMode = "silent"
	c.BlockChain.ChainID = 5
	c.BlockChain.Provider = "https://rpc.ankr.com/eth_goerli"
	// test contract address
	c.Contract = &config.Contract{
		Badge:       "0x0049770260b599Ecc2e2c0645450c965A44938b7",
		Quest:       "0xfE3e0366a52C6F668a1026dAF5e81162d34Ec38b",
		QuestMinter: "0xbE866FE4BAFC11ae886238772AFBD24570f9B530",
	}
	c.Pgsql.Prefix = "test_" // add test prefix
	d = dao.New(c)
	fmt.Println(d)
	b = New(c, d)

	result := m.Run()
	//d.DB().Migrator().DropTable(
	//	model.Users{},
	//	model.ClaimBadgeTweet{},
	//	model.Quest{},
	//	model.UserChallenges{},
	//	model.Transaction{},
	//)

	os.Exit(result)
}

func deleteQuest() {
	tx := b.dao.DB().Begin()
	err := tx.Exec("truncate test_quest").Error
	if tx.Commit().Error != nil {
		panic(err)
	}
}

func deleteChallenges() {
	err := b.dao.DB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.UserChallenges{}).Error
	if err != nil {
		panic(err)
	}
}

func deleteBadgeTweet() {
	err := b.dao.DB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.ClaimBadgeTweet{}).Error
	if err != nil {
		panic(err)
	}
}

func deleteTransaction() {
	err := b.dao.DB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Transaction{}).Error
	if err != nil {
		panic(err)
	}
}

func waitForQuestCreated(tokenId int64) {
	var count int64
	for i := 0; i < 20; i++ {
		_ = b.dao.DB().Model(&model.Quest{}).Where("token_id = ?", tokenId).Count(&count).Error
		if count != 0 {
			return
		}
		time.Sleep(time.Second)
	}
}

func waitForClaimed(tokenId int64, address string) {
	var count int64
	for i := 0; i < 20; i++ {
		_ = b.dao.DB().Model(&model.UserChallenges{}).Where("token_id", tokenId).Where("address", address).Count(&count).Error
		if count != 0 {
			return
		}
		time.Sleep(time.Second)
	}
}

func TestNew(t *testing.T) {
	config := initialize.Viper("../../../../config/config.yaml")
	// 初始化日志框架
	config.Log.Save = false
	config.Log.Level = "silent"
	config.Log.LogInConsole = false
	log.Init(config.Log)
	config.Pgsql.LogMode = "silent"
	config.Pgsql.Prefix = "test_" // add test prefix
	d = dao.New(config)
	config.BlockChain.Provider = "htt://test"
	assert.Panics(t, func() { New(config, d) }, "should panic when err Provider")
}
