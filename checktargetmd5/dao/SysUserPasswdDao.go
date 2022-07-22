package dao

import (
	"allinone/model"
	// "allinone/param"
	// "allinone/po"
	"allinone/tool"
	// "fmt"
)

type SysUserPasswdDao struct {
	IUD uint8
	*tool.Orm
}

func (td *SysUserPasswdDao) SelectAll(page int, offset int) ([]byte, error) {
	listb := make([]model.Sys_userpasswd, 0)

	err := td.Engine.Table("sys_userpasswd").Cols("id,user,passwd,sys_ip_id").OrderBy(
		"id DESC").Limit(page, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (td *SysUserPasswdDao) SelectByQuery(size int, offset int, v map[string]interface{}) ([]byte, error) {

	listb := make([]model.Sys_userpasswd, 0)

	err := td.Engine.Table("sys_userpasswd").Cols(
		"id,user,passwd,sys_ip_id").GroupBy("id DESC").Where(v).Limit(size, offset).Find(&listb)
	if err != nil {
		return nil, err
	}

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (td *SysUserPasswdDao) InsertUpdateDelete(FromSysUserPasswd *model.Sys_userpasswd) (int64, error) {
	var resultid int64
	var err error
	session := td.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	// 返回的条数
	if td.IUD == 1 {
		FromSysUserPasswd.Id = 0
		resultid, err = session.Table("sys_userpasswd").Insert(FromSysUserPasswd)
	}
	if td.IUD == 2 {
		resultid, err = session.Table("sys_userpasswd").ID(FromSysUserPasswd.Id).Update(FromSysUserPasswd)
	}
	if td.IUD == 3 {
		resultid, err = session.Table("sys_userpasswd").ID(FromSysUserPasswd.Id).Delete(FromSysUserPasswd)
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
