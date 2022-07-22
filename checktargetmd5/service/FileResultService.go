package service

import (
	"allinone/dao"
	"allinone/model"
	"allinone/param"
	"allinone/po"
	"allinone/tool"

	"errors"
	"fmt"
	"strconv"

	"github.com/satori/go.uuid"
)

type FileResult struct {
}

func (fr *FileResult) DiffBatch(AddFileResultParam *param.FileResultParam) (string, error) {

	if tool.GV.DiffMd5State == true {
		err := errors.New("程序正在运行")
		return "", err
	}

	tool.GV.DiffMd5State = true

	result_uuid := uuid.NewV4().String()

	go func(result_uuid string) {
		frdao := dao.FileResultDao{tool.DbEngine}
		stdao := dao.FileInfoDao{tool.DbEngine}
		BatchA, _ := stdao.SelectbySysBatchId(AddFileResultParam.Sys_batch_id_a)
		BatchB, _ := stdao.SelectbySysBatchId(AddFileResultParam.Sys_batch_id_b)

		// fmt.Println(BatchA)
		// fmt.Println(&BatchB)

		Mapaa := fr.SliceToMap(BatchA)
		Mapbb := fr.SliceToMap(BatchB)

		for key, aavalue := range Mapaa {

			// fmt.Println(key, value)
			if _, ok := Mapbb[key]; ok {

				if aavalue.Filemd5 == Mapbb[key].Filemd5 {
					delete(Mapaa, key)
					delete(Mapbb, key)
				} else {
					var Diffmd5 model.File_result
					Diffmd5.Result_uuid = result_uuid
					Diffmd5.Filemd5_a = aavalue.Filemd5
					Diffmd5.Filemd5_b = Mapbb[key].Filemd5
					Diffmd5.Sys_batch_id_a = AddFileResultParam.Sys_batch_id_a
					Diffmd5.Sys_batch_id_b = AddFileResultParam.Sys_batch_id_b
					Diffmd5.Filename = aavalue.Filename
					Diffmd5.Ipaddr = aavalue.Fileip
					_, err := frdao.AddFileResult(&Diffmd5)
					if err != nil {
						fmt.Println(err.Error())
					}
					delete(Mapaa, key)
					delete(Mapbb, key)
				}
			} else {
				var lastMapaa model.File_result
				lastMapaa.Result_uuid = result_uuid
				lastMapaa.Filemd5_a = aavalue.Filemd5
				lastMapaa.Sys_batch_id_a = AddFileResultParam.Sys_batch_id_a
				lastMapaa.Filename = aavalue.Filename
				lastMapaa.Ipaddr = aavalue.Fileip
				_, err := frdao.AddFileResult(&lastMapaa)
				if err != nil {
					fmt.Println(err.Error())
				}
				delete(Mapaa, key)
			}

		}
		if len(Mapbb) != 0 {
			var lastMapbb model.File_result
			for key, bbvalue := range Mapbb {
				lastMapbb.Result_uuid = result_uuid
				lastMapbb.Filemd5_b = bbvalue.Filemd5
				lastMapbb.Sys_batch_id_b = AddFileResultParam.Sys_batch_id_b
				lastMapbb.Filename = bbvalue.Filename
				lastMapbb.Ipaddr = bbvalue.Fileip
				_, err := frdao.AddFileResult(&lastMapbb)
				if err != nil {
					fmt.Println(err.Error())
				}
				delete(Mapbb, key)
			}
		}

		tool.GV.DiffMd5State = false
	}(result_uuid)

	return result_uuid, nil
}

func (fr *FileResult) SliceToMap(slicebatch *[]po.FileInfoByBatch) map[int64]po.FilResultInfo {

	mapData := make(map[int64]po.FilResultInfo)
	var result po.FilResultInfo
	for _, value := range *slicebatch {

		result.Fileip = value.Ipaddr
		result.Filemd5 = value.Filemd5
		result.Filename = value.Filename

		mapData[value.File_list_id] = result

	}

	return mapData

}

func (fr *FileResult) GetFileResult(GetFileResultParam *param.FileResultGetParam) ([]byte, error) {

	var result []byte
	var err error
	var id int
	stdao := dao.FileResultDao{tool.DbEngine}

	if GetFileResultParam.Result_uuid == "none" && GetFileResultParam.Sys_batch_id_a == "none" && GetFileResultParam.Sys_batch_id_b == "none" {
		if result, err = stdao.SelectAll(GetFileResultParam.PS.Page, GetFileResultParam.PS.Offset); err != nil {
			return result, err
		}

	} else {
		var temp = make(map[string]interface{})
		if GetFileResultParam.Result_uuid != "none" {
			temp["file_result.result_uuid"] = GetFileResultParam.Result_uuid
		}

		if GetFileResultParam.Sys_batch_id_a != "none" {
			id, err = strconv.Atoi(GetFileResultParam.Sys_batch_id_a)
			if err != nil {
				return result, err
			}
			temp["file_result.sys_batch_id_a"] = id
		}

		if GetFileResultParam.Sys_batch_id_b != "none" {
			id, err = strconv.Atoi(GetFileResultParam.Sys_batch_id_b)
			if err != nil {
				return result, err
			}
			temp["file_result.sys_batch_id_b"] = id
		}

		if result, err = stdao.SelectByQuery(GetFileResultParam.PS.Page, GetFileResultParam.PS.Offset, temp); err != nil {
			return result, err
		}
	}

	return result, nil

}
