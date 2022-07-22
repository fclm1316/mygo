package tool

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
	*xorm.Session
}

func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.DbInfo
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local", database.Dbuser, database.Dbpasswd,
		database.Dbip, database.Dbport, database.Dbname, database.Charset)

	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		return nil, err
	}

	err = engine.Ping()
	if err != nil {
		return nil, err
	}
	// defer engine.Close()
	engine.ShowSQL(database.Showsql)

	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(50)
	orm := new(Orm)
	orm.Engine = engine

	DbEngine = orm
	return orm, nil
}
