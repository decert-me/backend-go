package dao

import (
	"encoding/json"
)

// WechatBindAddress 处理地址绑定
func (d *Dao) WechatBindAddress(address, fromUserName string) (msg string, err error) {
	//
	// 绑定
	// 判断记录是否存在
	var count int
	err = d.db.Raw("SELECT count(1) FROM users WHERE address = ?", address).Scan(&count).Error
	if err != nil {
		return "服务器内部错误", err
	}
	wechatInfo := map[string]string{
		"openid": fromUserName,
	}
	wechatInfoBytes, err := json.Marshal(wechatInfo)
	if err != nil {
		return "服务器内部错误", err
	}
	wechatInfoStr := string(wechatInfoBytes)
	if count == 0 {
		// 插入
		err = d.db.Exec("INSERT INTO users (address, socials) VALUES (?, jsonb_set('{}', '{\"wechat\"}', ?))", address, wechatInfoStr).Error
		return "服务器内部错误", err
	} else {
		err = d.db.Exec(
			"UPDATE users SET socials = jsonb_set(COALESCE(socials,'{}'), '{\"wechat\"}', ?) WHERE address = ?",
			wechatInfoStr,
			address,
		).Error
		return "服务器内部错误", err
	}
}

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
