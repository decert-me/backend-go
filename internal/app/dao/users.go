package dao

import (
	"backend-go/internal/app/model"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/tidwall/gjson"
	"gorm.io/datatypes"
	"time"
)

func (d *Dao) nonceKeyRedis(key string) string {
	return fmt.Sprintf("%snonce_%s", d.c.Redis.Prefix, key)
}
func getRebindKey(bindType string, address string) string {
	return fmt.Sprintf("rebind_%s_%s", bindType, address)
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
	err = d.db.Model(&model.Users{}).Select("COALESCE(socials,'{}')").
		Where("address = ?", user.Address).
		First(&socials).Error
	return
}

// HasBindSocialAccount 判断是否已经绑定
func (d *Dao) HasBindSocialAccount(address string) (data map[string]bool, err error) {
	var discord, wechat, github, email bool
	socials, err := d.GetSocialsInfo(&model.Users{Address: address})
	if err != nil {
		return data, nil
	}
	id := gjson.Get(socials, "discord.id").String()
	if id != "" {
		discord = true
	}
	openid := gjson.Get(socials, "wechat.openid").String()
	if openid != "" {
		wechat = true
	}
	githubID := gjson.Get(socials, "github.id").String()
	if githubID != "" {
		github = true
	}
	if gjson.Get(socials, "email").String() != "" {
		email = true
	}

	data = map[string]bool{"wechat": wechat, "discord": discord, "github": github, "email": email}
	return data, err
}

// GetNeedRebindInfo 获取需要换绑信息
func (d *Dao) GetNeedRebindInfo(address string, bindType string) (currentBindingAddress string, err error) {
	key := getRebindKey(bindType, address)
	// 从 Redis 获取
	currentBindingAddress, err = d.redis.Get(context.Background(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		}
		return currentBindingAddress, err
	}
	return currentBindingAddress, err
}

// SaveRebindInfo 保存换绑信息
func (d *Dao) SaveRebindInfo(address string, bindType string, currentBindingAddress string) (err error) {
	key := getRebindKey(bindType, address)
	return d.redis.Set(context.Background(), key, currentBindingAddress, time.Duration(5)*time.Minute).Err()
}

// ConfirmBindChange 确认换绑
func (d *Dao) ConfirmBindChange(address string, bindType string) (err error) {
	key := getRebindKey(bindType, address)
	// 从 Redis 获取
	currentBindingAddress, err := d.redis.Get(context.Background(), key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return errors.New("操作失败！请重新操作")
		}
		return err
	}
	tx := d.db.Begin()
	selectSQL := fmt.Sprintf("SELECT COALESCE(socials->'%s', '{}') FROM users WHERE address = ? LIMIT 1", bindType)
	var rawData string
	err = tx.Raw(selectSQL, currentBindingAddress).Scan(&rawData).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	updateSQL := fmt.Sprintf("UPDATE users SET socials = jsonb_set(COALESCE(socials,'{}'), '{\"%s\"}', ?) WHERE address = ?", bindType)
	err = tx.Exec(
		updateSQL,
		rawData,
		address,
	).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Exec("UPDATE users SET socials = socials - ? WHERE address = ?", bindType, currentBindingAddress).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Commit().Error
	if err != nil {
		return
	}
	// 从 Redis 删除
	return d.redis.Del(context.Background(), key).Err()
}

// CancelBindChange 取消绑定变更
func (d *Dao) CancelBindChange(address string, bindType string) (err error) {
	key := getRebindKey(bindType, address)
	// 从 Redis 删除
	err = d.redis.Del(context.Background(), key).Err()
	if err != nil {
		return
	}
	return nil
}

// ParticleUpdateSocialsInfo 更新社交信息
func (d *Dao) ParticleUpdateSocialsInfo(address string, particleUserinfo datatypes.JSON) (err error) {
	provider := gjson.Get(particleUserinfo.String(), "thirdparty_user_info.provider").String()
	if provider == "github" {
		githubID := gjson.Get(particleUserinfo.String(), "thirdparty_user_info.user_info.id").String()
		username := gjson.Get(particleUserinfo.String(), "thirdparty_user_info.user_info.name").String()
		return d.GithubBindAddress(githubID, username, address)
	} else if provider == "discord" {
		discordID := gjson.Get(particleUserinfo.String(), "thirdparty_user_info.user_info.id").String()
		username := gjson.Get(particleUserinfo.String(), "thirdparty_user_info.user_info.name").String()
		return d.DiscordBindAddress(discordID, username, address)
	}
	if gjson.Get(particleUserinfo.String(), "email").String() != "" {
		email := gjson.Get(particleUserinfo.String(), "email").String()
		return d.EmailBindAddress(address, email)
	}
	return nil
}
