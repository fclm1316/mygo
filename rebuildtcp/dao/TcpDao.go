package dao

import (
	"log"
	"mygo/rebuildtcp/model"
	"mygo/rebuildtcp/tool"
)

type TcpDao struct {
	*tool.Orm
}

func (td *TcpDao) InsertParam(tcpinfo model.Tcpinfo) (int64, int64) {
	result, err := td.InsertOne(&tcpinfo)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	return result, tcpinfo.Id
}
