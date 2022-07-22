package controller

import (
	"allinone/param"
	"allinone/service"
	"allinone/tool"

	// "fmt"

	"github.com/gin-gonic/gin"
)

type SysEnvController struct {
}

func (sysenv *SysEnvController) Router(engine *gin.Engine) {
	engine.GET("/sysenv", sysenv.GetEnv)
	engine.POST("/sysenv", sysenv.AddEnv)
	engine.PUT("/sysenv", sysenv.UpdateEnv)
	engine.DELETE("/sysenv", sysenv.DeleteEnv)

}

func (sysenv *SysEnvController) GetEnv(context *gin.Context) {
	var respose = tool.SuccessListStruct{}
	var ts = service.SysEnv{}
	var GetSysEnvParam param.SysEnvGetParam
	GetSysEnvParam.Id = context.DefaultQuery("id", "none")
	GetSysEnvParam.EnvName = context.DefaultQuery("envname", "none")
	GetSysEnvParam.PS = tool.NewPageOffset()
	GetSysEnvParam.PS.Convert(context.DefaultQuery("page", "none"), context.DefaultQuery("offset", "none"))

	result, err := ts.GetEnv(&GetSysEnvParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	respose.SuccessList(context, result)

}

func (sysenv *SysEnvController) AddEnv(context *gin.Context) {
	var ts = service.SysEnv{}
	var AddSysEnvParam param.SysEnvParam
	if err := context.ShouldBindJSON(&AddSysEnvParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.AddEnv(&AddSysEnvParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}

func (sysenv *SysEnvController) UpdateEnv(context *gin.Context) {
	var ts = service.SysEnv{}
	var UpdateSysEnvParam param.SysEnvParam
	if err := context.ShouldBindJSON(&UpdateSysEnvParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.UpdateEnv(&UpdateSysEnvParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}

func (sysenv *SysEnvController) DeleteEnv(context *gin.Context) {
	var ts = service.SysEnv{}
	var DeleteSysEnvParam param.SysEnvParam
	if err := context.ShouldBindJSON(&DeleteSysEnvParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.DeleteEnv(&DeleteSysEnvParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}
