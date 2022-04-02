package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS int = 0
	FALSE   int = 1
)

func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "成功",
		"data": v,
	})
}

func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": FALSE,
		"msg":  "失败",
		"data": v,
	})
}
