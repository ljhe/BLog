package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// json 返回格式
type JsonResult struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// 结构体转化为 json 字符串
func (j JsonResult) MapOrStructToJson() []byte {
	bytes, e := json.Marshal(j)
	CheckErr(e)
	return bytes
}

// base64 转码
func Base64Decode(str string) []byte {
	bytes, e := base64.StdEncoding.DecodeString(str)
	CheckErr(e)
	return bytes
}

// 返回错误信息
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

// md5 加密
func GetMD5String(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
