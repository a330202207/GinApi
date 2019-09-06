package service

import (
	"GinApi/middleware/error"
	"GinApi/models"
)

type AdminLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

//登陆
func (service *AdminLoginService) Login() (models.Admin, int) {
	var admin models.Admin

	if err := models.DB.Where("user_name = ?", service.UserName).First(&admin).Error; err != nil {
		return admin, error.ERROR_NOT_EXIST_USER
	}

	//密码不对
	if admin.CheckPassword(service.Password) == false {
		return admin, error.ERROR_NOT_EXIST_USER
	}

	//禁用账户
	if admin.CheckStatus() == false {
		return admin, error.ERROR_DISABLE_USER
	}

	return admin, error.SUCCESS
}
