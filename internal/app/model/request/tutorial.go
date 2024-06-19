package request

import "backend-go/internal/app/model"

type ProgressData struct {
	DocId    string `json:"docId"`
	IsFinish bool   `json:"is_finish"`
}
type GetProgressRequest struct {
	CatalogueName string         `json:"catalogueName"  binding:"required"`
	Data          []ProgressData `json:"data"`
}

type UpdateProgressRequest struct {
	CatalogueName string         `json:"catalogueName"  binding:"required"`
	Data          []ProgressData `json:"data"  binding:"required"`
}

type GetProgressListRequest struct {
	CatalogueNameList []string `json:"catalogueNameList"  binding:"required"`
}

type GetTutorialListStatusRequest struct {
	PageInfo
	SearchKey string `json:"search_key"`
	model.Tutorial
}

type GetLabelRequest struct {
	Type string `json:"type" binding:"required"`
}
