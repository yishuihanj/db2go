package utils

import (
	// _ "github.com/lib/pq"
	"strings"
)

const BLANK = ""

// 数组中是否包含该字符
func In(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

// 转换为帕斯卡命名
//  如: userName => UserName
//     user_name => UserName
func Pascal(title string) string {
	arr := strings.FieldsFunc(title, func(c rune) bool { return c == '_' })
	RangeStringsFunc(arr, func(s string) string { return strings.Title(s) })
	return strings.Join(arr, BLANK)
}

// 遍历处理集合成员
func RangeStringsFunc(arr []string, f func(string) string) {
	for k, v := range arr {
		arr[k] = f(v)
	}
}

func PathTrim(path string) string {
	return strings.ReplaceAll(path, "//", "/")
}
