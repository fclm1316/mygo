package param

import (
	"allinone/tool"
)

type FileListGetParam struct {
	SysIpId string           `json:"id"`
	Status  string           `json:"status"`
	PS      *tool.PageOffset //结构体继承
}

type FileListParam struct {
	Id                int64  `json:"id" binding:"required"`
	Filepath          string `json:"filepath" binding:"required"`
	Filename          string `json:"filename" binding:"required"`
	Retype            int64  `json:"retype" binding:"required"`
	Sys_ip_id         int64  `json:"sys_ip_id" binding:"required"`
	Sys_userpasswd_id int64  `json:"sys_userpasswd_id" binding:"required"`
	Status            int64  `json:"status" binding:"required"`
}
