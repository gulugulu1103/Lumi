package main

import (
	"Lumi/pkg/database"
	"Lumi/pkg/logger"
	"Lumi/pkg/model"
	"log"
)

// 导入dsn
func init() {

}

func main() {
	logger.Log.Info("原神！启动！")
	database.AutoMigrate()
	user := &model.User{}
	log.Print(user)
}
