package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"errors"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (s *Service) AnswerCheck(key, answerUser string, userScore int64, quest *model.Quest) (pass bool, err error) {
	res := string(quest.MetaData)
	answerU := gjson.Get(utils.AnswerDecode(key, gjson.Get(res, "answers").String()), "@this").Array() // 标准答案
	passingScore := gjson.Get(res, "passingScore").Int()                                               // 通过分数
	scoreList := gjson.Get(res, "questions.#.score").Array()                                           // 题目分数
	answerS := gjson.Get(answerUser, "@this").Array()                                                  // 用户答案
	var totalScore int64
	for _, s := range scoreList {
		totalScore += s.Int()
	}
	if len(answerU) != len(answerS) || len(scoreList) != len(answerS) {
		return false, errors.New("unexpect error")
	}
	var score int64
	for i, _ := range answerS {
		_, isUUID := uuid.Parse(answerS[i].String())
		if isUUID == nil {
			if s.JudgeResultCheck(answerS[i].String(), quest, uint8(i)) {
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

func (s *Service) JudgeResultCheck(uuid string, quest *model.Quest, index uint8) (pass bool) {
	res, err := s.dao.FilterJudgeResult(
		model.JudgeResult{
			ID:         uuid,
			TokenID:    quest.TokenId,
			QuestIndex: index,
		})
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Errorv("FilterJudgeResult error", zap.Error(err))
	}
	return res.Pass
}

func AnswerScore(key, answerUser, uri string) (userScore int64, pass bool, err error) {
	res, err := utils.GetDataFromUri(uri)
	if err != nil || !gjson.Valid(res) {
		return userScore, pass, errors.New("tokenID invalid")
	}
	answerU := gjson.Get(utils.AnswerDecode(key, gjson.Get(res, "answers").String()), "@this").Array() // 标准答案
	passingScore := gjson.Get(res, "passingScore").Int()                                               // 通过分数
	scoreList := gjson.Get(res, "questions.#.score").Array()                                           // 题目分数
	answerS := gjson.Get(answerUser, "@this").Array()                                                  // 用户答案
	var totalScore int64
	for _, s := range scoreList {
		totalScore += s.Int()
	}
	if len(answerU) != len(answerS) || len(scoreList) != len(answerS) {
		return userScore, false, errors.New("unexpect error")
	}
	var score int64
	for i, _ := range answerS {
		if answerS[i].String() == answerU[i].String() {
			score += scoreList[i].Int()
		}
	}
	if score >= passingScore {
		return score, true, nil
	}
	return score, false, nil
}