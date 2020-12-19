package utils

import (
	"strings"
)

//类型转换，没有的类型在这里面添加

func TypeConvert(s string) string {
	if strings.Contains(s, "[]") {
		if strings.Contains(s, "char") || strings.Contains(s, "text") {
			return "pq.StringArray"
		}
		if strings.Contains(s, "integer") {
			return "pq.Int64Array"
		}
	}
	if strings.Contains(s, "char") || In(s, []string{"text"}) {
		return "string"
	}
	if In(s, []string{"bigserial", "serial", "big serial"}) {
		return "int"
	}
	if In(s, []string{"bigint"}) {
		return "int64"
	}
	if In(s, []string{"integer"}) {
		return "int32"
	}
	if In(s, []string{"smallint"}) {
		return "int8"
	}
	if In(s, []string{"numeric", "decimal", "real"}) {
		return "decimal.Decimal"
	}
	if In(s, []string{"bytea"}) {
		return "[]byte"
	}
	if strings.Contains(s, "time") || In(s, []string{"date"}) {
		return "time.Time"
	}

	return "interface{}"
}
