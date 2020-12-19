package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySql struct {
	DB     *sql.DB
	DBName string
}

//连接pgsql数据库
func (this *MySql) Init(userName, pwd, host string, port int, dbName string) error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", userName, pwd, host, port, dbName)
	_db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	this.DB = _db
	this.DBName = dbName
	return nil
}
func (this *MySql) FindTableString() string {
	return fmt.Sprintf("select table_name from information_schema.tables where table_schema='%s'", this.DBName)
}
func (*MySql) FindColumnsString() string {
	return findColumnSql
}

func (this *MySql) DBNameString() string {
	return this.DBName
}

func (this *MySql) GetDB() *sql.DB {
	return this.DB
}
