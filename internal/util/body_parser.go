package util

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
)

func BodyParser(c *fiber.Ctx, out interface{}) error {
	return sonic.Unmarshal((*c).Body(), out)
}
