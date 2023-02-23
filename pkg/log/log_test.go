package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestMain(m *testing.M) {
	Init(&Config{
		Level:       "info",
		Save:        false,
		ShowLine:    true,
		Format:      "json",
		EncodeLevel: "CapitalLevelEncoder",
	})
	m.Run()
}
func TestLog(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		Error("hello %s", "world")
		Errorv(zap.Any("key", 2222222), zap.Any("test2", "test"))
	})
	t.Run("Warn", func(t *testing.T) {
		Warn("hello %s", "world")
		Warnv(zap.Any("key", 2222222), zap.Any("test2", "test"))
	})
	t.Run("Info", func(t *testing.T) {
		Info("hello %s", "world")
		Infov(zap.Any("key", 2222222), zap.Any("test2", "test"))
	})
}
