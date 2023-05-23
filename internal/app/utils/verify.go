package utils

import "backend-go/internal/app/model/request"

func VerifyUploadJSONChallenge(uploadJSONChallenge request.UploadJSONChallenge) bool {
	// 校验题目正确性
	for _, quest := range uploadJSONChallenge.Questions {
		// 多选题
		if quest.Type == "multiple_response" {
			if len(quest.Options) < 2 {
				return false
			}
		}
		// 单选题
		if quest.Type == "multiple_choice" {
			if len(quest.Options) != 1 {
				return false
			}
		}
		// 填空题
		if quest.Type == "fill_blank" {
			if quest.Options[0] == "" {
				return false
			}
		}
		// 普通编程题
		if quest.Type == "coding" {
			if len(quest.Input) != len(quest.Output) {
				return false
			}
		}
		// 特殊编程题
		if quest.Type == "special_judge_coding" {
			if quest.SpjCode == "" {
				return false
			}
		}
	}
	return true
}
