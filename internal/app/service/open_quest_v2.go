package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/pkg/log"
	"fmt"
	"go.uber.org/zap"
	"sort"
)

func (s *Service) GetUserOpenQuestListV2(address string, r request.GetUserOpenQuestListRequest) (list []response.UserOpenQuestJsonElementsV2, total int64, totalToReview int64, err error) {
	offset := (r.Page - 1) * r.PageSize
	limit := r.PageSize
	db := s.dao.DB().Model(&model.UserOpenQuest{})
	dataSQL := `
		SELECT
			t.json_element ->> 'title' as title,quest.title as challenge_title,quest.uuid,quest.token_id,(idx::int - 1)  AS index,quest.add_ts as add_ts
		FROM
			quest,
			jsonb_array_elements (quest.quest_data -> 'questions') WITH ORDINALITY AS t(json_element, idx)
		WHERE
			t.json_element ->> 'type' = 'open_quest' AND quest.creator = ? AND quest.status = 1
	`
	err = db.Raw(dataSQL, address).Scan(&list).Error
	if err != nil {
		return
	}
	for i := 0; i < len(list); i++ {
		// 待评分数量
		toReviewCountSQL := `
		SELECT 
			count(1)
		FROM
			user_open_quest
		JOIN
			jsonb_array_elements(user_open_quest.answer) WITH ORDINALITY AS t(json_element, idx) ON true
		JOIN 
			quest ON quest.token_id = user_open_quest.token_id
		WHERE
			user_open_quest.deleted_at IS NULL AND quest.status = 1 AND json_element->>'type' = 'open_quest' AND quest.creator = ? AND quest.token_id = ? AND idx= ?
			AND json_element->>'score' IS NULL AND json_element->>'correct' IS NULL  AND json_element->>'score' IS NULL AND json_element->>'correct' IS NULL
		`
		err = s.dao.DB().Raw(toReviewCountSQL, address, list[i].TokenId, list[i].Index+1).Scan(&list[i].ToReviewCount).Error
		if err != nil {
			continue
		}
		// 已评分数量
		reviewedCountSQL := `
		SELECT 
			count(1)
		FROM
			user_open_quest
		JOIN
			jsonb_array_elements(user_open_quest.answer) WITH ORDINALITY AS t(json_element, idx) ON true
		JOIN 
			quest ON quest.token_id = user_open_quest.token_id
		WHERE
			user_open_quest.deleted_at IS NULL AND quest.status = 1 AND json_element->>'type' = 'open_quest' AND quest.creator = ? AND quest.token_id = ? AND idx= ?
			AND (json_element->>'score' IS NOT NULL OR json_element->>'correct' IS NOT NULL)
		`
		err = s.dao.DB().Raw(reviewedCountSQL, address, list[i].TokenId, list[i].Index+1).Scan(&list[i].ReviewedCount).Error
		if err != nil {
			continue
		}
		// 查询最新提交时间
		err = s.dao.DB().Model(&model.UserOpenQuest{}).Select("created_at").Where("token_id = ?", list[i].TokenId).Order("id desc").First(&list[i].LastSummitTime).Error
		if err != nil {
			continue
		}
		// 查询最新审核时间
		selectSQL := `
		SELECT
			to_timestamp(t.json_element ->> 'open_quest_review_time','YYYY-MM-DD HH24:MI:SS')
		FROM
			user_open_quest,
			jsonb_array_elements (user_open_quest.answer) WITH ORDINALITY AS t(json_element, idx)
		WHERE
		user_open_quest.token_id = ? AND idx= ? AND t.json_element ->> 'open_quest_review_time' != ''
		ORDER BY t.json_element ->> 'open_quest_review_time' desc
		limit 1
		`
		err = s.dao.DB().Model(&model.UserOpenQuest{}).Raw(selectSQL, list[i].TokenId, list[i].Index+1).Scan(&list[i].LastReviewTime).Error
		if err != nil {
			continue
		}
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Addts > list[j].Addts
	})
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].ToReviewCount > 0
	})
	sort.SliceStable(list, func(i, j int) bool {
		if list[i].ToReviewCount != 0 {
			if list[i].TokenId == list[j].TokenId {
				return list[i].Index < list[j].Index
			}
			return list[i].Addts > list[j].Addts
		}
		return false
	})
	// 过滤
	temp := make([]response.UserOpenQuestJsonElementsV2, 0)
	for _, v := range list {
		if v.LastSummitTime.IsZero() {
			continue
		}
		temp = append(temp, v)
	}
	for i := 0; i < len(list); i++ {
		// 先按照ToReviewCount倒序排序
		totalToReview += list[i].ToReviewCount
	}
	total = int64(len(temp))
	// limit offset
	result := make([]response.UserOpenQuestJsonElementsV2, 0)
	for i := offset; i < (offset+limit) && i < len(temp); i++ {
		result = append(result, temp[i])
	}
	return result, total, totalToReview, nil
}

