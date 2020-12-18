package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
	host,
	userName,
	pwd,
	dbName,
	tableName string //前缀
	port    int
	tables  []*Table
	columns []*Column
)

//load PgSql db
func initPgSql() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", userName, pwd, host, port, dbName)
	_db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	return _db, nil
}

func main() {
	host = "127.0.0.1"
	userName = "postgres"
	pwd = "postgres"
	dbName = "db_han"
	port = 5432

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("recover from a fatal error : %v \n ", e)
		}
	}()

	db, err = initPgSql()
	if err != nil {
		fmt.Println("ERROR! Connect PgSql err:", err.Error())
		return
	}
	defer db.Close()
	fmt.Println("Connect PgSql Success...")

	tables, err = FindTables()
	if err != nil {
		fmt.Println("ERROR! Lookup your tables  err:", err.Error())
		return
	}
	tableName = "c1_user"

	columns, err = FindColumns(tableName)
	if err != nil {
		fmt.Printf("ERROR! Lookup table '%s'  contains column error:%v", tableName, err.Error())
		return
	}

	fmt.Printf("%v", ColumnsToStruct())
}
