package dao

import (
	"allinone/model"
	"allinone/tool"
)

type FileResultDao struct {
	*tool.Orm
}

func (fr *FileResultDao) SelectAll(page int, offset int) ([]byte, error) {
	listb := make([]model.File_result, 0)

	err := fr.Engine.Table("file_result").OrderBy("creat_at DESC").Limit(page, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (fr *FileResultDao) SelectByQuery(size int, offset int, v map[string]interface{}) ([]byte, error) {

	listb := make([]model.File_result, 0)

	err := fr.Engine.Table("file_result").Where(v).Limit(size, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (td *FileResultDao) AddFileResult(AddFileResult *model.File_result) (int64, error) {
	session := td.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	// 返回插入的条数
	resultid, err := session.Table("file_result").Insert(AddFileResult)

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
