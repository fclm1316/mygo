package po

import (
	"allinone/model"
)

type FileListList struct {
	Id       int64  `json:"id"`
	Retype   int64  `json:"retype"`
	Status   int64  `json:"status"`
	Filepath string `json:"filepath"`
	Filename string `json:"filename"`
	Ipaddr   string `json:"ipaddr"`
	User     string `json:"name"`
	Passwd   string `json:"passwd"`
}

type FileListAllin struct {
	model.File_list      `xorm:"extends"`
	model.Sys_ip         `xorm:"extends"`
	model.Sys_userpasswd `xorm:"extends"`
}
