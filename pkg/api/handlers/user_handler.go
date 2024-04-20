package handlers

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"lumi/pkg/database"
	"lumi/pkg/logger"
	"lumi/pkg/model"
)

func RegisterHandler(c fiber.Ctx) (err error) {
	var user model.User
	err = c.Bind().JSON(&user)
	if err != nil {
		logger.Log.Error("解析注册请求失败", zap.Error(err))
		c.JSON("参数错误")
		return err
	}
	logger.Log.Info("接收到注册请求", zap.Any("user", user))
	db := database.DB
	if err = db.Create(&user).Error; err != nil {
		logger.Log.Error("注册失败", zap.Error(err))
		return err
	}
	logger.Log.Info("注册成功", zap.Any("user", user))
	c.JSON("注册成功")

	return nil
}

func DeleteUserHandler(c fiber.Ctx) (err error) {
	name := c.Query("name")
	if name == "" {
		c.JSON("用户名不能为空")
		return
	}

	user := model.User{Username: name}
	if rows := database.DB.Where(&user).Delete(&user).RowsAffected; rows == 0 {
		c.JSON("用户不存在")
		return
	}
	logger.Log.Info("已删除用户", zap.Any("user", user))
	c.JSON("删除成功")
	return
}
