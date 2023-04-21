package utils

import (
	"encoding/base64"
	"fmt"
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