// GetUserOpenQuestDetailListV2 获取用户开放题详情
func (s *Service) GetUserOpenQuestDetailListV2(address string, loginAddress string, r request.GetUserOpenQuestDetailListRequest) (list []response.UserOpenQuestJsonElements, total int64, err error) {
	offset := (r.Page - 1) * r.PageSize
	limit := r.PageSize
	db := s.dao.DB().Model(&model.UserOpenQuest{})
	// 是否管理员
	isAdmin, err := s.dao.IsAdmin(loginAddress)
	if err != nil {
		return
	}
	// OpenQuestReviewStatus 1 未审核 2 已审核
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
			user_open_quest.deleted_at IS NULL AND quest.status = 1 AND json_element->>'type' = 'open_quest' AND quest.creator = ? AND quest.token_id = ? AND idx= ?
	`
	if r.OpenQuestReviewStatus == 2 {
		countSQL += " AND (json_element->>'score' IS NOT NULL OR json_element->>'correct' IS NOT NULL)"
	} else {
		countSQL += " AND json_element->>'score' IS NULL AND json_element->>'correct' IS NULL"
	}
	err = db.Raw(countSQL, address, r.TokenID, *r.Index+1).Scan(&total).Error
	if err != nil {
		return
	}
	dataSQL := `
				SELECT 
					user_open_quest.id,
					user_open_quest.address,
					quest.uuid,
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
					json_element AS answer,
					quest.quest_data->>'passingScore' AS pass_score,
					quest.quest_data AS quest_data,
					quest.meta_data AS meta_data,  
					user_open_quest.answer AS user_answer
				FROM
					user_open_quest
				JOIN
					jsonb_array_elements(user_open_quest.answer) WITH ORDINALITY AS t(json_element, idx) ON true
				JOIN 
					quest ON quest.token_id = user_open_quest.token_id
				WHERE
					user_open_quest.deleted_at IS NULL AND quest.status = 1 AND json_element->>'type' = 'open_quest' AND quest.creator = ? AND quest.token_id = ? AND idx= ?
		`
	if r.OpenQuestReviewStatus == 2 {
		dataSQL += " AND (json_element->>'score' IS NOT NULL OR json_element->>'correct' IS NOT NULL)"
	} else {
		dataSQL += " AND json_element->>'score' IS NULL AND json_element->>'correct' IS NULL"
	}
	dataSQL += " ORDER BY updated_at asc OFFSET ? LIMIT ?"
	err = db.Raw(dataSQL, address, r.TokenID, *r.Index+1, offset, limit).Scan(&list).Error
	if err != nil {
		return
	}
	// 计算分数
	for i := 0; i < len(list); i++ {
		// 提交次数
		submitCountSQL := `
		SELECT 
			count(1)
		FROM
			user_open_quest
		WHERE
			user_open_quest.deleted_at IS NULL AND user_open_quest.token_id = ? AND user_open_quest.address = ? AND user_open_quest.id <= ?
		`
		err = db.Raw(submitCountSQL, list[i].TokenId, list[i].Address, list[i].ID).Scan(&list[i].SubmitCount).Error
		if err != nil {
			log.Error("获取提交次数失败", "error", err)
			continue
		}
		// 上次分数
		lastAnswerSQL := `
		SELECT
			COALESCE(json_element->>'score', '0') AS score
		FROM
			user_open_quest
		JOIN
			jsonb_array_elements(user_open_quest.answer) WITH ORDINALITY AS t(json_element, idx) ON true
		WHERE
			user_open_quest.deleted_at IS NULL AND json_element->>'type' = 'open_quest' AND user_open_quest.token_id = ? AND idx = ? AND user_open_quest.address = ? AND open_quest_review_status=2  AND user_open_quest.id < ?
		ORDER BY
			updated_at DESC
		LIMIT 1
		`
		err = db.Raw(lastAnswerSQL, list[i].TokenId, list[i].Index+1, list[i].Address, list[i].ID).Scan(&list[i].LastScore).Error
		if err != nil {
			log.Error("获取上次分数失败", "error", err)
			continue
		}
		quest := model.Quest{
			TokenId:   list[i].TokenId,
			MetaData:  list[i].MetaData,
			QuestData: list[i].QuestData,
		}
		list[i].TotalScore, list[i].UserScore, _, _, err = s.AnswerCheck(s.c.Quest.EncryptKey, list[i].UserAnswer.String(), address, 0, &quest, false)
		if err != nil {
			log.Errorv("AnswerCheck error", zap.Error(err))
			return
		}
		var showStr string
		showStr = fmt.Sprintf("%s...%s", list[i].Address[:6], list[i].Address[len(list[i].Address)-4:])
		if isAdmin {
			// 显示标签
			nickname, name, tags, err := s.dao.GetUserNameTagsByAddress(list[i].Address)
			if err == nil {
				if nickname != "" {
					showStr = nickname
				}
				if name != "" {
					showStr += "-" + name
				}
				for i := 0; i < len(tags); i++ {
					showStr += "，" + tags[i]
				}
			}
		}
		list[i].NickName = &showStr
	}
	return
}
