package dictApi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
	"vue3-bashItem/model"
	"vue3-bashItem/model/accountAdminModel/dictModel"
	"vue3-bashItem/pkg/logger"
	"vue3-bashItem/pkg/request"
	"vue3-bashItem/pkg/response"
	"vue3-bashItem/pkg/utils"
	"vue3-bashItem/services/accountAdminMethod/dictMethod"
)

// GetDictSon 获取字典子级
func GetDictSon(c *gin.Context) {
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
	validFilterFields := []string{"type_code", "name"}
	likeExpr, likeValues := request.GetLikeFilterQuery(c, validFilterFields, 4)

	// 查询db获取数据
	dictData := make([]*dictModel.SysDict, 0)
	dataNum, getErr := model.GetListPage(&dictData, offset, limit, orderExpr, likeExpr, likeValues...)
	if getErr != nil {
		response.BaseError(c, fmt.Sprintf("查询db获取数据失败 %v", err.Error()))
		return
	}

	var dictPyteList = make([]map[string]interface{}, 0)
	for _, dictY := range dictData {
		oneDictPyte := map[string]interface{}{
			"id":        dictY.Id,
			"type_code": dictY.TypeCode,
			"name":      dictY.Name,
			"value":     dictY.Value,
			"sort":      dictY.Sort,
			"status":    dictY.Status,
			//"defaulted":   dictY.Defaulted,
			"remark":      dictY.Remark,
			"create_time": dictY.CreateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 更改时间
			"update_time": dictY.UpdateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 更改时间
		}
		dictPyteList = append(dictPyteList, oneDictPyte)
	}
	var dictPyteResult = make(map[string]interface{})

	intData, intErr := utils.SortListInt64(dictPyteList, "sort", "ask") // ask正序 desc倒序
	if intErr != nil {
		logger.FileLogger.Error(fmt.Sprintf("数据没排序 因数字排序处理异常: %v", intErr))
		dictPyteResult["list"] = dictPyteList
	} else {
		dictPyteResult["list"] = intData
	}

	dictPyteResult["total"] = dataNum

	response.Success(c, dictPyteResult)
}

// CreateDictSon 创建字典子级
func CreateDictSon(c *gin.Context) {
	reqBody := &dictMethod.DictSonData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("创建接口-获取数据失败%v", err.Error()))
		return
	}

	// 字段值不能重复
	catDict := dictModel.SysDict{}
	dictErr := model.GetOneLast(&catDict, "type_code = ? and value = ?", reqBody.TypeCode, reqBody.Value)
	if dictErr == nil {
		response.ParamsError(c, "字典值已存在 请更改")
		return
	}

	nowTime := time.Now()
	item := &dictModel.SysDict{
		TypeCode:   reqBody.TypeCode,
		Name:       reqBody.Name,
		Value:      reqBody.Value,
		Sort:       reqBody.Sort,
		Status:     reqBody.Status,
		Defaulted:  reqBody.Defaulted,
		Remark:     reqBody.Remark,
		CreateTime: nowTime,
		UpdateTime: nowTime,
	}
	if err := model.Create(item); err != nil {
		response.BaseError(c, fmt.Sprintf("DictData-创建数据错误：%v", err.Error()))
		return
	}
	response.Success(c, item)
}

// PutDictSon 更新字典子级
func PutDictSon(c *gin.Context) {
	reqBody := &dictMethod.DictSonData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("获取参数错误：%v", err.Error()))
		return
	}

	// 字段值不能重复
	catDict := dictModel.SysDict{}
	dictErr := model.GetOneLast(&catDict, "id != ? and type_code = ? and value = ?", reqBody.Id, reqBody.TypeCode, reqBody.Value)
	if dictErr == nil {
		response.ParamsError(c, "字典值已存在 请更改")
		return
	}

	// 创建时间 string转time 东八区
	local, _ := time.LoadLocation("Asia/Shanghai")
	createTime, _ := time.ParseInLocation("2006-01-02 15:04:05", reqBody.CreateTime, local) // 转成东8区time
	updateTime := time.Now()                                                                // 更新时间为当前时间

	updateColumns := map[string]interface{}{
		"TypeCode":   reqBody.TypeCode,
		"Name":       reqBody.Name,
		"Value":      reqBody.Value,
		"Sort":       reqBody.Sort,
		"Status":     reqBody.Status,
		"Defaulted":  reqBody.Defaulted,
		"Remark":     reqBody.Remark,
		"CreateTime": createTime,
		"UpdateTime": updateTime,
	}
	// 已软删除的id 不会更改  但会返回更改成功
	if err := model.Update(&dictModel.SysDict{}, updateColumns, "id = ?", reqBody.Id); err != nil {
		response.BaseError(c, fmt.Sprintf("SysDict表更新错误：%v", err.Error()))
		return
	}

	response.Success(c, "更新成功!")
}

// DelDictSon 删除字典子级
func DelDictSon(c *gin.Context) {
	delId := &dictMethod.DelId{}
	if err := c.ShouldBindJSON(&delId); err != nil {
		response.ParamsError(c, fmt.Sprintf("删除接口，获取数据失败%v", err.Error()))
		return
	}

	// 批量删除
	idList := strings.Split(delId.Id, ",")
	for iNum := 0; iNum < len(idList); iNum++ {
		idInt, intErr := strconv.ParseInt(idList[iNum], 10, 64) //string转int64
		if intErr != nil {
			response.BaseError(c, fmt.Sprintf("：%v 转译失败 %v", idList[iNum], intErr.Error()))
			return
		}

		if err := model.DeleteByQuery(&dictModel.SysDict{}, map[string]interface{}{"id": idInt}); err != nil {
			response.BaseError(c, fmt.Sprintf("SysDictType-表删除数据错误：%v", err.Error()))
			return
		}
	}

	response.Success(c, "删除成功!")
}
