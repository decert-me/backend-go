package request

import "time"

type SpjCodeList struct {
	Frame string `json:"frame"`
	Code  string `json:"code"`
}

type UploadJSONNFT struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Image       string `json:"image" binding:"required"`
	Attributes  struct {
		ChallengeIpfsURL string `json:"challenge_ipfs_url" binding:"required"`
		ChallengeURL     string `json:"challenge_url" binding:"required"`
		ChallengeTitle   string `json:"challenge_title" binding:"required"`
		Difficulty       any    `json:"difficulty"`
	} `json:"attributes" binding:"required"`
	ExternalURL string  `json:"external_url" binding:"required"`
	Version     float64 `json:"version" binding:"required"`
}

type UploadJSONChallenge struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Creator     string `json:"creator" binding:"required"`
	Content     string `json:"content"`
	Questions   []struct {
		Input        []string `json:"input,omitempty"`
		Output       []string `json:"output,omitempty"`
		Title        string   `json:"title" binding:"required"`
		Description  string   `json:"description,omitempty"`
		Type         string   `json:"type" binding:"required"`
		Score        int      `json:"score" binding:"required"`
		Languages    []string `json:"languages,omitempty"`
		CodeSnippets []struct {
			Lang          string `json:"lang"`
			Code          string `json:"code"`
			CorrectAnswer string `json:"correctAnswer"`
		} `json:"code_snippets,omitempty"`
		SpjCode []SpjCodeList `json:"spj_code,omitempty"`
		Options []string      `json:"options,omitempty"`
	} `json:"questions" binding:"required"`
	Answers      string    `json:"answers" binding:"required"`
	StartTime    time.Time `json:"startTime"`
	EndTIme      any       `json:"endTIme"`
	EstimateTime any       `json:"estimateTime"`
	PassingScore int       `json:"passingScore"`
	Version      float64   `json:"version" binding:"required"`
}
