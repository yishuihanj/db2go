package findSql

//find table sql
var findTableSql = `SELECT 
	A.relname AS NAME,
	COALESCE(b.description,'') AS COMMENT 
FROM
	pg_class
	A LEFT OUTER JOIN pg_description b ON b.objsubid = 0
	AND A.oid = b.objoid
WHERE
	A.relnamespace = ( SELECT oid FROM pg_namespace WHERE nspname = 'public' ) --用户表一般存储在public模式下

	AND A.relkind = 'r'
ORDER BY
	A.relname`
