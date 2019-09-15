package model

import (
	"GinApi/util"
	"fmt"
	"strconv"
	"strings"
)

type AdminRole struct {
	AdminID int `gorm:"default:0" json:"admin_id"`
	RoleID  int `gorm:"default:0" json:"role_id"`
}

//添加
func AddAdminRole(adminId int, roleIds []string) (err error) {
	str := []string{"admin_id", "role_id"}
	newArr := [][]string{}
	for _, v := range roleIds {
		Arr := append([]string{strconv.Itoa(adminId)}, v)
		newArr = append(newArr, Arr)
	}

	var newStr string
	for _, v := range newArr {

		newStr += fmt.Sprintf("('%s'),", strings.Join(v, "','"))

	}
	key := strings.Join(str, ",")
	val := strings.TrimRight(newStr, ",")

	sql := util.BatchInsert("api_admin_role", key, val)

	err = DB.Exec(sql).Error

	return
}

//获取全部
func GetAdminRoles(adminID int) (roles []AdminRole, err error) {
	err = DB.Unscoped().Where("admin_id = ?", adminID).Find(&roles).Error
	return
}

//删除
func DelAdminRole(adminID int) (err error) {
	err = DB.Model(&AdminRole{}).Where("admin_id = ?", adminID).Unscoped().Delete(&AdminRole{}).Error
	return
}
