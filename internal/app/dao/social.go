package dao

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func (d *Dao) WechatQueryByAddress(address string) (wechatData string, err error) {
	err = d.db.Raw("SELECT COALESCE(socials->'wechat','{}') FROM users WHERE address = ? LIMIT 1", address).Scan(&wechatData).Error
	if err != nil {
		return wechatData, err
	}
	return wechatData, err
}

// WechatIsBinding 判断是否已经绑定过
func (d *Dao) WechatIsBinding(fromUserName string) (bool, error) {
	var count int
	err := d.db.Raw("SELECT count(1) FROM users WHERE socials->'wechat'->>'openid' = ?", fromUserName).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// WechatBindAddress 处理地址绑定
func (d *Dao) WechatBindAddress(address, fromUserName string) (err error) {
	//
	// 绑定
	// 判断记录是否存在
	var count int
	err = d.db.Raw("SELECT count(1) FROM users WHERE address = ?", address).Scan(&count).Error
	if err != nil {
		return errors.New("服务器内部错误")
	}
	wechatInfo := map[string]string{
		"openid": fromUserName,
	}
	wechatInfoBytes, err := json.Marshal(wechatInfo)
	if err != nil {
		return errors.New("服务器内部错误")
	}
	wechatInfoStr := string(wechatInfoBytes)
	if count == 0 {
		// 插入
		err = d.db.Exec("INSERT INTO users (address, socials) VALUES (?, jsonb_set('{}', '{\"wechat\"}', ?))", address, wechatInfoStr).Error
		return err
	} else {
		err = d.db.Exec(
			"UPDATE users SET socials = jsonb_set(COALESCE(socials,'{}'), '{\"wechat\"}', ?) WHERE address = ?",
			wechatInfoStr,
			address,
		).Error
		return err
	}
}

// DiscordBindAddress 处理地址绑定
func (d *Dao) DiscordBindAddress(discordID, username, address string) (err error) {
	// 判断记录是否存在
	var count int
	err = d.db.Raw("SELECT count(1) FROM users WHERE address = ?", address).Scan(&count).Error
	if err != nil {
		return err
	}
	discordInfo := map[string]string{
		"id":       discordID,
		"username": username,
	}
	discordInfoBytes, err := json.Marshal(discordInfo)
	if err != nil {
		return err
	}
	discordInfoStr := string(discordInfoBytes)
	if count == 0 {
		// 插入
		err = d.db.Exec("INSERT INTO users (address, socials) VALUES (?, jsonb_set('{}', '{\"discord\"}', ?))", address, discordInfoStr).Error
		return err
	} else {
		err = d.db.Exec(
			"UPDATE users SET socials = jsonb_set(COALESCE(socials,'{}'), '{\"discord\"}', ?) WHERE address = ?",
			discordInfoStr,
			address,
		).Error
		return err
	}
}

// DiscordQueryByAddress 查询地址绑定的discord信息
func (d *Dao) DiscordQueryByAddress(address string) (discordData string, err error) {
	err = d.db.Raw("SELECT COALESCE(socials->'discord', '{}') FROM users WHERE address = ? LIMIT 1", address).Scan(&discordData).Error
	if err != nil {
		return discordData, err
	}
	return discordData, err
}

// DiscordIsBinding 判断Discord是否已经绑定过
func (d *Dao) DiscordIsBinding(discordID string) (bool, error) {
	var count int
	err := d.db.Raw("SELECT count(1) FROM users WHERE socials->'discord'->>'id' = ?", discordID).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// EmailIsBinding 判断邮箱是否已经绑定过
func (d *Dao) EmailIsBinding(email string) (bool, error) {
	var count int
	err := d.db.Raw("SELECT count(1) FROM users WHERE socials->>'email' = ?", email).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// EmailQueryByAddress 查询地址绑定的邮箱信息
func (d *Dao) EmailQueryByAddress(address string) (emailData string, err error) {
	err = d.db.Raw("SELECT COALESCE(socials->'email', '{}') FROM users WHERE address = ? LIMIT 1", address).Scan(&emailData).Error
	if err != nil {
		return emailData, err
	}
	return emailData, err
}

// EmailBindAddress 处理地址绑定
func (d *Dao) EmailBindAddress(address, email string) (err error) {
	// 判断记录是否存在
	var count int
	err = d.db.Raw("SELECT count(1) FROM users WHERE address = ?", address).Scan(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		// 插入
		err = d.db.Exec("INSERT INTO users (address, socials) VALUES (?, jsonb_set('{}', '{\"email\"}', to_jsonb(?::text)))", address, email).Error
		return err
	} else {
		err = d.db.Exec(
			"UPDATE users SET socials = jsonb_set(COALESCE(socials,'{}'), '{\"email\"}', to_jsonb(?::text)) WHERE address = ?",
			email,
			address,
		).Error
		return err
	}
}

// EmailGetCode 获取邮箱验证码
func (d *Dao) EmailGetCode(email string) (code string, err error) {
	// 随机生成验证码
	rand.Seed(time.Now().UnixNano())
	code = fmt.Sprintf("%06d", rand.Intn(1000000))
	// 写入 Redis
	err = d.redis.Set(context.Background(), d.c.Redis.Prefix+email, code, time.Minute*30).Err()
	if err != nil {
		return code, err
	}
	return code, err
}

// EmailQueryCode 查询邮箱验证码
func (d *Dao) EmailQueryCode(email string) (code string, err error) {
	// Redis
	code, err = d.redis.Get(context.Background(), d.c.Redis.Prefix+email).Result()
	if err != nil {
		return code, err
	}
	return code, err
}

// GithubQueryByAddress 查询地址绑定的github信息
func (d *Dao) GithubQueryByAddress(address string) (githubData string, err error) {
	err = d.db.Raw("SELECT COALESCE(socials->'github', '{}') FROM users WHERE address = ? LIMIT 1", address).Scan(&githubData).Error
	if err != nil {
		return githubData, err
	}
	return githubData, err
}

// GithubIsBinding 判断Github是否已经绑定过
func (d *Dao) GithubIsBinding(githubID string) (bool, error) {
	var count int
	err := d.db.Raw("SELECT count(1) FROM users WHERE socials->'github'->>'id' = ?", githubID).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// GithubBindAddress 处理地址绑定
func (d *Dao) GithubBindAddress(githubID, username, address string) (err error) {
	// 判断记录是否存在
	var count int
	err = d.db.Raw("SELECT count(1) FROM users WHERE address = ?", address).Scan(&count).Error
	if err != nil {
		return err
	}
	githubInfo := map[string]string{
		"id":       githubID,
		"username": username,
	}
	githubInfoBytes, err := json.Marshal(githubInfo)
	if err != nil {
		return err
	}
	githubInfoStr := string(githubInfoBytes)
	if count == 0 {
		// 插入
		err = d.db.Exec("INSERT INTO users (address, socials) VALUES (?, jsonb_set('{}', '{\"github\"}', ?))", address, githubInfoStr).Error
		return err
	} else {
		err = d.db.Exec(
			"UPDATE users SET socials = jsonb_set(COALESCE(socials,'{}'), '{\"github\"}', ?) WHERE address = ?",
			githubInfoStr,
			address,
		).Error
		return err
	}
}
