package dao

import (
	"backend-go/internal/app/model"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm/clause"
	"time"
)

func (d *Dao) nonceKeyRedis(key string) string {
	return fmt.Sprintf("%snonce_%s", d.c.Redis.Prefix, key)
}
func (d *Dao) HasNonce(c context.Context, nonce string) (has bool, err error) {
	if err = d.redis.Get(c, d.nonceKeyRedis(nonce)).Err(); err != nil {
		if err == redis.Nil {
			err = nil
		}
		return
	}
	return true, nil
}

func (d *Dao) SetNonce(c context.Context, nonce string) (err error) {
	return d.redis.Set(c, d.nonceKeyRedis(nonce), "4", time.Duration(d.c.BlockChain.SignatureExp)).Err()
}

func (d *Dao) DelNonce(c context.Context, nonce string) (err error) {
	return d.redis.Del(c, d.nonceKeyRedis(nonce)).Err()
}

func (d *Dao) SaveUser(user *model.Users) (err error) {
	if err = d.db.Model(&model.Users{}).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "address"}},
		UpdateAll: true,
	}).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (d *Dao) GetSocialsInfo(user *model.Users) (socials string, err error) {
	err = d.db.Model(&model.Users{}).Select("socials").
		Where("address = ?", user.Address).
		First(&socials).Error
	return
}
