package model
type FibCache struct {
	Num	int32	`gorm:"column:num;not null;primaryKey;type:integer"`
	Fib	int32	`gorm:"column:fib;not null;type:integer"`
}