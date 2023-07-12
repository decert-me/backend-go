package request

type ProgressData struct {
	DocId    string `json:"docId"`
	IsFinish bool   `json:"is_finish"`
}
type GetProgressRequest struct {
	CatalogueName string         `json:"catalogueName"`
	Data          []ProgressData `json:"data"`
}

type UpdateProgressRequest struct {
	CatalogueName string         `json:"catalogueName"`
	Data          []ProgressData `json:"data"`
}
