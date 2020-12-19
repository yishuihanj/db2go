package findSql

//find table sql
var findTableSql = `SELECT 
	A.relname AS NAME
FROM
	pg_class A
WHERE
	A.relnamespace = ( SELECT oid FROM pg_namespace WHERE nspname = 'public' ) --用户表一般存储在public模式下

	AND A.relkind = 'r'
ORDER BY
	A.relname`
