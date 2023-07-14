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
	"vue3-bashItem/services/accountAdminMethod/dictMethod"
)

// GetDictType 获取字典类型
func GetDictType(c *gin.Context) {
	// 分页参数 offset是偏移 limit是条数
	offset, limit, err := request.GetPagerParams(c)
	if err != nil {
		response.ParamsError(c, err.Error())
		return
	}
	// 排序条件 orderExpr=id desc
	orderExpr, err := request.GetOrderParams(c)
	if err != nil {
		response.ParamsError(c, err.Error())
		return
	}

	// 多条件like查询 数据拼接成string
	validFilterFields := []string{"name"}
	likeExpr, likeValues := request.GetLikeFilterQuery(c, validFilterFields, 1)

	// 查询db获取数据
	dictData := make([]*dictModel.SysDictType, 0)
	dataNum, getErr := model.GetListPage(&dictData, offset, limit, orderExpr, likeExpr, likeValues...)
	if getErr != nil {
		response.BaseError(c, fmt.Sprintf("查询db获取数据失败 %v", err.Error()))
		return
	}

	var dictPyteList = make([]map[string]interface{}, 0)
	for _, dictY := range dictData {
		oneDictPyte := map[string]interface{}{
			"id":          dictY.Id,
			"name":        dictY.Name,
			"code":        dictY.Code,
			"status":      dictY.Status,
			"remark":      dictY.Remark,
			"create_time": dictY.CreateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 更改时间
			"update_time": dictY.UpdateTime.Format("2006-01-02 15:04:05"), // 时间类型转字符串, 更改时间
		}
		dictPyteList = append(dictPyteList, oneDictPyte)
	}
	var dictPyteResult = make(map[string]interface{})
	dictPyteResult["list"] = dictPyteList
	dictPyteResult["total"] = dataNum

	response.Success(c, dictPyteResult)
}

// CreateDictType 创建字典类型
func CreateDictType(c *gin.Context) {
	reqBody := &dictMethod.DictData{}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		response.ParamsError(c, fmt.Sprintf("创建接口-获取数据失败%v", err.Error()))
		return
	}

	nowTime := time.Now()
	item := &dictModel.SysDictType{
		Name:       reqBody.Name,
		Code:       reqBody.Code,
		Status:     reqBody.Status,
		Remark:     reqBody.Remark,
		CreateTime: nowTime,
		UpdateTime: nowTime,
	}
	if err := model.Create(item); err != nil {
		response.BaseError(c, fmt.Sprintf("SysDictType-创建数据错误：%v", err.Error()))
		return
	}
	response.Success(c, item)
}

// PutDictType 更新字典类型
func PutDictType(c *gin.Context) {
	reqBody := &dictMethod.DictData{}
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
		"Code":       reqBody.Code,
		"Status":     reqBody.Status,
		"Remark":     reqBody.Remark,
		"CreateTime": createTime,
		"UpdateTime": updateTime,
	}
	// 已软删除的id 不会更改  但会返回更改成功
	if err := model.Update(&dictModel.SysDictType{}, updateColumns, "id = ?", reqBody.Id); err != nil {
		response.BaseError(c, fmt.Sprintf("SysDictType表更新错误：%v", err.Error()))
		return
	}

	response.Success(c, "更新成功!")
}

// DelDictType 删除字典类型
func DelDictType(c *gin.Context) {
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

		dictTypeOther := dictModel.SysDictType{}
		if dictType := model.GetOneLast(&dictTypeOther, "id = ?", idInt); dictType != nil {
			logger.FileLogger.Info("SysDictType表id: %v  不存在 不用删除")
		} else {
			// 删除字典关联子级
			if err := model.DeleteByQuery(&dictModel.SysDictType{}, map[string]interface{}{"id": idInt}); err != nil {
				response.BaseError(c, fmt.Sprintf("SysDictType-表删除数据错误：%v", err.Error()))
				return
			}
			// 删除字典
			typeCode := map[string]interface{}{"type_code": dictTypeOther.Code}
			if err := model.DeleteByQuery(&dictModel.SysDict{}, typeCode); err != nil {
				response.BaseError(c, fmt.Sprintf("SysDict表code值：%v 删除异常：%v", dictTypeOther.Code, err.Error()))
				return
			}
		}
	}

	response.Success(c, "删除成功!")
}
