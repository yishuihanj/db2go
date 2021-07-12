// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"github.com/yishuihanj/db2go/core"
)

type Generator struct {
	Host   string
	Port   int
	User   string
	Auth   string
	DbName string
}

func New(config *core.Config) *Generator {
	gen := &Generator{
		Host:   "",
		Port:   0,
		User:   "",
		Auth:   "",
		DbName: "",
	}
	return gen
}


func (g *Generator) Gen() error {

	panic("implement me")
}
