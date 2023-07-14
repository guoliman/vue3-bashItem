package roleModel

// SysRoleMenu 角色和菜单关联表
type SysRoleMenu struct {
	RoleID int64 `gorm:"column:role_id;not null;comment:角色ID" json:"role_id"`
	MenuID int64 `gorm:"column:menu_id;not null;comment:菜单ID" json:"menu_id"`
}

// TableName SysRoleMenu的表名
func (t *SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
