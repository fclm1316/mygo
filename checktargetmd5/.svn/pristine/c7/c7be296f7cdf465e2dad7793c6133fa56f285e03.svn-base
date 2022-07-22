package model

import "time"

type File_result struct {
	Id             int64     `xorm:"pk autoincr notnull" json:"id" `
	Sys_batch_id_a int64     `xorm:"int(11)" json:"sys_batch_id_a" `
	Sys_batch_id_b int64     `xorm:"int(11)" json:"sys_batch_id_b" `
	Result_uuid    string    `xorm:"int(11)" json:"result_id" `
	Filename       string    `xorm:"varchar(255)" json:"filename" `
	Ipaddr         string    `xorm:"varchar(255)" json:"ipaddr" `
	Filemd5_a      string    `xorm:"varchar(255)" json:"filemd5_a" `
	Filemd5_b      string    `xorm:"varchar(255)" json:"filemd5_a" `
	CreatAt        time.Time `xorm:"created" json:"creatat" `
}
