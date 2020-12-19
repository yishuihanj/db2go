package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

//连接pgsql数据库
func InitPgSql(userName, pwd, host string, port int, dbName string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", userName, pwd, host, port, dbName)
	_db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return _db, nil
}

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
