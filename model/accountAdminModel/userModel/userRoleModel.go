package userModel

// SysUserRole 用户和角色关联表
type SysUserRole struct {
	// 联合索引   因2个id都写了primaryKey 实现UserID+RoleID组成联合索引
	UserID int64 `gorm:"column:user_id;not null;comment:用户ID;primaryKey" json:"user_id"`
	RoleID int64 `gorm:"column:role_id;not null;comment:角色ID;primaryKey" json:"role_id"`
}

// TableName SysUserRole的表名
func (t *SysUserRole) TableName() string {
	return "sys_user_role"
}
