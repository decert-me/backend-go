package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// Config log config.
type Config struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                            // 级别
	Save          bool   `mapstructure:"save" json:"save" yaml:"save"`                               // 是否保存到文件
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                         // 输出
	LogInConsole  bool   `mapstructure:"log-in-console" json:"log-in-console" yaml:"log-in-console"` // 输出到控制台
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                  // 日志文件夹
	ShowLine      bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`                // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encode-level" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktrace-key" yaml:"stacktrace-key"` // 栈名
}

const (
	// common log filed.
	_log = "log"
)

var (
	l *zap.Logger
	c *Config
)

// Init create logger with context.
func Init(conf *Config) {
	c = conf
	if ok, _ := PathExists(conf.Director); !ok && c.Save == true { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", conf.Director)
		_ = os.Mkdir(conf.Director, os.ModePerm)
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
		return lev == zap.ErrorLevel
	})
	// Panic级别
	panicPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.DPanicLevel
	})

	now := time.Now().Format("2006-01-02")

	cores := [...]zapcore.Core{
		getEncoderCore(c, fmt.Sprintf("./%s/%s/debug.log", conf.Director, now), debugPriority),
		getEncoderCore(c, fmt.Sprintf("./%s/%s/info.log", conf.Director, now), infoPriority),
		getEncoderCore(c, fmt.Sprintf("./%s/%s/warn.log", conf.Director, now), warnPriority),
		getEncoderCore(c, fmt.Sprintf("./%s/%s/error.log", conf.Director, now), errorPriority),
		getEncoderCore(c, fmt.Sprintf("./%s/%s/panic.log", conf.Director, now), panicPriority),
	}
	l = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller(), zap.AddCallerSkip(1))
	if conf.ShowLine {
		l = l.WithOptions(zap.AddCaller())
	}
}

// Info logs a message at the info log level.
func Info(format string, args ...interface{}) {
	l.Log(zap.InfoLevel, fmt.Sprintf(format, args...))
}

// Warn logs a message at the warning log level.
func Warn(format string, args ...interface{}) {
	l.Log(zap.WarnLevel, fmt.Sprintf(format, args...))
}

// Error logs a message at the error log level.
func Error(format string, args ...interface{}) {
	l.Log(zap.ErrorLevel, fmt.Sprintf(format, args...))
}

// Infov logs a message at the info log level.
func Infov(message string, field ...zap.Field) {
	l.Log(zap.InfoLevel, message, field...)
}

// Warnv logs a message at the warning log level.
func Warnv(message string, field ...zap.Field) {
	l.Log(zap.WarnLevel, message, field...)
}

// Errorv logs a message at the error log level.
func Errorv(message string, field ...zap.Field) {
	l.Log(zap.ErrorLevel, message, field...)
}
