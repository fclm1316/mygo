package controller

import (
	"allinone/param"
	"allinone/service"
	"allinone/tool"

	// "fmt"

	"github.com/gin-gonic/gin"
)

type SysUserPasswdController struct {
}

func (SysUserPasswd *SysUserPasswdController) Router(engine *gin.Engine) {
	engine.GET("/sysuserpasswd", SysUserPasswd.GetSUP)
	engine.POST("/sysuserpasswd", SysUserPasswd.AddSUP)
	engine.PUT("/sysuserpasswd", SysUserPasswd.UpdateSUP)
	engine.DELETE("/sysuserpasswd", SysUserPasswd.DeleteSUP)

}

func (SysUserPasswd *SysUserPasswdController) GetSUP(context *gin.Context) {
	var respose = tool.SuccessListStruct{}
	var ts = service.SysUserPasswd{}
	var AddSysUserPasswdParam param.SysUserPasswdGetParam
	AddSysUserPasswdParam.Id = context.DefaultQuery("id", "none")
	AddSysUserPasswdParam.User = context.DefaultQuery("user", "none")
	// AddSysUserPasswdParam.Passwd = context.DefaultQuery("passwd", "none")
	AddSysUserPasswdParam.Sys_ip_id = context.DefaultQuery("sys_ip_id", "none")
	AddSysUserPasswdParam.PS = tool.NewPageOffset()
	AddSysUserPasswdParam.PS.Convert(context.DefaultQuery("page", "none"), context.DefaultQuery("offset", "none"))

	result, err := ts.GetSUP(&AddSysUserPasswdParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	respose.SuccessList(context, result)

}

func (SysUserPasswd *SysUserPasswdController) AddSUP(context *gin.Context) {
	var ts = service.SysUserPasswd{1}
	var AddSysUserPasswdParam param.SysUserPasswdParam
	if err := context.ShouldBindJSON(&AddSysUserPasswdParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.InsertUpdateDelete(&AddSysUserPasswdParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}

func (SysUserPasswd *SysUserPasswdController) UpdateSUP(context *gin.Context) {
	var ts = service.SysUserPasswd{2}
	var UpdateSysUserPasswdParam param.SysUserPasswdParam
	if err := context.ShouldBindJSON(&UpdateSysUserPasswdParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.InsertUpdateDelete(&UpdateSysUserPasswdParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}

func (SysUserPasswd *SysUserPasswdController) DeleteSUP(context *gin.Context) {
	var ts = service.SysUserPasswd{3}
	var DeleteSysUserPasswdParam param.SysUserPasswdParam
	if err := context.ShouldBindJSON(&DeleteSysUserPasswdParam); err != nil {
		tool.Failed(context, err.Error())

		return
	}

	result, err := ts.InsertUpdateDelete(&DeleteSysUserPasswdParam)
	if err != nil {
		tool.Failed(context, err.Error())
		return
	}

	tool.Success(context, result)
}
