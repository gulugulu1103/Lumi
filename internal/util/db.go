package util

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

func init() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	dsn := os.Getenv("LUMI_DSN")
	if dsn == "" {
		log.Println("No DSN provided")
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别
	})

	if err != nil {
		log.Println("Failed to connect to database:", err)
	}

	// 迁移数据库模型
	err = DB.AutoMigrate(&ExampleModel{})
	if err != nil {
		log.Println("Failed to auto-migrate:", err)
	}

	return
}
