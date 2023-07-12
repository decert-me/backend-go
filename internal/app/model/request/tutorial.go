package request

type ProgressData struct {
	DocId    string `json:"docId"`
	IsFinish bool   `json:"is_finish"`
}
type GetProgressRequest struct {
	CatalogueName string         `json:"catalogueName"  binding:"required"`
	Data          []ProgressData `json:"data"  binding:"required"`
}

type UpdateProgressRequest struct {
	CatalogueName string         `json:"catalogueName"  binding:"required"`
	Data          []ProgressData `json:"data"  binding:"required"`
}
