package findSql

import (
	"github.com/yishuihanj/db2go/interface_sql"
)

//查找该数据库的数据库表
func FindTables(model interface_sql.SqlInterface) ([]string, error) {
	db := model.GetDB()
	rows, err := db.Query(model.FindTableString())
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tablesName := make([]string, 0, 0)
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tablesName = append(tablesName, tableName)
	}
	return tablesName, nil
}
