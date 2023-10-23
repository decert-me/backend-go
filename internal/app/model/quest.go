package model

import "gorm.io/datatypes"

type Quest struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	UUID             string         `gorm:"column:uuid" json:"uuid"`
	Title            string         `gorm:"column:title;comment:标题;type:varchar" json:"title" form:"title"` // 标题
	Label            string         `gorm:"column:label;comment:标签;type:varchar" json:"-"`                  // 标签
	Disabled         bool           `gorm:"column:disabled" json:"-"`
	Description      string         `gorm:"column:description;type:varchar" json:"description" form:"description"`
	Dependencies     datatypes.JSON `gorm:"column:dependencies" json:"-"`
	IsDraft          bool           `gorm:"column:is_draft" json:"-"`
	AddTs            int64          `gorm:"column:add_ts;autoCreateTime" json:"addTs"`
	TokenId          int64          `gorm:"column:token_id;UNIQUE;not null;" json:"tokenId"`
	Type             uint8          `gorm:"column:type;comment:0:问答 1:编程" json:"type" form:"type"`                     // 0:问答;1:编程
	Difficulty       uint8          `gorm:"column:difficulty;comment:0 easy 1 moderate 2 difficult" json:"difficulty"` // 0:easy;1:moderate;2:difficult
	EstimateTime     uint8          `gorm:"column:estimate_time;comment:预估时间/min" json:"estimate_time"`                // 预估时间/min
	Creator          string         `gorm:"column:creator;type:varchar;comment:创建者地址" json:"creator" form:"creator"`   // 用户 address
	MetaData         datatypes.JSON `gorm:"column:meta_data;comment:元数据" json:"metadata"`                              // 元数据
	QuestData        datatypes.JSON `gorm:"column:quest_data;comment:挑战数据" json:"quest_data"`                          // 挑战数据
	ExtraData        datatypes.JSON `gorm:"column:extra_data;comment:额外数据" json:"-"`                                   // 额外数据
	Uri              string         `gorm:"column:uri;comment:uri" json:"uri"`
	PassScore        int64          `gorm:"column:pass_score;comment:通过分数" form:"pass_score" json:"pass_score"`  // 通过分数
	TotalScore       int64          `gorm:"column:total_score;comment:总分" form:"total_score" json:"total_score"` // 总分
	Recommend        string         `gorm:"column:recommend;type:text;comment:推荐" json:"recommend"`              // 推荐
	Status           uint8          `gorm:"column:status;default:1;comment:状态 1 上架 2 未上架" json:"status"`         // 状态 1 上架 2 未上架
	Style            int            `gorm:"style;default:1;comment:1 单独 2 合辑" json:"style" `                     // 1:单独;2:合辑
	Cover            string         `gorm:"column:cover;comment:封面图" json:"cover"`
	Author           string         `gorm:"column:author;type:varchar(64);comment:合辑作者" json:"author"`
	Sort             int            `gorm:"column:sort;default:0;comment:排序" json:"sort"`                       // 	排序
	CollectionStatus uint8          `gorm:"column:collection_status;default:1;comment:合辑状态 1 独立 2 合辑" json:"-"` // 	合辑状态 1 独立 2 合辑
}
