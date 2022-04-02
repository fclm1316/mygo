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

func (tcp *TcpController) Info(context *gin.Context) {
	var ReqTcpParam param.RequestTcp
	err := tool.Decode(context.Request.Body, &ReqTcpParam)
	if err != nil {
		tool.Success(context, "解析JSON失败")
		return
	}
	ts := service.TcpService{}

	result := ts.Save(ReqTcpParam)
	if result {
		tool.Success(context, "插入数据库成功")
	} else {
		tool.Failed(context, "插入数据库失败")
	}

}
