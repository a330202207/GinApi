package model

type AdminRole struct {
	AdminID int `gorm:"default:0" json:"admin_id"`
	RoleID  int `gorm:"default:0" json:"role_id"`
}
