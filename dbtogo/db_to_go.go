package dbtogo

import (
	"db2go/findSql"
	"db2go/gormtogo"
	"db2go/utils"
	"fmt"
)

//将字段名转换成结构体字段   不包含tag
func ColumnsToStruct(_tableName string, columns []*findSql.Column) string {
	columnString := ""
	for _, column := range columns {
		singleString := fmt.Sprintf("\t%s\t%s", utils.SplitUnderline(column.ColumnName), utils.TypeConvert(column.ColumnType))

		//
		singleString = singleString + gormtogo.AddGormTag(column) + "\n"
		columnString += singleString

	}
	return fmt.Sprintf("package main\ntype %s struct {\n%s}", utils.SplitUnderline(_tableName), columnString)
}
