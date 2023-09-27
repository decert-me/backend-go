package service

import (
	"fmt"
	"github.com/tidwall/gjson"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (s *Service) HandleMetaRequest(_tokenID string) []byte {
	fileContent, err := os.ReadFile(filepath.Join(s.c.System.Website, "index.html"))
	if err != nil {
		return fileContent
	}
	// 查询信息
	tokenID, err := strconv.Atoi(_tokenID)
	if err != nil {
		return fileContent
	}
	quest, err := s.dao.GetQuestByTokenID(int64(tokenID))
	if err != nil {
		return fileContent
	}
	_ = quest
	// 需要替换的内容
	ipfs := "https://ipfs.io/ipfs/" + strings.Replace(gjson.Get(string(quest.MetaData), "image").String(), "ipfs://", "", 1)
	replaceList := map[string]string{
		"https://decert.me/": fmt.Sprintf("https://decert.me/quests/%s", _tokenID),
		"DeCert.Me":          quest.Title,
		"Decentralized skills learning and certification platform · Education in the Age of AI · SBT - Proof of Learn.": quest.Description,
		"/decert-social-card.png": ipfs,
	}
	// 修改文件内容
	temp := string(fileContent)
	for k, v := range replaceList {
		temp = strings.Replace(temp, k, v, -1)
	}
	return []byte(temp)
}
