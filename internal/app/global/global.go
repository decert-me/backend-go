package global

import (
	"backend-go/internal/app/config"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB            *gorm.DB               // 数据库链接
	LOG           *zap.Logger            // 日志框架
	CONFIG        config.Server          // 配置信息
	ContractEvent map[common.Hash]string // 合约事件
)
