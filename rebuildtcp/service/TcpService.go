package service

import (
	"log"
	"mygo/rebuildtcp/dao"
	"mygo/rebuildtcp/model"
	"mygo/rebuildtcp/param"
	"mygo/rebuildtcp/tool"
)

type TcpService struct {
}

func (sv *TcpService) Save(reqtcpparam param.RequestTcp) int64 {
	tcpinfo := model.Tcpinfo{}
	tcpinfo.Uuid = reqtcpparam.Uuid
	tcpinfo.Ack_num = reqtcpparam.Ack_num
	tcpinfo.Body = reqtcpparam.Body
	tcpinfo.Ip_dst = reqtcpparam.Ip_dst
	tcpinfo.Seq_num = reqtcpparam.Seq_num
	tcpinfo.Th_dport = reqtcpparam.Th_dport
	tcpinfo.Th_sport = reqtcpparam.Th_sport
	tcpinfo.Timestamp = reqtcpparam.Timestamp
	tcpinfo.Ip_src = reqtcpparam.Ip_src

	svdao := dao.TcpDao{tool.DbEngine}
	result, id := svdao.InsertParam(tcpinfo)
	if result > 0 {
		id
	}
	return 0
}
