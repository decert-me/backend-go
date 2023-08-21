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
	"strings"
	"time"
)

func (s *Service) AnswerCheck(key, answerUser, address string, userScore int64, quest *model.Quest) (userReturnScore int64, pass bool, err error) {
	res := string(quest.MetaData)
	questData := string(quest.QuestData)
	version := gjson.Get(res, "version").Float()

	answerU, scoreList, answerS, passingScore := utils.GetAnswers(version, key, res, questData, answerUser)
	var totalScore int64
	for _, s := range scoreList {
		totalScore += s.Int()
	}
	if len(answerU) != len(answerS) || len(scoreList) != len(answerS) {
		return userReturnScore, false, errors.New("unexpect error")
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
			reqMap["lang"] = gjson.Get(v.String(), "language").String()
			reqMap["token_id"] = quest.TokenId
			reqMap["quest_index"] = i
			reqMap["quest"] = quest
			reqMap["address"] = address
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
	if userScore == 0 {
		if score >= passingScore {
			return score * 10000 / totalScore, true, nil
		} else {
			return score * 10000 / totalScore, false, nil
		}
	}
	if userScore == (score*10000/totalScore) && score >= passingScore {
		return userScore, true, nil
	} else {
		return userScore, true, errors.New("not enough scores")
	}
	return
}

func (s *Service) CodingCheck(body interface{}) (correct bool) {
	client := req.C().SetTimeout(180 * time.Second)
	i := 0
	var item string
	// 存活检测
	for {
		if i > 2 {
			break
		}
		w := s.W.Next()
		item = w.Item
		res, err := req.C().SetTimeout(5 * time.Second).R().SetBody(body).Get(strings.Replace(item, "v1", "", 1) + "health")
		if err == nil && res.String() == "\"ok\"" {
			w.OnInvokeSuccess()
			break
		} else {
			w.OnInvokeFault()
		}
		i++
	}
	url := item + "/run/tryRun"
	res, err := client.R().SetBody(body).Post(url)
	if err != nil {
		log.Errorv("Post error", zap.Error(err))
	}
	if gjson.Get(res.String(), "data.correct").Bool() {
		return true
	}
	return false
}
