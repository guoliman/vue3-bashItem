package userMethod

import (
	"errors"
	"fmt"
	"vue3-bashItem/model"
	"vue3-bashItem/model/accountAdminModel/userModel"
	"vue3-bashItem/pkg/logger"
)

type MeData struct {
	UserId   int      `json:"userId"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	Avatar   string   `json:"avatar"`
	Roles    []string `json:"roles"`
	Perms    []string `json:"perms"`
}

// UserData 创建用户
type UserData struct {
	Id       int64   `gorm:"comment:主键" json:"id"`
	Username string  `gorm:"comment:用户名" json:"username"`
	Nickname string  `gorm:"comment:昵称" json:"nickname"`
	Gender   uint8   `gorm:"comment:性别((1:男;2:女))" json:"gender"`
	DeptID   int64   `gorm:"comment:部门ID" json:"deptId"`
	RoleIds  []int64 `gorm:"comment:角色列表" json:"roleIds"`
	//Avatar     string `gorm:"comment:用户头像" json:"avatar"`
	Mobile     string `gorm:"comment:联系方式" json:"mobile"`
	Status     uint8  `gorm:"column:status;default:1;comment:用户状态((1:正常;0:禁用))" json:"status"`
	Email      string `gorm:"comment:用户邮箱" json:"email"`
	CreateTime string `gorm:"comment:创建时间" json:"createTime"`
	UpdateTime string `gorm:"comment:更新时间" json:"updateTime"`
}

// UserPass 用户密码
type UserPass struct {
	Id       int64  `gorm:"comment:主键" json:"id"`
	Password string `gorm:"comment:密码" json:"password"`
}

// DelId 删除id
type DelId struct {
	Id string `json:"id"`
}

// CaptchaResponse 认证图片
type CaptchaResponse struct {
	VerifyCodeKey    string `json:"verifyCodeKey"`
	VerifyCodeBase64 string `json:"verifyCodeBase64"`
}

// ChangeUserRole 更新用户角色
func ChangeUserRole(item int64, reqBody UserData) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	// 删除该用户存在的权限 因没有权限id 所以可以这么操作
	delErr := model.DeleteByQuery(&userModel.SysUserRole{}, map[string]interface{}{"user_id": item})
	if delErr != nil {
		return errors.New(fmt.Sprintf("用户id:%v 删除数据错误：%v", item, delErr.Error()))
	}

	// 批量新增
	var insertUserRole = make([]userModel.SysUserRole, 0)
	for _, roleY := range reqBody.RoleIds {
		insertUserRole = append(insertUserRole, userModel.SysUserRole{UserID: item, RoleID: roleY})
	}
	result := model.Db.Create(&insertUserRole)
	for _, user := range insertUserRole {
		if result.Error != nil {
			return errors.New(fmt.Sprintf("新增异常 UserID:%v  RoleID:%v  错误: %v",
				user.UserID, user.RoleID, result.Error))
		} else {
			logger.FileLogger.Debug(fmt.Sprintf("新增成功 UserID:%v  RoleID:%v  成功插入: %v",
				user.UserID, user.RoleID, result.RowsAffected))
		}
	}
	return nil
}
