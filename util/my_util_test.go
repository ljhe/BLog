package util

import (
	"fmt"
	"testing"
)

func TestJsonResult_MapOrStructToJson(t *testing.T) {
	var res []map[string]interface{}
	var one = make(map[string]interface{})
	one["id"] = 1
	one["password"] = "e10adc3949ba59abbe56e057f20f883e"
	one["status"] = 0
	one["user"] = "admin"
	res = append(res, one)
	fmt.Printf("res: %s\n", res)
	var data = JsonResult{}
	data.Code = 0
	data.Msg = "success"
	data.Data = res
	fmt.Printf("data: %s\n", data.MapOrStructToJson())
}

func TestJsonResult_Base64Decode(t *testing.T) {
	var one = make(map[string]interface{})
	one["id"] = 1
	one["password"] = "ZTEwYWRjMzk0OWJhNTlhYmJlNTZlMDU3ZjIwZjg4M2U="
	one["status"] = 0
	one["user"] = "admin"
	fmt.Printf("%s", Base64Decode(one["password"].(string)))
}
