package pgsql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresSql struct {
	DB     *sql.DB
	DBName string
}

//连接pgsql数据库
func (this *PostgresSql) Init(userName, pwd, host string, port int, dbName string) error {
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", userName, pwd, host, port, dbName)
	_db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	this.DB = _db
	this.DBName = dbName
	return nil
}
func (this *PostgresSql) FindTableString() string {
	return findTableSql
}
func (*PostgresSql) FindColumnsString() string {
	return findColumnSql
}

func (this *PostgresSql) DBNameString() string {
	return this.DBName
}
func (this *PostgresSql) GetDB() *sql.DB {
	return this.DB
}
