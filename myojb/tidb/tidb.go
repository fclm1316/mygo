package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"mygo/myojb/tools"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	conf, err := ini.Load("myojb/tidb/tidb.ini")
	if err != nil {
		fmt.Println("配置文件失败 err = ", err)
	}

	sysThread, _ := conf.Section("bing").Key("thread").Int()
	sysCount, _ := conf.Section("bing").Key("count").Int()

	mysqlName := conf.Section("tidb").Key("db_name").String()
	mysqlUser := conf.Section("tidb").Key("db_user").String()
	mysqlPwd := conf.Section("tidb").Key("db_pwd").String()
	mysqlHost := conf.Section("tidb").Key("db_host").String()
	mysqlPort, _ := conf.Section("tidb").Key("db_port").Int()
	mysqlCharset := conf.Section("tidb").Key("db_charset").String()

	execSql := conf.Section("tidb").Key("sqlInsert").String()
	constr := fmt.Sprintf("%s:%s@tcp(%s:%d/%s?charset=%s&parseTime=True&loc=Local)",
		mysqlUser, mysqlPwd, mysqlHost, mysqlPort, mysqlName, mysqlCharset)

	var begin = time.Now()
	fmt.Println("开始时间：", begin)
	for i := 0; i < sysThread; i++ {
		wg.Add(1)
		fmt.Println("协程启动:", i)
		go InsertSql(&constr, &execSql, &sysCount)
	}
	wg.Wait()
	var end = time.Now().Sub(begin)
	var timestamp float64 = end.Seconds()
	tps := float64(sysThread) * float64(sysCount) / timestamp
	fmt.Println("结束时间：", time.Now())
	fmt.Println("耗时：", timestamp)
	fmt.Printf("TPS: %.2f \n", tps)

}

func InsertSql(constr *string, dosql *string, count *int) {
	defer wg.Done()
	db, err := sql.Open("mysql", *constr)
	checkError(err)
	defer db.Close()
	for i := 0; i < *count; i++ {
		round := tools.RandStingBytesMaskImprSrcUnsafe(10)
		id := tools.SumHash(round)
		stmt, err := db.Prepare(*dosql)
		checkError(err)
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}

func SelectSql(constr *string, dosql *string) {
	db, err := sql.Open("mysql", *constr)
	checkError(err)
	defer db.Close()
	rows, err := db.Query(*dosql)
	fmt.Println(*dosql)
	var id, name, age []byte // null 值
	for rows.Next() {
		err = rows.Scan(&id, &name, &age)
		checkError(err)
		fmt.Println(string(id), string(name), string(age))
	}
}
func checkError(e error) {
	if e != nil {
		panic(e.Error())
	}
}
