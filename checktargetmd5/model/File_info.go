package model

import "time"

type File_info struct {
	Id           int64     `xorm:"pk autoincr notnull" json:"id" `
	Sys_batch_id int64     `xorm:"int(11)" json:"sys_batch_id" `
	File_list_id int64     `xorm:"int(11)" json:"file_list_id" `
	Filesize     int       `xorm:"int(11)" json:"filesize" `
	Filemodtime  int       `xorm:"int(11)" json:"filemodtime" `
	Status       int       `xorm:int(1) json:"status"`
	Statusmsg    string    `xorm varcgar(255) json:"statusmsg"`
	Filename     string    `xorm:"varchar(255)" json:"filename" `
	Filemd5      string    `xorm:"varchar(255)" json:"filemd5" `
	Ipaddr       string    `xorm:"varchar(255)" json:"ipaddr"`
	User         string    `xorm:"varchar(255)"json:"name"`
	CreatAt      time.Time `xorm:"created" json:"creatat" `
}
