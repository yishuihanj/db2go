// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package generator

// database driver type
type Driver int

const (
	Unknown Driver = iota
	MySQL
	PostgreSQL
	SQLite
	MongoDB
	MariaDB
	SQLServer
)

// driver name
func (d Driver) String() string {
	return []string{"MySQL", "PostgreSQL", "SQLite", "MongoDB", "MariaDB", "SQLServer"}[d-1]
}

// driver short name
func (d Driver) Name() string {
	return []string{"mysql", "pgsql", "sqlite", "mongo", "maria", "mssql"}[d-1]
}

// get driver value by driver short name
func DriverValue(name string) Driver {
	if d, ok := map[string]Driver{
		"mysql":  MySQL,
		"pgsql":  PostgreSQL,
		"sqlite": SQLite,
		"mongo":  MongoDB,
		"maria":  MariaDB,
		"mssql":  SQLServer,
	}[name]; ok {
		return d
	}
	return Unknown
}

// Determine whether the driver is supported
func Supported(name string) bool {
	ds := []Driver{
		MySQL,
		PostgreSQL,
		// SQLite,
		// MongoDB,
		// MariaDB,
		// SQLServer,
	}
	for _, v := range ds {
		if DriverValue(name) == v {
			return true
		}
	}
	return false
}
