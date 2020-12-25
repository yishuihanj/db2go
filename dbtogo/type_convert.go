package dbtogo

import (
	"github.com/yishuihanj/db2go/utils"
	"strings"
)

//类型转换，没有的类型在这里面添加

func typeConvert(s string) string {
	if strings.Contains(s, "[]") {
		if strings.Contains(s, "char") || strings.Contains(s, "text") {
			return "pq.StringArray"
		}
		if strings.Contains(s, "integer") {
			return "pq.Int64Array"
		}
	}
	if strings.Contains(s, "char") || utils.In(s, []string{"text", "longtext"}) {
		return "string"
	}
	if utils.In(s, []string{"bigserial", "serial", "big serial", "int"}) {
		return "int"
	}
	if utils.In(s, []string{"bigint"}) {
		return "int64"
	}
	if utils.In(s, []string{"integer"}) {
		return "int32"
	}
	if utils.In(s, []string{"smallint"}) {
		return "int16"
	}
	if utils.In(s, []string{"numeric", "decimal", "real"}) {
		return "decimal.Decimal"
	}
	if utils.In(s, []string{"bytea"}) {
		return "[]byte"
	}
	if strings.Contains(s, "time") || utils.In(s, []string{"date"}) {
		return "time.Time"
	}

	return "interface{}"
}
