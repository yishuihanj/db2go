package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"pgtogo/utils"
)

var (
	db  *sql.DB
	err error
	host,
	userName,
	pwd,
	dbName,
	tableName,
	outDir string
	port    int
	tables  []*Table
	columns []*Column
	gorm    bool
)

func init() {
	flag.StringVar(&host, "host", "localhost", "数据库ip，默认为localhost")
	flag.IntVar(&port, "port", 5432, "数据库端口，默认为5432")
	flag.StringVar(&userName, "user", "postgres", "数据库用户名，默认为postgres")
	flag.StringVar(&pwd, "pwd", "postgres", "数据库密码，默认为postgres")
	flag.StringVar(&dbName, "dbname", "", "数据库名称，必填，否则会报错")
	flag.StringVar(&tableName, "table", "", "需要导出的数据库表名称，如果不设置的话会将该数据库所有的表导出")
	flag.BoolVar(&gorm, "gorm", false, "是否添加 gorm tag，true添加，false不添加，默认不添加")
	flag.StringVar(&outDir, "outdir", "./pg_output", ".go 文件输出路径，不设置的话会输出到当前程序所在路径")
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("recover from a fatal error : %v \n ", e)
		}
	}()
	flag.Parse()
	if dbName == "" {
		fmt.Println("错误! dbname 没有被设置，程序退出!!!")
		return
	}
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	db, err = utils.InitPgSql(userName, pwd, host, port, dbName)
	if err != nil {
		fmt.Println("错误! 连接数据库失败：", err.Error())
		return
	}
	defer db.Close()

	tables, err = FindTables()
	if err != nil {
		fmt.Println("错误! 查看数据库表失败：", err.Error())
		return
	}
	if tableName == "" {
		fmt.Println("警告：没有设置table，将要导出数据库所有的表...")
		for _, table := range tables {
			columns, err = FindColumns(table.Name)
			if err != nil {
				fmt.Printf("错误! 查找数据库表 '%s'  包含的列失败：%v", tableName, err.Error())
				return
			}
			utils.CreateFile(table.Name, ColumnsToStruct(table.Name), outDir)
		}
	} else {
		columns, err = FindColumns(tableName)
		if err != nil {
			fmt.Printf("错误! 查找数据库表 '%s'  包含的列失败：%v", tableName, err.Error())
			return
		}
		utils.CreateFile(tableName, ColumnsToStruct(tableName), outDir)
	}
}
