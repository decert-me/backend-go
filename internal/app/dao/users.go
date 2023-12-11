package dao

import (
	"backend-go/internal/app/model"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/tidwall/gjson"
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
	return d.redis.Set(c, d.nonceKeyRedis(nonce), "", time.Duration(d.c.BlockChain.SignatureExp)*time.Second).Err()
}

func (d *Dao) DelNonce(c context.Context, nonce string) (err error) {
	return d.redis.Del(c, d.nonceKeyRedis(nonce)).Err()
}

func (d *Dao) CreateUser(user *model.Users) (err error) {
	return d.db.Create(&user).Error
}

func (d *Dao) GetUser(address string) (user model.Users, err error) {
	err = d.db.Model(&model.Users{}).Where("address", address).First(&user).Error
	return
}

func (d *Dao) UpdateUser(address string, user model.Users) error {
	raw := d.db.Model(&model.Users{}).Where("address = ?", address).Updates(&user)
	if raw.Error != nil {
		return raw.Error
	}
	if raw.RowsAffected == 0 {
		return errors.New("UpdateUser failed")
	}
	return nil
}

func (d *Dao) UpdateAvatar(address string, avatar string) error {
	raw := d.db.Model(&model.Users{}).Where("address = ?", address).Update("avatar", avatar)
	if raw.Error != nil {
		return raw.Error
	}
	if raw.RowsAffected == 0 {
		return errors.New("UpdateUser failed")
	}
	return nil
}

func (d *Dao) GetSocialsInfo(user *model.Users) (socials string, err error) {
	err = d.db.Model(&model.Users{}).Select("socials").
		Where("address = ?", user.Address).
		First(&socials).Error
	return
}

// HasBindSocialAccount 判断是否已经绑定
func (d *Dao) HasBindSocialAccount(address string) (wechat bool, discord bool, err error) {
	socials, err := d.GetSocialsInfo(&model.Users{Address: address})
	if err != nil {
		return wechat, discord, nil
	}
	id := gjson.Get(socials, "discord.id").String()
	if id != "" {
		discord = true
	}
	openid := gjson.Get(socials, "wechat.openid").String()
	if openid != "" {
		wechat = true
	}
	return wechat, discord, err
}
