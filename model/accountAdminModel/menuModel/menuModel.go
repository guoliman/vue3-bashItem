package menuModel

import "time"

// SysMenu 菜单管理
type SysMenu struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement;comment:主键" json:"id"`
	ParentID   int64     `gorm:"column:parent_id;not null;comment:父节点id" json:"parent_id"`
	Name       string    `gorm:"column:name;type:varchar(64);not null;default:'';comment:菜单名称" json:"name"`
	Type       uint8     `gorm:"column:type;not null;comment:菜单类型(1:菜单；2:目录；3:外链；4:按钮)" json:"type"`
	Path       string    `gorm:"column:path;type:varchar(128);default:'';comment:路由路径(浏览器地址栏路径)" json:"path"`
	Component  string    `gorm:"column:component;type:varchar(128);default:NULL;comment:组件路径(vue页面完整路径)" json:"component"`
	Perm       string    `gorm:"column:perm;default:'';comment:权限标识" json:"perm"`
	Visible    uint8     `gorm:"column:visible;not null;default:1;comment:显示状态(1-显示;0-隐藏)" json:"visible"`
	Sort       int64     `gorm:"column:sort;default:0;comment:排序" json:"sort"`
	Icon       string    `gorm:"column:icon;type:varchar(64);default:NULL;comment:菜单图标" json:"icon"`
	Redirect   string    `gorm:"column:redirect;type:varchar(128);default:NULL;comment:跳转路径" json:"redirect"`
	ApiType    string    `gorm:"column:api_type;type:varchar(64);default:'';comment:后端接口类型" json:"apiType"`
	ApiPath    string    `gorm:"column:api_path;type:varchar(255);default:'';comment:后端请求地址" json:"apiPath"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`
	// ALTER TABLE sys_menu ADD COLUMN api_path VARCHAR(255) DEFAULT '' COMMENT '后端请求地址'; // 新增2个后端字段
	// ALTER TABLE sys_menu ADD COLUMN api_type VARCHAR(64) DEFAULT '' COMMENT '后端接口类型';
}

// TableName 菜单管理设置表名
func (t *SysMenu) TableName() string {
	return "sys_menu"
}
