package database

import (
	"fmt"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"lumi/pkg/logger"
	"os"
)

var RegisterModels []interface{}

var DB *gorm.DB

type lumiNameStrategy struct {
	schema.NamingStrategy
}

// TableName 将逻辑表名映射到物理表名。
func (l lumiNameStrategy) TableName(table string) string {
	return "e_" + l.NamingStrategy.TableName(table) // 给实体表名添加前缀 "e_"
}

// JoinTableName 生成多对多关系的联接表名。
func (l lumiNameStrategy) JoinTableName(relation string) string {
	return "r_" + l.NamingStrategy.JoinTableName(relation) // 给联接表名添加前缀 "r_"
}

func init() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Warn("无法加载 .env 文件", zap.Error(err))
	}

	dsn := os.Getenv("LUMI_MYSQL_DSN")
	if dsn == "" {
		log.Error("LUMI_MYSQL_DSN 未设置")
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false, // 禁用默认事务
		NamingStrategy: lumiNameStrategy{schema.NamingStrategy{
			SingularTable: true,
		}}, // 设置命名策略
		Logger:      gorm_logger.Default.LogMode(gorm_logger.Info), // 设置日志级别
		PrepareStmt: true,                                          // 启用预编译语句
	})

	if err != nil {
		logger.Log.Error("数据库连接失败", zap.Error(err))
	}

	return
}

func AutoMigrate() {
	logger.Log.Info("正在执行数据库迁移...", zap.Int("models", len(RegisterModels)))
	for _, model := range RegisterModels {
		modelType := fmt.Sprintf("%T", model) // 获取模型的类型名称
		logger.Log.Info("正在迁移模型:", zap.String("model", modelType))
		if err := DB.AutoMigrate(model); err != nil {
			logger.Log.Error("模型迁移失败", zap.String("model", modelType), zap.Error(err))
		}
	}
	logger.Log.Info("数据库迁移完成")
}
