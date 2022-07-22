package controller

import (
	"allinone/param"
	"allinone/service"
	"allinone/tool"

	// "fmt"

	"github.com/gin-gonic/gin"
)

type SysNameController struct {
}

func (sysname *SysNameController) Router(engine *gin.Engine) {
	engine.GET("/sysname", sysname.GetName)
	engine.POST("/sysname", sysname.AddName)
	engine.PUT("/sysname", sysname.UpdateName)
	engine.DELETE("/sysname", sysname.DeleteName)

}

func (sysname *SysNameController) GetName(context *gin.Context) {
	var respose = tool.SuccessListStruct{}
	var ts = service.SysName{}
	var AddSysNameParam param.SysNameGetParam
	AddSysNameParam.Id = context.DefaultQuery("id", "none")
	AddSysNameParam.Name = context.DefaultQuery("name", "none")
	AddSysNameParam.PS = tool.NewPageOffset()
	AddSysNameParam.PS.Convert(context.DefaultQuery("page", "none"), context.DefaultQuery("offset", "none"))

	result, err := ts.GetName(&AddSysNameParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	respose.SuccessList(context, result)

}

func (sysname *SysNameController) AddName(context *gin.Context) {
	var ts = service.SysName{1}
	var AddSysNameParam param.SysNameParam
	if err := context.ShouldBindJSON(&AddSysNameParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.InsertUpdateDelete(&AddSysNameParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}

func (sysname *SysNameController) UpdateName(context *gin.Context) {
	var ts = service.SysName{2}
	var UpdateSysNameParam param.SysNameParam
	if err := context.ShouldBindJSON(&UpdateSysNameParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.InsertUpdateDelete(&UpdateSysNameParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}

func (sysname *SysNameController) DeleteName(context *gin.Context) {
	var ts = service.SysName{3}
	var DeleteSysNameParam param.SysNameParam
	if err := context.ShouldBindJSON(&DeleteSysNameParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.InsertUpdateDelete(&DeleteSysNameParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}
