package utils

import (
	"errors"
	"fmt"
	"sort"
)

// SortListString 字符串排序
// [map[aa:bb sort:2] map[aa:cc sort:1] map[aa:aa sort:3]]
// intData, intErr := SortListNum(results, "aa", "ask") // ask正序 desc倒序  用aa排序
func SortListString(listData []map[string]interface{}, keyName string, order string) (result []map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("SortListString: %v", r)
		}
	}()
	if order == "ask" { // 字符串排序
		sort.Slice(listData, func(i, j int) bool { return listData[i][keyName].(string) < listData[j][keyName].(string) })
	} else if order == "desc" {
		sort.Slice(listData, func(i, j int) bool { return listData[i][keyName].(string) > listData[j][keyName].(string) })
	} else {
		return nil, errors.New("order填写ask或desc")
	}
	return listData, nil

}

// SortListNum 数字排序
// [map[aa:bb sort:2] map[aa:cc sort:1] map[aa:aa sort:3]]
// strData, strErr := SortListString(results, "sort", "ask") // ask正序 desc倒序  用sort排序
func SortListNum(listData []map[string]interface{}, keyName string, order string) (result []map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("SortListNum: %v", r)
		}
	}()
	if order == "ask" { // 字符串排序
		sort.Slice(listData, func(i, j int) bool { return listData[i][keyName].(int) < listData[j][keyName].(int) })
	} else if order == "desc" {
		sort.Slice(listData, func(i, j int) bool { return listData[i][keyName].(int) > listData[j][keyName].(int) })
	} else {
		return nil, errors.New("order填写ask或desc")
	}
	return listData, nil
}

// SortListInt64 int64排序
func SortListInt64(listData []map[string]interface{}, keyName string, order string) (result []map[string]interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("SortListNum: %v", r)
		}
	}()
	if order == "ask" { // 字符串排序
		sort.Slice(listData, func(i, j int) bool { return listData[i][keyName].(int64) < listData[j][keyName].(int64) })
	} else if order == "desc" {
		sort.Slice(listData, func(i, j int) bool { return listData[i][keyName].(int64) > listData[j][keyName].(int64) })
	} else {
		return nil, errors.New("order填写ask或desc")
	}
	return listData, nil
}
