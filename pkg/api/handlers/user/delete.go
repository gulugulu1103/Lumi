package user

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"lumi/pkg/database"
	"lumi/pkg/logger"
	"lumi/pkg/model"
)

func DeleteUserHandler(c fiber.Ctx) (err error) {
	name := c.Query("name")
	if name == "" {
		return c.JSON("用户名不能为空")
	}

	user := model.User{Username: name}
	if rows := database.DB.Where(&user).Delete(&user).RowsAffected; rows == 0 {
		return c.JSON("不存在该用户")
	}
	logger.Log.Info("已删除用户", zap.Any("user", user))
	return c.JSON("删除成功")
}
