package initialize

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/utils"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLog(c *config.Config) (logger *zap.Logger) {
	if ok, _ := utils.PathExists(c.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", c.Zap.Director)
		_ = os.Mkdir(c.Zap.Director, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	now := time.Now().Format("2006-01-02")

	cores := [...]zapcore.Core{
		getEncoderCore(c, fmt.Sprintf("./%s/%s/debug.log", c.Zap.Director, now), debugPriority),
		getEncoderCore(c, fmt.Sprintf("./%s/%s/info.log", c.Zap.Director, now), infoPriority),
		getEncoderCore(c, fmt.Sprintf("./%s/%s/warn.log", c.Zap.Director, now), warnPriority),
		getEncoderCore(c, fmt.Sprintf("./%s/%s/error.log", c.Zap.Director, now), errorPriority),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if c.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig(c *config.Config) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  c.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case c.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case c.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case c.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case c.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder(c *config.Config) zapcore.Encoder {
	if c.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig(c))
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig(c))
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(c *config.Config, fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := utils.GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(c), writer, level)
}

// customTimeEncoder 自定义日志输出时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05.000"))
}
