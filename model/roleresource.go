package model

type RoleResource struct {
	RoleID     int `gorm:"default:0" json:"role_id"`
	ResourceID int `gorm:"default:0" json:"resource_id"`
}
