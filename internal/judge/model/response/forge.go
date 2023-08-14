package response

type TestRes struct {
	Output         string `json:"output"`
	Status         uint8  `json:"status"` // 0 成功 1 异常
	TotalTestcases int    `json:"total_testcases"`
	TotalCorrect   int    `json:"total_correct"`
}
