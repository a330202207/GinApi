package model

type Resource struct {
	Model

	PID            int    `gorm:"default:0" json:"pid"`              //上级ID
	Name           string `gorm:"default:''" json:"name"`            //资源名称
	ResourceRouter string `gorm:"default:''" json:"resource_router"` //资源访问路由
}
