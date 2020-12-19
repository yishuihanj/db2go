package findSql

import (
	"database/sql"
)

type Table struct {
	Name    string
	Comment sql.NullString
}

func FindTables(db *sql.DB) ([]*Table, error) {
	rows, err := db.Query(findTableSql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tables := make([]*Table, 0, 0)
	for rows.Next() {
		var table Table
		err = rows.Scan(&table.Name, &table.Comment)
		if err != nil {
			return nil, err
		}
		tables = append(tables, &table)
	}
	return tables, nil
}
