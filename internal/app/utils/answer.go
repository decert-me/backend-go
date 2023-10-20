package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/tidwall/gjson"
)

// xorStrings ...
func xorStrings(key string, input string) string {
	output := ""
	runeArray := []rune(input)
	for i := 0; i < len(runeArray); i++ {
		c := runeArray[i]
		k := int32(key[i%len(key)])
		output += fmt.Sprintf("%c", c^k)
	}
	return output
}

// AnswerEncode 加密答案
func AnswerEncode(key, data string) string {
	if key == "" {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(xorStrings(key, data)))
}

// AnswerDecode 解密答案
func AnswerDecode(key, data string) string {
	if key == "" {
		return ""
	}
	str, _ := base64.StdEncoding.DecodeString(data)
	return xorStrings(key, string(str))
}

func GetAnswers(version float64, key, res, questData, answerUser string) (answerU, scoreList, answerS []gjson.Result, passingScore int64) {
	if version == 1 {
		answerU = gjson.Get(AnswerDecode(key, gjson.Get(res, "properties.answers").String()), "@this").Array() // 标准答案
		passingScore = gjson.Get(res, "properties.passingScore").Int()                                         // 通过分数
		scoreList = gjson.Get(res, "properties.questions.#.score").Array()                                     // 题目分数
		answerS = gjson.Get(answerUser, "@this").Array()                                                       // 用户答案
	} else if version == 1.1 || version == 1.2 {
		answerU = gjson.Get(AnswerDecode(key, gjson.Get(questData, "answers").String()), "@this").Array() // 标准答案
		passingScore = gjson.Get(questData, "passingScore").Int()                                         // 通过分数
		scoreList = gjson.Get(questData, "questions.#.score").Array()                                     // 题目分数
		answerS = gjson.Get(answerUser, "@this").Array()                                                  // 用户答案
	}
	return
}
