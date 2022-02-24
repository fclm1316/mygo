package main

import (
	"fmt"
	"mygo/myojb/pgsql/Dao"
	"mygo/myojb/pgsql/MyTools"
)

func main() {
	db := Dao.ConnectDB()
	aa := Dao.Query(db, "adm_acms")
	for _, v := range aa.Allstring {
		fmt.Println("==========")
		MyTools.CutClearString("adm_acms", v.Tablename, v.Sql)
	}
}
