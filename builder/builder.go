// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package builder

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/yishuihanj/db2go/builder/findSql"
	"github.com/yishuihanj/db2go/core"
	"github.com/yishuihanj/db2go/dialect/pgsql"
	"github.com/yishuihanj/db2go/utils"
)

var Tables []string
var Columns []*findSql.Column

// 生成器工厂
func schemaFactory(driver core.DatabaseDriver, cfg *core.Config) core.Schema {
	switch driver {
	// case core.MySQL:
	// 	return mysql.New(cfg)
	// case core.PostgreSQL:
	// 	return pgsql.New(cfg)
	// case core.SQLite:
	// 	return new(mysql.Generator)
	// case core.MariaDB:
	// 	return new(mysql.Generator)
	// case core.MongoDB:
	// 	return new(mysql.Generator)
	// case core.Oracle:
	// 	return new(mysql.Generator)
	default:
		return pgsql.New(cfg)
	}
}

// 执行生成命令
func Generate(driver core.DatabaseDriver, cfg *core.Config) error {
	schema := schemaFactory(driver, cfg)
	tables, err := schema.Tables()
	if err != nil {
		return err
	}
	if len(tables) == 0 {
		return errors.New("the count of tables in the database is 0")
	}

	// // 解析模板
	// t := template.New("text")                             // 定义模板对象
	// t = t.Funcs(template.FuncMap{"ifImports": IfImports}) // 控制自定义元素
	// t = t.Funcs(template.FuncMap{"ifComment": IfComment}) // 控制自定义元素
	// t, err = t.Parse(_template)                           //
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(cfg)

	if err := os.MkdirAll(cfg.Out, os.ModePerm); err != nil {
		return err
	}
	// fmt.Printf(" 📁 输出目录：%s\n", cfg.Out)
	// 单文件输出
	if cfg.Pile {
		gofile := utils.PathTrim(fmt.Sprintf("%s/%s.go", cfg.Out, cfg.DbName))
		f, err := os.Create(gofile)
		if err != nil {
			return err
		}
		data := Schema2Template(driver, cfg, tables...)
		if err := Execute(f, data); err != nil {
			return err
		}
		fmt.Printf(" 📖 生成文件：%s\n", gofile)
		return nil
	}

	// 多文件输出
	for _, table := range tables {
		gofile := utils.PathTrim(fmt.Sprintf("%s/%s.go", cfg.Out, table.Name))
		f, err := os.Create(gofile)
		if err != nil {
			return err
		}

		data := Schema2Template(driver, cfg, table)
		if err := Execute(f, data); err != nil {
			return err
		}
		fmt.Printf(" 📖 生成文件：%s\n", gofile)
	}
	return nil
}

func Schema2Template(driver core.DatabaseDriver, cfg *core.Config, tables ...core.Table) *GenTemplate {
	t := &GenTemplate{
		Generator: cfg.AppName,
		Version:   cfg.Version,
		Source:    fmt.Sprintf("%s://%s:%d/%s", driver, cfg.Host, cfg.Port, cfg.DbName),
		Date:      time.Now().Format("2006-01-02"),
		Package:   cfg.Package,
	}
	if len(tables) == 1 {
		t.Source = fmt.Sprintf("%s://%s:%d/%s/%s", driver, cfg.Host, cfg.Port, cfg.DbName, tables[0].Name)
	}

	for _, table := range tables {
		obj := Struct{
			Name:    utils.Pascal(table.Name),
			Comment: table.Comment,
		}
		for _, column := range table.Columns {
			obj.Fields = append(obj.Fields, Field{
				Name:    utils.Pascal(column.Name),
				Type:    column.Type, // fixme 转换
				Tag:     column.Default,
				Comment: column.Comment,
			})
		}
		t.Imports = table.SpecialPack
		t.Structs = append(t.Structs, obj)
	}
	return t
}
