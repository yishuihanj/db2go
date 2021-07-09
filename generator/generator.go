// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package generator

// 代码生成器
type Generator interface {
	Driver() Driver
	Host() string     // 主机
	Port() int        // 端口
	User() string     // 用户名
	Password() string // 密码
	DbName() string   // 数据库名称
	Check() error     // 参数校验
	Ping() error      // 测试数据库连通性
	Gen() error       // 执行生成
}
