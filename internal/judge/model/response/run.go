package response

type TryRunRes struct {
	Input          string `json:"input"`
	Output         string `json:"output"`
	ExceptOutput   string `json:"except_output"`
	Msg            string `json:"msg"`
	Gas            string `json:"gas"`
	Status         uint8  `json:"status"` // 状态码 1 编译失败 2 运行失败 3 运行成功
	Correct        bool   `json:"correct"`
	TotalTestcases int    `json:"total_testcases"`
	TotalCorrect   uint   `json:"total_correct"`
	LastInput      string `json:"last_input"`
	LastOutput     string `json:"last_output"`
	LastExpect     string `json:"last_expect"`
	JudgeID        string `json:"-"`
}

type TryTestRunRes struct {
	Msg            string `json:"msg"`
	Status         uint8  `json:"status"` // 状态码 1 编译失败 2 运行失败 3 运行成功
	Correct        bool   `json:"correct"`
	TotalTestcases int    `json:"total_testcases"`
	TotalCorrect   uint   `json:"total_correct"`
	LastInput      string `json:"last_input"`
	LastOutput     string `json:"last_output"`
}
