package db

import (
	"Blog/util"
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
)

var conn Connector

type Connector struct {
	Db *sql.DB
}

func Constructor() *Connector {
	return &conn
}

// 连接数据库
func (conn *Connector) Connect() {
	var e error
	// 这里使用的 dataSourceName 格式为 user:password@tcp(localhost:5555)/dbname?charset=utf8
	conn.Db, e = sql.Open("mysql", USERNAME + ":" + PASSWORD + "@tcp(" + HOST + ":" + PORT +
		")/" + DATEBASE + "?charset=utf8")
	util.CheckErr(e)

	// 检查与数据库的连接是否有效
	e = conn.Db.Ping()
	util.CheckErr(e)
}
