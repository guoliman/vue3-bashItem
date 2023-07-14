package roleMethod

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"vue3-bashItem/model"
	"vue3-bashItem/model/accountAdminModel/menuModel"
	"vue3-bashItem/model/accountAdminModel/roleModel"
	"vue3-bashItem/model/accountAdminModel/userModel"
	"vue3-bashItem/pkg/utils"
)

// RoleGetMenu 用户拥有的菜单权限 动态路由使用
func RoleGetMenu(c *gin.Context) (userData map[string]interface{}, roleCodeListInfo []string, menuIdListInfo []int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("RoleGetMenu: %v", r)
		}
	}()

	getUser, _ := c.Get("username") // 取用户名
	// 获取用户信息
	userInfo := userModel.SysUser{}
	if userInfoErr := model.GetOneLast(&userInfo, "username = ?", getUser.(string)); userInfoErr != nil {
		return map[string]interface{}{}, []string{}, []int64{}, errors.New(fmt.Sprintf("用户不存在 %v ", userInfoErr.Error()))
	} else if userInfo.Status == 0 {
		return map[string]interface{}{}, []string{}, []int64{}, errors.New(fmt.Sprintf("用户已禁用 请联系管理员"))
	}

	userResult := make(map[string]interface{})
	userResult["userId"] = userInfo.Id
	userResult["username"] = userInfo.Username
	userResult["nickname"] = userInfo.Nickname
	userResult["avatar"] = userInfo.Avatar

	// 获取关联数据 用户的角色
	userRole := make([]userModel.SysUserRole, 0)
	relationErr := model.GetMany(&userRole, "user_id = ?", userInfo.Id)
	if relationErr != nil {
		return map[string]interface{}{}, []string{}, []int64{}, errors.New(fmt.Sprintf("用户的角色不存在 %v ", relationErr.Error()))
	}

	// 取角色id列表
	var userRoleList []int64
	for _, userRoleV := range userRole {
		userRoleList = append(userRoleList, userRoleV.RoleID)
	}

	// 取角色code列表
	roleInfo := make([]roleModel.SysRole, 0)
	if roleErr := model.Db.Where("id In ?", userRoleList).Find(&roleInfo).Error; roleErr != nil {
		return map[string]interface{}{}, []string{}, []int64{}, errors.New(fmt.Sprintf("角色不存在 %v ", roleErr.Error()))
	}
	var roleCodeList []string
	for _, roleV := range roleInfo {
		roleCodeList = append(roleCodeList, roleV.Code)
	}
	//logger.FileLogger.Debug("角色code列表===", roleCodeList)

	// 取菜单权限 菜单id列表
	var menuIdList []int64
	if existBool := utils.StringInSlice(roleCodeList, "ROOT"); existBool { // 是超级管理员ROOT权限 取全部菜单id
		menuInfo := make([]menuModel.SysMenu, 0)
		if roleIdErr := model.GetAll(&menuInfo); roleIdErr != nil {
			return map[string]interface{}{}, []string{}, []int64{}, errors.New(fmt.Sprintf("角色id不存在 %v ", roleIdErr.Error()))
		}
		for _, menuV := range menuInfo {
			menuIdList = append(menuIdList, menuV.Id)
		}
	} else { // 不是 就根据角色查询子菜单
		menuInfo := make([]roleModel.SysRoleMenu, 0)
		if roleIdErr := model.Db.Where("role_id In ?", userRoleList).Find(&menuInfo).Error; roleIdErr != nil {
			return map[string]interface{}{}, []string{}, []int64{}, errors.New(fmt.Sprintf("角色id不存在 %v ", roleIdErr.Error()))
		}
		for _, menuV := range menuInfo {
			menuIdList = append(menuIdList, menuV.MenuID)
		}
	}

	menuIdList = utils.RemoveListInt64(menuIdList) // 去重
	//logger.FileLogger.Debug("去重后的菜单id==", menuIdList)

	return userResult, roleCodeList, menuIdList, nil
}

// 后端接口认证
func UrlAuth(c *gin.Context, urlType, urlAddr string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("UrlAuth: %v", r)
		}
	}()
	// 获取用户拥有权限的菜单id列表
	//userResult, roleCodeList, menuIdList, menuIdListErr := RoleGetMenu(c)
	_, _, menuIdList, menuIdListErr := RoleGetMenu(c)
	if menuIdListErr != nil {
		return errors.New(fmt.Sprintf("获取菜单异常: %v ", menuIdListErr.Error()))
	}

	// 获取菜单表中接口id
	menuInfo := menuModel.SysMenu{}
	menuInfoErr := model.GetOneLast(&menuInfo, "api_type = ? and api_path = ?", urlType, urlAddr)
	if menuInfoErr != nil {
		menuInfoErrData := fmt.Sprintf("接口不存在 %v", menuInfoErr.Error())
		return errors.New(menuInfoErrData)
	}

	// 对比是否存在
	//logger.FileLogger.Debug("menuIdList===", menuIdList)
	if idStatus := utils.JudgeInt64InList(menuInfo.Id, menuIdList); !idStatus { // 不存在 表示没权限
		return errors.New(fmt.Sprintf("菜单id:%v 菜单name:%v", menuInfo.Id, menuInfo.Name))
	}
	return nil

}
