package dao

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/pkg/log"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"math"
)

func (d *Dao) ProgressExists(userID uint, req request.GetProgressRequest) (exists bool, change bool, err error) {
	var progress model.ReadProgress
	err = d.db.Model(&model.ReadProgress{}).Where("user_id = ? AND catalogue_name = ?", userID, req.CatalogueName).First(&progress).Error
	if err != nil {
		return exists, change, nil
	}
	jsonData, err := json.Marshal(req.Data)
	if err != nil {
		log.Errorv("json.Marshal error", zap.Error(err))
		return exists, change, err
	}
	// 计算JSON数据的MD5哈希值
	hash := md5.Sum(jsonData)
	md5Hash := hex.EncodeToString(hash[:])
	if md5Hash != progress.Hash {
		change = true
	}
	return true, change, err
}

func (d *Dao) CreateProgress(userID uint, req request.GetProgressRequest) error {
	// Data
	jsonData, err := json.Marshal(req.Data)
	if err != nil {
		log.Errorv("json.Marshal error", zap.Error(err))
		return err
	}
	// 计算JSON数据的MD5哈希值
	hash := md5.Sum(jsonData)
	md5Hash := hex.EncodeToString(hash[:])
	// 创建数据
	err = d.db.Model(&model.ReadProgress{}).Create(&model.ReadProgress{
		UserID:        userID,
		CatalogueName: req.CatalogueName,
		Data:          jsonData,
		Hash:          md5Hash,
	}).Error
	return err
}

func (d *Dao) GetProgress(userID uint, catalogueName string) (req model.ReadProgress, err error) {
	err = d.db.Model(&model.ReadProgress{}).Where("user_id = ? AND catalogue_name = ?", userID, catalogueName).First(&req).Error
	return
}

func (d *Dao) ChangeProgress(userID uint, req request.GetProgressRequest) (err error) {
	var progress model.ReadProgress
	err = d.db.Model(&model.ReadProgress{}).Where("user_id = ? AND catalogue_name = ?", userID, req.CatalogueName).First(&progress).Error
	if err != nil {
		return err
	}
	// Data
	var data []request.ProgressData
	err = json.Unmarshal(progress.Data, &data)
	if err != nil {
		log.Errorv("json.Unmarshal error", zap.Error(err))
		return err
	}
	// 比较数据
	temp := req.Data
	for i, v := range req.Data {
		for _, v2 := range data {
			if v.DocId == v2.DocId {
				temp[i].IsFinish = v2.IsFinish
			}
		}
	}
	// Data
	jsonData, err := json.Marshal(temp)
	if err != nil {
		log.Errorv("json.Marshal error", zap.Error(err))
		return err
	}
	// 计算JSON数据的MD5哈希值
	hash := md5.Sum(jsonData)
	md5Hash := hex.EncodeToString(hash[:])
	// 更新数据
	err = d.db.Model(&model.ReadProgress{}).Where("user_id = ? AND catalogue_name = ?", userID, req.CatalogueName).Updates(&model.ReadProgress{
		Data: jsonData,
		Hash: md5Hash,
	}).Error
	return err
}

func (d *Dao) UpdateProgress(userID uint, req request.UpdateProgressRequest) (err error) {
	var progress model.ReadProgress
	err = d.db.Model(&model.ReadProgress{}).Where("user_id = ? AND catalogue_name = ?", userID, req.CatalogueName).First(&progress).Error
	if err != nil {
		return err
	}
	// Data
	var data []request.ProgressData
	err = json.Unmarshal(progress.Data, &data)
	if err != nil {
		log.Errorv("json.Unmarshal error", zap.Error(err))
		return err
	}
	// 比较数据
	temp := data
	for i, v := range data {
		for _, v2 := range req.Data {
			if v.DocId == v2.DocId {
				if v2.IsFinish == false {
					break
				}
				temp[i].IsFinish = v2.IsFinish
			}
		}
	}
	// Data
	jsonData, err := json.Marshal(temp)
	if err != nil {
		log.Errorv("json.Marshal error", zap.Error(err))
		return err
	}
	// 计算JSON数据的MD5哈希值
	hash := md5.Sum(jsonData)
	md5Hash := hex.EncodeToString(hash[:])
	// 更新数据
	err = d.db.Model(&model.ReadProgress{}).Where("user_id = ? AND catalogue_name = ?", userID, req.CatalogueName).Updates(&model.ReadProgress{
		Data: jsonData,
		Hash: md5Hash,
	}).Error
	return err
}

// GetProgressList 获取列表
func (d *Dao) GetProgressList(userID uint, catalogueNameList []string) (res []response.GetProgressListRes, err error) {
	for _, catalogueName := range catalogueNameList {
		// 获取阅读人数
		var readNum int64
		if err := d.db.Model(&model.ReadProgress{}).Where("catalogue_name = ?", catalogueName).Count(&readNum).Error; err != nil {
			log.Errorv("Count error", zap.Error(err))
		}
		// 	获取课程数量
		var docNum int
		var readProgress model.ReadProgress
		if err := d.db.Model(&model.ReadProgress{}).Where("catalogue_name = ? AND data != 'null'", catalogueName).Order("created_at desc").First(&readProgress).Error; err != nil {
			log.Errorv("First error", zap.Error(err))
		}
		docNum = len(gjson.Get(string(readProgress.Data), "@this").Array())
		// 阅读进度
		var percent float64
		var userReadProgress model.ReadProgress
		if userID != 0 {
			if err := d.db.Model(&model.ReadProgress{}).Where("user_id = ? AND catalogue_name = ?", userID, catalogueName).First(&userReadProgress).Error; err != nil {
				log.Errorv("First error", zap.Error(err))
			}
		}
		finishArr := gjson.Get(string(userReadProgress.Data), "#.is_finish").Array()
		total := len(finishArr)
		var read int
		for _, v := range finishArr {
			if v.Bool() == true {
				read++
			}
		}
		temp := float64(read) / float64(total)
		percent = math.Round(temp*100) / 100
		res = append(res, response.GetProgressListRes{
			CatalogueName: catalogueName,
			ReadNum:       readNum,
			DocNum:        docNum,
			Percent:       percent,
		})
	}
	return res, nil
}
