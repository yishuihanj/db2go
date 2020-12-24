package findSql

import (
	"github.com/yishuihanj/db2go/interface_sql"
)

type Column struct {
	ColumnName   string //字段名
	NotNull      string //是否为空
	DefaultValue string //默认值
	IsPrimaryKey string //是否是主键
	Comment      string //注释
	ColumnType   string //字段类型
}

//查找表中的字段
func FindColumns(sqlInterface interface_sql.SqlInterface, _tableName string) ([]*Column, error) {
	db := sqlInterface.GetDB()
	rows, err := db.Query(sqlInterface.FindColumnsString(), _tableName)
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
