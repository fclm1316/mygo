package param

import (
	"allinone/tool"
)

type SysEnvGetParam struct {
	Id      string           `json:"id"`
	EnvName string           `json:"envname"`
	PS      *tool.PageOffset //结构体继承
}

type SysEnvParam struct {
	Id      int64  `json:"id" binding:"required"`
	EnvName string `json:"envname" binding:"required"`
}
