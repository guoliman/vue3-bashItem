package menuMethod

import (
	"errors"
	"fmt"
	"vue3-bashItem/model"
	"vue3-bashItem/model/accountAdminModel/menuModel"
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/utils"
	"vue3-bashItem/services/accountAdminMethod/deptMethod"
)

// SysMenu 菜单管理
type MenuData struct {
	Id         int64  `comment:"主键" json:"id"`
	ParentID   int64  `gorm:"comment:父节点id" json:"parentId"`
	Name       string `gorm:"comment:菜单名称" json:"name"`
	Type       string `gorm:"comment:菜单类型(1:菜单；2:目录；3:外链；4:按钮)" json:"type"`
	Path       string `gorm:"comment:路由路径(浏览器地址栏路径)" json:"path"`
	Component  string `gorm:"comment:组件路径(vue页面完整路径)" json:"component"`
	Perm       string `gorm:"comment:权限标识" json:"perm"`
	Visible    uint8  `gorm:"comment:显示状态(1-显示;0-隐藏)" json:"visible"`
	Sort       int64  `gorm:"comment:排序" json:"sort"`
	Icon       string `gorm:"comment:菜单图标" json:"icon"`
	Redirect   string `gorm:"comment:跳转路径" json:"redirect"`
	ApiType    string `gorm:"comment:后端接口类型" json:"apiType"`
	ApiPath    string `gorm:"comment:后端请求地址" json:"apiPath"`
	CreateTime string `gorm:"comment:创建时间" json:"createTime"`
	UpdateTime string `gorm:"comment:更新时间" json:"updateTime"`
}

// DelId 删除id
type DelId struct {
	Id int64 `comment:"主键" json:"id"`
}

// TypeUintToStr 菜单类型uint8转换string
func TypeUintToStr(typeData uint8) (result string) {
	if typeData == uint8(1) {
		return "MENU"
	} else if typeData == uint8(2) {
		return "CATALOG"
	} else if typeData == uint8(3) {
		return "EXTLINK"
	} else if typeData == uint8(4) {
		return "BUTTON"
	}
	return "未知"
}

// TypeStrToUint 菜单类型string转换uint8
func TypeStrToUint(typeData string) (result uint8) {
	if typeData == string("MENU") {
		return uint8(1)
	} else if typeData == string("CATALOG") {
		return uint8(2)
	} else if typeData == string("EXTLINK") {
		return uint8(3)
	} else if typeData == string("BUTTON") {
		return uint8(4)
	}
	return uint8(4)
}

// GetMenuTree 获取菜单递归数据
func GetMenuTree(parentMenuMdl int64, roleInt []int64, roleList []string) (r []map[string]interface{}) {
	menuQs := make([]*menuModel.SysMenu, 0)
	model.GetList(&menuQs, "parent_id = ? And id In ?", parentMenuMdl, roleInt)

	results := make([]map[string]interface{}, 0)
	for _, menuMdl := range menuQs {
		//logger.FileLogger.Debug("menuMdl.Component===", menuMdl.Component)
		menuData := map[string]interface{}{
			"component": menuMdl.Component,
			"path":      menuMdl.Path,
			"redirect":  menuMdl.Redirect,
			"name":      menuMdl.Name,
			"sort":      menuMdl.Sort,
			"meta":      make(map[string]interface{}),
		}
		hiddenData := true
		if menuMdl.Visible == 1 {
			hiddenData = false
		}
		menuData["meta"] = map[string]interface{}{
			"title":     menuMdl.Name,
			"hidden":    hiddenData,
			"icon":      menuMdl.Icon,
			"keepAlive": true, // 是否缓存 以后加字段 true缓存 false不缓存
			"roles":     roleList,
		}

		// 判断是否有子级
		menuConut, menuConutErr := model.Count(&menuModel.SysMenu{},
			"parent_id = ? AND id IN ?", menuMdl.Id, roleInt)
		if menuConutErr == nil && menuConut >= 1 { //没有报错并且有值
			childrenList := make([]map[string]interface{}, 0) // children值
			for _, sonM := range GetMenuTree(menuMdl.Id, roleInt, roleList) {
				if sonM["type"] != 4 {
					childrenList = append(childrenList, sonM)
				}
			}
			menuData["children"] = childrenList // 增加子级
		}
		results = append(results, menuData)

		// 左侧菜单排序 根据sort字段排序  格式 [{"sort":2,"aa":"a1"},{"sort":1,"bb":"b1"}]
		var intErr error
		results, intErr = utils.SortListInt64(results, "sort", "ask")
		if intErr != nil {
			logger.FileLogger.Error("排序错误")
		}
	}
	return results
}

// LevelJudgment 判断目录层级能否能存放
func LevelJudgment(reqBody *MenuData) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	// 取pid数据
	menuPid := menuModel.SysMenu{}
	if menuPidErr := model.GetOneLast(&menuPid, "id = ?", reqBody.ParentID); menuPidErr != nil {
		return errors.New(fmt.Sprintf("pid:%v不存在 err: %v ", reqBody.ParentID, menuPidErr.Error()))
	}

	// 取自己子级 生成列表
	menuAllData := make([]*menuModel.SysMenu, 0)
	menuAllDataErr := model.GetAll(&menuAllData)
	if menuAllDataErr != nil {
		return errors.New(fmt.Sprintf("查询SysMenu失败 %v", menuAllDataErr.Error()))
	}
	var menuList = make([]map[string]interface{}, 0)
	for _, menuY := range menuAllData {
		menuList = append(menuList, map[string]interface{}{"id": menuY.Id, "parentId": menuY.ParentID})
	}

	createType := TypeStrToUint(reqBody.Type) // 类型取对应int

	// 类型 1:菜单；2:目录；3:外链；4:按钮
	if createType == 4 {
		if menuPid.Type != 1 {
			return errors.New("父级必须是菜单,请更改")
		}
	} else if createType == 2 {
		if menuPid.Type != 2 {
			return errors.New("父级必须是目录,请更改")
		}
	} else if createType == 1 || createType == 3 {
		if menuPid.Type == 3 || menuPid.Type == 4 {
			return errors.New("父级必须是目录或菜单,请更改")
		}
	}

	// 父级不能是自身子级
	var idList []int64
	if reqBody.Id != 0 { // 创建接口 id不存在默认是0
		idList = append(idList, reqBody.Id)
		idList = append(idList, deptMethod.RecDeptDelList(reqBody.Id, menuList)...)
	}
	idListStatus := utils.JudgeInt64InList(menuPid.Id, idList)
	logger.FileLogger.Debug(fmt.Sprintf("父级id:%v, 父级类型:%v, 类型:%v, 子级id列表:%v, 是否为自身子级:%v",
		menuPid.Id, menuPid.Type, createType, idList, idListStatus))
	if idListStatus { // 父级不能是自身子级
		return errors.New("父级不能是自身子级,请更改")
	}

	return nil
}
