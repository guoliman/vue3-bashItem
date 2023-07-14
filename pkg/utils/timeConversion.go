package utils

import (
	"errors"
	"fmt"
	"time"
)

// UtcToStc 时间类型转换 标准时区utc 转 东八区cst
func UtcToStc(utcTime time.Time) (result time.Time) {
	uctUnix := utcTime.Unix() // utc时间戳
	// 时间戳转字符串
	bb := time.Unix(uctUnix, 0)
	var cc = bb.Format("2006-01-02 15:04:05")
	// 字符串转东八区stc
	local, _ := time.LoadLocation("Asia/Shanghai")
	stcTime, _ := time.ParseInLocation("2006-01-02 15:04:05", cc, local)
	return stcTime
}

// ElmentTimeToCst elment-ui时间选择器类型 2023-03-21T10:54:49.000Z 转time类型Cst
func ElmentTimeToCst(elmentTime string) (result time.Time, err error) {
	// string转Utc
	t, err := time.Parse(time.RFC3339Nano, elmentTime)
	if err != nil {
		return time.Time{}, errors.New(fmt.Sprintf("时间类型转换异常：%v", err.Error()))
	}
	//Utc转CST
	//logger.FileLogger.Info(fmt.Sprintf("=====1111==t:%v 时区%v", t, time.Now()))
	loc, _ := time.LoadLocation("Asia/Shanghai")
	//logger.FileLogger.Info(fmt.Sprintf("=====222=="))
	startTiemCST := t.In(loc)
	return startTiemCST, nil
}
