package roleModel

import "time"

// SysRole 角色表
type SysRole struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement;comment:主键" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(64);unique;not null;default:'';comment:角色名称" json:"name"`
	Code       string    `gorm:"column:code;type:varchar(32);default:NULL;comment:角色编码" json:"code"`
	Sort       int64     `gorm:"column:sort;default:0;comment:排序" json:"sort"`
	Status     uint8     `gorm:"column:status;default:1;comment:角色状态(1-正常；0-停用)" json:"status"`
	DataScope  uint8     `gorm:"column:data_scope;default:1;comment:数据权限(0-所有数据；1-部门及子部门数据；2-本部门数据；3-本人数据)" json:"data_scope"`
	Deleted    uint8     `gorm:"column:deleted;not null;default:0;comment:逻辑删除标识(0-未删除；1-已删除)" json:"deleted"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`
}

// TableName SysRole的表名
func (t *SysRole) TableName() string {
	return "sys_role"
}
