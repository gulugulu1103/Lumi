package main

import (
	"Lumi/internal/util"
	"Lumi/pkg/logger"
	"Lumi/pkg/model"
	"log"
)

// 导入dsn
func init() {

}

func main() {
	logger.Info("原神！启动！")
	util.AutoMigrate()
	user := &model.User{}
	log.Print(user)
}
