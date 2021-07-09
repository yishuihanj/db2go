package pgsql

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
			AND ad.adnum = A.attnum) THEN
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
