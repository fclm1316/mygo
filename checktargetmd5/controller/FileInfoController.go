package controller

import (
	"allinone/param"
	"allinone/service"
	"allinone/tool"

	"github.com/gin-gonic/gin"
)

type FileInfoController struct {
}

func (FileInfo *FileInfoController) Router(engine *gin.Engine) {
	engine.POST("/fileinfo", FileInfo.AddFileInfo)
	engine.GET("/fileinfo", FileInfo.GetFileInfo)

}

func (sysname *FileInfoController) GetFileInfo(context *gin.Context) {

	var respose = tool.SuccessListStruct{}
	var ts = service.FileInfo{}
	var GetSysFileInfoParam param.FileInfoGetParam
	GetSysFileInfoParam.Ipaddr = context.DefaultQuery("ipaddr", "none")
	GetSysFileInfoParam.Sys_batch_id = context.DefaultQuery("sys_batch_id", "none")
	GetSysFileInfoParam.File_list_id = context.DefaultQuery("file_list_id", "none")
	GetSysFileInfoParam.Status = context.DefaultQuery("status", "none")
	GetSysFileInfoParam.PS = tool.NewPageOffset()
	GetSysFileInfoParam.PS.Convert(context.DefaultQuery("page", "none"), context.DefaultQuery("offset", "none"))

	result, err := ts.GetFileInfo(&GetSysFileInfoParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	respose.SuccessList(context, result)

}

func (sysname *FileInfoController) AddFileInfo(context *gin.Context) {
	ts := service.FileInfo{}
	var FileInfoParam param.FileInfoParam
	if err := context.ShouldBindJSON(&FileInfoParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.GetFileMd5()
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}
