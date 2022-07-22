package controller

import (
	"allinone/param"
	"allinone/service"
	"allinone/tool"

	// "fmt"

	"github.com/gin-gonic/gin"
)

type FileListController struct {
}

func (FileList *FileListController) Router(engine *gin.Engine) {
	engine.GET("/filelist", FileList.GetFile)
	engine.POST("/filelist", FileList.AddFile)
	engine.PUT("/filelist", FileList.UpdateFile)
	engine.DELETE("/filelist", FileList.DeleteFile)

}

func (FileList *FileListController) GetFile(context *gin.Context) {
	var respose = tool.SuccessListStruct{}
	var ts = service.FileList{}
	var GetFileListParam param.FileListGetParam
	GetFileListParam.SysIpId = context.DefaultQuery("sysipid", "none")
	GetFileListParam.Status = context.DefaultQuery("status", "none")
	GetFileListParam.PS = tool.NewPageOffset()
	GetFileListParam.PS.Convert(context.DefaultQuery("page", "none"), context.DefaultQuery("offset", "none"))

	result, err := ts.GetFile(&GetFileListParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	respose.SuccessList(context, result)

}

func (sysname *FileListController) AddFile(context *gin.Context) {
	var ts = service.FileList{1}
	var AddFileListParam param.FileListParam
	if err := context.ShouldBindJSON(&AddFileListParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.InsertUpdateDelete(&AddFileListParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}

func (sysname *FileListController) UpdateFile(context *gin.Context) {
	var ts = service.FileList{2}
	var UpdateFileListParam param.FileListParam
	if err := context.ShouldBindJSON(&UpdateFileListParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.InsertUpdateDelete(&UpdateFileListParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}

func (sysname *FileListController) DeleteFile(context *gin.Context) {
	var ts = service.FileList{3}
	var DeleteFileListParam param.FileListParam
	if err := context.ShouldBindJSON(&DeleteFileListParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.InsertUpdateDelete(&DeleteFileListParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}
