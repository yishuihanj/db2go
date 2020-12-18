package main

import (
	"fmt"
	"strings"
)

var findColumnSql = `SELECT A.attname AS COLUMN_NAME,
	A.attnotnull AS not_null,
	COALESCE ( pg_get_expr ( ad.adbin, ad.adrelid ), '' ) AS default_value,
	COALESCE ( ct.contype = 'p', FALSE ) AS is_primary_key,
	COALESCE(b.description,'') AS comment,
CASE
		
		WHEN A.atttypid = ANY ( '{int,int8,int2}' :: regtype [] ) 
		AND EXISTS (
		SELECT
			1 
		FROM
			pg_attrdef ad 
		WHERE
			ad.adrelid = A.attrelid 
			AND ad.adnum = A.attnum 
			AND ad.adsrc = 'nextval(''' || ( pg_get_serial_sequence ( A.attrelid :: regclass :: TEXT, A.attname ) ) :: regclass || '''::regclass)' 
			) THEN
		CASE
				A.atttypid 
				WHEN 'int' :: regtype THEN
				'serial' 
				WHEN 'int8' :: regtype THEN
				'bigserial' 
				WHEN 'int2' :: regtype THEN
				'smallserial' 
			END 
				WHEN A.atttypid = ANY ( '{uuid}' :: regtype [] ) 
				AND COALESCE ( pg_get_expr ( ad.adbin, ad.adrelid ), '' ) != '' THEN
					'autogenuuid' ELSE format_type ( A.atttypid, A.atttypmod ) 
					END AS column_type 
			FROM
				pg_attribute
				A JOIN ONLY pg_class C ON C.oid = A.attrelid
				JOIN ONLY pg_namespace n ON n.oid = C.relnamespace
				LEFT JOIN pg_constraint ct ON ct.conrelid = C.oid 
				AND A.attnum = ANY ( ct.conkey ) 
				AND ct.contype = 'p'
				LEFT JOIN pg_attrdef ad ON ad.adrelid = C.oid 
				AND ad.adnum = A.attnum 
				LEFT JOIN pg_description b ON a.attrelid=b.objoid AND a.attnum = b.objsubid
			WHERE
				A.attisdropped = FALSE 
				AND n.nspname = 'public' 
				AND C.relname = $1 
				AND A.attnum > 0 
		ORDER BY 
	A.attnum`

type Column struct {
	ColumnName   string
	NotNull      string
	DefaultValue string
	IsPrimaryKey string
	Comment      string
	ColumnType   string
}

func FindColumns(_tableName string) ([]*Column, error) {
	rows, err := db.Query(findColumnSql, _tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns := make([]*Column, 0, 0)
	for rows.Next() {
		var column Column
		err = rows.Scan(&column.ColumnName,
			&column.NotNull,
			&column.DefaultValue,
			&column.IsPrimaryKey,
			&column.Comment,
			&column.ColumnType)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil

}

func ColumnsToStruct() string {
	columnString := ""
	for _, column := range columns {
		singleString := fmt.Sprintf("\t%s\t%s", splitUnderline(column.ColumnName), typeConvert(column.ColumnType))

		//
		singleString = singleString + addGormTag(column) + "\n"
		columnString += singleString

	}
	return fmt.Sprintf("package main\ntype %s struct {\n%s}", splitUnderline(tableName), columnString)
}

func addGormTag(column *Column) string {
	if !gorm {
		return ""
	}
	tag := fmt.Sprintf("\t`gorm:\"column:%s", column.ColumnName)
	if column.NotNull == "true" {
		tag += fmt.Sprintf(";not null")
	}
	if column.DefaultValue != "" {
		tag += fmt.Sprintf(";default:%s", column.DefaultValue)
	}
	if column.IsPrimaryKey == "true" {
		tag += fmt.Sprintf(";primaryKey")
	}
	if column.ColumnType != "" {
		tag += fmt.Sprintf(";type:%s", column.ColumnType)
	}
	if column.Comment != "" {
		tag += fmt.Sprintf(";commnet:'%s'", column.Comment)
	}
	end := fmt.Sprintf("\"`")
	return tag + end
}

//Convert the SQL type to go type
func typeConvert(s string) string {

	if strings.Contains(s, "[]") {
		if strings.Contains(s, "char") || strings.Contains(s, "text") {
			return "pq.StringArray"
		}
		if strings.Contains(s, "integer") {
			return "pq.Int64Array"
		}
	}

	if strings.Contains(s, "char") || in(s, []string{"text"}) {
		return "string"
	}
	if in(s, []string{"bigserial", "serial", "big serial"}) {
		return "int"
	}
	if in(s, []string{"bigint"}) {
		return "int64"
	}
	if in(s, []string{"integer"}) {
		return "int32"
	}
	if in(s, []string{"smallint"}) {
		return "int8"
	}
	if in(s, []string{"numeric", "decimal", "real"}) {
		return "decimal.Decimal"
	}
	if in(s, []string{"bytea"}) {
		return "[]byte"
	}
	if strings.Contains(s, "time") || in(s, []string{"date"}) {
		return "time.Time"
	}

	return "interface{}"
}

//split underline to ""
func splitUnderline(s string) string {
	arr := strings.Split(s, "_")
	ret := ""
	for _, v := range arr {
		ret += strings.Title(v)
	}
	return ret
}

//是否包含
func in(s string, arr []string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
