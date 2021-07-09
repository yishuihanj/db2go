// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"flag"

	"github.com/yishuihanj/db2go/dbtogo"
	"github.com/yishuihanj/db2go/dialect/gorm"
)

// gorm模式或raw模式
func pgFlag() *flag.FlagSet {
	fs := flag.NewFlagSet("pgsql", flag.ExitOnError)
	fs.StringVar(&Host, "host", "localhost", "主机名")
	fs.IntVar(&Port, "port", 5432, "端口")
	fs.StringVar(&User, "user", "postgres", "用户名")
	fs.StringVar(&Pwd, "pwd", "postgres", "密码")
	fs.StringVar(&DbName, "dbname", "", "数据库名称")
	fs.BoolVar(&gorm.Gorm, "gorm", false, "是否添加gorm标签")
	fs.StringVar(&Out, "out", "./model", "输出路径")
	fs.StringVar(&dbtogo.Pkg, "package", "model", "go文件包名")
	fs.BoolVar(&Help, "help", false, "帮助文档")
	return fs
}