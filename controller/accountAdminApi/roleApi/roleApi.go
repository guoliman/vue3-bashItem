package roleApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
	"vue3-bashItem/model"
	"vue3-bashItem/model/accountAdminModel/roleModel"
	"vue3-bashItem/model/accountAdminModel/userModel"
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/request"
	"vue3-bashItem/pkg/response"
	"vue3-bashItem/pkg/utils"
	"vue3-bashItem/services/accountAdminMethod/roleMethod"
)

// SelectRole 获取角色下拉菜单
func SelectRole(c *gin.Context) {

	roleData := make([]*roleModel.SysRole, 0)
	roleDataErr := model.GetAll(&roleData)
	if roleDataErr != nil {
		response.BaseError(c, fmt.Sprintf("查询SysDept失败 %v", roleDataErr.Error()))
	}

	var roleList = make([]map[string]interface{}, 0)
	for _, roleY := range roleData {
		oneDept := map[string]interface{}{
			"value": roleY.Id,
			"label": roleY.Name,
		}
		roleList = append(roleList, oneDept)
	}
	response.Success(c, roleList)
}

// GetRole 获取角色
func GetRole(c *gin.Context) {
	// 分页参数 offset是偏移 limit是条数
	offset, limit, err := request.GetPagerParams(c)
	if err != nil {
		response.ParamsError(c, err.Error())
		return
	}
	// 排序条件
	orderExpr, err := request.GetOrderParams(c)
	if err != nil {
		response.ParamsError(c, err.Error())
		return
	}

	// 多条件like查询 数据拼接成string
	validFilterFields := []string{"name"}
	likeExpr, likeValues := request.GetLikeFilterQuery(c, validFilterFields, 0)

	// 查询db获取数据
	dictData := make([]*roleModel.SysRole, 0)
	dataNum, getErr := model.GetListPage(&dictData, offset, limit, orderExpr, likeExpr, likeValues...)
	if getErr != nil {
		response.BaseError(c, fmt.Sprintf("查询db获取数据失败 %v", err.Error()))
		return
	}

	var roleList = make([]map[string]interface{}, 0)
	for _, dictY := range dictData {
		oneDictPyte := map[string]interface{}{
			"id":          dictY.Id,
			"name":        dictY.Name,
			"code":        dictY.Code,
			"sort":        dictY.Sort,
			"status":      dictY.Status,
			"dataScope":   dictY.DataScope,
			"create_time": dictY.CreateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 更改时间
			"update_time": dictY.UpdateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 更改时间
		}
		roleList = append(roleList, oneDictPyte)
	}
	var roleResult = make(map[string]interface{})
	roleResult["list"] = roleList
	roleResult["total"] = dataNum

	response.Success(c, roleResult)
}

// CreateRole 创建角色
func CreateRole(c *gin.Context) {
	reqBody := &roleMethod.RoleData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("创建接口-获取数据失败%v", err.Error()))
		return
	}
	// 角色名唯一判断
	roleData := roleModel.SysRole{}
	if roleErr := model.GetOneLast(&roleData, "name = ?", reqBody.Name); roleErr == nil {
		response.ParamsError(c, fmt.Sprintf("%v 角色名称已存在 请更换", reqBody.Name))
		return
	}
	// code唯一判断
	if roleCodeErr := model.GetOneLast(&roleData, "code = ?", reqBody.Code); roleCodeErr == nil {
		response.ParamsError(c, fmt.Sprintf("%v 角色编码已存在 请更换", reqBody.Code))
		return
	}

	nowTime := time.Now()
	item := &roleModel.SysRole{
		Name:       reqBody.Name,
		Code:       reqBody.Code,
		Sort:       reqBody.Sort,
		Status:     reqBody.Status,
		DataScope:  reqBody.DataScope,
		Deleted:    reqBody.Deleted,
		CreateTime: nowTime,
		UpdateTime: nowTime,
	}
	if err := model.Create(item); err != nil {
		response.BaseError(c, fmt.Sprintf("SysRole-创建数据错误：%v", err.Error()))
		return
	}
	response.Success(c, item)
}

// UpdateRole 更新角色
func UpdateRole(c *gin.Context) {
	reqBody := &roleMethod.RoleData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("获取参数错误：%v", err.Error()))
		return
	}

	// 角色名唯一判断
	roleData := roleModel.SysRole{}
	if roleErr := model.GetOneLast(&roleData, "id != ? and name = ?", reqBody.Id, reqBody.Name); roleErr == nil {
		response.ParamsError(c, fmt.Sprintf("%v 角色名称已存在 请更换", reqBody.Name))
		return
	}
	// code唯一判断
	roleCodeErr := model.GetOneLast(&roleData, "id != ? and code = ?", reqBody.Id, reqBody.Code)
	if roleCodeErr == nil {
		response.ParamsError(c, fmt.Sprintf("%v 角色编码已存在 请更换", reqBody.Code))
		return
	}

	// 创建时间 string转time 东八区
	local, _ := time.LoadLocation("Asia/Shanghai")
	createTime, _ := time.ParseInLocation("2006-01-02 15:04:05", reqBody.CreateTime, local) // 转成东8区time

	updateColumns := map[string]interface{}{
		"Name":       reqBody.Name,
		"Code":       reqBody.Code,
		"Sort":       reqBody.Sort,
		"Status":     reqBody.Status,
		"DataScope":  reqBody.DataScope,
		"Deleted":    reqBody.Deleted,
		"CreateTime": createTime,
		"UpdateTime": time.Now(),
	}
	// 已软删除的id 不会更改  但会返回更改成功
	if err := model.Update(&roleModel.SysRole{}, updateColumns, "id = ?", reqBody.Id); err != nil {
		response.BaseError(c, fmt.Sprintf("SysRole表更新错误：%v", err.Error()))
		return
	}

	response.Success(c, "更新成功!")
}

