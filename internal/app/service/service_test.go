package service

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/dao"
	"backend-go/internal/app/initialize"
	"backend-go/internal/app/model"
	jobConfigStruct "backend-go/internal/job/config"
	jobInit "backend-go/internal/job/initialize"
	jobService "backend-go/internal/job/service"
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

// init const
const QuestCreatedHash = "0x1ec6cee6b706ecfaaf1d691873dee9765c15217982d1d591b4e7e26a3fcfed2e"
const TOKENID = 10032
const SCORE = 10000
const ANSWER = "[0,[0,1],\"true\"]"
const TWEETURL = "https://twitter.com/liangjies/status/1633028821715927041?s=20"
const TWEETID = "1633028821715927041"
const ADDRESS = "0x7d32D1DE76acd73d58fc76542212e86ea63817d8"

func TestMain(m *testing.M) {
	c = initialize.Viper("../../../bin/app/config.yaml")
	// 初始化日志框架
	c.Log.Save = false
	log.Init(c.Log)
	c.Pgsql.AutoMigrate = true
	// test contract address
	c.Contract = &config.Contract{
		Badge:       "0x0049770260b599Ecc2e2c0645450c965A44938b7",
		Quest:       "0xfE3e0366a52C6F668a1026dAF5e81162d34Ec38b",
		QuestMinter: "0xbE866FE4BAFC11ae886238772AFBD24570f9B530",
	}
	c.BlockChain.SignPrivateKey = "94e0a5961679f979d86020010513d0825d1cb6905e6ae0bf31f41e7fc23dd272" // for testing
	c.Quest.EncryptKey = "eb5a5bb2-ebbd-45cc-9d37-77a9377f2aca"
	c.Pgsql.Prefix = "test_" // add test prefix
	d = dao.New(c)
	s = New(c)
	time.Sleep(time.Second)
	/*
	 * Job
	 */
	jobConfig := jobInit.Viper("../../../bin/job/config.yaml")
	// 初始化日志框架
	jobConfig.Log.Save = false
	jobConfig.Pgsql.AutoMigrate = false
	jobConfig.BlockChain.ChainID = 80001
	jobConfig.BlockChain.Provider = []jobConfigStruct.Provider{
		{Url: "https://rpc.ankr.com/polygon_mumbai"},
	}
	jobConfig.Pgsql.Prefix = "test_" // add test prefix
	log.Init(jobConfig.Log)

	jobS := jobService.New(jobConfig)
	_ = jobS

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
