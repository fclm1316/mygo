package param

import (
	"allinone/tool"
)

type SysInfoGetParam struct {
	Id     string           `json:"id"`
	EnvId  string           `json:"envid"`
	Ip     string           `json:"ip"`
	NameId string           `json:"nameid"`
	PS     *tool.PageOffset //结构体继承
}

type SysInfoParam struct {
	Id     int64  `json:"id" binding:"required"`
	EnvId  int64  `json:"envid" binding:"required"`
	NameId int64  `json:"nameid" binding:"required"`
	Ip     string `json:"ip" binding:"required"`
}
