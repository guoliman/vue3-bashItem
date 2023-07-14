package deptModel

import "time"

// SysDept 部门表
type SysDept struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement;comment:主键" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(64);not null;default:'';comment:部门名称" json:"name"`
	ParentID   int64     `gorm:"column:parent_id;not null;default:0;comment:父节点id" json:"parent_id"`
	Sort       int64     `gorm:"column:sort;default:0;comment:显示顺序" json:"sort"`
	Status     uint8     `gorm:"column:status;not null;default:1;comment:状态(1:正常;0:禁用)" json:"status"`
	Deleted    uint8     `gorm:"column:deleted;default:0;comment:逻辑删除标识(1:已删除;0:未删除)" json:"deleted"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`
}

func (t *SysDept) TableName() string {
	return "sys_dept"
}
