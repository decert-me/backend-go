package utils

import (
	"backend-go/internal/app/model/request"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
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

// HashData 生成校验Hash
func HashData(data interface{}, key string) (timestamp int64, hashValue string) {
	// 将结构体数据转换为 JSON 格式的字符串
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	// 获取当前时间戳
	timestamp = time.Now().Unix()

	// 将时间戳与字符串拼接
	hashData := string(jsonData) + key + fmt.Sprintf("%d", timestamp)
	// 对字符串进行哈希计算
	hasher := sha256.New()
	hasher.Write([]byte(hashData))
	hashValue = hex.EncodeToString(hasher.Sum(nil))
	return timestamp, hashValue
}

// VerifyData 校验数据正确性
func VerifyData(data interface{}, key, hash, timestampStr string) (verify bool) {
	// 将结构体数据转换为 JSON 格式的字符串
	jsonData, err := json.Marshal(data)
	if err != nil {
		return
	}
	// 类型转换
	timestamp, err := strconv.Atoi(timestampStr)
	if err != nil {
		return
	}
	// 将时间戳与 JSON 字符串拼接
	hashData := string(jsonData) + key + fmt.Sprintf("%d", timestamp)
	// 对拼接后的字符串进行哈希计算
	hasher := sha256.New()
	hasher.Write([]byte(hashData))
	hashValue := hex.EncodeToString(hasher.Sum(nil))
	if hashValue == hash {
		return true
	}
	return false
}
