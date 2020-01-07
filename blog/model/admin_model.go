package model

import (
	"Blog/blog/dao"
	"Blog/db"
	"Blog/util"
	"time"
)

type AdminModel struct {
}

func (am AdminModel) SelectByNameAndPassword(ad dao.Admin) []map[string]interface{} {
	field := "id"
	where := map[string]interface{}{
		"user=":     ad.User,
		"password=": ad.Password,
	}
	res, err := db.Constructor().Select("admin", field, where, "")
	if err != nil {
		util.CheckErr(err)
	}
	return res
}

func (am AdminModel) SelectById(ad dao.Admin) []map[string]interface{} {
	field := "id,user,password,status,time"
	where := map[string]interface{}{
		"id>=": 1,
	}
	res, err := db.Constructor().Select("admin", field, where, "id desc")
	if err != nil {
		util.CheckErr(err)
	}
	return res
}

func (am AdminModel) Add(ad dao.Admin) int64 {
	args := map[string]interface{}{
		"user=":     ad.User,
		"password=": util.GetMD5String(ad.Password),
		"time=":     time.Now().Format("2006-01-02 15:04:05"),
	}
	res, err := db.Constructor().Insert("admin", args)
	if err != nil {
		util.CheckErr(err)
	}
	return res
}

func (am AdminModel) ChangeStatus(ad dao.Admin) int64 {
	args := map[string]interface{}{
		"status=": ad.Status,
	}
	where := map[string]interface{}{
		"id=": ad.Id,
	}
	res, err := db.Constructor().Update("admin", args, where)
	if err != nil {
		util.CheckErr(err)
	}
	return res
}

func (am AdminModel) Delete(ad dao.Admin) int64 {
	where := map[string]interface{}{
		"id=": ad.Id,
	}
	res, err := db.Constructor().Delete("admin", where)
	if err != nil {
		util.CheckErr(err)
	}
	return res
}
