package menuApi

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"vue3-bashItem/model"
	"vue3-bashItem/model/accountAdminModel/menuModel"
	"vue3-bashItem/model/accountAdminModel/roleModel"
	"vue3-bashItem/pkg/request"
	"vue3-bashItem/pkg/response"
	"vue3-bashItem/pkg/utils"
	"vue3-bashItem/services/accountAdminMethod/deptMethod"
	"vue3-bashItem/services/accountAdminMethod/menuMethod"
	"vue3-bashItem/services/accountAdminMethod/roleMethod"
)

// MenuRoutes 左侧动态路由 假数据
func MenuRoutes(c *gin.Context) {

	jsonData := `
[{"path":"/system","component":"Layout","redirect":"/system/user","meta":{"title":"系统管理","icon":"system","hidden":false,"roles":["ADMIN"],"keepAlive":true},"children":[{"path":"user","component":"system/user/index","name":"User","meta":{"title":"用户管理","icon":"user","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"role","component":"system/role/index","name":"Role","meta":{"title":"角色管理","icon":"role","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"menu","component":"system/menu/index","name":"Menu","meta":{"title":"菜单管理","icon":"menu","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"dept","component":"system/dept/index","name":"Dept","meta":{"title":"部门管理","icon":"tree","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"dict","component":"system/dict/index","name":"Dict","meta":{"title":"字典管理","icon":"dict","hidden":false,"roles":["ADMIN"],"keepAlive":true}}]},{"path":"/api","component":"Layout","meta":{"title":"接口","icon":"api","hidden":false,"roles":["ADMIN"],"keepAlive":true},"children":[{"path":"apidoc","component":"demo/api-doc","name":"Apidoc","meta":{"title":"接口文档","icon":"api","hidden":false,"roles":["ADMIN"],"keepAlive":true}}]},{"path":"/external-link","component":"Layout","redirect":"noredirect","meta":{"title":"外部链接","icon":"link","hidden":false,"roles":["ADMIN"],"keepAlive":true},"children":[{"path":"https://juejin.cn/post/7228990409909108793","meta":{"title":"document","icon":"document","hidden":false,"roles":["ADMIN"],"keepAlive":true}}]},{"path":"/multi-level","component":"Layout","redirect":"/multi-level/multi-level1","meta":{"title":"多级菜单","icon":"multi_level","hidden":false,"roles":["ADMIN"],"keepAlive":true},"children":[{"path":"multi-level1","component":"demo/multi-level/level1","redirect":"/multi-level/multi-level2","meta":{"title":"菜单一级","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true},"children":[{"path":"multi-level2","component":"demo/multi-level/children/level2","redirect":"/multi-level/multi-level2/multi-level3-1","meta":{"title":"菜单二级","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true},"children":[{"path":"multi-level3-1","component":"demo/multi-level/children/children/level3-1","name":"MultiLevel31","meta":{"title":"菜单三级-1","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"multi-level3-2","component":"demo/multi-level/children/children/level3-2","name":"MultiLevel32","meta":{"title":"菜单三级-2","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true}}]}]}]},{"path":"/component","component":"Layout","meta":{"title":"组件封装","icon":"menu","hidden":false,"roles":["ADMIN"],"keepAlive":true},"children":[{"path":"wang-editor","component":"demo/wang-editor","name":"WangEditor","meta":{"title":"富文本编辑器","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"upload","component":"demo/upload","name":"Upload","meta":{"title":"图片上传","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"icon-selector","component":"demo/icon-selector","name":"IconSelector","meta":{"title":"图标选择器","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"taginput","component":"demo/taginput","name":"Taginput","meta":{"title":"标签输入框","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"signature","component":"demo/signature","name":"Signature","meta":{"title":"签名","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"table","component":"demo/table","name":"Table","meta":{"title":"表格","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true}}]},{"path":"/function","component":"Layout","meta":{"title":"功能演示","icon":"menu","hidden":false,"roles":["ADMIN"],"keepAlive":true},"children":[{"path":"websocket","component":"demo/websocket","name":"Websocket","meta":{"title":"Websocket","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true}},{"path":"other","component":"demo/other","meta":{"title":"敬请期待...","icon":"","hidden":false,"roles":["ADMIN"],"keepAlive":true}}]}]
`

	var resultData []interface{}
	if err := json.Unmarshal([]byte(jsonData), &resultData); err != nil {
		response.BaseError(c, fmt.Sprintf("json转译失败 %v", err))
		return
	}

	//msg := "一切ok"
	response.Vue3Response(c, resultData)

}

