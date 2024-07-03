package service

import (
	"backend-go/internal/app/model"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
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

func (s *Service) GetProgressList(userID uint, req request.GetProgressListRequest) (res []response.GetProgressListRes, err error) {
	return s.dao.GetProgressList(userID, req.CatalogueNameList)
}

func (s *Service) GetTutorialList(info request.GetTutorialListStatusRequest) (list interface{}, total int64, err error) {
	return s.dao.GetTutorialList(info)

}

// LabelLangList 获取语言列表
func (s *Service) LabelLangList() (language []model.Language, err error) {
	return s.dao.LabelLangList()
}

// LabelCategoryList 获取分类标签列表
func (s *Service) LabelCategoryList(class string) (res []model.Category, err error) {
	category, err := s.dao.LabelCategoryList()
	if err != nil {
		return nil, err
	}
	// 查询 category 数量
	var categoryList []model.Category
	for i, v := range category {
		exist, err := s.IsExistLabelCategory(v.ID, class)
		if err != nil {
			return nil, err
		}
		if exist {
			categoryList = append(categoryList, category[i])
		}
	}

	return categoryList, nil
}

// IsExistLabelCategory 按照分类判断分类标签是否存在
func (s *Service) IsExistLabelCategory(categoryID uint, class string) (exist bool, err error) {
	if class == "quest" {
		return s.dao.IsExistQuestByCatalogueID(categoryID)
	}
	return s.dao.IsExistTutorialByCatalogueID(categoryID)
}

// LabelThemeList 获取分类标签列表
func (s *Service) LabelThemeList() (theme []model.Theme, err error) {
	return s.dao.LabelThemeList()
}
