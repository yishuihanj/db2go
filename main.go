package main

import (
	"fmt"

	"github.com/yishuihanj/db2go/cmd"
	"github.com/yishuihanj/db2go/dbtogo"
	"github.com/yishuihanj/db2go/findSql"
	"github.com/yishuihanj/db2go/generator"
	"github.com/yishuihanj/db2go/interface_sql"
)

var Tables []string
var Columns []*findSql.Column

//go:generate go build
//  ./db2go -driver=pgsql  -host=localhost -port=5432 -user=postgres -auth=123456 -dbname=deeplink -gorm=true -package=hello
func main() {
	cfg := cmd.Launch()
	dosometing(cfg)
}

func dosometing(cfg *cmd.DriverConfig) {
	model, err := interface_sql.SelectDriver(cfg.Driver)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = model.Init(cfg.User, cfg.Auth, cfg.Host, cfg.Port, cfg.DbName)
	if err != nil {
		fmt.Println("错误，连接数据库错误：", err.Error())
		return
	}

	defer model.GetDB().Close()

	Tables, err = findSql.FindTables(model)
	if err != nil {
		fmt.Println("错误! 查看数据库表失败：", err.Error())
		return
	}
	if len(Tables) == 0 {
		fmt.Println("警告：当前数据库中数据库表的数量为0，程序退出...")
		return
	}

	fmt.Println("警告：没有设置table，将要导出数据库所有的表...")
	for _, tName := range Tables {
		Columns, err = findSql.FindColumns(model, tName)
		if err != nil {
			fmt.Printf("错误! 查找数据库表 '%s'  包含的列失败：%v", "tableName", err.Error())
			return
		}
		generator.CreateFile(tName, dbtogo.ColumnsToStruct(tName, Columns), cfg.Out)
	}
}
