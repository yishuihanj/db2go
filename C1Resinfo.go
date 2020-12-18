package main

import "github.com/lib/pq"

type C1Resinfo struct {
	Id            int64          `gorm:"column:id;not null;primaryKey;type:bigint"`
	ResName       string         `gorm:"column:res_name;not null;type:character varying(128);commnet:'资源名称'"`
	FileName      string         `gorm:"column:file_name;not null;type:character varying(20);commnet:'文件名'"`
	ThumbFileName string         `gorm:"column:thumb_file_name;not null;type:character varying(20)"`
	UserGuid      int64          `gorm:"column:user_guid;not null;type:bigint"`
	CreateName    string         `gorm:"column:create_name;not null;type:character varying(64)"`
	FileType      int8           `gorm:"column:file_type;not null;type:smallint"`
	ResTimestamp  int64          `gorm:"column:res_timestamp;not null;type:bigint"`
	FileCrc       int64          `gorm:"column:file_crc;not null;type:bigint"`
	ThumbFileCrc  int64          `gorm:"column:thumb_file_crc;not null;type:bigint"`
	MachineType   int32          `gorm:"column:machine_type;not null;type:integer"`
	ResContext    string         `gorm:"column:res_context;not null;type:character varying(1024)"`
	EditTimestamp int64          `gorm:"column:edit_timestamp;not null;type:bigint"`
	ResTags       int32          `gorm:"column:res_tags;not null;type:integer"`
	ResPrice      int32          `gorm:"column:res_price;not null;type:integer"`
	RecId         int64          `gorm:"column:rec_id;not null;type:bigint"`
	FileNameArray pq.StringArray `gorm:"column:file_name_array;type:character varying(20)[]"`
	St            int8           `gorm:"column:st;not null;type:smallint"`
	PriceType     int32          `gorm:"column:price_type;not null;type:integer"`
	HaveOther     int8           `gorm:"column:have_other;not null;type:smallint"`
	ResState      int8           `gorm:"column:res_state;not null;type:smallint"`
	CheckState    int8           `gorm:"column:check_state;not null;type:smallint"`
	ScriptName    string         `gorm:"column:script_name;not null;type:character varying(128)"`
	Src           int8           `gorm:"column:src;not null;type:smallint"`
}
