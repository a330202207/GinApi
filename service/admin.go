package service

import (
	"GinApi/model"
	"GinApi/package/error"
	"golang.org/x/crypto/bcrypt"
)

type AdminLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

//登陆
func (service *AdminLoginService) Login() (model.Admin, int) {

	admin, err := model.GetAdmin(service.UserName)
	if err != nil {
		return admin, error.ERROR_NOT_EXIST_USER
	}

	//密码不对
	if CheckPassword(admin.Password, service.Password) == false {
		return admin, error.ERROR_NOT_EXIST_USER
	}

	//禁用账户
	if CheckStatus(admin.State) == false {
		return admin, error.ERROR_DISABLE_USER
	}

	return admin, error.SUCCESS
}

//检查用户密码
func CheckPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

//检查用户状态
func CheckStatus(status int) bool {
	if status == 2 || status == 3 {
		return false
	}
	return true
}
