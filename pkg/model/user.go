package model

import (
	"gorm.io/gorm"
	"lumi/pkg/database"
	"time"
)

func init() {
	database.RegisterModels = append(database.RegisterModels, User{})
}

// User 定义了用户模型的基本结构
type User struct {
	gorm.Model            // 内嵌 gorm.Model，包含 ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string     `gorm:"size:255;index:,unique,where:DeletedAt IS NULL"` // 用户名，唯一且不为空
	Password   string     `gorm:"size:255"`                                       // 密码
	Email      string     `gorm:"size:100;index:,unique,where:DeletedAt IS NULL"` // 电子邮箱，唯一且不为空
	Avatar     *string    `gorm:"size:255"`                                       // 头像
	IsActive   bool       `gorm:"default:true"`                                   // 是否激活
	LastLogin  *time.Time // 最后一次登录时间
}
