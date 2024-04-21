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
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Username  string         `gorm:"size:255;index" validate:"required" json:"username,omitempty"`          // 用户名，唯一且不为空
	Password  string         `gorm:"size:255" validate:"required,min=6,max=30" json:"password,omitempty"`   // 密码
	Email     string         `gorm:"size:100;index:,unique,where:DeletedAt IS NULL" json:"email,omitempty"` // 电子邮箱，唯一且不为空
	Avatar    *string        `gorm:"size:255" json:"avatar,omitempty"`                                      // 头像
	IsActive  bool           `gorm:"default:true" json:"is_active,omitempty"`                               // 是否激活
	LastLogin *time.Time     `json:"last_login,omitempty"`                                                  // 最后一次登录时间
}
