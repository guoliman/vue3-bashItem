package utils

import (
	"errors"
	"fmt"
	"os"
)

// StringInSlice 判断字符串是否在列表内
func StringInSlice(slice []string, item string) bool {
	for _, value := range slice {
		if value == item {
			return true
		}
	}
	return false
}

// JudgeInt64InList 判断int64是否在列表内 utils.JudgeInt64InList(1,[]int64{1,2,3,4,5,6})
func JudgeInt64InList(intData int64, data []int64) bool {
	for _, y := range data {
		if intData == y {
			return true
		}
	}
	return false
}

// UintInSlice 判断uint是否在列表内
func UintInSlice(slice []uint, item uint) int {
	for index, value := range slice {
		if value == item {
			return index
		}
	}
	return -1
}

// JudgeStringInList 判断int是否在列表内
//utils.JudgeStringInList(1,[]string{1,2,3,4,5,6})
func JudgeStringInList(intData string, data []string) (string, error) {
	for _, y := range data {
		if intData == y {
			return y, nil
		}
	}
	return "", errors.New("notExist")
}

// JudgeIntInList 判断int是否在列表内
// utils.JudgeIntInList(1,[]int{1,2,3,4,5,6})
func JudgeIntInList(intData int, data []int) (int, error) {
	for _, y := range data {
		if intData == y {
			return y, nil
		}
	}
	return 0, errors.New("notExist")
}

// JudgeInSliceMap 判断int是否在切片map列表内
//oldRoleList :=[]map[string]interface{}
//oldRoleList = append(oldRoleList, map[string]interface{}{"roleId":3,"id":1})
//oldRoleList = append(oldRoleList, map[string]interface{}{"roleId":4,"id":2})
//utils.JudgeInList(1, "roleId",oldRoleList)
func JudgeInSliceMap(intData int, query string, data []map[string]interface{}) (int, error) {
	for _, y := range data {
		if intData == y[query] {
			return y["id"].(int), nil
		}
	}
	return 0, errors.New("notExist")
}

// CheckDirExist 检查目录是否存在 不存在就创建
func CheckDirExist(path string) error {
	var exist = false
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = true
	}
	if exist {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return errors.New(fmt.Sprintf("%v目录创建失败%v", path, err))
		}
	}
	return nil
}

// CheckFileExist 检查文件是否存在
func CheckFileExist(filepath string) error {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("%v 文件不存在 %v", filepath, err))
	}
	return nil
}

// RemoveFile 判断文件是否存在,存在就删除
func RemoveFile(filepath string) error {
	if fileExist := CheckFileExist(filepath); fileExist == nil {
		err := os.Remove(filepath)
		if err != nil {
			return err
		}
	}
	return nil
}
