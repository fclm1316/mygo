package dao

import (
	"allinone/model"
	// "allinone/param"
	// "allinone/po"
	"allinone/tool"
	// "fmt"
)

type SysNameDao struct {
	IUD uint8
	*tool.Orm
}

func (td *SysNameDao) SelectAll(page int, offset int) ([]byte, error) {
	listb := make([]model.Sys_name, 0)

	err := td.Engine.Table("sys_name").Cols("id,name").OrderBy("id DESC").Limit(page, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (td *SysNameDao) SelectByQuery(size int, offset int, v map[string]interface{}) ([]byte, error) {

	listb := make([]model.Sys_name, 0)

	err := td.Engine.Table("sys_name").Cols("id,name").Where(v).Limit(size, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (td *SysNameDao) InsertUpdateDelete(FromSysName *model.Sys_name) (int64, error) {
	var resultid int64
	var err error
	session := td.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	// 返回的条数
	if td.IUD == 1 {
		FromSysName.Id = 0
		resultid, err = session.Table("sys_name").Insert(FromSysName)
	}
	if td.IUD == 2 {
		resultid, err = session.Table("sys_name").ID(FromSysName.Id).Update(FromSysName)
	}
	if td.IUD == 3 {
		resultid, err = session.Table("sys_name").ID(FromSysName.Id).Delete(FromSysName)
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
