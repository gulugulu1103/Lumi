package api

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"log"
	"lumi/pkg/api/handlers/user"
)

func Serve() {
	// 创建一个新的 Fiber 实例
	app := fiber.New(
		fiber.Config{
			AppName:         "lumi v0.1",
			JSONEncoder:     sonic.Marshal,
			JSONDecoder:     sonic.Unmarshal,
			StructValidator: nil,
			UnescapePath:    true,
		},
	)

	// 添加一个路由处理函数，当访问根 URL ("/") 时调用
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// 登录模块
	auth := v1.Group("/auth")
	auth.Post("/register", user.RegisterHandler)
	auth.Delete("/user", user.DeleteUserHandler)

	// 启动 HTTP 服务器在 3000 端口
	err := app.Listen(":3000", fiber.ListenConfig{
		EnablePrefork:     true,   // 开启协程池
		EnablePrintRoutes: true,   // 打印路由
		ListenerNetwork:   "tcp4", // 监听网络v4\v6
	})
	if err != nil {
		log.Fatal("启动 HTTP 服务器失败")
		return
	}
}
