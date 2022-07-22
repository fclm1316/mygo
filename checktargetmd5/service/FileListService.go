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

type FileList struct {
	IUD uint8
}

func (sn *FileList) GetFile(GetFileListParam *param.FileListGetParam) ([]byte, error) {

	var result []byte
	var err error
	var id int
	stdao := dao.FileListDao{0, tool.DbEngine}

	if GetFileListParam.SysIpId == "none" && GetFileListParam.Status == "none" {
		if result, err = stdao.SelectAll(GetFileListParam.PS.Page, GetFileListParam.PS.Offset); err != nil {
			return result, err
		}

	} else {
		var temp = make(map[string]interface{})
		if GetFileListParam.SysIpId != "none" {
			id, err = strconv.Atoi(GetFileListParam.SysIpId)
			if err != nil {
				return result, err
			}
			temp["file_list.sys_ip_id"] = id
		}

		if GetFileListParam.Status != "none" {
			id, err = strconv.Atoi(GetFileListParam.Status)
			if err != nil {
				return result, err
			}
			temp["file_list.status"] = id
		}

		if result, err = stdao.SelectByQuery(GetFileListParam.PS.Page, GetFileListParam.PS.Offset, temp); err != nil {
			return result, err
		}
	}

	return result, nil

}

func (sn *FileList) InsertUpdateDelete(FromFileListParam *param.FileListParam) (int64, error) {
	var resultid int64
	var err error
	FromFileList := &model.File_list{
		Id:                FromFileListParam.Id,
		Filepath:          FromFileListParam.Filepath,
		Filename:          FromFileListParam.Filename,
		Retype:            FromFileListParam.Retype,
		Sys_ip_id:         FromFileListParam.Sys_ip_id,
		Sys_userpasswd_id: FromFileListParam.Sys_userpasswd_id,
		Status:            FromFileListParam.Status,
	}
	stdao := dao.FileListDao{0, tool.DbEngine}
	if sn.IUD == 1 {
		stdao.IUD = 1
	}
	if sn.IUD == 2 {
		stdao.IUD = 2
	}
	if sn.IUD == 3 {
		stdao.IUD = 3
		// FromFileList.Filepath = ""
		// FromFileList.Filename = ""
		// FromFileList.Retype = 0
		// FromFileList.Sys_ip_id = 0
		// FromFileList.Sys_userpasswd_id = 0
		// FromFileList.Status = 0
	}
	if resultid, err = stdao.InsertUpdateDelete(FromFileList); err != nil {
		return resultid, err
	}
	return resultid, nil

}
