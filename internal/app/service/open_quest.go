package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"gorm.io/datatypes"
	"math/big"
	"time"
)

// GetUserOpenQuestList 获取用户开放题列表
func (s *Service) GetUserOpenQuestList(address string, r request.GetUserOpenQuestListRequest) (list []response.UserOpenQuestJsonElements, total int64, err error) {
	offset := (r.Page - 1) * r.PageSize
	limit := r.PageSize
	// 查询用户是否有其它权限
	var questIDList []uint
	err = s.dao.DB().Model(&model.OpenQuestUserPerm{}).Where("address_list @> ?", pq.Array([]string{address})).Pluck("quest_id", &questIDList).Error
	if err != nil {
		return
	}
	db := s.dao.DB().Model(&model.UserOpenQuest{})
	if r.OpenQuestReviewStatus != 0 {
		countSQL := `
				SELECT 
					count(1)
				FROM
					user_open_quest
				JOIN
					jsonb_array_elements(user_open_quest.answer) WITH ORDINALITY AS t(json_element, idx) ON true
				JOIN 
					quest ON quest.token_id = user_open_quest.token_id
				WHERE
					user_open_quest.deleted_at IS NULL AND quest.status = 1 AND json_element->>'type' = 'open_quest'
		`
		if r.OpenQuestReviewStatus == 2 {
			countSQL += " AND (quest.creator = ? OR (quest.id IN ? AND json_element->>'open_quest_review_address' IS NOT NULL))"
			countSQL += " AND (json_element->>'score' IS NOT NULL OR json_element->>'correct' IS NOT NULL)"
		} else {
			countSQL += " AND (quest.creator = ? OR quest.id IN ? )"
			countSQL += " AND json_element->>'score' IS NULL AND json_element->>'correct' IS NULL"
		}
		err = db.Raw(countSQL, address, questIDList).Scan(&total).Error
		if err != nil {
			return
		}
		dataSQL := `
				SELECT 
					user_open_quest.id,
					user_open_quest.address,
					user_open_quest.token_id,
					CASE 
						WHEN json_element->>'score' IS NOT NULL OR json_element->>'correct' IS NOT NULL THEN 2
						ELSE 1
					END AS open_quest_review_status,
					json_element->>'open_quest_review_time' AS open_quest_review_time,
					user_open_quest.updated_at,
					(idx::int - 1)  AS index,
					json_element->>'type' AS type,
					json_element->>'value' AS value,
					quest.title AS challenge_title,
					(quest.quest_data->'questions')->(idx::int - 1)->>'title' AS title,
					(quest.quest_data->'questions')->(idx::int - 1)->>'score' AS score,
					(quest.quest_data->'questions')->(idx::int - 1)->>'correct' AS correct,
					json_element AS answer
				FROM
					user_open_quest
				JOIN
					jsonb_array_elements(user_open_quest.answer) WITH ORDINALITY AS t(json_element, idx) ON true
				JOIN 
					quest ON quest.token_id = user_open_quest.token_id
				WHERE
					user_open_quest.deleted_at IS NULL AND quest.status = 1 AND json_element->>'type' = 'open_quest'
		`
		if r.OpenQuestReviewStatus == 2 {
			// 评阅开放题状态 1 未审核 2 已审核
			dataSQL += " AND (quest.creator = ? OR (quest.id IN ? AND json_element->>'open_quest_review_address' = ?))"
			dataSQL += " AND (json_element->>'score' IS NOT NULL OR json_element->>'correct' IS NOT NULL)"
			dataSQL += " ORDER BY updated_at asc OFFSET ? LIMIT ?"
			err = db.Raw(dataSQL, address, questIDList, address, offset, limit).Scan(&list).Error
		} else {
			dataSQL += " AND (quest.creator = ? OR quest.id IN ? )"
			dataSQL += " AND json_element->>'score' IS NULL AND json_element->>'correct' IS NULL"
			dataSQL += " ORDER BY updated_at asc OFFSET ? LIMIT ?"
			err = db.Raw(dataSQL, address, questIDList, offset, limit).Scan(&list).Error
		}

	} else {
		err = db.Raw(`
				SELECT 
					count(1)
				FROM
					user_open_quest
				JOIN
					jsonb_array_elements(user_open_quest.answer) WITH ORDINALITY AS t(json_element, idx) ON true
				JOIN 
					quest ON quest.token_id = user_open_quest.token_id
				WHERE
					user_open_quest.deleted_at IS NULL AND quest.status = 1 AND json_element->>'type' = 'open_quest'  AND (quest.creator = ? OR (
					 quest.id IN ? AND ((
						((json_element->>'score' IS NOT NULL OR json_element->>'correct' IS NOT NULL) AND json_element->>'open_quest_review_address' = ?))
						OR
						(json_element->>'score' IS NULL AND json_element->>'correct' IS NULL)
					 ))
					)
		`, address, questIDList, address).Scan(&total).Error
		if err != nil {
			return
		}
		err = db.Raw(`
				SELECT 
					user_open_quest.id,
					user_open_quest.address,
					user_open_quest.token_id,
					CASE 
						WHEN json_element->>'score' IS NOT NULL OR json_element->>'correct' IS NOT NULL THEN 2
						ELSE 1
					END AS open_quest_review_status,
					json_element->>'open_quest_review_time' AS open_quest_review_time,
					user_open_quest.updated_at,
					(idx::int - 1) AS index,
					json_element->>'type' AS type,
					json_element->>'value' AS value,
					quest.title AS challenge_title,
					(quest.quest_data->'questions')->(idx::int - 1)->>'title' AS title,
					(quest.quest_data->'questions')->(idx::int - 1)->>'score' AS score,
					(quest.quest_data->'questions')->(idx::int - 1)->>'correct' AS correct,
					json_element AS answer
				FROM
					user_open_quest
				JOIN
					jsonb_array_elements(user_open_quest.answer) WITH ORDINALITY AS t(json_element, idx) ON true
				JOIN 
					quest ON quest.token_id = user_open_quest.token_id
				WHERE
					user_open_quest.deleted_at IS NULL AND quest.status = 1 AND json_element->>'type' = 'open_quest'  AND (quest.creator = ? OR (
					 quest.id IN ? AND ((
						((json_element->>'score' IS NOT NULL OR json_element->>'correct' IS NOT NULL) AND json_element->>'open_quest_review_address' = ?))
						OR
						(json_element->>'score' IS NULL AND json_element->>'correct' IS NULL)
					 ))
					)
				ORDER BY updated_at asc
				OFFSET ? LIMIT ?
		`, address, questIDList, address, offset, limit).Scan(&list).Error
	}
	return
}