// DeleteRole 删除用户
func DeleteRole(c *gin.Context) {
	delId := &roleMethod.DelId{}
	if err := c.ShouldBindJSON(&delId); err != nil {
		response.ParamsError(c, fmt.Sprintf("删除接口，获取数据失败%v", err.Error()))
		return
	}

	// 批量删除
	idList := strings.Split(delId.Id, ",")
	for iNum := 0; iNum < len(idList); iNum++ {
		idInt, intErr := strconv.ParseInt(idList[iNum], 10, 64) //string转int64
		if intErr != nil {
			response.BaseError(c, fmt.Sprintf("：%v string转int64转译失败 %v", idList[iNum], intErr.Error()))
			return
		}

		// 删除该用户关联的角色
		delErr := model.DeleteByQuery(&roleModel.SysRoleMenu{}, map[string]interface{}{"role_id": idInt})
		if delErr != nil {
			response.ParamsError(c, fmt.Sprintf("删除角色关联菜单失败：%v", delErr.Error()))
			return
		}

		// 删除该用户关联的角色
		delUserRoleErr := model.DeleteByQuery(&userModel.SysUserRole{}, map[string]interface{}{"role_id": idInt})
		if delUserRoleErr != nil {
			response.ParamsError(c, fmt.Sprintf("删除用户关联角色失败：%v", delUserRoleErr.Error()))
			return
		}

		// 删除用户
		if err := model.DeleteByQuery(&roleModel.SysRole{}, map[string]interface{}{"id": idInt}); err != nil {
			response.BaseError(c, fmt.Sprintf("删除用户数据错误：%v", err.Error()))
			return
		}
	}

	response.Success(c, "删除成功!")
}

// GetRoleSonMenu 获取角色拥有的菜单权限
func GetRoleSonMenu(c *gin.Context) {
	// 接收数据
	roleSonMenu := &roleMethod.RoleSonMenu{}
	if err := c.ShouldBindJSON(&roleSonMenu); err != nil {
		response.ParamsError(c, fmt.Sprintf("获取角色id数据失败%v", err.Error()))
		return
	}
	logger.FileLogger.Debug("roleSonMenu====", roleSonMenu)
	// 查询数据
	roleData := make([]*roleModel.SysRoleMenu, 0)
	if roleErr := model.GetMany(&roleData, "role_id = ?", roleSonMenu.RoleId); roleErr != nil {
		response.ParamsError(c, fmt.Sprintf("role_id:%v 不存在SysRoleMenu表", roleSonMenu.RoleId))
		return
	}
	logger.FileLogger.Debug("role管理menu表数据====", roleData)
	var menuList []int64
	for _, roleY := range roleData {
		menuList = append(menuList, roleY.MenuID)
	}

	response.Success(c, menuList)
}

// PostRoleSonMenu 更改角色拥有的菜单权限
func PostRoleSonMenu(c *gin.Context) {
	// 获取前端数据
	roleSonMenu := &roleMethod.RoleSonMenu{}
	if err := c.ShouldBindJSON(&roleSonMenu); err != nil {
		response.ParamsError(c, fmt.Sprintf("获取参数错误：%v", err.Error()))
		return
	}
	logger.FileLogger.Debug("角色id: ", roleSonMenu.RoleId)
	logger.FileLogger.Debug("新菜单id列表: ", roleSonMenu.MenuIds)

	// 根据角色id查数据
	roleData := make([]*roleModel.SysRoleMenu, 0)
	if roleErr := model.GetMany(&roleData, "role_id = ?", roleSonMenu.RoleId); roleErr != nil {
		response.ParamsError(c, fmt.Sprintf("role_id:%v 不存在SysRoleMenu表", roleSonMenu.RoleId))
		return
	}
	var oldMenuIds []int64
	for _, idY := range roleData {
		oldMenuIds = append(oldMenuIds, idY.MenuID)
	}
	logger.FileLogger.Debug("旧菜单id列表: ", oldMenuIds)

	// 差集处理 得到新增数据和删除数据
	delMenuIds := utils.SetInt64XJ(oldMenuIds, roleSonMenu.MenuIds)
	logger.FileLogger.Debug(fmt.Sprintf("删除数据 roleId: %v mune-id: %v ", roleSonMenu.RoleId, delMenuIds))
	addMenuIds := utils.SetInt64XJ(roleSonMenu.MenuIds, oldMenuIds)
	logger.FileLogger.Debug(fmt.Sprintf("新增数据 roleId: %v mune-id: %v ", roleSonMenu.RoleId, addMenuIds))

	// 删除操作
	for _, delY := range delMenuIds {
		var delInfo = map[string]interface{}{"role_id": roleSonMenu.RoleId, "menu_id": delY}
		if err := model.DeleteByQuery(&roleModel.SysRoleMenu{}, delInfo); err != nil {
			response.BaseError(c, fmt.Sprintf("SysRoleMenu-删除异常 role_id:%v menu_id:%v 错误：%v", roleSonMenu.RoleId, delY, err.Error()))
			return
		}
	}

	// 新增
	for _, addY := range addMenuIds {
		item := &roleModel.SysRoleMenu{RoleID: roleSonMenu.RoleId, MenuID: addY}
		if err := model.Create(item); err != nil {
			response.BaseError(c, fmt.Sprintf("SysRoleMenu-新增异常 role_id:%v menu_id:%v 错误：%v", roleSonMenu.RoleId, addY, err.Error()))
			return
		}
	}

	response.Success(c, "更新成功!")
}
