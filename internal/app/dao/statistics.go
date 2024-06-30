package dao

import "backend-go/internal/app/model/response"

func (d *Dao) GetAddressChallengeCount(address string) (res response.GetAddressChallengeCountRes, err error) {
	err = d.db.Table("users").
		Select("users.id as user_id, users.address, users.name,string_agg(tag.name, ',') as tags").
		Joins("LEFT JOIN users_tag ON users_tag.user_id = users.id").
		Joins("LEFT JOIN tag ON tag.id = users_tag.tag_id").
		Where("users.address = ?", address).
		Group("users.id").
		Find(&res).Error
	if err != nil {
		return res, err
	}
	// 领取NFT数量
	d.db.Raw(`SELECT COUNT(1) FROM
			(
			SELECT
				quest.ID 
			FROM
				"user_challenges"
			LEFT JOIN quest ON quest.token_id = user_challenges.token_id 
			LEFT JOIN user_challenge_log ON user_challenge_log.token_id = user_challenges.token_id AND user_challenge_log.address = user_challenges.address
			WHERE
				user_challenges.address = ? AND quest.token_id IS NOT NULL AND user_challenge_log.token_id IS NOT NULL 
			GROUP BY
				quest.ID 
			UNION
			SELECT
				quest_id 
			FROM
				zcloak_card 
			LEFT JOIN quest ON quest.id = zcloak_card.quest_id 
						LEFT JOIN user_challenge_log ON user_challenge_log.token_id = quest.token_id AND user_challenge_log.address = zcloak_card.address
			WHERE
				zcloak_card.address = ? AND quest.token_id IS NOT NULL AND user_challenge_log.token_id IS NOT NULL 
			GROUP BY
			quest_id ) AS f`, address, address).Scan(&res.ClaimNum)
	// 挑战成功/失败数量
	type CountResult struct {
		TokenId      string `json:"token_id"`
		PassCount    int    `json:"pass_count"`
		NotPassCount int    `json:"not_pass_count"`
	}
	var countResult []CountResult
	if err := d.db.Raw(`
		SELECT 
			token_id,
			sum(pass_count) as pass_count,
			sum(not_pass_count) as not_pass_count
		FROM (
			(SELECT user_challenge_log.token_id, sum(case when pass then 1 else 0 end) as pass_count, sum(case when not pass then 1 else 0 end) as not_pass_count 
			 FROM user_challenge_log
			 LEFT JOIN quest ON user_challenge_log.token_id=quest.token_id
			 WHERE address = ? AND quest.token_id IS NOT NULL
			 GROUP BY user_challenge_log.token_id)
			UNION ALL
			(SELECT user_open_quest.token_id, sum(case when pass then 1 else 0 end) as pass_count, sum(case when not pass then 1 else 0 end) as not_pass_count 
			 FROM user_open_quest
			 LEFT JOIN quest ON user_open_quest.token_id=quest.token_id
			 WHERE address = ? AND quest.token_id IS NOT NULL
			 GROUP BY user_open_quest.token_id)
		) as combined
		GROUP BY token_id
		`, address, address).Scan(&countResult).Error; err != nil {
		return res, err
	}
	for _, result := range countResult {
		if result.PassCount == 0 {
			res.FailNum += 1
			continue
		}
		res.SuccessNum += 1
	}
	res.NotClaimNum = res.SuccessNum - res.ClaimNum
	if res.NotClaimNum < 0 {
		res.NotClaimNum = 0
	}
	return res, nil
}
