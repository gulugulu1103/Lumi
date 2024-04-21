package api

import (
	"github.com/gofiber/fiber/v3"
	"lumi/pkg/api/handlers/user"
)

func Serve() {
	// 创建一个新的 Fiber 实例
	app := fiber.New()

	// 添加一个路由处理函数，当访问根 URL ("/") 时调用
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// 登录模块
	auth := v1.Group("/auth")
	auth.Post("/register", user.RegisterHandler)
	auth.Delete("/user", user.DeleteUserHandler)

	// 启动 HTTP 服务器在 3000 端口
	app.Listen(":3000")
}
