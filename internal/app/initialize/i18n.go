package initialize

import (
	"backend-go/internal/app/assets"
	"backend-go/internal/app/config"
	"backend-go/pkg/log"
	"encoding/json"
	"go.uber.org/zap"
	"io/fs"
)

// InitI18n loads messages from json
func InitI18n(c *config.Config) (res map[string]map[string]string) {
	res = make(map[string]map[string]string)

	data, err := fs.ReadFile(assets.Assets, "locale.json")
	if err != nil {
		log.Errorv("InitI18n ReadFile error", zap.String("path", c.System.I18n), zap.Error(err))
		panic("InitI18n ReadFile error")
		return res
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Errorv("InitI18n Unmarshal error", zap.Error(err))
		return res
	}

	return res
}
