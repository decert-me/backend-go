package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"errors"
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
		// 编程题目
		if gjson.Get(v.String(), "type").String() == "coding" || gjson.Get(v.String(), "type").String() == "special_judge_coding" {
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
		if answerS[i].String() == answerU[i].String() {
			score += scoreList[i].Int()
		}
	}
	if userScore == (score*10000/totalScore) && score >= passingScore {
		return true, nil
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
			if gjson.Get(v.String(), "type").String() == "coding" || gjson.Get(v.String(), "type").String() == "special_judge_coding" {
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
	res, err := client.R().SetBody(body).Post(s.c.Judge.SolidityAPI[0])
	if err != nil {
		log.Errorv("Post error", zap.Error(err))
	}
	if gjson.Get(res.String(), "data.correct").Bool() {
		return true
	}
	return false
}
