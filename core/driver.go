// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

/*
代码生成器所支持的数据库驱动项
*/
package core

// database driver type
type DatabaseDriver int

const (
	Unknown DatabaseDriver = iota
	MySQL
	PostgreSQL
	SQLite
	MongoDB
	MariaDB
	Oracle
	SQLServer
)

// driver name
func (d DatabaseDriver) String() string {
	return []string{"MySQL", "PostgreSQL", "SQLite", "MongoDB", "MariaDB", "Oracle", "SQLServer"}[d-1]
}

// driver short name
func (d DatabaseDriver) Name() string {
	return []string{"mysql", "pgsql", "sqlite", "mongo", "maria", "oracle", "mssql"}[d-1]
}

// get driver value by driver short name
func DriverValue(name string) DatabaseDriver {
	if d, ok := map[string]DatabaseDriver{
		"mysql":  MySQL,
		"pgsql":  PostgreSQL,
		"sqlite": SQLite,
		"mongo":  MongoDB,
		"maria":  MariaDB,
		"oracle": Oracle,
		"mssql":  SQLServer,
	}[name]; ok {
		return d
	}
	return Unknown
}

// Determine whether the driver is supported
func Supported(name string) bool {
	ds := []DatabaseDriver{
		MySQL,
		PostgreSQL,
		// SQLite,
		// MongoDB,
		// MariaDB,
		// Oracle,
		// SQLServer,
	}
	for _, v := range ds {
		if DriverValue(name) == v {
			return true
		}
	}
	return false
}
