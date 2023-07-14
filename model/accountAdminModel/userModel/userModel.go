package userModel

import "time"

// SysUser 用户信息表
type SysUser struct {
	Id         int64     `gorm:"column:id;primaryKey;autoIncrement;comment:主键" json:"id"`
	Username   string    `gorm:"column:username;type:varchar(64);default:NULL;comment:用户名" json:"username"`
	Nickname   string    `gorm:"column:nickname;type:varchar(64);default:NULL;comment:昵称" json:"nickname"`
	Gender     uint8     `gorm:"column:gender;default:1;comment:性别((1:男;2:女))" json:"gender"`
	Password   string    `gorm:"column:password;type:varchar(100);default:NULL;comment:密码" json:"password"`
	DeptID     int64     `gorm:"column:dept_id;default:NULL;comment:部门ID" json:"dept_id"`
	Avatar     string    `gorm:"column:avatar;default:'';comment:用户头像" json:"avatar"`
	Mobile     string    `gorm:"column:mobile;type:varchar(20);default:NULL;comment:联系方式" json:"mobile"`
	Status     uint8     `gorm:"column:status;default:1;comment:用户状态((1:正常;0:禁用))" json:"status"`
	Email      string    `gorm:"column:email;type:varchar(128);default:NULL;comment:用户邮箱" json:"email"`
	Deleted    uint8     `gorm:"column:deleted;default:0;comment:逻辑删除标识(0:未删除;1:已删除)" json:"deleted"`
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`
}

// TableName SysUser的表名
func (t *SysUser) TableName() string {
	return "sys_user"
}
