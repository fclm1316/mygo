package service

import (
	"allinone/dao"
	"allinone/model"
	"allinone/param"

	// "allinone/po"
	"allinone/tool"
	// "fmt"
	"strconv"
)

type SysUserPasswd struct {
	IUD uint8
}

func (sn *SysUserPasswd) GetSUP(GetSysUserPasswdParam *param.SysUserPasswdGetParam) ([]byte, error) {

	var result []byte
	var err error
	var id int
	stdao := dao.SysUserPasswdDao{0, tool.DbEngine}

	if GetSysUserPasswdParam.Id == "none" && GetSysUserPasswdParam.User == "none" && GetSysUserPasswdParam.Sys_ip_id == "none" {
		if result, err = stdao.SelectAll(GetSysUserPasswdParam.PS.Page, GetSysUserPasswdParam.PS.Offset); err != nil {
			return result, err
		}

	} else {
		var temp = make(map[string]interface{})
		if GetSysUserPasswdParam.Id != "none" {
			id, err = strconv.Atoi(GetSysUserPasswdParam.Id)
			if err != nil {
				return result, err
			}
			temp["id"] = id
		}

		if GetSysUserPasswdParam.User != "none" {
			temp["user"] = GetSysUserPasswdParam.User
		}

		if GetSysUserPasswdParam.Sys_ip_id != "none" {
			id, err = strconv.Atoi(GetSysUserPasswdParam.Sys_ip_id)
			if err != nil {
				return result, err
			}
			temp["sys_ip_id"] = id
		}
		if result, err = stdao.SelectByQuery(GetSysUserPasswdParam.PS.Page, GetSysUserPasswdParam.PS.Offset, temp); err != nil {
			return result, err
		}
	}

	return result, nil

}

func (sn *SysUserPasswd) InsertUpdateDelete(FromSysUserPasswdParam *param.SysUserPasswdParam) (int64, error) {
	var resultid int64
	var err error
	FromSysUserPasswd := &model.Sys_userpasswd{
		Id:        FromSysUserPasswdParam.Id,
		User:      FromSysUserPasswdParam.User,
		Passwd:    FromSysUserPasswdParam.Passwd,
		Sys_ip_id: FromSysUserPasswdParam.Sys_ip_id,
	}
	stdao := dao.SysUserPasswdDao{0, tool.DbEngine}
	if sn.IUD == 1 {
		stdao.IUD = 1
	}
	if sn.IUD == 2 {
		stdao.IUD = 2
	}
	if sn.IUD == 3 {
		stdao.IUD = 3
		// FromSysUserPasswd.User = ""
		// FromSysUserPasswd.Passwd = ""
		// FromSysUserPasswd.Sys_ip_id = 0
	}
	if resultid, err = stdao.InsertUpdateDelete(FromSysUserPasswd); err != nil {
		return resultid, err
	}
	return resultid, nil

}
