package model

type Resource struct {
	ID             int    `gorm:"primary_key" json:"id"`
	PID            int    `gorm:"default:0" json:"p_id"`             //上级ID
	Name           string `gorm:"default:''" json:"name"`            //资源名称
	ResourceRouter string `gorm:"default:''" json:"resource_router"` //资源访问路由
	CreatedAt      int    `json:"created_at"`
	UpdatedAt      int    `json:"updated_at"`
	Children       []*Resource
}

type TreeResources struct {
	ID             int              `json:"id"`
	PID            int              `json:"p_id"`            //上级ID
	Name           string           `json:"text"`            //资源名称
	ResourceRouter string           `json:"resource_router"` //资源访问路由
	Children       []*TreeResources `json:"children"`
}

//添加资源
func AddResource(resource *Resource) (err error) {
	err = DB.Create(resource).Error
	return
}

//获取资源列表
func GetResourceList(Limit, Offset int, order string, query interface{}, args ...interface{}) (list []Resource, count int, err error) {
	err = DB.Unscoped().Where(query, args...).Order(order).Limit(Limit).Offset(Offset).Find(&list).Error
	DB.Unscoped().Model(&Resource{}).Where(query, args...).Count(&count)
	return
}

//获取资源
func GetResource(maps interface{}) (resource Resource, err error) {
	err = DB.Unscoped().Where(maps).First(&resource).Error
	return
}

//获取多个资源
func GetResources(maps interface{}) (resource []Resource, err error) {
	err = DB.Unscoped().Where(maps).Find(&resource).Error
	return
}

//获取资源树
func (r *Resource) GetTreeResources(pid interface{}) []*TreeResources {
	resources, _ := GetResources(map[string]interface{}{"p_id": pid})

	treeList := []*TreeResources{}
	for _, v := range resources {
		child := v.GetTreeResources(v.ID)
		node := &TreeResources{
			ID:             v.ID,
			PID:            v.PID,
			Name:           v.Name,
			ResourceRouter: v.ResourceRouter,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}

//资源是否存在
func ExistResource(maps interface{}) bool {
	var resource Resource
	DB.Unscoped().Where(maps).First(&resource)
	if resource.ID > 0 {
		return true
	}
	return false
}

//删除资源
func DelResource(maps interface{}) (err error) {
	err = DB.Where(maps).Unscoped().Delete(&Resource{}).Error
	return
}

//保存资源
func SaveResource(id int, resource Resource) (err error) {
	err = DB.Model(&resource).Where("id = ?", id).Updates(resource).Error
	return
}
