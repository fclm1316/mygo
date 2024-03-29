package dao

import (
	"allinone/model"
	// "allinone/param"
	// "allinone/po"
	"allinone/tool"
	// "fmt"
)

type SysEnvDao struct {
	*tool.Orm
}

func (td *SysEnvDao) SelectAll(page int, offset int) ([]byte, error) {
	listb := make([]model.Sys_env, 0)

	err := td.Engine.Table("sys_env").Cols("id,envname").OrderBy("id DESC").Limit(page, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (td *SysEnvDao) SelectByQuery(size int, offset int, v map[string]interface{}) ([]byte, error) {

	listb := make([]model.Sys_env, 0)

	err := td.Engine.Table("sys_env").Cols("id,envname").Where(v).Limit(size, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (td *SysEnvDao) AddSysEnv(AddSysEnv *model.Sys_env) (int64, error) {
	session := td.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	// 返回插入的条数
	resultid, err := session.Table("sys_env").Insert(AddSysEnv)

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
func (td *SysEnvDao) UpdateSysEnv(UpdateSysEnv *model.Sys_env) (int64, error) {
	session := td.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	// 返回插入的条数
	resultid, err := session.Table("sys_env").ID(UpdateSysEnv.Id).Update(UpdateSysEnv)

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
func (td *SysEnvDao) DeleteSysEnv(DeleteSysEnv *model.Sys_env) (int64, error) {
	session := td.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	// 返回插入的条数
	resultid, err := session.Table("sys_env").ID(DeleteSysEnv.Id).Delete(DeleteSysEnv)

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
