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

type SysInfo struct {
}

func (si *SysInfo) GetInfo(GetSysInfoParam *param.SysInfoGetParam) ([]byte, error) {

	var result []byte
	var err error
	var id, env int
	stdao := dao.SysInfoDao{tool.DbEngine}

	// page, size, err = tool.PageSize(GetSysInfoParam.Page, GetSysInfoParam.Size)
	// if err != nil {
	// 	return result, err
	// }

	if GetSysInfoParam.EnvId == "none" && GetSysInfoParam.Ip == "none" &&
		GetSysInfoParam.NameId == "none" && GetSysInfoParam.Id == "none" {
		if result, err = stdao.SelectAll(GetSysInfoParam.PS.Page, GetSysInfoParam.PS.Offset); err != nil {
			return result, err
		}

	} else {
		var temp = make(map[string]interface{})
		if GetSysInfoParam.Id != "none" {
			id, err = strconv.Atoi(GetSysInfoParam.Id)
			if err != nil {
				return result, err
			}
			temp["sys_ip.id"] = id
		}
		if GetSysInfoParam.EnvId != "none" {
			env, err = strconv.Atoi(GetSysInfoParam.EnvId)
			if err != nil {
				return result, err
			}
			temp["sys_env.id"] = env
		}
		if GetSysInfoParam.Ip != "none" {
			temp["sys_ip.ipaddr"] = GetSysInfoParam.Ip
		}
		if GetSysInfoParam.NameId != "none" {
			temp["sys_name.id"] = GetSysInfoParam.NameId
		}
		if result, err = stdao.SelectByQuery(GetSysInfoParam.PS.Page, GetSysInfoParam.PS.Offset, temp); err != nil {
			return result, err
		}
	}

	return result, nil

}

func (si *SysInfo) AddInfo(AddInfoParam *param.SysInfoParam) (int64, error) {
	var resultid int64
	var err error
	AddSysIp := &model.Sys_ip{
		Ipaddr:      AddInfoParam.Ip,
		Sys_env_id:  AddInfoParam.EnvId,
		Sys_name_id: AddInfoParam.NameId,
	}

	stdao := dao.SysInfoDao{tool.DbEngine}
	if resultid, err = stdao.AddInfoIP(AddSysIp); err != nil {
		return resultid, err
	}

	return resultid, nil
}

func (si *SysInfo) UpdateInfo(UpdateInfoParam *param.SysInfoParam) (int64, error) {
	var resultid int64
	var err error
	UpdateSysIp := &model.Sys_ip{
		Id:          UpdateInfoParam.Id,
		Ipaddr:      UpdateInfoParam.Ip,
		Sys_env_id:  UpdateInfoParam.EnvId,
		Sys_name_id: UpdateInfoParam.NameId,
	}

	stdao := dao.SysInfoDao{tool.DbEngine}
	if resultid, err = stdao.UpdateInfoIP(UpdateSysIp); err != nil {
		return resultid, err
	}

	return resultid, nil
}

func (si *SysInfo) DeleteInfo(DeleteInfoParam *param.SysInfoParam) (int64, error) {
	var resultid int64
	var err error
	DeleteSysIp := &model.Sys_ip{
		Id: DeleteInfoParam.Id,
		// Ipaddr:      DeleteInfoParam.Ip,
		// Sys_env_id:  DeleteInfoParam.EnvId,
		// Sys_name_id: DeleteInfoParam.NameId,
	}

	stdao := dao.SysInfoDao{tool.DbEngine}
	if resultid, err = stdao.DeleteInfoIP(DeleteSysIp); err != nil {
		return resultid, err
	}

	return resultid, nil
}
