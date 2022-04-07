package controller

import (
	"log"
	"mygo/rebuildtcp/param"
	"mygo/rebuildtcp/service"
	"mygo/rebuildtcp/tool"

	"github.com/gin-gonic/gin"
)

type TcpController struct {
}

func (tcp (TcpController) Router(engine *gin.Engine){
	engine.Post("api/tcp",tcp.Info)
}

func (tcp *TcpController) Info(context *gin.Context) {
	var ReqTcpParam param.RequestTcp
	err := tool.Decode(context.Request.Body, &ReqTcpParam)
	if err != nil {
		tool.Failed(context, "解析JSON失败")
		return
	}
	ts := service.TcpService{}

	result := ts.Save(ReqTcpParam)
	if result {
		tool.Success(context, result)
		return
	}
	tool.Failed(context, "插入数据库失败")

}
