// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/yishuihanj/db2go/findSql"
	"github.com/yishuihanj/db2go/generator"
)

var (
	DB *sql.DB
	Host,
	User,
	Pwd,
	DbName,
	Out string
	Port    int
	Tables  []string
	Help    bool
	Columns []*findSql.Column // 模式
)

func check() error {
	if DbName == "" {
		return errors.New("")
	}
	return nil
}

func InitDriver() (gen generator.Generator,err error) {
	defer func() {
		if err != nil {
			usage()
		}
	}()

	return nil, nil
}

//go:generate go build
func InitCommand() (driver generator.Driver, err error) {
	defer func() {
		if err != nil {
			usage()
		}
	}()

	if len(os.Args) == 1 {
		return generator.Invalid, fmt.Errorf("args is missing")
	}

	if os.Args[1] == "mysql" {
		fs := mysqlFlag()
		args := os.Args[2:]
		if err := fs.Parse(args); err != nil || len(args) == 0 || Help {
			fs.Usage()
			return generator.Invalid, err
		}
		if err := check(); err != nil {
			fs.Usage()
			return generator.Invalid, err
		}
		// next
		return generator.Mysql, nil
	}

	if os.Args[1] == "pgsql" {
		fs := pgFlag()
		args := os.Args[2:]
		if err := fs.Parse(args); err != nil || len(args) == 0 || Help {
			fs.Usage()
			return generator.Invalid, err
		}
		if err := check(); err != nil {
			fs.Usage()
			return generator.Invalid, err
		}
		// next
		return generator.Postgres, nil
	}
	return generator.Invalid, fmt.Errorf("unknown args")
}

func usage() {
	fmt.Fprintf(os.Stderr, `「Golang」一个数据库表实体生成工具,支持mysql和postgres
Usage:
    db2go <command> dbname=<dbName> [option]...

e.g. ./db2go pgsql  -host=localhost -port=5432 -user=postgres -pwd=123456 -dbname=deeplink -gorm=true -package=hello

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

更多详情，请参考 https://github.com/hollson/db2go

`)
}
