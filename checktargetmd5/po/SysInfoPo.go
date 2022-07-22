package po

import (
	"allinone/model"
)

type SysInfoList struct {
	Id      int64  `json:"id"`
	Ipaddr  string `json:"ipaddr"`
	Name    string `json:"name"`
	Envname string `json:"envname"`
}

type SysInfoAllin struct {
	model.Sys_ip   `xorm:"extends"`
	model.Sys_name `xorm:"extends"`
	model.Sys_env  `xorm:"extends"`
}
