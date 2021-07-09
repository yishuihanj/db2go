package dbtogo

import (
	"fmt"
	"strings"

	"github.com/yishuihanj/db2go/dialect/gorm"
	"github.com/yishuihanj/db2go/findSql"

)

var Pkg string

//将字段名转换成结构体字段   不包含tag
func ColumnsToStruct(_tableName string, columns []*findSql.Column) string {
	columnString := ""
	for _, column := range columns {
		singleString := fmt.Sprintf("\t%s\t%s", splitUnderline(column.ColumnName), typeConvert(column.ColumnType))

		//
		singleString = singleString + gorm.AddGormTag(column) + "\n"
		columnString += singleString

	}
	return fmt.Sprintf("package %s\ntype %s struct {\n%s}", Pkg,splitUnderline(_tableName), columnString)
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