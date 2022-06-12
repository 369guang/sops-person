package models

import (
	"fmt"
	"gorm.io/gorm"
	"person/core/auth"
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

func (u *User) GetStatus() string {
	data := map[uint8]string{
		0: "",
		1: "正常",
		2: "停用",
	}
	return data[u.Status]
}

func (u *User) Encrypt() (err error) { // 加密处理
	if u.Password != "" {
		u.Password, err = auth.Encrypt(u.Password)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	fmt.Println("执行了 BeforeCreate")
	return u.Encrypt()
}
