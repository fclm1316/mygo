package dao

import (
	"allinone/model"
	"allinone/po"
	"allinone/tool"
)

type FileInfoDao struct {
	*tool.Orm
}

func (fl *FileInfoDao) SelectAll(page int, offset int) ([]byte, error) {
	listb := make([]model.File_info, 0)

	err := fl.Engine.Table("file_info").OrderBy("creat_at DESC").Limit(page, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (fl *FileInfoDao) SelectByQuery(size int, offset int, v map[string]interface{}) ([]byte, error) {

	listb := make([]model.File_info, 0)

	err := fl.Engine.Table("file_info").Where(v).Limit(size, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (fl *FileInfoDao) InsertFileInfo(AddFileinfo model.File_info) (int64, error) {
	session := fl.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	// 返回插入的条数
	resultid, err := session.Table("file_info").Insert(AddFileinfo)

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

func (fl *FileInfoDao) SelectbySysBatchId(id int64) (*[]po.FileInfoByBatch, error) {
	result := make([]po.FileInfoByBatch, 0)
	err := fl.Engine.Table("file_info").Cols("file_list_id,filename,filemd5,ipaddr").Where("sys_batch_id=?", id).Find(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
