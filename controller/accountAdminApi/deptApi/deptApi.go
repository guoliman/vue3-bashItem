package deptApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
	"vue3-bashItem/model"
	"vue3-bashItem/model/accountAdminModel/deptModel"
	"vue3-bashItem/model/accountAdminModel/userModel"
	"vue3-bashItem/pkg/request"
	"vue3-bashItem/pkg/response"
	"vue3-bashItem/pkg/utils"
	"vue3-bashItem/services/accountAdminMethod/deptMethod"
	"vue3-bashItem/services/accountAdminMethod/dictMethod"
)

// SelectDept 获取部门下拉菜单 递归
func SelectDept(c *gin.Context) {

	deptData := make([]*deptModel.SysDept, 0)
	deptDataErr := model.GetAll(&deptData)
	if deptDataErr != nil {
		response.BaseError(c, fmt.Sprintf("查询SysDept失败 %v", deptDataErr.Error()))
	}

	var deptList = make([]map[string]interface{}, 0)
	for _, deptY := range deptData {
		oneDept := map[string]interface{}{
			"value":    deptY.Id,
			"label":    deptY.Name,
			"parentId": deptY.ParentID,
			"sort":     deptY.Sort,
		}
		deptList = append(deptList, oneDept)
	}
	deptResult := deptMethod.RecDeptSelect(0, deptList)
	response.Success(c, deptResult)
}

// GetDept 获取部门 递归
func GetDept(c *gin.Context) {
	// 分页参数 offset是偏移 limit是条数
	offset, limit := 0, 10000
	// 排序条件
	orderExpr := "id asc"

	// 多条件like查询 数据拼接成string
	validFilterFields := []string{"name", "status"}
	likeExpr, likeValues := request.GetLikeFilterQuery(c, validFilterFields, 1)

	// 查询db获取数据
	deptData := make([]*deptModel.SysDept, 0)
	_, getErr := model.GetListPage(&deptData, offset, limit, orderExpr, likeExpr, likeValues...)
	if getErr != nil {
		response.BaseError(c, fmt.Sprintf("查询db获取数据失败 %v", getErr.Error()))
		return
	}

	var deptList = make([]map[string]interface{}, 0)
	for _, deptY := range deptData {
		oneDept := map[string]interface{}{
			"id":         deptY.Id,
			"name":       deptY.Name,
			"parentId":   deptY.ParentID,
			"sort":       deptY.Sort,
			"status":     deptY.Status,
			"deleted":    deptY.Deleted,
			"createTime": deptY.CreateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 创建时间
			"updateTime": deptY.UpdateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 更改时间
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

	deptResult := deptMethod.RecDept(0, deptList)

	response.Success(c, deptResult)
}

// CreateDept 创建部门
func CreateDept(c *gin.Context) {
	reqBody := &deptMethod.DeptData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("创建接口-获取数据失败%v", err.Error()))
		return
	}

	nowTime := time.Now()
	item := &deptModel.SysDept{
		Name:       reqBody.Name,
		ParentID:   reqBody.ParentID,
		Sort:       reqBody.Sort,
		Status:     reqBody.Status,
		Deleted:    reqBody.Deleted,
		CreateTime: nowTime,
		UpdateTime: nowTime,
	}
	if err := model.Create(item); err != nil {
		response.BaseError(c, fmt.Sprintf("SysDept-创建数据错误：%v", err.Error()))
		return
	}
	response.Success(c, item)
}

// PutDept 更新部门
func PutDept(c *gin.Context) {
	reqBody := &deptMethod.DeptData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("获取参数错误：%v", err.Error()))
		return
	}
	// 创建时间 string转time 东八区
	local, _ := time.LoadLocation("Asia/Shanghai")
	createTime, _ := time.ParseInLocation("2006-01-02 15:04:05", reqBody.CreateTime, local) // 转成东8区time
	updateTime := time.Now()                                                                // 更新时间为当前时间

	updateColumns := map[string]interface{}{
		"Name":       reqBody.Name,
		"ParentID":   reqBody.ParentID,
		"Sort":       reqBody.Sort,
		"Status":     reqBody.Status,
		"Deleted":    reqBody.Deleted,
		"CreateTime": createTime,
		"UpdateTime": updateTime,
	}
	// 已软删除的id 不会更改  但会返回更改成功
	if err := model.Update(&deptModel.SysDept{}, updateColumns, "id = ?", reqBody.Id); err != nil {
		response.BaseError(c, fmt.Sprintf("SysDept表更新错误：%v", err.Error()))
		return
	}

	response.Success(c, "更新成功!")
}

// DelDept 删除部门
func DelDept(c *gin.Context) {
	// 获取参数
	delId := &dictMethod.DelId{}
	if err := c.ShouldBindJSON(&delId); err != nil {
		response.ParamsError(c, fmt.Sprintf("删除接口，获取数据失败%v", err.Error()))
		return
	}

	// 获取部门全部数据
	deptData := make([]*deptModel.SysDept, 0)
	deptDataErr := model.GetAll(&deptData)
	if deptDataErr != nil {
		response.BaseError(c, fmt.Sprintf("查询SysDept失败 %v", deptDataErr.Error()))
	}
	var deptList = make([]map[string]interface{}, 0)
	for _, deptY := range deptData {
		deptList = append(deptList, map[string]interface{}{"id": deptY.Id, "parentId": deptY.ParentID})
	}
	// 取要删除的id列表
	var delDeptData []int64
	idList := strings.Split(delId.Id, ",") // string转[]int
	for iNum := 0; iNum < len(idList); iNum++ {
		idInt, intErr := strconv.ParseInt(idList[iNum], 10, 64) //string转int64
		if intErr != nil {
			response.BaseError(c, fmt.Sprintf("：%v 转译失败 %v", idList[iNum], intErr.Error()))
			return
		}
		delDeptData = append(delDeptData, idInt)
		delDeptData = append(delDeptData, deptMethod.RecDeptDelList(idInt, deptList)...)
	}
	delDeptDataList := utils.RemoveListInt64(delDeptData) // 去重

	var userData = make([]*userModel.SysUser, 0)
	if userDataErr := model.Db.Where("dept_id in ?", delDeptDataList).Find(&userData).Error; userDataErr != nil {
		response.BaseError(c, fmt.Sprintf("dept_id:%v SysUser查询异常：%v", delDeptDataList, userDataErr.Error()))
		return
	} else if len(userData) != 0 {
		response.BaseError(c, "请清空部门内用户 再删除部门")
		return
	}
	//logger.FileLogger.Debug("userData====", userData)

	// 删除
	for iData := 0; iData < len(delDeptDataList); iData++ {
		err := model.DeleteByQuery(&deptModel.SysDept{}, map[string]interface{}{"id": delDeptDataList[iData]})
		if err != nil {
			response.BaseError(c, fmt.Sprintf("SysDept表 id:%v 删除数据错误：%v", iData, err.Error()))
			return
		}
	}
	response.Success(c, "删除成功!")
}
