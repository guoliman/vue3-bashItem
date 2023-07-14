package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

// StringToSliceInt string转[]int
// aa := "15,25,64,65,129,218,219,220,221,222,254,255,12"
// listData :=stringToSliceInt(aa)
func StringToSliceInt(strData string) []int {
	listData := strings.Split(strData, ",")
	result := make([]int, 0)
	for _, i := range listData {
		var int_two, _ = strconv.ParseInt(i, 10, 32)
		int_free := int(int_two)
		result = append(result, int_free)
	}
	return result
}

// IntSliceToStrSlice []int 转 []string
func IntSliceToStrSlice(intSlice []int) []string {
	var strSlice []string
	for i := 0; i < len(intSlice); i++ {
		strSlice = append(strSlice, fmt.Sprintf("%v", intSlice[i]))
	}
	return strSlice
}
func GetRandomUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UrlSplit(url string) string {
	sList := strings.Split(url, "//")
	if len(sList) > 1 {
		return sList[1]
	}
	return sList[0]
}
