package models

import "time"

type Admin struct {
	Model

	Username  string    `gorm:"default:''" json:"user_name"` //用户名
	Password  string    `gorm:"default:''" json:"password"`  //密码
	CreateIp  string    `gorm:"default:''" json:"create_ip"` //创建时IP
	LoginIp   string    `gorm:"default:''" json:"login_ip"`  //登录时IP
	LoginDate time.Time `json:"login_date"`                  //登录日期
	LoginsCnt int       `gorm:"default:0" json:"logins_cnt"` //登录次数
	State     int       `json:"state"`
}

//获取
func GetAdmin(maps interface{}) (admin Admin, err error) {
	err = db.Preload("Group").Model(&Admin{}).Where(maps).First(&admin).Error
	return
}