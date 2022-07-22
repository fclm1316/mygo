package param

import "allinone/tool"

type FileInfoParam struct {
	Id int64 `json:"id"`
}

type FileInfoGetParam struct {
	Ipaddr       string           `json:"ipaddr"`
	Sys_batch_id string           `json:"sys_batch_id"`
	File_list_id string           `json:"file_list_id"`
	Status       string           `json:"status"`
	PS           *tool.PageOffset //结构体继承
}
