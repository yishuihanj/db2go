// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package generator

type Driver int

const (
	Invalid  Driver = 0 // 无效的
	Mysql           = 1
	Postgres        = 2
)

func (d Driver) String() string {
	return []string{"Invalid", "Mysql", "Postgres"}[d]
}