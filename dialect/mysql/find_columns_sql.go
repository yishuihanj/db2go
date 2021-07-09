package mysql

var findColumnSql = `SELECT
	COLUMN_NAME AS column_name, 
	#IS_NULLABLE as not_null 
	(CASE IS_NULLABLE  WHEN 'YES' THEN  'true' WHEN 'NO' THEN 'false'  END) as not_null, 
	(CASE ISNULL(COLUMN_DEFAULT) WHEN true THEN '' ELSE COLUMN_DEFAULT END) as default_value,
	(CASE COLUMN_KEY WHEN 'PRI' THEN 'true' ELSE 'false' END ) as is_primary_key,
	COLUMN_COMMENT as comment,
	COLUMN_TYPE as column_type 
FROM
	information_schema.COLUMNS 
WHERE
	table_name = ?`
