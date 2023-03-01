package service

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/dao"
	"backend-go/internal/app/initialize"
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"context"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
	"time"
)

var c *config.Config
var d *dao.Dao
var s *Service

func TestMain(m *testing.M) {
	c = initialize.Viper("../../../config/config.yaml")
	// 初始化日志框架
	c.Log.Save = false
	c.Log.Level = "silent"
	c.Log.LogInConsole = false
	log.Init(c.Log)
	c.Pgsql.LogMode = "silent"
	c.Pgsql.AutoMigrate = true
	c.BlockChain.ChainID = 5
	c.BlockChain.Provider = "https://rpc.ankr.com/eth_goerli"
	// test contract address
	c.Contract = &config.Contract{
		Badge:       "0x0049770260b599Ecc2e2c0645450c965A44938b7",
		Quest:       "0xfE3e0366a52C6F668a1026dAF5e81162d34Ec38b",
		QuestMinter: "0xbE866FE4BAFC11ae886238772AFBD24570f9B530",
	}
	c.Pgsql.Prefix = "test_" // add test prefix
	c.Scheduler.Active = false
	d = dao.New(c)
	s = New(c)

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

func TestService_Ping(t *testing.T) {
	assert.Nil(t, s.Ping(context.Background()))
}

func TestService_Close(t *testing.T) {
	testService := New(c)
	testService.Close()
	db, _ := testService.dao.DB().DB()
	assert.EqualErrorf(t, db.Ping(), "sql: database is closed", "")
}

func TestNew(t *testing.T) {
	testConfig := *c
	testConfig.Scheduler.Active = true
	testConfig.Scheduler.AirdropBadge = "46 */6 * * *"
	testService := New(&testConfig)
	assert.Equal(t, 1, len(testService.cron.Entries()), "should have 1 entry")
	// Scheduler rules error
	testConfig.Scheduler.AirdropBadge = "46 * * *"
	testServiceScheduler := New(&testConfig)
	assert.Equal(t, 0, len(testServiceScheduler.cron.Entries()), "should have 0 entry")
	// Scheduler off
	testConfig.Scheduler.Active = false
	testConfig.Scheduler.AirdropBadge = "46 */6 * * *"
	testServiceNone := New(&testConfig)
	assert.Nil(t, testServiceNone.cron)
}

func deleteQuest() {
	err := s.dao.DB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Quest{}).Error
	if err != nil {
		panic(err)
	}
}

func deleteChallenges() {
	err := s.dao.DB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.UserChallenges{}).Error
	if err != nil {
		panic(err)
	}
}

func deleteBadgeTweet() {
	err := s.dao.DB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.ClaimBadgeTweet{}).Error
	if err != nil {
		panic(err)
	}
}

func waitForQuestCreated(tokenId int64) {
	var count int64
	for i := 0; i < 20; i++ {
		_ = s.dao.DB().Model(&model.Quest{}).Where("token_id = ?", tokenId).Count(&count).Error
		if count != 0 {
			return
		}
		time.Sleep(time.Second)
	}
}

func waitForClaimed(tokenId int64, address string) {
	var count int64
	for i := 0; i < 20; i++ {
		_ = s.dao.DB().Model(&model.UserChallenges{}).Where("token_id", tokenId).Where("address", address).Count(&count).Error
		if count != 0 {
			return
		}
		time.Sleep(time.Second)
	}
}
