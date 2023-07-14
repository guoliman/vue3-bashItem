package request

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"vue3-bashItem/pkg/convert"
	"vue3-bashItem/pkg/utils"
)

// 分页处理
func GetPagerParams(c *gin.Context) (offset, limit int, err error) {
	// 分页参数 currentPage是页 pageSize是是页的条数
	offset, err = convert.StrTo(c.DefaultQuery("pageNum", "0")).Int()
	if err != nil {
		return offset, limit, errors.New("offset参数非法. ")
	}

	limit, err = convert.StrTo(c.DefaultQuery("pageSize", "10")).Int()
	if err != nil || (offset != 0 && limit == 0) {
		return offset, limit, errors.New("limit参数非法. ")
	}

	// 减1是第一天是0-10 所以减1 乘limit是右移的条数
	if offset != 0 {
		offset = (offset - 1) * limit
	}
	return offset, limit, err
}

// table页排序
func GetOrderParams(c *gin.Context) (orderExpr string, err error) {
	if orderExpr = c.Query("order"); orderExpr != "" {
		orderExprSlice := strings.Split(orderExpr, " ")
		if len(orderExprSlice) != 2 {
			err = errors.New("order参数非法. ")
			return "", err
		}
		if ix := utils.StringInSlice([]string{"asc", "desc"}, orderExprSlice[1]); ix == false {
			err = errors.New("order参数非法. ")
		}
	}
	return orderExpr, err
}

// sql拼接  1两边包含 2前面包含 3后边包含 4完全匹配
func GetLikeFilterQuery(c *gin.Context, fields []string, condition int) (string, []interface{}) {
	queryFieldList := make([]string, 0)
	queryValueList := make([]interface{}, 0)
	for _, field := range fields {
		if value := c.Query(field); len(value) != 0 {
			// 默认 全查询
			queryField := fmt.Sprintf(" %s like ?", field)
			queryValue := fmt.Sprintf("%%%s%%", value)

			if condition == 2 { // 包含前面
				queryValue = fmt.Sprintf("%%%s", value)
			} else if condition == 3 { // 包含后面
				queryValue = fmt.Sprintf("%s%%", value)
			} else if condition == 4 { // 包含后面
				queryField = fmt.Sprintf(" %s = ?", field)
				queryValue = fmt.Sprintf("%s", value)
			}
			queryFieldList = append(queryFieldList, queryField)
			queryValueList = append(queryValueList, queryValue)
		}
	}
	if len(queryFieldList) != 0 {
		likeExpr := strings.Join(queryFieldList, " and ")
		return likeExpr, queryValueList
	}
	return "", nil
}
