package logger

import (
	"Lumi/internal/util"
	"github.com/gofiber/fiber/v3/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Log *zap.Logger

func init() {
	Log = newLogger()
}

// newLogger 创建并配置一个 zap.Logger 实例
func newLogger() *zap.Logger {
	// 配置 zap
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.TimeKey = time.DateTime
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

	// 挂载mongo写入
	logger = logger.WithOptions(
		zap.Hooks(
			func(entry zapcore.Entry) error {
				_, err := util.MongoDB.Collection("logs").InsertOne(nil, entry)
				if err != nil {
					log.Errorf("日志写入mongo失败", err.Error())
				}
				return err
			}))

	return logger
}
