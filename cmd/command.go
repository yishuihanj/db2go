// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"


	"github.com/yishuihanj/db2go/dbtogo"
	"github.com/yishuihanj/db2go/dialect/gorm"
	"github.com/yishuihanj/db2go/generator"
	"github.com/yishuihanj/db2go/utils/flag"
)

const VERSION = "db2go v1.0.0"

var (
	// fs       *flag.FlagSet
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
)

type DriverConfig struct {
	Driver  string
	Host    string
	Port    int
	User    string
	Auth    string
	DbName  string
	Package string
	Out     string
}

func initFlag() {
	flag.Usage = usage
	flag.StringVar(&_driver, "driver", "", "DB驱动")
	flag.StringVar(&_host, "host", "localhost", "主机名")
	flag.IntVar(&_port, "port", 0, "端口")
	flag.StringVar(&_user, "user", "", "用户名")
	flag.StringVar(&_auth, "auth", "", "密码")
	flag.StringVar(&_dbName, "dbname", "", "数据库名称")
	flag.BoolVar(&gorm.Gorm, "gorm", false, "是否添加gorm标签")
	flag.StringVar(&_out, "out", "./model", "输出路径")
	flag.StringVar(&dbtogo.Pkg, "package", "model", "go文件包名")
	flag.BoolVar(&_clean, "clean", false, "是否清理输出目录")
	flag.BoolVar(&_version, "version", false, "查看版本")
	flag.BoolVar(&_help, "help", false, "查看帮助")
}

//go:generate go build
// start command line
func Launch() *DriverConfig {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("\033[%dmAn error occurred: %v\033[0m\n\n", FgRed, err)
			usage()
			os.Exit(1)
		}
	}()

	if len(os.Args) == 1 {
		panic("Command line params is missing")
	}

	initFlag()
	if err := flag.Parse(); err != nil {
		panic(err)
	}

	if _help {
		usage()
		os.Exit(0)
	}

	if _version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if err := Check(); err != nil {
		panic(err)
	}
	if _clean {
		os.RemoveAll(_out)
		os.Exit(0)
	}
	return &DriverConfig{
		Driver:  _driver,
		Host:    _host,
		Port:    _port,
		User:    _user,
		Auth:    _auth,
		DbName:  _dbName,
		Package: _package,
		Out:     _out,
	}

	//
	// switch {
	// case generator.DriverValue(_driver) == generator.MySQL:
	//
	// }
	//
	// if os.Args[1] == "mysql" {
	// 	fs := mysqlFlag()
	// 	args := os.Args[2:]
	// 	if err := fs.Parse(args); err != nil || len(args) == 0 || _help {
	// 		fs.Usage()
	// 		return generator.Invalid, err
	// 	}
	// 	if err := check(); err != nil {
	// 		fs.Usage()
	// 		return generator.Invalid, err
	// 	}
	// 	// next
	// 	return generator.Mysql, nil
	// }
	// 
	// return generator.Invalid, fmt.Errorf("unknown args")
}

// check the command line params
func Check() error {
	if len(_driver) == 0 {
		return fmt.Errorf("driver is needed")
	}

	if !generator.Supported(_driver) {
		return fmt.Errorf("the driver named %v is not supported", _driver)
	}

	if len(_dbName) == 0 {
		return fmt.Errorf("dbname is needed")
	}
	return nil
}

func usage() {

	fmt.Fprintf(os.Stderr, `「Golang」一个实用的数据库对象代码生成器
 ______  ______     ______  _____ 
 |     \ |_____]   |  ____ |     |
 |_____/ |_____] 2 |______||_____|

Usage:
    db2go <command> dbname=<dbName> [option]...

e.g. ./db2go pgsql  -host=localhost -port=5432 -user=postgres -auth=123456 -dbname=deeplink -gorm=true -package=hello

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

更多详情，请参考 https://github.com/hollson/db2go

`)
}
