package roleMethod

// RoleData 角色数据
type RoleData struct {
	Id         int64  `gorm:"comment:主键" json:"id"`
	Name       string `gorm:"comment:角色名称" json:"name"`
	Code       string `gorm:"comment:角色编码" json:"code"`
	Sort       int64  `gorm:"comment:排序" json:"sort"`
	Status     uint8  `gorm:"comment:角色状态(1-正常；0-停用)" json:"status"`
	DataScope  uint8  `gorm:"comment:数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)" json:"dataScope"`
	Deleted    uint8  `gorm:"comment:逻辑删除标识(0-未删除；1-已删除)" json:"deleted"`
	CreateTime string `gorm:"comment:创建时间" json:"create_time"`
	UpdateTime string `gorm:"comment:更新时间" json:"update_time"`
}

// RoleId 角色id
type RoleSonMenu struct {
	RoleId  int64   `comment:"角色id" json:"roleId"`
	MenuIds []int64 `comment:"菜单列表" json:"menuIds"`
}

// DelId 删除id
type DelId struct {
	Id string `json:"id"`
}
