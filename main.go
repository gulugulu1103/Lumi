package main

import (
	"lumi/pkg/api"
	"lumi/pkg/logger"
)

// 导入dsn
func init() {

}

func main() {
	logger.Log.Info("原神！启动！")
	api.Serve()
}
