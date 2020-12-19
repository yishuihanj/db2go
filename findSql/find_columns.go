package findSql

import (
	"database/sql"
	"fmt"
	"pgtogo/utils"
)

var Gorm bool

type Column struct {
	ColumnName   string //字段名
	NotNull      string //是否为空
	DefaultValue string //默认值
	IsPrimaryKey string //是否是主键
	Comment      string //注释
	ColumnType   string //字段类型
}

//查找表中的字段
func FindColumns(db *sql.DB, _tableName string) ([]*Column, error) {
	rows, err := db.Query(findColumnSql, _tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns := make([]*Column, 0, 0)
	for rows.Next() {
		var column Column
		err = rows.Scan(&column.ColumnName,
			&column.NotNull,
			&column.DefaultValue,
			&column.IsPrimaryKey,
			&column.Comment,
			&column.ColumnType)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil

}

//将字段名转换成结构体字段   不包含tag
func ColumnsToStruct(_tableName string, columns []*Column) string {
	columnString := ""
	for _, column := range columns {
		singleString := fmt.Sprintf("\t%s\t%s", utils.SplitUnderline(column.ColumnName), utils.TypeConvert(column.ColumnType))

		//
		singleString = singleString + addGormTag(column) + "\n"
		columnString += singleString

	}
	return fmt.Sprintf("package main\ntype %s struct {\n%s}", utils.SplitUnderline(_tableName), columnString)
}

//向字段中添加 Gorm tag
func addGormTag(column *Column) string {
	//如果没有开启gorm 则 不需要转换
	if !Gorm {
		return ""
	}
	tag := fmt.Sprintf("\t`gorm:\"column:%s", column.ColumnName)
	if column.NotNull == "true" {
		tag += fmt.Sprintf(";not null")
	}
	if column.DefaultValue != "" && column.IsPrimaryKey != "true" {
		tag += fmt.Sprintf(";default:%s", column.DefaultValue)
	}
	if column.IsPrimaryKey == "true" {
		tag += fmt.Sprintf(";primaryKey")
	}
	if column.ColumnType != "" {
		tag += fmt.Sprintf(";type:%s", column.ColumnType)
	}
	if column.Comment != "" {
		tag += fmt.Sprintf(";commnet:'%s'", column.Comment)
	}
	end := fmt.Sprintf("\"`")
	return tag + end
}
