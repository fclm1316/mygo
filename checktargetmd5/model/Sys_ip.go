package model

type Sys_ip struct {
	Id          int64  `xorm:"pk autoincr notnull int(11)" json:"id" `
	Ipaddr      string `xorm:"varchar(255)" json:"ipaddr" `
	Sys_env_id  int64  `xorm:"int(11)" json:"sys_env_type"`
	Sys_name_id int64  `xorm:"int(11)" json:"sys_name_id"`
}
