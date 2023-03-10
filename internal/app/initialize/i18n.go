package initialize

import (
	"backend-go/internal/app/config"
	"backend-go/pkg/log"
	"encoding/json"
	"go.uber.org/zap"
	"os"
)

// InitI18n loads messages from json
func InitI18n(c *config.Config) (res map[string]map[string]string) {
	res = make(map[string]map[string]string)

	data, err := os.ReadFile(c.System.I18n)
	if err != nil {
		log.Errorv("InitI18n ReadFile error", zap.String("path", c.System.I18n), zap.Error(err))
		return res
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Errorv("InitI18n Unmarshal error", zap.Error(err))
		return res
	}

	return res
}
