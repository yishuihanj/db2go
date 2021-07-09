package hel
type Student struct {
	Id	int	`gorm:"column:id;not null;primaryKey;type:serial"`
	Name	string	`gorm:"column:name;type:character varying"`
}