func (s *Service) ReviewOpenQuest(address string, req []request.ReviewOpenQuestRequest) (err error) {
	// 开启事务
	db := s.dao.DB().Begin()
	// 用户开放题
	userOpenQuestTimeMap := make(map[uint]time.Time)
	// 题目
	questMap := make(map[int64]model.Quest)
	for _, r := range req {
		var userOpenQuest model.UserOpenQuest
		if err = db.Model(&model.UserOpenQuest{}).Where("id = ? AND open_quest_review_status = 1", r.ID).First(&userOpenQuest).Error; err != nil {
			db.Rollback()
			return errors.New("获取答案失败")
		}
		// 如果不存在
		if _, ok := userOpenQuestTimeMap[r.ID]; !ok {
			userOpenQuestTimeMap[r.ID] = userOpenQuest.UpdatedAt
		}
		// 检查是否有变动，跳过
		if r.UpdatedAt != nil && !userOpenQuestTimeMap[r.ID].Equal(*r.UpdatedAt) {
			continue
		}
		// 如果不存在
		if _, ok := questMap[userOpenQuest.TokenId]; !ok {
			var quest model.Quest
			if err = db.Model(&model.Quest{}).Where("token_id = ?", userOpenQuest.TokenId).First(&quest).Error; err != nil {
				db.Rollback()
				return errors.New("获取题目失败")
			}
			questMap[userOpenQuest.TokenId] = quest
		}
		// 获取题目
		quest := questMap[userOpenQuest.TokenId]
		// 检查是否用户发布的题目
		if quest.Creator != address {
			var questIDList []uint
			// 查询用户是否有其它权限
			err = s.dao.DB().Model(&model.OpenQuestUserPerm{}).Where("address_list @> ?", pq.Array([]string{address})).Pluck("quest_id", &questIDList).Error
			if err != nil {
				return
			}
			var isPerm bool
			for _, v := range questIDList {
				if v == quest.ID {
					isPerm = true
					break
				}
			}
			if !isPerm {
				db.Rollback()
				return errors.New("非法操作")
			}
		}
		// 填入审核人
		answerRaw, err := sjson.Set(string(r.Answer), "open_quest_review_address", address)
		if err != nil {
			db.Rollback()
			return errors.New("写入审核结果失败")
		}
		// 填入原有答案
		answer := userOpenQuest.Answer
		answerRes, err := sjson.Set(string(answer), fmt.Sprintf("%d", r.Index), gjson.Parse(answerRaw).Value())
		if err != nil {
			db.Rollback()
			return errors.New("写入审核结果失败")
		}
		// 判断所有开放题是否审核完成
		var openQuestReviewStatus uint8 = 2 // 已审核
		for _, v := range gjson.Get(answerRes, "@this").Array() {
			// 跳过不是开放题
			if v.Get("type").String() != "open_quest" {
				continue
			}
			// 判断分数是否为空
			if v.Get("score").String() == "" {
				openQuestReviewStatus = 1 // 未审核
				break
			}
		}
		var openQuestReviewTime time.Time // 审核时间
		var pass bool                     // 是否通过
		var userReturnScore int64         // 分数
		if openQuestReviewStatus == 2 {
			openQuestReviewTime = time.Now()
		}
		// 判断是否通过
		if openQuestReviewStatus == 2 {
			userReturnScore, pass, err = s.AnswerCheck(s.c.Quest.EncryptKey, answerRes, address, 0, &quest, false)
			if err != nil {
				db.Rollback()
				return errors.New("服务器错误")
			}
		}
		userReturnScoreFloat := big.NewFloat(float64(userReturnScore))
		scoreFloat := new(big.Float).Quo(userReturnScoreFloat, big.NewFloat(100))
		score := fmt.Sprintf("%.0f", scoreFloat)
		openQuestScore, _ := scoreFloat.Int64()
		// 写入审核结果
		err = db.Model(&model.UserOpenQuest{}).Where("id = ? AND open_quest_review_status = 1", r.ID).Updates(&model.UserOpenQuest{
			OpenQuestReviewTime:   openQuestReviewTime,
			OpenQuestReviewStatus: openQuestReviewStatus,
			OpenQuestScore:        openQuestScore,
			Answer:                datatypes.JSON(answerRes),
			Pass:                  pass,
		}).Error
		if err != nil {
			db.Rollback()
			return errors.New("写入结果失败")
		}
		// 审核完成发送消息
		if openQuestReviewStatus == 2 {
			// 写入Message
			var message model.UserMessage
			if pass {
				message = model.UserMessage{
					Title:     "恭喜通过挑战",
					TitleEn:   "Congratulations on passing the challenge!",
					Content:   "你在《" + quest.Title + "》的挑战成绩为 " + cast.ToString(score) + " 分，可领取一枚NFT！",
					ContentEn: "Your score for the challenge \"" + quest.Title + "\" is " + cast.ToString(score) + " points, and you can claim an NFT!",
				}
				// 创建证书
				go func() {
					s.GenerateCardInfo(userOpenQuest.Address, openQuestScore, request.GenerateCardInfoRequest{
						TokenId: userOpenQuest.TokenId,
						Answer:  answerRes,
					})
				}()
			} else {
				message = model.UserMessage{
					Title:     "挑战未通过",
					TitleEn:   "Challenge failed",
					Content:   "你在《" + quest.Title + "》的挑战成绩为 " + cast.ToString(score) + " 分，请继续加油吧！",
					ContentEn: "Your score for the challenge \"" + quest.Title + "\" is " + cast.ToString(score) + " points, please continue to working hard.",
				}
			}
			message.TokenId = quest.TokenId
			message.Address = userOpenQuest.Address
			err = db.Model(&model.UserMessage{}).Create(&message).Error
			if err != nil {
				db.Rollback()
				return errors.New("发送消息失败")
			}
		}
	}
	return db.Commit().Error
}
