package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"errors"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"time"
)

func (s *Service) AnswerCheck(key, answerUser string, userScore int64, quest *model.Quest) (pass bool, err error) {
	res := string(quest.MetaData)
	questData := string(quest.QuestData)
	version := gjson.Get(res, "version").Float()

	answerU, scoreList, answerS, passingScore := utils.GetAnswers(version, key, res, questData, answerUser)
	var totalScore int64
	for _, s := range scoreList {
		totalScore += s.Int()
	}
	if len(answerU) != len(answerS) || len(scoreList) != len(answerS) {
		return false, errors.New("unexpect error")
	}
	var score int64
	for i, v := range answerS {
		if v.String() == "" {
			continue
		}
		questType := gjson.Get(v.String(), "type").String()
		questValue := gjson.Get(v.String(), "value").String()
		// 编程题目
		if questType == "coding" || questType == "special_judge_coding" {
			// 跳过不正确
			if gjson.Get(v.String(), "correct").Bool() == false {
				continue
			}
			reqMap := make(map[string]interface{})
			reqMap["code"] = gjson.Get(v.String(), "code").String()
			reqMap["lang"] = gjson.Get(v.String(), "lang").String()
			reqMap["token_id"] = quest.TokenId
			reqMap["quest_index"] = i
			// 检查答案
			if s.CodingCheck(reqMap) {
				score += scoreList[i].Int()
			}
			continue
		}
		// 单选题
		if questType == "multiple_choice" || questType == "fill_blank" {
			if questValue == answerU[i].String() {
				score += scoreList[i].Int()
			}
			continue
		}
		// 多选题
		if questType == "multiple_response" {
			answerArray := gjson.Get(questValue, "@this").Array()
			fmt.Println(len(answerArray))
			fmt.Println(len(answerU[i].Array()))
			// 数量
			if len(answerArray) != len(answerU[i].Array()) {
				continue
			}
			// 内容
			allRight := true
			for _, v := range answerArray {
				var right bool
				for _, item := range answerU[i].Array() {
					if item.String() == v.String() {
						right = true
						break
					}
				}
				if !right {
					allRight = false
					break
				}
			}
			if allRight {
				score += scoreList[i].Int()
			}
		}
	}

	if userScore == (score*10000/totalScore) && score >= passingScore {
		return true, nil
	} else {
		return true, errors.New("not enough scores")
	}
	return
}

func (s *Service) AnswerScore(key, answerUser, uri string, quest model.Quest) (userScore int64, pass bool, err error) {
	res := quest.MetaData.String()
	questData := string(quest.QuestData)
	version := gjson.Get(res, "version").Float()
	answerU, scoreList, answerS, passingScore := utils.GetAnswers(version, key, res, questData, answerUser)

	var totalScore int64
	for _, s := range scoreList {
		totalScore += s.Int()
	}
	if len(answerU) != len(answerS) || len(scoreList) != len(answerS) {
		return userScore, false, errors.New("unexpect error")
	}
	var score int64
	for i, v := range answerS {
		if answerS[i].String() == answerU[i].String() {
			// 编程题目
			if gjson.Get(v.String(), "type").String() == "coding" || gjson.Get(v.String(), "type").String() == "special_judge_coding" || answerU[i].String() == "" {
				// 跳过不正确
				if gjson.Get(v.String(), "correct").Bool() == false {
					continue
				}
				reqMap := make(map[string]interface{})
				reqMap["code"] = gjson.Get(v.String(), "code").String()
				reqMap["lang"] = gjson.Get(v.String(), "lang").String()
				reqMap["token_id"] = quest.TokenId
				reqMap["quest_index"] = i
				// 检查答案
				if s.CodingCheck(reqMap) {
					score += scoreList[i].Int()
				}
				continue
			}
			score += scoreList[i].Int()
		}
	}
	if score >= passingScore {
		return score, true, nil
	}
	return score, false, nil
}

func (s *Service) CodingCheck(body interface{}) (correct bool) {
	client := req.C().SetTimeout(180 * time.Second)
	url := s.W.Next().Item + "/run/tryRun"
	res, err := client.R().SetBody(body).Post(url)
	if err != nil {
		log.Errorv("Post error", zap.Error(err))
	}
	if gjson.Get(res.String(), "data.correct").Bool() {
		return true
	}
	return false
}
