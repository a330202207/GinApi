package model

type Resource struct {
	ID             int    `gorm:"primary_key" json:"id"`
	PID            int    `gorm:"default:0" json:"p_id"`             //上级ID
	Name           string `gorm:"default:''" json:"name"`            //资源名称
	ResourceRouter string `gorm:"default:''" json:"resource_router"` //资源访问路由
	CreatedAt      int    `json:"created_at"`
	UpdatedAt      int    `json:"updated_at"`
}

func AddResource(resource *Resource) (err error) {
	err = DB.Create(resource).Error
	return
}

func GetResourceList(Limit, Offset int, order string, query interface{}, args ...interface{}) (list []Resource, count int, err error) {
	err = DB.Unscoped().Where(query, args...).Order(order).Limit(Limit).Offset(Offset).Find(&list).Error
	DB.Unscoped().Model(&Resource{}).Where(query, args...).Count(&count)
	return
}

func GetResource(maps interface{}) (resource Resource, err error) {
	err = DB.Unscoped().Where(maps).First(&resource).Error
	return
}

func GetResources(maps interface{}) (resource []Resource, err error) {
	err = DB.Unscoped().Where(maps).Find(&resource).Error
	return
}

func ExistResource(maps interface{}) bool {
	var resource Resource
	DB.Unscoped().Where(maps).First(&resource)
	if resource.ID > 0 {
		return true
	}
	return false
}

//删除角色
func DelResource(maps interface{}) (err error) {
	err = DB.Where(maps).Unscoped().Delete(&Resource{}).Error
	return
}

//保存角色
func SaveResource(id int, resource Resource) (err error) {
	err = DB.Model(&resource).Where("id = ?", id).Updates(resource).Error
	return
}
