package Dao

import (
	"database/sql"
	"fmt"
	"mygo/myojb/pgsql/Models"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "pgadmin"
	password = "gpadmin"
	dbname   = "dw"
)

func ConnectDB() *sql.DB {
	pgsqlInfo := fmt.Sprintf("host=%s port=%d password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", pgsqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func Query(db *sql.DB, schema string) *Models.AllTablenameSql {
	Tbsql := Models.NewAllTableSql()
	var pg_tablename, pg_sql string
	rows, err := db.Query("select  routine_name ,routine_definition from information_schema.routines"+
		"where routine_schema = $1", schema)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&pg_tablename, &pg_sql)
		if err != nil {
			panic(err)
		}
		tablesql := Models.NewSingleTablenameSql(pg_tablename, pg_sql)
		Tbsql.AddAllTableSql(tablesql)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return Tbsql
}
