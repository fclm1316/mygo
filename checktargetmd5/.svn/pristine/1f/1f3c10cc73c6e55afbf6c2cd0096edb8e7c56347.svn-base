package model

type File_list struct {
	Id                int64  `xorm:"pk autoincr notnull" json:"id" `
	Retype            int64  `xorm:"int(1)" json:"retype" `
	Status            int64  `xorm:"int(1)" json:"status" `
	Sys_ip_id         int64  `xorm:"varchar(255)" json:"sys_ip_id" `
	Sys_userpasswd_id int64  `xorm:"datetime" json:"sys_userpasswd_id" `
	Filepath          string `xorm:"varchar(255)" json:"filepath" `
	Filename          string `xorm:"varchar(255)" json:"filename" `
}
