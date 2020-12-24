package dbtogo

import (

	"fmt"

	"github.com/yishuihanj/db2go/findSql"
	"github.com/yishuihanj/db2go/gormtogo"
	"github.com/yishuihanj/db2go/utils"
)

var Pkg string

//将字段名转换成结构体字段   不包含tag
func ColumnsToStruct(_tableName string, columns []*findSql.Column) string {
	columnString := ""
	for _, column := range columns {
		singleString := fmt.Sprintf("\t%s\t%s", utils.SplitUnderline(column.ColumnName), utils.TypeConvert(column.ColumnType))

		//
		singleString = singleString + gormtogo.AddGormTag(column) + "\n"
		columnString += singleString

	}
	return fmt.Sprintf("package %s\ntype %s struct {\n%s}", Pkg,utils.SplitUnderline(_tableName), columnString)
}
