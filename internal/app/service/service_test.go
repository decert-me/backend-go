package service

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/dao"
	"backend-go/internal/app/initialize"
	"backend-go/internal/app/model"
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

func TestMain(m *testing.M) {
	c = initialize.Viper("../cmd/config.yaml")
	// 初始化日志框架
	c.Log.Save = false
	c.Log.Level = "silent"
	c.Log.LogInConsole = true
	log.Init(c.Log)
	c.Pgsql.LogMode = "silent"
	c.Pgsql.AutoMigrate = true
	c.Twitter.ClaimContent = "我在 @DecertMe 上完成了一个挑战并获得了链上能力认证的徽章。\nhttps://decert.me/quests/\n#DecertMe"
	// test contract address
	c.Contract = &config.Contract{
		Badge:       "0x0049770260b599Ecc2e2c0645450c965A44938b7",
		Quest:       "0xfE3e0366a52C6F668a1026dAF5e81162d34Ec38b",
		QuestMinter: "0xbE866FE4BAFC11ae886238772AFBD24570f9B530",
	}
	c.Pgsql.Prefix = "test_" // add test prefix
	d = dao.New(c)
	s = New(c)
	/*
	 * Job
	 */
	jobConfig := jobInit.Viper("../../job/cmd/config.yaml")
	// 初始化日志框架
	jobConfig.Log.Save = false
	jobConfig.Log.Level = "silent"
	jobConfig.Log.LogInConsole = true
	jobConfig.Pgsql.AutoMigrate = false

	jobConfig.Pgsql.Prefix = "test_" // add test prefix
	log.Init(jobConfig.Log)

	jobS := jobService.New(jobConfig)
	_ = jobS

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
