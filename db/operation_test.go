package db

import (
	"Blog/util"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestConnector_Insert(t *testing.T) {
	conn := Constructor()
	conn.Connect()
	args := map[string]interface{}{
		"ip=":       "127.0.0.1",
		"nickname=": "test",
		"time=":     time.Now().Unix(),
	}
	i, e := conn.Insert("users", args)
	if e != nil || i == 0 {
		t.Errorf("insert error: %s, id:%v", e, i)
	}
}

func TestConnector_Update(t *testing.T) {
	conn := Constructor()
	conn.Connect()
	args := map[string]interface{}{
		"ip=":       "180.169.221.150",
		"nickname=": "test11",
		"time=":     time.Now().Unix(),
	}
	where := map[string]interface{}{
		"id=": "13",
	}
	i, e := conn.Update("users", args, where)
	if e != nil || i == 0 {
		t.Errorf("update error: %s, id:%v", e, i)
	}
}

func TestConnector_Select(t *testing.T) {
	conn := Constructor()
	conn.Connect()
	field := "id, user, password, time"
	where := map[string]interface{}{
		"id=":    "1",
	}
	i, e := conn.Select("admin", field, where,"")
	if e != nil {
		t.Errorf("insert error:%s", e)
	}
	var data = util.JsonResult{}
	data.Data = i
	//fmt.Printf("data: %s\n", data.MapOrStructToJson())
	for _, v := range i {
		for _, v1 := range v {
			fmt.Println(reflect.TypeOf(v1))
		}
	}
}

func TestConnector_Delete(t *testing.T) {
	conn := Constructor()
	conn.Connect()
	where := map[string]interface{}{
		"id=":     "13",
		"status=": "0",
	}
	_, e := conn.Delete("users", where)
	if e != nil {
		t.Errorf("insert error:%s", e)
	}
}
