package param

import (
	"allinone/tool"
)

type SysUserPasswdGetParam struct {
	Id   string `json:"id"`
	User string `json:"user"`
	// Passwd    string           `json:"passwd"`
	Sys_ip_id string           `json:"sys_ip_id"`
	PS        *tool.PageOffset //结构体继承
}

type SysUserPasswdParam struct {
	Id        int64  `json:"id" binding:"required"`
	User      string `json:"user" binding:"required"`
	Passwd    string `json:"passwd" binding:"required"`
	Sys_ip_id int64  `json:"sys_ip_id" binding:"required"`
}
