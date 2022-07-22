package model

type Sys_userpasswd struct {
	Id        int64  `xorm:"pk autoincr notnull" json:"id" `
	User      string `xorm:"varchar(255)" json:"user" `
	Passwd    string `xorm:"varchar(255)" json:"passwd" `
	Sys_ip_id int64  `xorm:"varchar(255)" json:"sys_ip_id" `
}
