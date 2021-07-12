// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"fmt"
)

// 数据库代码生成器配置项
type Config struct {
	AppName string // 应用名称
	Version string // 应用版本
	Host    string // 主机
	Port    int    // 端口，具化默认值
	User    string // 用户，具化默认值
	Auth    string // 密码，具化默认值
	DbName  string // 数据库
	Package string // 包名
	Out     string // 输出路径
	Pile    bool   // 单文件输出
}

func (c *Config) String()string  {
	return fmt.Sprintf("%+v",*c)
}