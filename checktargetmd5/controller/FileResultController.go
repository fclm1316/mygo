package controller

import (
	"allinone/param"
	"allinone/service"
	"allinone/tool"

	"github.com/gin-gonic/gin"
)

type FileResultController struct {
}

func (FileResult *FileResultController) Router(engine *gin.Engine) {
	engine.POST("/fileresult", FileResult.AddFileResult)
	engine.GET("/fileresult", FileResult.GetFileResult)

}

func (FileResult *FileResultController) AddFileResult(context *gin.Context) {
	var ts = service.FileResult{}
	var AddFileResultParam param.FileResultParam
	if err := context.ShouldBindJSON(&AddFileResultParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.DiffBatch(&AddFileResultParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}

func (FileResult *FileResultController) GetFileResult(context *gin.Context) {
	var respose = tool.SuccessListStruct{}
	var ts = service.FileResult{}
	var GetFileResultGetParam param.FileResultGetParam
	GetFileResultGetParam.Sys_batch_id_a = context.DefaultQuery("sys_batch_id_a", "none")
	GetFileResultGetParam.Sys_batch_id_b = context.DefaultQuery("sys_batch_id_b", "none")

	GetFileResultGetParam.Result_uuid = context.DefaultQuery("uuid", "none")
	GetFileResultGetParam.PS = tool.NewPageOffset()
	GetFileResultGetParam.PS.Convert(context.DefaultQuery("page", "none"), context.DefaultQuery("offset", "none"))

	result, err := ts.GetFileResult(&GetFileResultGetParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	respose.SuccessList(context, result)
}
