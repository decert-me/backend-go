package response

type GetProgressListRes struct {
	CatalogueName string  `json:"catalogueName"`
	ReadNum       int64   `json:"readNum"`
	DocNum        int     `json:"docNum"`
	Percent       float64 `json:"percent"`
}
