// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package builder

import (
	"fmt"
	"os"

	"github.com/yishuihanj/db2go/core"
	"github.com/yishuihanj/db2go/dialect/gorm"
	"github.com/yishuihanj/db2go/utils"
)

// 版本信息
const AppName = "db2go"
const VERSION = "v1.0.0"

var (
	_driver  string
	_host    string
	_port    int
	_user    string
	_auth    string
	_dbName  string
	_package string
	_out     string
	_clean   bool
	_version bool
	_help    bool
	_pile    bool
)

// 初始化Flag
func initFlag() {
	utils.Usage = Usage
	utils.StringVar(&_driver, "driver", "", "DB驱动")
	utils.StringVar(&_host, "host", "localhost", "主机名")
	utils.IntVar(&_port, "port", 0, "端口")
	utils.StringVar(&_user, "user", "", "用户名")
	utils.StringVar(&_auth, "auth", "", "密码")
	utils.StringVar(&_dbName, "dbname", "", "数据库名称")
	utils.BoolVar(&gorm.Gorm, "gorm", false, "是否添加gorm标签")
	utils.StringVar(&_out, "out", "./_gen", "输出路径")
	utils.StringVar(&_package, "package", "model", "go文件包名")
	utils.BoolVar(&_clean, "clean", false, "是否清理输出目录")
	utils.BoolVar(&_version, "version", false, "查看版本")
	utils.BoolVar(&_pile, "pile", false, "单文件输出")
	utils.BoolVar(&_help, "help", false, "查看帮助")
}

// 加载命令行参数
func Load() (core.DatabaseDriver, *core.Config) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("\033[%dmAn error occurred: %v\033[0m\n\n", utils.FgRed, err)
			Usage()
			os.Exit(1)
		}
	}()

	initFlag()
	if err := utils.Parse(); err != nil {
		panic(err)
	}

	if _help || len(os.Args) == 1 {
		Usage()
		os.Exit(0)
	}

	if _version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if err := check(); err != nil {
		panic(err)
	}
	if _clean {
		os.RemoveAll(_out)
		os.Exit(0)
	}
	return core.DriverValue(_driver), &core.Config{
		AppName: AppName,
		Version: VERSION,
		Host:    _host,
		Port:    _port,
		User:    _user,
		Auth:    _auth,
		DbName:  _dbName,
		Package: _package,
		Out:     _out,
		Pile:    _pile,
	}
}

// 检查命令行参数
func check() error {
	if len(_driver) == 0 {
		return fmt.Errorf("driver is needed")
	}

	if !core.Supported(_driver) {
		return fmt.Errorf("the driver named %v is not supported", _driver)
	}

	if len(_dbName) == 0 {
		return fmt.Errorf("dbname is needed")
	}
	return nil
}

// 打印帮助信息
func Usage() {
	fmt.Println("\033[1;34m Welcome to db2go\033[0m")
	fmt.Printf("\033[1;34m  ______  ______     ______  _____ \n  |     \\ |_____]   |  ____ |     |\n  |_____/ |_____] 2 |______||_____| (%v)\033[0m\n", VERSION)
	fmt.Printf(`
Usage:
    db2go <command> dbname=<dbName> [option]...

Command:
    mysql	从mysql数据库生成表实体
    pgsql	从postgres数据库生成表实体
    help	查看帮助

Option:
    -host	主机名
    -host	主机名
    -host	主机名
    -host	主机名
    -host	主机名
    -host	主机名

Default param:
    mysql: -host=localhost -port=3306 -user=root -auth=""
    pgsql: -host=localhost -port=5432 -user=postgres -auth=postgres

Example:
    db2go -driver=pgsql -dbname=testdb
    db2go -driver=pgsql -host=localhost -port=5432 -user=postgres -auth=postgres -dbname=testdb -gorm -package=entity

更多详情，请参考 https://github.com/hollson/db2go

`)
}
