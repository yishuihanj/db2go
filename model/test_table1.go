package model
type TestTable1 struct {
	Id	int	`gorm:"column:id;not null;primaryKey;type:serial"`
	NickName	string	`gorm:"column:nick_name;type:character varying(20);commnet:'我是列注释'"`
	Addr	string	`gorm:"column:addr;type:character varying"`
	Age	interface{}	`gorm:"column:age;default:1;type:smallserial"`
	Asset	interface{}	`gorm:"column:asset;type:numeric(18,2);commnet:'我是列注释'"`
	Index	interface{}	`gorm:"column:index;default:1.0;type:double precision"`
	Amount	interface{}	`gorm:"column:amount;default:(0)::numeric;type:money"`
	Nonce	int	`gorm:"column:nonce;default:0;type:bigserial"`
	Birth	time.Time	`gorm:"column:birth;default:now();type:date;commnet:'我是列注释'"`
	CreateTime	time.Time	`gorm:"column:create_time;default:CURRENT_TIMESTAMP;type:timestamp without time zone"`
	UpdateTime	time.Time	`gorm:"column:update_time;default:(CURRENT_TIMESTAMP)::date;type:timestamp(6) without time zone;commnet:'我是列注释'"`
	Tels	pq.StringArray	`gorm:"column:tels;default:ARRAY['182'::text, '156'::text];type:character varying(11)[]"`
	Tags	pq.Int64Array	`gorm:"column:tags;default:ARRAY[ARRAY[1, 2, 3], ARRAY[4, 5, 6]];type:integer[]"`
}