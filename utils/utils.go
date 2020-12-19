package utils

import (
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

//如果字符串有下滑线的则将下划线除去 并开头字母大写  例如   v1_user 变为  V1User
func SplitUnderline(s string) string {
	arr := strings.Split(s, "_")
	ret := ""
	for _, v := range arr {
		ret += strings.Title(v)
	}
	return ret
}

//数组中是否包含该字符
func In(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

func CheckFlagParse(port int, user, pwd, dbname, driver string) string {
	if port <= 0 {
		return fmt.Sprintf("错误：port 必填，程序退出...")
	}
	if user == "" {
		return fmt.Sprintf("错误：user 必填，程序退出...")
	}
	if pwd == "" {
		return fmt.Sprintf("错误：pwd 必填，程序退出...")
	}
	if dbname == "" {
		return fmt.Sprintf("错误：dbname 必填，程序退出...")
	}
	if driver == "" {
		return fmt.Sprintf("错误：driver 必填，程序退出...")
	}
	return ""
}
