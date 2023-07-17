package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/pkg/log"
	"go.uber.org/zap"
)

func (s *Service) GetProgress(userID uint, req request.GetProgressRequest) (res model.ReadProgress, err error) {
	// 查询是否存在
	exists, change, err := s.dao.ProgressExists(userID, req)
	if err != nil {
		log.Errorv("ProgressExists error", zap.Error(err))
	}
	// 保存数据
	if !exists {
		if err = s.dao.CreateProgress(userID, req); err != nil {
			log.Errorv("CreateProgress error", zap.Error(err))
			return
		}
	}
	if len(req.Data) == 0 {
		return s.dao.GetProgress(userID, req.CatalogueName)
	}
	// 修改数据
	if change {
		if err = s.dao.ChangeProgress(userID, req); err != nil {
			log.Errorv("UpdateProgress error", zap.Error(err))
			return
		}
	}
	// 返回数据
	return s.dao.GetProgress(userID, req.CatalogueName)
}

func (s *Service) UpdateProgress(userID uint, req request.UpdateProgressRequest) (err error) {
	return s.dao.UpdateProgress(userID, req)
}
