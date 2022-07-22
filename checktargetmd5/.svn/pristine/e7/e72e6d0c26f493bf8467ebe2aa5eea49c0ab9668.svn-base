package dao

import (
	"allinone/model"
	// "allinone/param"
	// "allinone/po"
	"allinone/tool"
	// "fmt"
)

type SysBatchDao struct {
	*tool.Orm
}

func (sb *SysBatchDao) SelectAll(page int, offset int) ([]byte, error) {
	listb := make([]model.Sys_batch, 0)

	err := sb.Engine.Table("sys_batch").OrderBy("id DESC").Limit(page, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (sb *SysBatchDao) SelectByQuery(size int, offset int, v map[string]interface{}) ([]byte, error) {

	listb := make([]model.Sys_batch, 0)

	err := sb.Engine.Table("sys_batch").Where(v).Limit(size, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (sb *SysBatchDao) InsertSysBatch(AddSysBatch *model.Sys_batch) (int64, error) {
	session := sb.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	// 返回插入的条数
	_, err := session.Table("sys_batch").InsertOne(&AddSysBatch)

	if err != nil {
		session.Rollback()
		return 0, err
	}
	err = session.Commit()
	if err != nil {
		return 0, err
	}

	return AddSysBatch.Id, nil
}

func (sb *SysBatchDao) UionDeleteSysBatchFileIfo(DeleteSysBatchFileIfo *model.Sys_batch) (int64, error) {

	session := sb.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}

	DeteteFileInfo := &model.File_info{
		Sys_batch_id: DeleteSysBatchFileIfo.Id,
	}
	rows_file_info, err := session.Table("file_info").Delete(DeteteFileInfo)
	rows_sys_batch, err := session.Table("sys_batch").Delete(DeleteSysBatchFileIfo)

	if err != nil {
		session.Rollback()
		return 0, err
	}
	err = session.Commit()
	if err != nil {
		return 0, err
	}

	return rows_file_info + rows_sys_batch, nil
}
