package utils

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

// FileRead 读取excel文件内容
/*
fileData, err := FileRead("/Users/guoliman/Desktop/1yong/6.cmdb/项目部署记录表.xlsx", "Sheet1")
if err != nil {fmt.Println("============")}
fmt.Println("长度", len(fileData)
*/
func FileRead(filePath, sheet string) ([][]string, error) {
	var result [][]string

	f, err := excelize.OpenFile(filePath) //Book1.xlsx是打开的文件
	if err != nil {
		return result, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows(sheet)
	if err != nil {
		return result, err
	}

	countRow := 0
	for _, row := range rows {
		countRow += 1
		//fmt.Println(fmt.Sprintf("===========循环行 当前行%v 行数据:%v", countRow, row))
		countCol := 0
		var colList []string
		for _, colData := range row {
			colList = append(colList, colData)
			countCol += 1
			//fmt.Println(fmt.Sprintf("当前第%v列 值:%v", countCol, colData))
		}
		result = append(result, colList)
	}
	return result, nil
}
