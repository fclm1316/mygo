package tool

import (
	"fmt"
	"mygo/rebuildtcp/model"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {
	database := cfg.DbInfo
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=Local", database.DbUser,
		database.DbPwd, database.DbHost, database.DbPort, database.DbName, database.Charset)

	engine, err := xorm.NewEngine(database.DbDriver, conn)
	if err != nil {
		return nil, err
	}

	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(50)
	orm := new(Orm)
	orm.Engine = engine
	DbEngine = orm
	return orm, nil
}