// GetMoveRoute 左侧动态路由 真数据
func GetMoveRoute(c *gin.Context) {
	// 角色获取菜单列表
	_, roleCodeList, menuIdList, menuIdListErr := roleMethod.RoleGetMenu(c)
	if menuIdListErr != nil {
		response.BaseError(c, fmt.Sprintf("角色获取菜单异常: %v ", menuIdListErr.Error()))
		return
	}

	// 根据菜单id查菜单数据
	menuM := make([]menuModel.SysMenu, 0)
	roleIdErr := model.Db.Where("type NOT IN ? AND id In ?", []int64{4}, menuIdList).Find(&menuM).Error
	if roleIdErr != nil {
		response.BaseError(c, fmt.Sprintf("按钮查询异常 %v ", roleIdErr.Error()))
		return
	}

	var menuList []int64
	for _, menuI := range menuM {
		menuList = append(menuList, menuI.Id)
	}
	menuList = utils.RemoveListInt64(menuList) // 去重

	// 递归权限 获取路由
	menuresult := menuMethod.GetMenuTree(0, menuList, roleCodeList)

	response.Success(c, menuresult)
}

// GetMenu 获取全部菜单
func GetMenu(c *gin.Context) {
	// 分页参数 offset是偏移 limit是条数
	offset, limit := 0, 10000
	// 排序条件
	orderExpr := "id asc"

	// 多条件like查询 数据拼接成string
	validFilterFields := []string{"name"}
	likeExpr, likeValues := request.GetLikeFilterQuery(c, validFilterFields, 1)

	// 查询db获取数据
	deptData := make([]*menuModel.SysMenu, 0)
	_, getErr := model.GetListPage(&deptData, offset, limit, orderExpr, likeExpr, likeValues...)
	if getErr != nil {
		response.BaseError(c, fmt.Sprintf("查询db获取数据失败 %v", getErr.Error()))
		return
	}

	var deptList = make([]map[string]interface{}, 0)
	for _, menuY := range deptData {
		oneDept := map[string]interface{}{
			"id":         menuY.Id,
			"parentId":   menuY.ParentID,
			"name":       menuY.Name,
			"type":       menuMethod.TypeUintToStr(menuY.Type), // 类型取对应string
			"path":       menuY.Path,
			"component":  menuY.Component,
			"perm":       menuY.Perm,
			"visible":    menuY.Visible,
			"sort":       menuY.Sort,
			"icon":       menuY.Icon,
			"redirect":   menuY.Redirect,
			"apiType":    menuY.ApiType,
			"apiPath":    menuY.ApiPath,
			"createTime": menuY.CreateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 创建时间
			"updateTime": menuY.UpdateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 更改时间
		}
		deptList = append(deptList, oneDept)
	}

	selectStatus := false // 判断是否查询 有查询就用不递归了
	for _, field := range validFilterFields {
		if value := c.Query(field); len(value) != 0 {
			selectStatus = true
		}
	}
	if selectStatus {
		response.Success(c, deptList)
		return
	}

	//deptResult := menuMethod.RecMenu(0, deptList)
	deptResult := deptMethod.RecDept(0, deptList) // 公用方法 递归获取全部数据

	response.Success(c, deptResult)
}

// DirAndMenuSelect 获取目录和菜单不取按钮(用于新增或修改菜单时使用) 递归
func DirAndMenuSelect(c *gin.Context) {
	deptData := make([]*menuModel.SysMenu, 0)
	deptDataErr := model.Db.Where("type in ?", []int64{1, 2}).Find(&deptData).Error // 4是按钮
	if deptDataErr != nil {
		response.BaseError(c, fmt.Sprintf("查询SysMenu失败 %v", deptDataErr.Error()))
		return
	}

	var deptList = make([]map[string]interface{}, 0)
	for _, menuY := range deptData {
		oneDept := map[string]interface{}{
			"value":    menuY.Id,
			"label":    menuY.Name,
			"parentId": menuY.ParentID,
			"sort":     menuY.Sort,
		}
		deptList = append(deptList, oneDept)
	}
	deptResult := deptMethod.RecDeptSelect(0, deptList)
	response.Success(c, deptResult)
}

// RoleGetMenuSelect 角色分配权限时获取下拉菜单 递归
func RoleGetMenuSelect(c *gin.Context) {

	deptData := make([]*menuModel.SysMenu, 0)
	deptDataErr := model.GetAll(&deptData)
	if deptDataErr != nil {
		response.BaseError(c, fmt.Sprintf("查询SysMenu失败 %v", deptDataErr.Error()))
	}

	var deptList = make([]map[string]interface{}, 0)
	for _, menuY := range deptData {
		oneDept := map[string]interface{}{
			"value":    menuY.Id,
			"label":    menuY.Name,
			"parentId": menuY.ParentID,
			"sort":     menuY.Sort,
		}
		deptList = append(deptList, oneDept)
	}
	deptResult := deptMethod.RecDeptSelect(0, deptList)
	response.Success(c, deptResult)
}

