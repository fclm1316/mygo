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

type SysName struct {
	IUD uint8
}

func (sn *SysName) GetName(GetSysNameParam *param.SysNameGetParam) ([]byte, error) {

	var result []byte
	var err error
	var id int
	stdao := dao.SysNameDao{0, tool.DbEngine}

	if GetSysNameParam.Id == "none" && GetSysNameParam.Name == "none" {
		if result, err = stdao.SelectAll(GetSysNameParam.PS.Page, GetSysNameParam.PS.Offset); err != nil {
			return result, err
		}

	} else {
		var temp = make(map[string]interface{})
		if GetSysNameParam.Id != "none" {
			id, err = strconv.Atoi(GetSysNameParam.Id)
			if err != nil {
				return result, err
			}
			temp["id"] = id
		}

		if GetSysNameParam.Name != "none" {
			temp["name"] = GetSysNameParam.Name
		}

		if result, err = stdao.SelectByQuery(GetSysNameParam.PS.Page, GetSysNameParam.PS.Offset, temp); err != nil {
			return result, err
		}
	}

	return result, nil

}

func (sn *SysName) InsertUpdateDelete(FromSysNameParam *param.SysNameParam) (int64, error) {
	var resultid int64
	var err error
	FromSysName := &model.Sys_name{
		Id:   FromSysNameParam.Id,
		Name: FromSysNameParam.Name,
	}
	stdao := dao.SysNameDao{0, tool.DbEngine}
	if sn.IUD == 1 {
		stdao.IUD = 1
	}
	if sn.IUD == 2 {
		stdao.IUD = 2
	}
	if sn.IUD == 3 {
		stdao.IUD = 3
		// FromSysName.Name = ""
	}
	if resultid, err = stdao.InsertUpdateDelete(FromSysName); err != nil {
		return resultid, err
	}
	return resultid, nil

}
