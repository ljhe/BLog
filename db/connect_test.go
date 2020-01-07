package db

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	conn := Constructor()
	// 查看打印结果是否为 *sql.DB
	conn.Connect()
	fmt.Println(reflect.TypeOf(conn.Db))
}
