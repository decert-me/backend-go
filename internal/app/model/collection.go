package model

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
)

// Collection 合辑
type Collection struct {
	ID               uint           `gorm:"primarykey" json:"id"`
	UUID             string         `gorm:"column:uuid" json:"uuid"`
	Title            string         `gorm:"column:title;comment:标题;type:varchar" json:"title" form:"title"` // 标题
	Label            string         `gorm:"column:label;comment:标签;type:varchar" json:"-"`                  // 标签
	Disabled         bool           `gorm:"column:disabled" json:"-"`
	Description      string         `gorm:"column:description;type:varchar" json:"description" form:"description"`
	Dependencies     datatypes.JSON `gorm:"column:dependencies" json:"-"`
	IsDraft          bool           `gorm:"column:is_draft" json:"-"`
	AddTs            int64          `gorm:"column:add_ts;autoCreateTime" json:"addTs"`
	TokenId          string         `gorm:"column:token_id;default:'';type:varchar(100)" json:"tokenId"`
	Type             uint8          `gorm:"column:type" json:"type" form:"type"`                       // 0:问答;1:编程
	Difficulty       *uint8         `gorm:"column:difficulty" json:"difficulty"`                       // 0:easy;1:moderate;2:difficult
	EstimateTime     uint8          `gorm:"column:estimate_time" json:"-"`                             // 预估时间/min
	Creator          string         `gorm:"column:creator;type:varchar" json:"creator" form:"creator"` // 用户 address
	MetaData         datatypes.JSON `gorm:"column:meta_data" json:"metadata"`                          // 元数据
	QuestData        datatypes.JSON `gorm:"column:quest_data" json:"quest_data"`                       // 元数据
	ExtraData        datatypes.JSON `gorm:"column:extra_data" json:"-"`                                // 额外数据
	Uri              string         `gorm:"column:uri" json:"uri"`
	PassScore        int64          `gorm:"column:pass_score" form:"pass_score" json:"pass_score"`    // 通过分数
	TotalScore       int64          `gorm:"column:total_score" form:"total_score" json:"total_score"` // 总分
	Recommend        string         `gorm:"column:recommend;type:text" json:"recommend"`              // 推荐
	Status           uint8          `gorm:"column:status;default:1" json:"status"`                    // 状态 1 上架 2 未上架
	Style            int            `gorm:"style;default:1" json:"style" `                            // 1:单独;2:合辑
	Cover            string         `gorm:"column:cover;comment:封面图" json:"cover"`
	Author           string         `gorm:"column:author;type:varchar(64);comment:合辑作者" json:"author"`
	Sort             int            `gorm:"column:sort;default:0" json:"sort"` // 	排序
	CollectionStatus uint8          `gorm:"column:collection_status;default:0" json:"-"`
	ChainID          int64          `gorm:"column:chain_id;comment:链ID;default:0" json:"chain_id"`              // 链ID
	Category         pq.Int64Array  `gorm:"column:category;type:int8[];comment:分类标签" json:"category,omitempty"` // 分类标签
}

func (Collection) TableName() string {
	return "collection"
}
