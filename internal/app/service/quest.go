package service

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
)

func GetQuestList(searchInfo request.GetQuestListRequest) (questList []response.GetQuestListRes, err error) {
	var quest []model.Quest
	db := global.DB.Model(&model.Quest{}).Where(searchInfo.Quest)
	if err != nil {
		return
	} else {
		err = db.Order("id desc").Find(&quest).Error
	}
	for _, v := range quest {
		questList = append(questList, response.GetQuestListRes{Quest: v})
	}
	return
}
