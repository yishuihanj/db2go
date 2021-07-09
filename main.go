package main

import (
	"fmt"

	"github.com/yishuihanj/db2go/cmd"
	"github.com/yishuihanj/db2go/dbtogo"
	"github.com/yishuihanj/db2go/findSql"
	"github.com/yishuihanj/db2go/generator"
	"github.com/yishuihanj/db2go/interface_sql"
)

//go:generate go build
//  ./db2go pgsql  -Host=localhost -port=5432 -User=postgres -pwd=123456 -dbname=deeplink -gorm=true -package=hello
func main() {
	driver, err := cmd.InitCommand()
	if err == nil && driver != generator.Invalid {
		fmt.Println("========", driver.String())
		dosometing("pgsql")
	}
}

func dosometing(driver string) {
	model, err := interface_sql.SelectDriver(driver)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = model.Init(cmd.User, cmd.Pwd, cmd.Host, cmd.Port, cmd.DbName)
	if err != nil {
		fmt.Println("错误，连接数据库错误：", err.Error())
		return
	}

	defer model.GetDB().Close()

	cmd.Tables, err = findSql.FindTables(model)
	if err != nil {
		fmt.Println("错误! 查看数据库表失败：", err.Error())
		return
	}
	if len(cmd.Tables) == 0 {
		fmt.Println("警告：当前数据库中数据库表的数量为0，程序退出...")
		return
	}

	fmt.Println("警告：没有设置table，将要导出数据库所有的表...")
	for _, tName := range cmd.Tables {
		cmd.Columns, err = findSql.FindColumns(model, tName)
		if err != nil {
			fmt.Printf("错误! 查找数据库表 '%s'  包含的列失败：%v", "tableName", err.Error())
			return
		}
		generator.CreateFile(tName, dbtogo.ColumnsToStruct(tName, cmd.Columns), cmd.Out)
	}
}
