package handlers

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"lumi/internal/util"
	"lumi/pkg/database"
	"lumi/pkg/logger"
	"lumi/pkg/model"
)

func RegisterHandler(c fiber.Ctx) (err error) {
	var user model.User
	util.BodyParser(&c, &user)
	logger.Log.Info("接收到注册请求: %v", zap.Any("user", user))
	db := database.DB
	if err := db.Create(&user).Error; err != nil {
		logger.Log.Error("注册失败: %v", zap.Error(err))
		return err
	}
	logger.Log.Info("注册成功: %v", zap.Any("user", user))
	c.JSON("注册成功")

	return nil
}
