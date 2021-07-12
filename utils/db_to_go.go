package utils

import (
	"fmt"
	"strings"

	"github.com/yishuihanj/db2go/builder/findSql"
	"github.com/yishuihanj/db2go/dialect/gorm"
)

var Pkg string

//将字段名转换成结构体字段   不包含tag
func ColumnsToStruct(_tableName string, columns []*findSql.Column) string {
	columnString := ""
	for _, column := range columns {
		singleString := fmt.Sprintf("\t%s\t%s", splitUnderline(column.ColumnName), typeConvert(column.ColumnType))
		singleString = singleString + gorm.AddGormTag(column) + "\n"
		columnString += singleString
	}
	return fmt.Sprintf("package %s\ntype %s struct {\n%s}", Pkg, splitUnderline(_tableName), columnString)
}

//如果字符串有下滑线的则将下划线除去 并开头字母大写  例如   v1_user 变为  V1User
func splitUnderline(s string) string {
	arr := strings.Split(s, "_")
	ret := ""
	for _, v := range arr {
		ret += strings.Title(v)
	}
	return ret
}


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
	if strings.Contains(s, "char") || In(s, []string{"text", "longtext"}) {
		return "string"
	}
	if In(s, []string{"bigserial", "serial", "big serial", "int"}) {
		return "int"
	}
	if In(s, []string{"bigint"}) {
		return "int64"
	}
	if In(s, []string{"integer"}) {
		return "int32"
	}
	if In(s, []string{"smallint"}) {
		return "int16"
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
