package migrate

import (
	"fmt"
	"vue3-bashItem/model"
	"vue3-bashItem/model/accountAdminModel/deptModel"
	"vue3-bashItem/model/accountAdminModel/dictModel"
	"vue3-bashItem/model/accountAdminModel/menuModel"
	"vue3-bashItem/model/accountAdminModel/roleModel"
	"vue3-bashItem/model/accountAdminModel/userModel"
	"vue3-bashItem/pkg/logger"
)

func MigrateTable() {
	// 如下方法指定表引擎和表备注
	//model.Db.Set("gorm:table_options", "ENGINE=InnoDB COMMENT='用户表'").AutoMigrate(&leakManage.User{})
	err := model.Db.AutoMigrate(
		// 在这里注册表
		//&accountModel.Menu{},             // 菜单表
		//&accountModel.Role{},             // 角色表
		//&accountModel.User{},             // 用户表
		//&accountModel.RelationUserRole{}, // 用户与角色关联表

		// 权限
		&dictModel.SysDictType{}, // 字典表
		&dictModel.SysDict{},     // 字典子级表
		&deptModel.SysDept{},     // 部门表
		&menuModel.SysMenu{},     // 部门表
		&roleModel.SysRole{},     // 角色表
		&roleModel.SysRoleMenu{}, // 角色关联菜单表
		&userModel.SysUser{},     // 用户表
		&userModel.SysUserRole{}, // 用户关联角色表
	)

	if err != nil {
		//log.Fatal("====================migrate table fail! :", err)
		migrateErr := fmt.Sprintf("====================migrate table fail! :%v", err)
		logger.FileLogger.Error(migrateErr)
		panic(migrateErr)
	} else {
		migrateInfo := "\n====================migrate table success===================="
		logger.Logger.Info(migrateInfo)
		logger.FileLogger.Info(migrateInfo)
	}

}
