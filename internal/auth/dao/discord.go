package dao

import "encoding/json"

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

// DiscordIsBinding 判断是否已经绑定过
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