// CreateMenu 创建菜单
func CreateMenu(c *gin.Context) {
	reqBody := &menuMethod.MenuData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("创建接口-获取数据失败%v", err.Error()))
		return
	}

	// 判断目录层级能否能存放
	if reqBody.ParentID != 0 { // 不是顶层目录
		if levelJudgmentErr := menuMethod.LevelJudgment(reqBody); levelJudgmentErr != nil {
			response.ParamsError(c, fmt.Sprintf("层级问题: %v", levelJudgmentErr.Error()))
			return
		}
	}

	// 判断类型是2(目录)时 compoment设为Layout path必须加/
	createType := menuMethod.TypeStrToUint(reqBody.Type) // 类型取对应int
	createComponent := reqBody.Component
	createPath := reqBody.Path
	if createType == 2 {
		createComponent = "Layout"
		if createPath[0] != '/' {
			createPath = "/" + createPath
		}
	}

	nowTime := time.Now()
	item := &menuModel.SysMenu{
		Type:       createType,
		Component:  createComponent,
		Path:       createPath,
		ParentID:   reqBody.ParentID,
		Name:       reqBody.Name,
		Perm:       reqBody.Perm,
		Visible:    reqBody.Visible,
		Sort:       reqBody.Sort,
		Icon:       reqBody.Icon,
		Redirect:   reqBody.Redirect,
		ApiType:    reqBody.ApiType,
		ApiPath:    reqBody.ApiPath,
		CreateTime: nowTime,
		UpdateTime: nowTime,
	}
	if err := model.Create(item); err != nil {
		response.BaseError(c, fmt.Sprintf("SysDept-创建数据错误：%v", err.Error()))
		return
	}
	response.Success(c, item)
}

// PutMenu 更新菜单
func PutMenu(c *gin.Context) {
	reqBody := &menuMethod.MenuData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("获取参数错误：%v", err.Error()))
		return
	}

	// 判断目录层级能否能存放
	if reqBody.ParentID != 0 { // 不是顶层目录
		if levelJudgmentErr := menuMethod.LevelJudgment(reqBody); levelJudgmentErr != nil {
			response.ParamsError(c, fmt.Sprintf("层级问题: %v", levelJudgmentErr.Error()))
			return
		}
	}

	// 创建时间 string转time 东八区
	local, _ := time.LoadLocation("Asia/Shanghai")
	createTime, _ := time.ParseInLocation("2006-01-02 15:04:05", reqBody.CreateTime, local) // 转成东8区time
	updateTime := time.Now()                                                                // 更新时间为当前时间

	// 判断类型是2(目录)时 compoment设为Layout path必须加/
	createType := menuMethod.TypeStrToUint(reqBody.Type) // 类型取对应int
	createComponent := reqBody.Component
	createPath := reqBody.Path
	if createType == 2 {
		createComponent = "Layout"
		if createPath[0] != '/' {
			createPath = "/" + createPath
		}
	}

	updateColumns := map[string]interface{}{
		"Type":       createType,
		"Path":       createPath,
		"Component":  createComponent,
		"ParentID":   reqBody.ParentID,
		"Name":       reqBody.Name,
		"Perm":       reqBody.Perm,
		"Visible":    reqBody.Visible,
		"Sort":       reqBody.Sort,
		"Icon":       reqBody.Icon,
		"Redirect":   reqBody.Redirect,
		"ApiType":    reqBody.ApiType,
		"ApiPath":    reqBody.ApiPath,
		"CreateTime": createTime,
		"UpdateTime": updateTime,
	}
	// 已软删除的id 不会更改  但会返回更改成功
	if err := model.Update(&menuModel.SysMenu{}, updateColumns, "id = ?", reqBody.Id); err != nil {
		response.BaseError(c, fmt.Sprintf("SysDept表更新错误：%v", err.Error()))
		return
	}

	response.Success(c, "更新成功!")
}

// DelMenu 删除菜单
func DelMenu(c *gin.Context) {
	// 获取参数
	delId := &menuMethod.DelId{}
	if err := c.ShouldBindJSON(&delId); err != nil {
		response.ParamsError(c, fmt.Sprintf("删除接口，获取数据失败%v", err.Error()))
		return
	}

	// 获取全部菜单数据
	menuData := make([]*menuModel.SysMenu, 0)
	menuDataErr := model.GetAll(&menuData)
	if menuDataErr != nil {
		response.BaseError(c, fmt.Sprintf("查询失败 %v", menuDataErr.Error()))
	}
	var menuList = make([]map[string]interface{}, 0)
	for _, menuY := range menuData {
		menuList = append(menuList, map[string]interface{}{"id": menuY.Id, "parentId": menuY.ParentID})
	}

	// 取要删除的id列表
	var delMenuData []int64
	delMenuData = append(delMenuData, delId.Id)
	delMenuData = append(delMenuData, deptMethod.RecDeptDelList(delId.Id, menuList)...)

	// 删除
	for iData := 0; iData < len(delMenuData); iData++ {
		// 删除菜单关联角色数据
		menuRoleErr := model.DeleteByQuery(&roleModel.SysRoleMenu{}, map[string]interface{}{"menu_id": delMenuData[iData]})
		if menuRoleErr != nil {
			response.BaseError(c, fmt.Sprintf("menu_id:%v 删除菜单关联角色数据异常：%v", iData, menuRoleErr.Error()))
			return
		}
		// 删除菜单
		menuErr := model.DeleteByQuery(&menuModel.SysMenu{}, map[string]interface{}{"id": delMenuData[iData]})
		if menuErr != nil {
			response.BaseError(c, fmt.Sprintf("id:%v 删除菜单异常：%v", iData, menuErr.Error()))
			return
		}
	}
	response.Success(c, "删除成功!")
}
