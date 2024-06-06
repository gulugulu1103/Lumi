package user

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"lumi/pkg/database"
	"lumi/pkg/logger"
	"lumi/pkg/model"
)

func RegisterHandler(c fiber.Ctx) (err error) {
	var user model.User
	if err = c.Bind().JSON(&user); err != nil {
		logger.Log.Error("解析注册请求失败", zap.Error(err))
		return c.JSON("参数错误")
	}

	if err = user.RegisterValidate(); err != nil {
		logger.Log.Error("参数错误", zap.Error(err))
		return c.JSON("参数错误")
	}

	logger.Log.Info("接收到注册请求", zap.Any("user", user))
	db := database.DB
	tx := db.Begin() // 开启事务，否则ID会乱序
	if err = tx.Create(&user).Error; err != nil {
		tx.Rollback()
		logger.Log.Error("注册失败", zap.Error(err))
		return c.JSON("注册失败")
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		logger.Log.Error("注册失败，数据库出错", zap.Error(err))
		return c.JSON("注册失败，数据库出错")
	}
	logger.Log.Info("注册成功", zap.Any("user", user))

	return c.JSON("注册成功")
}
