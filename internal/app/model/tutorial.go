package model

import (
	"github.com/lib/pq"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Tutorial struct {
	gorm.Model
	RepoUrl       string         `json:"repoUrl,omitempty"`
	Label         string         `json:"label,omitempty"`                                                    // 教程名称
	Category      pq.Int64Array  `gorm:"column:category;type:int8[];comment:分类标签" json:"category,omitempty"` // 分类标签
	Theme         pq.Int64Array  `gorm:"column:theme;type:int8[];comment:主题标签" json:"theme,omitempty"`       // 主题标签
	Language      uint           `json:"language,omitempty"`                                                 // 语言
	CatalogueName string         `gorm:"column:catalogue_name;UNIQUE;not null;" json:"catalogueName,omitempty"`
	DocType       string         `json:"docType,omitempty"` // 媒体类型
	Img           string         `json:"img,omitempty"`     // 教程封面图
	Desc          string         `json:"desc,omitempty"`
	Branch        *string        `json:"branch,omitempty"`
	DocPath       *string        `json:"docPath,omitempty"`
	StartPage     string         `json:"startPage,omitempty"`
	CommitHash    *string        `json:"commitHash,omitempty"`
	VideoCategory string         `json:"videoCategory,omitempty"`
	Video         datatypes.JSON `json:"video,omitempty"`                                            // 视频排序
	Sort          pq.StringArray `gorm:"column:sort;type:text[];comment:视频排序" json:"sort,omitempty"` // 视频排序
	Url           string         `json:"url,omitempty"`
	EstimateTime  uint           `json:"estimateTime,omitempty"`                                    // 预估时间
	Challenge     *string        `json:"challenge,omitempty"`                                       // 挑战
	VisitNum      uint           `json:"visitNum,omitempty"`                                        // 教程浏览量
	AddrNum       uint           `json:"addrNum,omitempty"`                                         // 参与人员数量
	Order         int            `json:"order"`                                                     // 排序
	Difficulty    uint8          `json:"difficulty"`                                                // 难度
	Status        uint8          `gorm:"column:status;default:1" json:"status,omitempty"`           // 状态 1 未上架 2 已上架
	PackStatus    uint8          `gorm:"column:pack_status;default:1" json:"pack_status,omitempty"` // 状态 1 未打包 2 打包成功 3 打包失败
	PackLog       string         `gorm:"column:pack_log;type:text" json:"pack_log"`                 // 打包日志
	Top           *bool          `gorm:"column:top;default:false" json:"top"`                       // 是否置顶
	TutorialSort  *int           `gorm:"column:tutorial_sort;default:0" json:"tutorial_sort"`       // 教程排序
}

func (Tutorial) TableName() string {
	return "admin_tutorial"
}
