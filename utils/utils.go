package utils

import (
	_ "github.com/lib/pq"
)

//数组中是否包含该字符
func In(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}


