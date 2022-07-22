package controller

import (
	"allinone/param"
	"allinone/service"
	"allinone/tool"

	// "fmt"

	"github.com/gin-gonic/gin"
)

type SysBatchController struct {
}

func (sysbatch *SysBatchController) Router(engine *gin.Engine) {
	engine.GET("/sysbatch", sysbatch.GetSysBatch)
	engine.DELETE("/sysbatch", sysbatch.DeleteSysBatch)

}

func (sysbatch *SysBatchController) GetSysBatch(context *gin.Context) {
	var respose = tool.SuccessListStruct{}
	var ts = service.SysBatch{}
	var GetSysBatchParam param.SysBatchGetParam
	GetSysBatchParam.Id = context.DefaultQuery("id", "none")
	GetSysBatchParam.PS = tool.NewPageOffset()
	GetSysBatchParam.PS.Convert(context.DefaultQuery("page", "none"), context.DefaultQuery("offset", "none"))

	result, err := ts.GetSysBatch(&GetSysBatchParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	respose.SuccessList(context, result)

}

func (sysbatch *SysBatchController) DeleteSysBatch(context *gin.Context) {

	var ts = service.SysBatch{}
	var SysBatchParam param.SysBatchParam
	if err := context.ShouldBindJSON(&SysBatchParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.DeleteSysBatch(&SysBatchParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)

}
