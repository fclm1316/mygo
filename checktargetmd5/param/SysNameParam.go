package param

import (
	"allinone/tool"
)

type SysNameGetParam struct {
	Id   string           `json:"id"`
	Name string           `json:"name"`
	PS   *tool.PageOffset //结构体继承
}

type SysNameParam struct {
	Id   int64  `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
