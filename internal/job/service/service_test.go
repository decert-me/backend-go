package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/job/config"
	"backend-go/internal/job/dao"
	"backend-go/internal/job/initialize"
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

const QuestCreatedHash = "0x86760fe2fd761b497c40df0f2c3ebbaaef911b5d238c7c77b483b78c7a64f57c"
const ClaimHash = "0x6314296c70327955f089734d4b67de7b174798e32424a0dd3874f85c8ff82e25"
const TOKENID = 10033
const ADDRESS = "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"
const FailHash = "0x438688940d6bcaf169dae9207e97872392f0b81717250596777f8713be37b945"
const WaitHash = "0x3e47e241a2b7a5bcaecacc89c563b1eb70231222b561369c82d9f951d39b75f1"

func TestMain(m *testing.M) {
	c = initialize.Viper("../../../bin/job/config.yaml")
	// 初始化日志框架
	c.Log.Save = false
	log.Init(c.Log)
	c.Pgsql.AutoMigrate = true
	c.BlockChain.ChainID = 5
	c.BlockChain.Attempt = 5
	// test contract address
	c.Contract = &config.Contract{
		Badge:       "0x66C54CB10Ef3d038aaBA2Ac06d2c25B326be8142",
		Quest:       "0x020ef5c45182019A5aa48A8dD089a3712ad491b4",
		QuestMinter: "0xEdC46868f04d482f04A8c29E915aBED72C03cD35",
	}
	c.BlockChain.Provider = []config.Provider{
		{Url: "https://rpc.ankr.com/polygon_mumbai"},
	}
	c.Pgsql.Prefix = "test_" // add test prefix
	c.Scheduler.Active = true
	d = dao.New(c)
	s = New(c)

	result := m.Run()
	d.DB().Migrator().DropTable(
		model.Users{},
		model.ClaimBadgeTweet{},
		model.Quest{},
		model.UserChallenges{},
		model.Transaction{},
	)

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
	tx := s.dao.DB().Begin()
	err := tx.Exec("truncate test_quest").Error
	if tx.Commit().Error != nil {
		panic(err)
	}
}

func deleteUser() {
	err := s.dao.DB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Users{}).Error
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
func deleteTransaction() {
	err := s.dao.DB().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Transaction{}).Error
	if err != nil {
		panic(err)
	}
}

func waitForQuestCreated(tokenId int64) {
	var count int64
	for i := 0; i < 20; i++ {
		_ = s.dao.DB().Model(&model.Quest{}).Where("token_id", tokenId).Count(&count).Error
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
