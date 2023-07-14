package dictModel

import "time"

// SysDictType 字典类型表
type SysDictType struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement;comment:主键" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(50);comment:类型名称" json:"name"`
	Code       string    `gorm:"column:code;type:varchar(50);unique;comment:类型编码 唯一" json:"code"`
	Status     uint8     `gorm:"column:status;comment:状态(0:正常;1:禁用)" json:"status"`
	Remark     string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`
}

// TableName 设置字典类型表名
func (t *SysDictType) TableName() string {
	return "sys_dict_type"
}

// SysDict 字典数据表
type SysDict struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement;comment:主键" json:"id"`
	TypeCode   string    `gorm:"column:type_code;type:varchar(64);comment:字典类型编码" json:"type_code"`
	Name       string    `gorm:"column:name;type:varchar(50);comment:字典项-名称" json:"name"`
	Value      string    `gorm:"column:value;type:varchar(50);comment:字典项-值" json:"value"`
	Sort       int64     `gorm:"column:sort;comment:排序" json:"sort"`
	Status     uint8     `gorm:"column:status;comment:状态(1:正常;0:禁用)" json:"status"`
	Defaulted  uint8     `gorm:"column:defaulted;default:0;comment:是否默认(1:是;0:否)" json:"defaulted"`
	Remark     string    `gorm:"column:remark;type:varchar(255);comment:备注" json:"remark"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`
}

func (t *SysDict) TableName() string {
	return "sys_dict"
}
