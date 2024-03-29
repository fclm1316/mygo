package dao

import (
	"allinone/model"
	// "allinone/param"
	"allinone/po"
	"allinone/tool"
	// "fmt"
)

type SysInfoDao struct {
	*tool.Orm
}

func (td *SysInfoDao) SelectAll(page int, offset int) ([]byte, error) {
	listb := make([]po.SysInfoList, 0)

	err := td.Engine.Table("sys_ip").Cols("sys_ip.id,sys_ip.ipaddr,sys_name.name,sys_env.envname").Join(
		"LEFT", "sys_name", "sys_name.id=sys_ip.sys_name_id").Join("LEFT",
		"sys_env", "sys_env.id=sys_ip.sys_env_id").OrderBy(
		"sys_ip.id DESC").Limit(page, offset).Find(&listb)
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

func (td *SysInfoDao) SelectByQuery(size int, offset int, v map[string]interface{}) ([]byte, error) {

	listb := make([]po.SysInfoList, 0)
	session := td.Engine.Table("sys_ip")
	// for key, value := range v {
	// 	// fmt.Println(key, value)
	// 	str := fmt.Sprintf("%s = ?", key)
	// 	session.And(str, value)
	// }
	session.Where(v)
	session.Cols("sys_ip.id,sys_ip.ipaddr,sys_name.name,sys_env.envname").Join(
		"LEFT", "sys_name", "sys_name.id=sys_ip.sys_name_id").Join("LEFT",
		"sys_env", "sys_env.id=sys_ip.sys_env_id").OrderBy("sys_ip.id DESC").Limit(size, offset).Find(&listb)

	listbyte, err := tool.SliceToByte(listb)
	if err != nil {
		return listbyte, err
	}

	return listbyte, nil
}

func (td *SysInfoDao) SelectById(id int64) ([]byte, error) {

	lista := make([]po.SysInfoList, 0)
	td.Engine.Table("sys_ip").Cols("sys_ip.id,sys_ip.ipaddr,sys_name.name,sys_env.envname").Join(
		"INNER", "sys_name", "sys_name.id=sys_ip.sys_name_id").Join("INNER",
		"sys_env", "sys_env.id=sys_ip.sys_env_id").Where("sys_ip.id = ?", id).Find(&lista)

	listabyte, err := tool.SliceToByte(lista)
	if err != nil {
		return listabyte, err
	}
	// fmt.Println(string(listabyte))
	return listabyte, nil
}

func (td *SysInfoDao) AddInfoIP(AddSysIp *model.Sys_ip) (int64, error) {
	session := td.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}
	// 返回插入的条数
	resultid, err := session.Table("sys_ip").Insert(AddSysIp)

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
func (td *SysInfoDao) UpdateInfoIP(UpdateSysIp *model.Sys_ip) (int64, error) {
	session := td.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}

	resultid, err := session.Table("sys_ip").ID(UpdateSysIp.Id).Update(UpdateSysIp)

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
func (td *SysInfoDao) DeleteInfoIP(DeleteSysIp *model.Sys_ip) (int64, error) {
	session := td.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return 0, err
	}

	resultid, err := session.Table("sys_ip").ID(DeleteSysIp.Id).Delete(DeleteSysIp)

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
