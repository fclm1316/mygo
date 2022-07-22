package service

import (
	"allinone/dao"
	"allinone/model"
	"allinone/param"
	"allinone/po"
	"allinone/tool"

	// "encoding/json"
	"errors"
	"fmt"
	"path"
	"strconv"
	"strings"
	"time"
)

type FileInfo struct {
	ValueChan chan po.FileListList
	OutChan   chan bool
}

type Result struct {
	Md5         string `json: md5`
	Filename    string `json: md5`
	Filesize    int    `json: md5`
	Filemodtime int    `json: md5`
}

// var ValueChan = make(chan po.FileListList, 2)
// var OutChan = make(chan bool)

func (fi *FileInfo) GetFileInfo(GetFIP *param.FileInfoGetParam) ([]byte, error) {

	var result []byte
	var err error
	var id int
	stdao := dao.FileInfoDao{tool.DbEngine}

	if GetFIP.Ipaddr == "none" && GetFIP.Sys_batch_id == "none" && GetFIP.Status == "none" && GetFIP.File_list_id == "none" {
		if result, err = stdao.SelectAll(GetFIP.PS.Page, GetFIP.PS.Offset); err != nil {
			return result, err
		}

	} else {
		var temp = make(map[string]interface{})
		if GetFIP.Sys_batch_id != "none" {
			id, err = strconv.Atoi(GetFIP.Sys_batch_id)
			if err != nil {
				return result, err
			}
			temp["file_info.sys_batch_id"] = id
		}

		if GetFIP.File_list_id != "none" {
			id, err = strconv.Atoi(GetFIP.File_list_id)
			if err != nil {
				return result, err
			}
			temp["file_info.file_list_id"] = id
		}

		if GetFIP.Status != "none" {
			id, err = strconv.Atoi(GetFIP.Status)
			if err != nil {
				return result, err
			}
			temp["file_info.status"] = id
		}

		if GetFIP.Ipaddr != "none" {
			temp["file_info.ipaddr"] = GetFIP.Ipaddr
		}

		if result, err = stdao.SelectByQuery(GetFIP.PS.Page, GetFIP.PS.Offset, temp); err != nil {
			return result, err
		}
	}

	return result, nil

}

func (fi *FileInfo) GetFileMd5() (int64, error) {

	if tool.GV.GetMd5State == true {
		err := errors.New("程序正在运行")
		return 0, err
	}

	tool.GV.GetMd5State = true
	var err error
	var resultid int64
	var chanleng int
	fldao := dao.FileListDao{4, tool.DbEngine}
	mimi, _ := fldao.SelectByFileInfo()

	if len(mimi) > 100 {
		chanleng = 100
	} else {
		chanleng = len(mimi) % 2
	}

	sbdao := dao.SysBatchDao{tool.DbEngine}
	batch := new(model.Sys_batch)
	batchid, err := sbdao.InsertSysBatch(batch)
	if err != nil {
		return 0, err
	}

	fi.ValueChan = make(chan po.FileListList, chanleng)
	fi.OutChan = make(chan bool)
	if err != nil {
		return resultid, err
	}

	go func() {
		for _, value := range mimi {
			fi.ValueChan <- value
		}
		fi.OutChan <- true

	}()

	go fi.DoGet(batchid)
	return 1, nil

}

func (rs *Result) ParseResult(result string) (Result, error) {
	fileresult := Result{}
	temp1 := strings.Split(result, "\n")
	Md5Filename := temp1[0]
	FilesizeFilemodtime := temp1[1]

	temp2 := strings.Split(Md5Filename, "  ")
	fileresult.Md5 = temp2[0]
	fileresult.Filename = temp2[1]

	temp3 := strings.Split(FilesizeFilemodtime, "  ")
	// fmt.Println(FilesizeFilemodtime)

	filesize, err := strconv.Atoi(temp3[0])
	if err != nil {
		return fileresult, err
	}

	filemodtime, err := strconv.Atoi(temp3[1])
	if err != nil {
		return fileresult, err
	}

	fileresult.Filesize = filesize
	fileresult.Filemodtime = filemodtime

	return fileresult, nil

}

func (fi *FileInfo) DoGet(batchid int64) {

	for {
		select {
		case value := <-fi.ValueChan:

			go func(batchid int64) {

				var cmd string
				var sshc tool.Sshinfo

				abspath := path.Join(value.Filepath, value.Filename)

				if value.Retype == 1 {

					cmd = fmt.Sprintf("md5sum %s && stat --printf='%%s  %%Y\\n' %s", abspath, abspath)

				} else {
					cmd = fmt.Sprintf("aa=`ls -t %s|head -n1` && md5sum $aa && stat --printf='%%s  %%Y\\n' $aa", abspath)

				}

				// fmt.Println(value.Ipaddr, value.User, value.Passwd)
				var file_info model.File_info
				file_info.File_list_id = value.Id
				file_info.Sys_batch_id = batchid
				file_info.Ipaddr = value.Ipaddr
				file_info.User = value.User

				sshc = tool.NewSSHClient(value.Ipaddr, value.User, value.Passwd, cmd)
				err := sshc.Run()

				if err != nil {
					file_info.Filename = abspath
					file_info.Status = 2
					file_info.Statusmsg = err.Error()

				} else {
					RP := Result{}
					parseresult, err := RP.ParseResult(sshc.Result)
					if err != nil {
						file_info.Filename = parseresult.Filename
						file_info.Status = 2
						file_info.Statusmsg = err.Error()
					} else {

						file_info.Filemd5 = parseresult.Md5
						file_info.Filemodtime = parseresult.Filemodtime
						file_info.Filename = parseresult.Filename
						file_info.Filesize = parseresult.Filesize
						file_info.Status = 1

					}
				}
				stdao := dao.FileInfoDao{tool.DbEngine}
				resultid, err := stdao.InsertFileInfo(file_info)
				if err != nil {
					//return resultid, err
					fmt.Println(resultid, err)
				}

			}(batchid)

		case <-fi.OutChan:
			goto gotofinish
		default:
			time.Sleep(1 * time.Second)

		}

	}
gotofinish:
	tool.GV.GetMd5State = false
	close(fi.ValueChan)
	close(fi.OutChan)
}
