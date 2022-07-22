package param

import (
	"allinone/tool"
)

type FileResultGetParam struct {
	Sys_batch_id_a string           `json:"sys_batch_id_a"`
	Sys_batch_id_b string           `json:"sys_batch_id_b"`
	Result_uuid    string           `json:"uuid"`
	PS             *tool.PageOffset //结构体继承
}

type FileResultParam struct {
	Sys_batch_id_a int64 `json:"sys_batch_id_a" binding:"required"`
	Sys_batch_id_b int64 `json:"sys_batch_id_b" binding:"required"`
}
