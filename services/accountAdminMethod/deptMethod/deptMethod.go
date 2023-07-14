package deptMethod

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
	"vue3-bashItem/model"
	"vue3-bashItem/model/accountAdminModel/deptModel"
	"vue3-bashItem/pkg/logger"
)

// DeptData 部门操作类型
type DeptData struct {
	Id         int64  `comment:"主键" json:"id"`
	Name       string `comment:"部门名称" json:"name"`
	ParentID   int64  `comment:"父节点id" json:"parentId"`
	Sort       int64  `comment:"显示顺序" json:"sort"`
	Status     uint8  `comment:"状态(1:正常;0:禁用)" json:"status"`
	Deleted    uint8  `comment:"逻辑删除标识(1:已删除;0:未删除)" json:"deleted"`
	CreateTime string `comment:"创建时间" json:"createTime"`
	UpdateTime string `comment:"更新时间" json:"updateTime"`
}

// DelId 删除id
type DelId struct {
	Id string `comment:"主键" json:"id"`
}

// RecDept 部门递归
func RecDept(deptPid int64, deptData []map[string]interface{}) (result []map[string]interface{}) {
	var changeData []map[string]interface{}
	//logger.FileLogger.Info(fmt.Sprintf("deptPid===%v,changeData===%v", deptPid, deptData))
	for deptI := 0; deptI < len(deptData); deptI++ {
		// 看下类型 可以类型不同 导致条件不成立
		//logger.FileLogger.Info(fmt.Sprintf("parent_id类型:%v, deptPid类型:%v",reflect.TypeOf(deptData[deptI]["parent_id"]),reflect.TypeOf(deptPid)))
		if deptData[deptI]["parentId"] == deptPid {
			nextPid := deptData[deptI]["id"]
			deptData[deptI]["children"] = RecDept(nextPid.(int64), deptData)
			changeData = append(changeData, deptData[deptI])
		}
	}
	// 左侧菜单排序 根据sort字段排序  格式 [{"sort":2,"aa":"a1"},{"sort":1,"bb":"b1"}]
	sort.Slice(changeData, func(i, j int) bool {
		//return results[i]["sort"].(int) > results[j]["sort"].(int) // 降序
		return changeData[i]["sort"].(int64) < changeData[j]["sort"].(int64) // 升序
	})
	return changeData
}

// RecDeptSelect 部门下拉框递归
func RecDeptSelect(deptPid int64, deptData []map[string]interface{}) (result []map[string]interface{}) {
	var changeData []map[string]interface{}
	for deptI := 0; deptI < len(deptData); deptI++ {
		if deptData[deptI]["parentId"] == deptPid {
			nextPid := deptData[deptI]["value"]
			deptData[deptI]["children"] = RecDeptSelect(nextPid.(int64), deptData)
			changeData = append(changeData, deptData[deptI])
		}
	}
	// 左侧菜单排序 根据sort字段排序  格式 [{"sort":2,"aa":"a1"},{"sort":1,"bb":"b1"}]
	sort.Slice(changeData, func(i, j int) bool {
		//return results[i]["sort"].(int) > results[j]["sort"].(int) // 降序
		return changeData[i]["sort"].(int64) < changeData[j]["sort"].(int64) // 升序
	})
	return changeData
}

// 递归要删除的id
func RecDeptDelList(pid int64, listData []map[string]interface{}) (result []int64) {
	var data []int64
	for x := 0; x < len(listData); x++ {
		if listData[x]["parentId"].(int64) == pid {
			nextPid := listData[x]["id"].(int64)
			data = append(data, nextPid)
			data = append(data, RecDeptDelList(nextPid, listData)...)
		}
	}
	return data
}

// GetDeptIdList 获取部门id列表 递归
func GetDeptIdList(c *gin.Context, ) (result []int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	// 取全部部门
	allDept := make([]*deptModel.SysDept, 0)
	deptErr := model.GetAll(&allDept)
	if deptErr != nil {
		return nil, errors.New(fmt.Sprintf("获取部门失败 %v", deptErr.Error()))
	}
	var deptAll = make([]map[string]interface{}, 0)
	for _, deptAlone := range allDept {
		deptAll = append(deptAll, map[string]interface{}{"id": deptAlone.Id, "parentId": deptAlone.ParentID})
	}

	// 取部门筛选id
	deptSelect := deptModel.SysDept{}
	var deptId int64 = 0
	if deptValue := c.Query("deptId"); len(deptValue) != 0 {
		if deptErr := model.GetOneLast(&deptSelect, "id = ?", deptValue); deptErr != nil {
			logger.FileLogger.Error(fmt.Sprintf("部门id:%v 不存在 %v", deptValue, deptErr.Error()))
		}
		deptId = deptSelect.Id
	}
	logger.FileLogger.Info("前端发送部门id: ", deptId)

	// 取部门列表 递归
	var deptIds []int64
	deptIds = append(deptIds, deptId)
	deptIds = append(deptIds, RecDeptDelList(deptId, deptAll)...)

	return deptIds, nil
}
