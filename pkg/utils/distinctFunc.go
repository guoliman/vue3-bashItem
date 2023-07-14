package utils

import "sort"

// RemoveListInt64 列表去重 []int64{}类型
//b1 :=[]int64{1,2,3,4,5,5,4,3,2,1,11,12,13,14,15}
//c1 :=RemoveListInt(b1)
func RemoveListInt64(userIDs []int64) []int64 {
	// 如果有0或1个元素，则返回切片本身。
	if len(userIDs) < 2 {
		return userIDs
	}
	//  使切片升序排序
	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	uniqPointer := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}
	return userIDs[:uniqPointer+1]
}

// RemoveListInt 列表去重 []int{}类型
//b1 :=[]int{1,2,3,4,5,5,4,3,2,1,11,12,13,14,15}
//c1 :=RemoveListInt(b1)
//func RemoveListInt(userIDs []int64) []int64 { // int64去重
func RemoveListInt(userIDs []int) []int {
	// 如果有0或1个元素，则返回切片本身。
	if len(userIDs) < 2 {
		return userIDs
	}
	//  使切片升序排序
	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	uniqPointer := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}
	return userIDs[:uniqPointer+1]
}

// RemoveListStr 列表去重 string类型
//b1 :=[]string{"aa","bb","cc","dd","aa","bb","cc","dd","ee"}
//c1 :=RemoveListStr(b1)
func RemoveListStr(stringList []string) []string {
	result := make([]string, 0, len(stringList))
	temp := map[string]struct{}{}
	for _, item := range stringList {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
