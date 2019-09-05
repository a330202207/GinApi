package service

import "GinApi/models"

type AdminLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=40"`
}

func (service *AdminLoginService) Login() models.Admin {
	var admin models.Admin

}
