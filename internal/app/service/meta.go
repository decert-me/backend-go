package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/utils"
	"fmt"
	"github.com/tidwall/gjson"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (s *Service) HandleMetaRequest(_id string) []byte {
	fileContent, err := os.ReadFile(filepath.Join(s.c.System.Website, "index.html"))
	if err != nil {
		return fileContent
	}
	// 查询信息
	var quest model.Quest
	if utils.IsUUID(_id) {
		quest, err = s.dao.GetQuestByUUID(_id)
		if err != nil {
			return fileContent
		}
	} else {
		quest, err = s.dao.GetQuestByTokenID(_id)
		if err != nil {
			return fileContent
		}
	}
	// 需要替换的内容
	ipfs := "https://ipfs.decert.me/" + strings.Replace(gjson.Get(string(quest.MetaData), "image").String(), "ipfs://", "", 1)
	replaceList := map[string]string{
		"https://decert.me/": fmt.Sprintf("https://decert.me/quests/%s", _id),
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

func (s *Service) HandleCollectionMetaRequest(_id string) []byte {
	fileContent, err := os.ReadFile(filepath.Join(s.c.System.Website, "index.html"))
	if err != nil {
		return fileContent
	}
	// 查询信息
	id, err := strconv.Atoi(_id)
	if err != nil {
		return fileContent
	}
	collection, err := s.dao.GetCollectionByID(id)
	if err != nil {
		return fileContent
	}
	// 需要替换的内容
	ipfs := "https://ipfs.decert.me/" + collection.Cover
	replaceList := map[string]string{
		"https://decert.me/": fmt.Sprintf("https://decert.me/collection/%s", _id),
		"DeCert.Me":          collection.Title,
		"Decentralized skills learning and certification platform · Education in the Age of AI · SBT - Proof of Learn.": collection.Description,
		"/decert-social-card.png": ipfs,
	}
	// 修改文件内容
	temp := string(fileContent)
	for k, v := range replaceList {
		temp = strings.Replace(temp, k, v, -1)
	}
	return []byte(temp)
}
