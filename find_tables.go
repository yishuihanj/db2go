package main

import (
	"database/sql"
	"pgtogo/findSql"
)

type Table struct {
	Name    string
	Comment sql.NullString
}

func FindTables() ([]*Table, error) {
	rows, err := db.Query(findSql.FindTableSql)
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
