// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package generator

// 代码生成器
type Generator interface {
	Driver() Driver
	Host() string   // 主机
	Port() int      // 端口
	User() string   // 用户
	Auth() string   // 密码
	Schema() string // 数据库
	Ping() error    // 测试连通性
	Gen() error     // 执行生成
}

type Database struct {}

func (d *Database) Driver() Driver {
	panic("implement me")
}

func (d *Database) Host() string {
	panic("implement me")
}

func (d *Database) Port() int {
	panic("implement me")
}

func (d *Database) User() string {
	panic("implement me")
}

func (d *Database) Auth() string {
	panic("implement me")
}

func (d *Database) Schema() string {
	panic("implement me")
}

func (d *Database) Ping() error {
	panic("implement me")
}

func (d *Database) Gen() error {
	panic("implement me")
}

