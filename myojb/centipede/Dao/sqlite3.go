package Dao

import (
	"fmt"
	_ "github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func InitDB(conf *ini.File) {
	var err error
	dbPath := conf.Section("sqlite3").Key("db_path").String()
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		fmt.Println("%s 数据库连接失败\n", dbPath)
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func Close() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func GetDB() *gorm.DB {
	return db
}
