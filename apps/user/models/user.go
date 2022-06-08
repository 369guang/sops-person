package models

import (
	"person/core/database"
	"time"
)

type User struct {
	database.BaseModel
	Username  string    `gorm:"unique;not null;comment('用户名')" json:"username" binding:"required"`
	Password  string    `gorm:"not null;comment('密码')" json:"password" serializers:"write;"`
	Email     string    `gorm:"comment('邮箱')" json:"email" binding:"required"`
	Mobile    string    `gorm:"comment('联系号码')" json:"mobile" binding:"required"`
	LoginIp   string    `gorm:"comment('登录IP')" json:"login_ip"`
	LastLogin time.Time `gorm:"comment('最后登录时间')" json:"last_login"`
	Status    uint8     `gorm:"comment('状态')" json:"status"`
	Avatar    string    `gorm:"comment('头像')" json:"avatar"`
}

func (User) TableName() string {
	return "system_user"
}
