package main

import (
	"database/sql"
	"db2go/dbtogo"
	"db2go/findSql"
	"db2go/gormtogo"
	"db2go/interface_sql"
	"db2go/utils"
	"flag"
	"fmt"
	"os"
)

var (
	db  *sql.DB
	err error
	host,
	userName,
	pwd,
	dbName,
	tableName,
	outDir,
	driver string
	port    int
	tables  []string
	columns []*findSql.Column
)

func init() {
	flag.StringVar(&host, "host", "localhost", "选填，数据库ip，默认为localhost")
	flag.IntVar(&port, "port", 0, "必填，数据库端口")
	flag.StringVar(&userName, "user", "", "必填，数据库用户名")
	flag.StringVar(&pwd, "pwd", "", "必填，数据库密码")
	flag.StringVar(&dbName, "dbname", "", "必填，数据库名称，否则会报错")
	flag.StringVar(&tableName, "table", "", "选填，需要导出的数据库表名称，如果不设置的话会将该数据库所有的表导出")
	flag.BoolVar(&gormtogo.Gorm, "gorm", false, "选填，是否添加 gorm tag，true添加，false不添加，默认不添加")
	flag.StringVar(&outDir, "outdir", "./go_output", "选填，go 文件输出路径，不设置的话会输出到当前程序所在路径")
	flag.StringVar(&driver, "driver", "", "必填，需要连接的数据库，现在只支持mysql、pgsql 例如 -driver=mysql，-driver=pgsql")
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("recover from a fatal error : %v \n ", e)
		}
	}()
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	ret := utils.CheckFlagParse(port, userName, pwd, dbName, driver)
	if ret != "" {
		fmt.Println(ret)
		return
	}

	model, err := interface_sql.SelectDriver(driver)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = model.Init(userName, pwd, host, port, dbName)
	if err != nil {
		fmt.Println("错误，连接数据库错误：", err.Error())
		return
	}

	defer model.GetDB().Close()

	tables, err = findSql.FindTables(model)
	if err != nil {
		fmt.Println("错误! 查看数据库表失败：", err.Error())
		return
	}
	if len(tables) == 0 {
		fmt.Println("警告：当前数据库中数据库表的数量为0，程序退出...")
		return
	}

	if tableName == "" {
		fmt.Println("警告：没有设置table，将要导出数据库所有的表...")
		for _, tName := range tables {
			columns, err = findSql.FindColumns(model, tName)
			if err != nil {
				fmt.Printf("错误! 查找数据库表 '%s'  包含的列失败：%v", tableName, err.Error())
				return
			}
			utils.CreateFile(tName, dbtogo.ColumnsToStruct(tName, columns), outDir)
		}
	} else {
		if !utils.In(tableName, tables) {
			fmt.Println("错误：数据库中没有您想要导出的数据库表，程序退出...")
			return
		}
		columns, err = findSql.FindColumns(model, tableName)
		if err != nil {
			fmt.Printf("错误! 查找数据库表 '%s'  包含的列失败：%v", tableName, err.Error())
			return
		}
		utils.CreateFile(tableName, dbtogo.ColumnsToStruct(tableName, columns), outDir)
	}
}
