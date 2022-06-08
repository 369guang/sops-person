package models

import "person/core/database"

type ActionLogs struct {
	database.BaseModel
	UserID     uint64 `gorm:"comment('用户ID')" json:"user_id"`
	Username   string `gorm:"comment('用户名')" json:"username"`
	Title      string `gorm:"comment('日志标题')" json:"title"`
	URI        string `gorm:"comment('路由地址')" json:"uri"`
	Method     string `gorm:"comment('http方法')" json:"method"` // GET/POST/HEAD/PUT/DELETE/OPTIONS
	StatusCode int    `gorm:"comment('http状态码')" json:"status_code"`
	Params     string `gorm:"comment('参数')" json:"params"`
	IP         string `gorm:"comment('客户端ip')" json:"ip"`
	Device     string `gorm:"comment('驱动')" json:"device"`       // Windows/Linux/MacOS/Android/IOS
	Browser    string `gorm:"comment('浏览器')" json:"browser"`     // Chrome/Firefox/IE/Edge/Safari
	FinishTime uint32 `gorm:"comment('总耗时')" json:"finish_time"` // 单位：millisecond
}

func (ActionLogs) TableName() string {
	return "system_action_logs"
}

type LoginLogs struct {
	database.BaseModel
	UserID   uint64 `gorm:"comment('用户ID')" json:"user_id"`
	Username string `gorm:"comment('用户名')" json:"username"`
	IP       string `gorm:"comment('客户端ip')" json:"ip"`
	Device   string `gorm:"comment('驱动')" json:"device"`   // Windows/Linux/MacOS/Android/IOS
	Browser  string `gorm:"comment('浏览器')" json:"browser"` // Chrome/Firefox/IE/Edge/Safari
}

func (LoginLogs) TableName() string {
	return "system_login_logs"
}
