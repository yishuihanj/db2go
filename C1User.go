package main
type C1User struct {
	Id	int32	`gorm:"column:id;not null;primaryKey;type:integer"`
	Name	string	`gorm:"column:name;not null;type:character varying(255);commnet:'姓名'"`
	Age	int8	`gorm:"column:age;not null;type:smallint;commnet:'年龄'"`
	Label	pq.Int64Array	`gorm:"column:label;not null;type:integer[];commnet:'标签'"`
	CreateTimestamp	time.Time	`gorm:"column:create_timestamp;not null;type:time(6) without time zone"`
}