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

type SysEnv struct {
}

func (se *SysEnv) GetEnv(GetSysEnvParam *param.SysEnvGetParam) ([]byte, error) {

	var result []byte
	var err error
	var id int
	stdao := dao.SysEnvDao{tool.DbEngine}

	if GetSysEnvParam.Id == "none" && GetSysEnvParam.EnvName == "none" {
		if result, err = stdao.SelectAll(GetSysEnvParam.PS.Page, GetSysEnvParam.PS.Offset); err != nil {
			return result, err
		}

	} else {
		var temp = make(map[string]interface{})
		if GetSysEnvParam.Id != "none" {
			id, err = strconv.Atoi(GetSysEnvParam.Id)
			if err != nil {
				return result, err
			}
			temp["id"] = id
		}

		if GetSysEnvParam.EnvName != "none" {
			temp["envname"] = GetSysEnvParam.EnvName
		}

		if result, err = stdao.SelectByQuery(GetSysEnvParam.PS.Page, GetSysEnvParam.PS.Offset, temp); err != nil {
			return result, err
		}
	}

	return result, nil

}
func (se *SysEnv) AddEnv(AddSysEnvParam *param.SysEnvParam) (int64, error) {
	var resultid int64
	var err error
	AddSysEnv := &model.Sys_env{
		Envname: AddSysEnvParam.EnvName,
	}

	stdao := dao.SysEnvDao{tool.DbEngine}
	if resultid, err = stdao.AddSysEnv(AddSysEnv); err != nil {
		return resultid, err
	}

	return resultid, nil
}
func (se *SysEnv) UpdateEnv(UpdateSysEnvParam *param.SysEnvParam) (int64, error) {
	var resultid int64
	var err error
	UpdateSysEnv := &model.Sys_env{
		Id:      UpdateSysEnvParam.Id,
		Envname: UpdateSysEnvParam.EnvName,
	}

	stdao := dao.SysEnvDao{tool.DbEngine}
	if resultid, err = stdao.UpdateSysEnv(UpdateSysEnv); err != nil {
		return resultid, err
	}

	return resultid, nil
}
func (se *SysEnv) DeleteEnv(DeleteSysEnvParam *param.SysEnvParam) (int64, error) {
	var resultid int64
	var err error
	DeleteSysEnv := &model.Sys_env{
		Id: DeleteSysEnvParam.Id,
		// Envname: DeleteSysEnvParam.EnvName,
	}

	stdao := dao.SysEnvDao{tool.DbEngine}
	if resultid, err = stdao.DeleteSysEnv(DeleteSysEnv); err != nil {
		return resultid, err
	}

	return resultid, nil
}
