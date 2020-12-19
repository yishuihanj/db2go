package findSql

import (
	"database/sql"
)

func FindTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query(findTableSql)
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
