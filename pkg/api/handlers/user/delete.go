package user

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"lumi/pkg/database"
	"lumi/pkg/logger"
	"lumi/pkg/model"
)

func DeleteUserHandler(c fiber.Ctx) (err error) {
	deletedID := fiber.Query[int](c, "id")
	if deletedID == 0 {
		return c.JSON("用户ID不能为空")
	}

	user := model.User{ID: uint(deletedID)}
	if rows := database.DB.Where(&user).Delete(&user).RowsAffected; rows == 0 {
		return c.JSON("不存在该用户")
	}
	logger.Log.Info("已删除用户", zap.Any("user", user))
	return c.JSON("删除成功")
}
