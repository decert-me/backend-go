package response

type GenerateShareResponse struct {
	ShareCode string `json:"share_code"`
}

type ShareCallbackResponse struct {
	ShareCode string `json:"share_code"`
	Params    string `json:"params"`
}

type ClickCallbackResponse struct {
	App       string `json:"app"`
	ShareCode string `json:"share_code"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}
