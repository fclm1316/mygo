package controller

import (
	"allinone/param"
	"allinone/service"
	"allinone/tool"

	// "fmt"

	"github.com/gin-gonic/gin"
)

type SysInfoController struct {
}

func (sysinfo *SysInfoController) Router(engine *gin.Engine) {
	engine.GET("/sysinfo", sysinfo.GetInfo)
	engine.POST("/sysinfo", sysinfo.AddInfo)
	engine.PUT("/sysinfo", sysinfo.UpdateInfo)
	engine.DELETE("/sysinfo", sysinfo.DeleteInfo)

}

func (sysinfo *SysInfoController) GetInfo(context *gin.Context) {
	var respose = tool.SuccessListStruct{}
	var ts = service.SysInfo{}
	var GetSysInfoParam param.SysInfoGetParam
	GetSysInfoParam.Id = context.DefaultQuery("id", "none")
	GetSysInfoParam.EnvId = context.DefaultQuery("envid", "none")
	GetSysInfoParam.Ip = context.DefaultQuery("ipaddr", "none")
	GetSysInfoParam.NameId = context.DefaultQuery("nameid", "none")
	GetSysInfoParam.PS = tool.NewPageOffset()
	GetSysInfoParam.PS.Convert(context.DefaultQuery("page", "none"), context.DefaultQuery("offset", "none"))

	result, err := ts.GetInfo(&GetSysInfoParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	respose.SuccessList(context, result)

}
func (sysinfo *SysInfoController) AddInfo(context *gin.Context) {
	var ts = service.SysInfo{}
	var AddSysInfoParam param.SysInfoParam
	if err := context.ShouldBindJSON(&AddSysInfoParam); err != nil {
		tool.Failed(context, err.Error())
		// fmt.Println(AddSysInfoParam)
		return
	}

	result, err := ts.AddInfo(&AddSysInfoParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}

func (sysinfo *SysInfoController) UpdateInfo(context *gin.Context) {
	var ts = service.SysInfo{}
	var UpdateSysinfoParam param.SysInfoParam
	if err := context.ShouldBindJSON(&UpdateSysinfoParam); err != nil {
		tool.Failed(context, err.Error())
		return
	}
	result, err := ts.UpdateInfo(&UpdateSysinfoParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}
	tool.Success(context, result)

}
func (sysinfo *SysInfoController) DeleteInfo(context *gin.Context) {
	var ts = service.SysInfo{}
	var DeleteSysinfoParam param.SysInfoParam
	if err := context.ShouldBindJSON(&DeleteSysinfoParam); err != nil {
		tool.Failed(context, err.Error())
		return
	}
	result, err := ts.DeleteInfo(&DeleteSysinfoParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}
	tool.Success(context, result)

}
