package models

import (
	// "fmt"
	"log"

	// "gorm.io/gorm/logger"

	//"github.com/spf13/viper"
	// "gorm.io/driver/mysql"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "gorm.io/gorm/schema"
)

var db *gorm.DB

func InitDB() {
	//viper.SetConfigName("conf")
	//viper.SetConfigType("toml")
	////viper.AddConfigPath("english/etc")
	//viper.AddConfigPath("./etc")
	//err := viper.ReadInConfig()
	//if err != nil {
	//	panic(err)
	//}
	//mysqlUser := viper.Get("mysql.username")
	//mysqlPassword := viper.Get("mysql.password")
	//mysqlHost := viper.Get("mysql.ip")
	//mysqlPort := viper.Get("mysql.port")
	//mysqlDb := viper.Get("mysql.dbname")

	// const (
	// 	mysqlUser     = "root"
	// 	mysqlPassword = "root123"
	// 	mysqlHost     = "127.0.0.1"
	// 	mysqlPort     = 3306
	// 	mysqlDb       = "myenglish"
	// )
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDb)
	// var err error
	// db, err = gorm.Open(mysql.New(mysql.Config{
	// 	DSN:                       dsn,
	// 	DefaultStringSize:         256,
	// 	DisableDatetimePrecision:  true,
	// 	DontSupportRenameColumn:   true,
	// 	DontSupportRenameIndex:    true,
	// 	SkipInitializeWithVersion: false,
	// }), &gorm.Config{NamingStrategy: schema.NamingStrategy{
	// 	SingularTable: true,
	// },
	// 	Logger: logger.Default.LogMode(logger.Error),
	// })

	// if err != nil {
	// 	log.Printf("%s 数据库连接失败\n", mysqlHost)
	// 	log.Fatal(err)
	// }
	// sqlDB, _ := db.DB()
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(20)
	// //return nil

	var err error
	db, err = gorm.Open(sqlite.Open("e.db"), &gorm.Config{})
	if err != nil {
		log.Printf("数据库连接失败\n")
		log.Fatal(err)
	}
}

func DbClose() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

func GetDB() *gorm.DB {
	return db
}
