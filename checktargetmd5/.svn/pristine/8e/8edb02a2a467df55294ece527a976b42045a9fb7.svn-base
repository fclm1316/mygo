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

type SysBatch struct {
}

func (sb *SysBatch) GetSysBatch(GetSysBatchParam *param.SysBatchGetParam) ([]byte, error) {

	var result []byte
	var err error
	var id int
	stdao := dao.SysBatchDao{tool.DbEngine}

	if GetSysBatchParam.Id == "none" {
		if result, err = stdao.SelectAll(GetSysBatchParam.PS.Page, GetSysBatchParam.PS.Offset); err != nil {
			return result, err
		}

	} else {
		var temp = make(map[string]interface{})

		id, err = strconv.Atoi(GetSysBatchParam.Id)
		if err != nil {
			return result, err
		}
		temp["id"] = id

		if result, err = stdao.SelectByQuery(GetSysBatchParam.PS.Page, GetSysBatchParam.PS.Offset, temp); err != nil {
			return result, err
		}
	}

	return result, nil

}
func (sb *SysBatch) DeleteSysBatch(DeleteSysBatchParam *param.SysBatchParam) (int64, error) {
	var resultid int64
	var err error

	DeleteSysBatchFileIfo := &model.Sys_batch{
		Id: DeleteSysBatchParam.Id,
	}

	stdao := dao.SysBatchDao{tool.DbEngine}
	if resultid, err = stdao.UionDeleteSysBatchFileIfo(DeleteSysBatchFileIfo); err != nil {
		return resultid, err
	}

	return resultid, nil
}
