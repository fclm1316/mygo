package tool

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	S int = 0
	F int = 1
)

// 普通返回
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": S,
		"smg":  "成功",
		"data": v,
	})
}

func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": F,
		"smg":  "失败",
		"data": v,
	})
}

type SuccessListStruct struct {
	Code int                      `"json:code"`
	Msg  string                   `"json:msg"`
	Data []map[string]interface{} `"json:data"`
}

func (sls *SuccessListStruct) SuccessList(ctx *gin.Context, v []byte) {
	sls.Code = F
	sls.Msg = "成功"
	sls.Data = ByteToJosn(v)
	ctx.JSON(http.StatusOK, sls)
}
