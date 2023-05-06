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
