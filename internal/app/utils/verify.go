package utils

import (
	"backend-go/internal/app/model/request"
	"github.com/tidwall/gjson"
)

func VerifyUploadJSONChallenge(key string, uploadJSONChallenge request.UploadJSONChallenge) bool {
	answers := gjson.Parse(AnswerDecode(key, uploadJSONChallenge.Answers)).Array()
	// 校验题目正确性
	for i, quest := range uploadJSONChallenge.Questions {
		// 多选题
		if quest.Type == "multiple_response" {
			if len(answers[i].Array()) < 2 {
				return false
			}
		}
		// 单选题
		if quest.Type == "multiple_choice" {
			if len(answers[i].Array()) != 1 {
				return false
			}
		}
		// 填空题
		if quest.Type == "fill_blank" {
			if answers[i].String() == "" {
				return false
			}
		}
	}
	return true
}
