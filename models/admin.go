package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Admin struct {
	Model

	UserName  string    `gorm:"default:''" json:"user_name"` //用户名
	Password  string    `gorm:"default:''" json:"password"`  //密码
	CreateIp  string    `gorm:"default:''" json:"create_ip"` //创建时IP
	LoginIp   string    `gorm:"default:''" json:"login_ip"`  //登录时IP
	LoginDate time.Time `json:"login_date"`                  //登录日期
	LoginsCnt int       `gorm:"default:0" json:"logins_cnt"` //登录次数
	State     int       `json:"state"`
}

//检查用户密码
func (admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	return err == nil
}

//检查用户状态
func (admin *Admin) CheckStatus() bool {
	if admin.State == 2 || admin.State == 3 {
		return false
	}
	return true
}
