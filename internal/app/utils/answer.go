package utils

import (
	"encoding/base64"
	"errors"
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
func answerEncode(key, data string) string {
	if key == "" {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(xorStrings(key, data)))
}

// AnswerDecode 解密答案
func answerDecode(key, data string) string {
	if key == "" {
		return ""
	}
	str, _ := base64.StdEncoding.DecodeString(data)
	return xorStrings(key, string(str))
}

func AnswerCheck(key, answerUser, uri string, userScore int64) (pass bool, err error) {
	res, err := GetDataFromUri(uri)
	if err != nil || !gjson.Valid(res) {
		return pass, errors.New("tokenID invalid")
	}
	answerU := gjson.Get(answerDecode(key, gjson.Get(res, "properties.answers").String()), "@this").Array() // 标准答案
	passingScore := gjson.Get(res, "properties.passingScore").Int()                                         // 通过分数
	scoreList := gjson.Get(res, "properties.questions.#.score").Array()                                     // 题目分数
	answerS := gjson.Get(answerUser, "@this").Array()                                                       // 用户答案
	var totalScore int64
	for _, s := range scoreList {
		totalScore += s.Int()
	}
	if len(answerU) != len(answerS) || len(scoreList) != len(answerS) {
		return false, errors.New("unexpect error")
	}
	var score int64
	for i, _ := range answerS {
		if answerS[i].String() == answerU[i].String() {
			score += scoreList[i].Int()
		}
	}
	if userScore == (score*10000/totalScore) && score >= passingScore {
		return true, nil
	}
	return
}

func AnswerScore(key, answerUser, uri string) (userScore int64, pass bool, err error) {
	res, err := GetDataFromUri(uri)
	if err != nil || !gjson.Valid(res) {
		return userScore, pass, errors.New("tokenID invalid")
	}
	answerU := gjson.Get(answerDecode(key, gjson.Get(res, "properties.answers").String()), "@this").Array() // 标准答案
	passingScore := gjson.Get(res, "properties.passingScore").Int()                                         // 通过分数
	scoreList := gjson.Get(res, "properties.questions.#.score").Array()                                     // 题目分数
	answerS := gjson.Get(answerUser, "@this").Array()                                                       // 用户答案
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
