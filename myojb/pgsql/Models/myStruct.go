package Models

import "fmt"

// 表名sql的结构体
type SingleTablenameSql struct {
	Tablename string
	Sql       string
}

// SingleTablenameSql 构造方法
func NewSingleTablenameSql(tablename string, sql string) *SingleTablenameSql {
	return &SingleTablenameSql{
		Tablename: tablename,
		Sql:       sql,
	}
}

// 所有 SingleTablenameSql 结构体集合
type AllTablenameSql struct {
	Allstring []*SingleTablenameSql
}

// 初始化结构体
func NewAllTableSql() *AllTablenameSql {
	return &AllTablenameSql{
		Allstring: make([]*SingleTablenameSql, 0, 10),
	}
}

// 新增的方法
func (s *AllTablenameSql) AddAllTableSql(newTbSql *SingleTablenameSql) {
	s.Allstring = append(s.Allstring, newTbSql)
}

// 查询方法

func (s *AllTablenameSql) showAllTableSql() {
	for _, v := range s.Allstring {
		fmt.Printf("tablename : %s \n", v.Tablename)
	}
}
