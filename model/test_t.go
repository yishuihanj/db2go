package model
type TestT struct {
	Id	int32	`gorm:"column:id;type:integer"`
	Name	string	`gorm:"column:name;type:character varying"`
}