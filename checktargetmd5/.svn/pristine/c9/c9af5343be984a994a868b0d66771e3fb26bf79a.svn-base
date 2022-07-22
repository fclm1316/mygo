package dao

import (
	"allinone/model"
	// "allinone/param"
	"allinone/po"
	"allinone/tool"
	// "encoding/json"
	// "fmt"
)

type FileListDao struct {
	IUD uint8
	*tool.Orm
}

func (td *FileListDao) SelectAll(page int, offset int) ([]byte, error) {
	listb := make([]po.FileListList, 0)

	err := td.Engine.Table("file_list").Cols(
		"file_list.id,file_list.retype,file_list.filepath,file_list.filename,file_list.status,sys_ip.ipaddr, sys_userpasswd.user,sys_userpasswd.passwd").Join(
		"LEFT", "sys_ip", "file_list.sys_ip_id = sys_ip.id").Join(
		"LEFT", "sys_userpasswd", "file_list.sys_userpasswd_id = sys_userpasswd.id").OrderBy(
		"file_list.id DESC").Limit(page, offset).Find(&listb)
	// aa, _ := json.Marshal(listb[0])
	// fmt.Print(string(aa))
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}
	// fmt.Println(listbyte)
	return listbyte, nil
}

func (td *FileListDao) SelectByQuery(size int, offset int, v map[string]interface{}) ([]byte, error) {

	listb := make([]po.FileListList, 0)

	err := td.Engine.Table("file_list").Cols(
		"file_list.id,file_list.retype,file_list.filepath,file_list.filename,file_list.status,sys_ip.ipaddr, sys_userpasswd.user,sys_userpasswd.passwd").Join(
		"LEFT", "sys_ip", "file_list.sys_ip_id = sys_ip.id").Join(
		"LEFT", "sys_userpasswd", "file_list.sys_userpasswd_id = sys_userpasswd.id").Where(v).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (td *FileListDao) InsertUpdateDelete(FromFileList *model.File_list) (int64, error) {
	var resultid int64
	var err error
	session := td.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	// 返回的条数
	if td.IUD == 1 {
		FromFileList.Id = 0
		resultid, err = session.Table("file_list").Insert(FromFileList)
	}
	if td.IUD == 2 {
		resultid, err = session.Table("file_list").ID(FromFileList.Id).Update(FromFileList)
	}
	if td.IUD == 3 {
		resultid, err = session.Table("file_list").ID(FromFileList.Id).Delete(FromFileList)
	}
	if err != nil {
		session.Rollback()
		return 0, err
	}
	err = session.Commit()
	if err != nil {
		return 0, err
	}

	return resultid, nil
}

func (td *FileListDao) SelectByFileInfo() ([]po.FileListList, error) {
	listb := make([]po.FileListList, 0)

	err := td.Engine.Table("file_list").Cols(
		"file_list.id,file_list.retype,file_list.filepath,file_list.filename,file_list.status,sys_ip.ipaddr, sys_userpasswd.user,sys_userpasswd.passwd").Join(
		"LEFT", "sys_ip", "file_list.sys_ip_id = sys_ip.id").Join(
		"LEFT", "sys_userpasswd", "file_list.sys_userpasswd_id = sys_userpasswd.id").OrderBy(
		"file_list.id DESC").Find(&listb)

	// listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return nil, err
	}

	return listb, nil
}
