package model

type JudgeResult struct {
	ID         string `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"id"` // 主键
	TokenID    int64  `json:"token_id"`
	QuestIndex uint8  `json:"quest_index"`
	ScoreRaw   int64  `gorm:"column:score_raw;default:0"`
	Pass       bool
}
