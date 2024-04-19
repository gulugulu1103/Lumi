package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var log *zap.Logger

func init() {
	log = newLogger()
}

// newLogger 创建并配置一个 zap.Logger 实例
func newLogger() *zap.Logger {
	// 配置 zap
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.MessageKey = "message"

	// 设置日志输出位置
	logPath := "app.log"
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		file, err := os.Create(logPath)
		if err != nil {
			panic(err)
		}
		file.Close()
	}
	config.OutputPaths = []string{logPath, "stderr"}

	// 初始化 logger
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

// Info logs a message at level Info on the standard logger.
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Error logs a message at level Error on the standard logger.
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(message string, fields ...zap.Field) {
	log.Warn(message, fields...)
}
