package utils

import (
	"fmt"
	"reflect"
	"sort"
)

// SetInt64JJ  交集 []int64 取aa和bb都有的 [2 3 5]
//aa := []int64{1, 2, 3, 5, 7}
//bb := []int64{2, 3, 4, 5, 6}
//a1 := SetInt64JJ(aa, bb)
func SetInt64JJ(a, b []int64) []int64 {
	intersection := make([]int64, 0)

	for _, valA := range a {
		for _, valB := range b {
			if valA == valB {
				intersection = append(intersection, valA)
				break
			}
		}
	}

	return intersection
}

// SetInt64XJ 差集 []int64 取aa中有 bb中没有的值  [1 7]
//aa := []int64{1, 2, 3, 5, 7}
//bb := []int64{2, 3, 4, 5, 6}
//a2 := SetInt64XJ(aa, bb)  // [1 7]
func SetInt64XJ(a, b []int64) []int64 {
	difference := make([]int64, 0)

	for _, valA := range a {
		found := false
		for _, valB := range b {
			if valA == valB {
				found = true
				break
			}
		}
		if !found {
			difference = append(difference, valA)
		}
	}
	return difference
}

// SetInt64BJ 求并集 int64 aa中没有的值在bb中取出 写入aa
//aa := []int64{1, 2, 3, 5, 7}
//bb := []int64{2, 3, 4, 5, 6}
//a3 := SetInt64BJ(aa, bb)      // [2 3 5 7 4 6 1]
func SetInt64BJ(a, b []int64) []int64 {
	unionSet := make(map[int64]bool)
	union := make([]int64, 0)

	for _, val := range a {
		unionSet[val] = true
	}

	for _, val := range b {
		unionSet[val] = true
	}

	for key := range unionSet {
		union = append(union, key)
	}

	return union
}

// UnionData 列表 求并集  slice1中没有的值在slice2中取出 写入slice1
//slice1 := []string{"1", "2", "3", "6", "8"}
//slice2 := []string{"2", "3", "5", "0"}
//un := union(slice1, slice2)
//fmt.Println("slice1与slice2的并集为：", un) //[1 2 3 6 8 5 0]
func UnionData(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

// IntersectData 列表求交集 slice1和slice2共有的值
//slice1 := []string{"1", "2", "3", "6", "8"}
//slice2 := []string{"2", "3", "5", "0"}
//in := intersect(slice1, slice2)
//fmt.Println("slice1与slice2的交集为：", in) // [2 3]
func IntersectData(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

// DifferenceData 列表求差集 slice1的数据存在 slice2不存在的值
//slice1 := []string{"1", "2", "3", "6", "8"}
//slice2 := []string{"2", "3", "5", "0"}
//di := difference(slice1, slice2)
//fmt.Println("slice1与slice2的差集为：", di) //  [1 6 8]
func DifferenceData(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := IntersectData(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

// SetIntersection 取交集  []map[string]interface类型   取共有的值   值内不能有时间类型
//a := []map[string]interface{}{{"aa": "a1", "status": 1, "name": "张三"}, {"c": "cc", "c1": "c11"},}
//b := []map[string]interface{}{{"c1": "cc", "c2": "c11"},{"c": "cc", "c1": "c11"},}
//c := SetIntersection(a, b)
func SetIntersection(a []map[string]interface{}, b []map[string]interface{}) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, amap := range a {
		for _, bmap := range b {
			if reflect.DeepEqual(amap, bmap) {
				result = append(result, amap)
			}
		}
	}
	return result
}

// SetDifference 取差集 []map[string]interface类型 取a中有 b中没有的值  值内不能有时间类型!!!!
//a := []map[string]interface{}{{"aa": "a1", "status": 1, "name": "张三"}, {"c": "cc", "c1": "c11"},}
//b := []map[string]interface{}{{"c1": "cc", "c2": "c11"},{"c": "cc", "c1": "c11"},}
//d := SetDifference(a, b)
func SetDifference(a []map[string]interface{}, b []map[string]interface{}) []map[string]interface{} {
	result := []map[string]interface{}{}
	sort.Slice(a, func(i, j int) bool {
		return fmt.Sprintf("%!v(MISSING)", a[i]) < fmt.Sprintf("%!v(MISSING)", a[j])
	})
	sort.Slice(b, func(i, j int) bool {
		return fmt.Sprintf("%!v(MISSING)", b[i]) < fmt.Sprintf("%!v(MISSING)", b[j])
	})
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if reflect.DeepEqual(a[i], b[j]) {
			i++
			j++
		} else if fmt.Sprintf("%!v(MISSING)", a[i]) < fmt.Sprintf("%!v(MISSING)", b[j]) {
			result = append(result, a[i])
			i++
		} else {
			j++
		}
	}
	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	return result
}
