package model

import (
	"GinApi/util"
	"fmt"
	"strconv"
	"strings"
)

type RoleResource struct {
	RoleID     int `gorm:"default:0" json:"role_id"`
	ResourceID int `gorm:"default:0" json:"resource_id"`
}

func AddRoleResource(roleID int, resourceIDs []string) (err error) {
	str := []string{"role_id", "resource_id"}
	newArr := [][]string{}
	for _, v := range resourceIDs {
		Arr := append([]string{strconv.Itoa(roleID)}, v)
		newArr = append(newArr, Arr)
	}

	var newStr string
	for _, v := range newArr {

		newStr += fmt.Sprintf("('%s'),", strings.Join(v, "','"))

	}
	key := strings.Join(str, ",")
	val := strings.TrimRight(newStr, ",")

	sql := util.BatchInsert("api_role_resource", key, val)

	err = DB.Exec(sql).Error

	return
}

//获取全部
func GetRoleResources(roleID int) (list []RoleResource, err error) {
	err = DB.Unscoped().Where("role_id = ?", roleID).Find(&list).Error
	return
}

//删除
func DelRoleResource(roleID int) (err error) {
	err = DB.Model(&AdminRole{}).Where("role_id = ?", roleID).Unscoped().Delete(&RoleResource{}).Error
	return
}
