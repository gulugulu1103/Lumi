package api

import "github.com/gofiber/fiber/v3"

func main() {
	// 创建一个新的 Fiber 实例
	app := fiber.New()

	// 添加一个路由处理函数，当访问根 URL ("/") 时调用
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!") // 发送响应到客户端
	})

	// 启动 HTTP 服务器在 3000 端口
	app.Listen(":3000")
}
