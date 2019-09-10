package service

import (
	"GinApi/model"
	"GinApi/package/error"
	"golang.org/x/crypto/bcrypt"
)

type AdminLoginInfo struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
}

//登陆
func (loginInfo *AdminLoginInfo) Login() (model.Admin, int) {
	userName := map[string]interface{}{"user_name": loginInfo.UserName}
	admin, err := model.GetAdmin(userName)
	if err != nil {
		return admin, error.ERROR_NOT_EXIST_USER
	}

	//密码不对
	if CheckPassword(admin.Password, loginInfo.Password) == false {
		return admin, error.ERROR_NOT_EXIST_USER
	}

	//禁用账户
	if CheckStatus(admin.Status) == false {
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

type AdminInfo struct {
	ID       int    `form:"id" json:"id"`
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=40"`
	CreateIp string `json:"login_ip"`
	Status   int    `form:"status" json:"status" binding:"required"`
}

//添加用户
func (adminInfo *AdminInfo) AdminAdd() int {

	userName := map[string]interface{}{"user_name": adminInfo.UserName, "status": 1}
	isExist := model.ExistAdmin(userName)

	if isExist == true {
		return error.ERROR_EXIST_USER
	}

	hashPassword, err := SetPassword(adminInfo.Password)
	if err == false {
		return error.ERROR_PASSWORD_USER
	}
	admin := model.Admin{
		UserName: adminInfo.UserName,
		CreateIp: adminInfo.CreateIp,
		Password: hashPassword,
	}

	if err := model.AddAdmin(&admin); err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}
	return error.SUCCESS
}

//设置密码
func SetPassword(password string) (string, bool) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", false
	}
	return string(hash), true
}

//删除管理员
func (admin *AdminInfo) AdminDel() int {
	userName := map[string]interface{}{"id": admin.ID, "status": 1}
	isExist := model.ExistAdmin(userName)
	if isExist == false {
		return error.ERROR_NOT_EXIST_USER
	}

	err := model.DelAdmin(map[string]interface{}{"id": admin.ID})
	if err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}
	return error.SUCCESS
}

//编辑管理员
func (admin *AdminInfo) AdminEdit() (model.Admin, int) {
	adminInfo, err := model.GetAdmin(map[string]interface{}{"id": admin.ID})
	if err != nil {
		return adminInfo, error.ERROR
	}
	return adminInfo, error.SUCCESS
}

//保存管理员
func (adminInfo *AdminInfo) AdminSave() int {
	id := adminInfo.ID
	admin := model.Admin{
		UserName: adminInfo.UserName,
		Status:   adminInfo.Status,
	}
	if err := model.SaveAdmin(id, admin); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}
	return error.SUCCESS
}
