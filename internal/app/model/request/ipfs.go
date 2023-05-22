package request

import "time"

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
		Title   string   `json:"title" binding:"required"`
		Type    string   `json:"type" binding:"required"`
		Score   int      `json:"score" binding:"required"`
		Options []string `json:"options,omitempty"`
	} `json:"questions" binding:"required"`
	Answers      string    `json:"answers" binding:"required"`
	StartTime    time.Time `json:"startTime"`
	EndTIme      any       `json:"endTIme"`
	EstimateTime any       `json:"estimateTime"`
	PassingScore int       `json:"passingScore"`
	Version      float64   `json:"version" binding:"required"`
}
