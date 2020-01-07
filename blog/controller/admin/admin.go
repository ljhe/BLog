package admin

import (
	"Blog/util"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func List(w http.ResponseWriter, r *http.Request) {
	CheckLogin(w, r)
	if r.Method == "GET" {
		f, _ := template.ParseFiles("view/admin/admin-list.html")
		_ = f.Execute(w, nil)
	} else if r.Method == "POST" {
		res := adminService.SelectById(adminDao)
		var data = util.JsonResult{}
		if res == nil {
			data.Msg = "fail"
			_, err := w.Write(data.MapOrStructToJson())
			util.CheckErr(err)
			return
		}
		fmt.Printf("res: %s\n", res)
		data.Code = 1
		data.Msg = "success"
		data.Data = res
		fmt.Printf("data: %s\n", data.MapOrStructToJson())
		_, err := w.Write(data.MapOrStructToJson())
		util.CheckErr(err)
	}
}

func Add(w http.ResponseWriter, r *http.Request) {
	CheckLogin(w, r)
	if r.Method == "GET" {
		f, _ := template.ParseFiles("view/admin/admin-add.html")
		_ = f.Execute(w, nil)
	} else if r.Method == "POST" {
		adminDao.User = r.PostFormValue("username")
		adminDao.Password = r.PostFormValue("pass")
		res := adminService.Add(adminDao)
		var data = util.JsonResult{}
		if res == 0 {
			data.Msg = "add fail"
			_, err := w.Write(data.MapOrStructToJson())
			util.CheckErr(err)
			return
		}
		data.Code = 1
		data.Msg = "success"
		_, err := w.Write(data.MapOrStructToJson())
		util.CheckErr(err)
	}
}

func ChangeStatus(w http.ResponseWriter, r *http.Request) {
	CheckLogin(w, r)
	if r.Method == "POST" {
		var err error
		adminDao.Id, err = strconv.ParseInt(r.PostFormValue("id"), 10, 64)
		adminDao.Status, err = strconv.Atoi(r.PostFormValue("status"))
		util.CheckErr(err)
		res := adminService.ChangeStatus(adminDao)
		var data = util.JsonResult{}
		if res == 0 {
			data.Msg = "changeStatus fail"
			_, err := w.Write(data.MapOrStructToJson())
			util.CheckErr(err)
			return
		}
		data.Code = 1
		data.Msg = "success"
		_, err = w.Write(data.MapOrStructToJson())
		util.CheckErr(err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	CheckLogin(w, r)
	if r.Method == "POST" {
		var err error
		adminDao.Id, err = strconv.ParseInt(r.PostFormValue("id"), 10, 64)
		util.CheckErr(err)
		res := adminService.Delete(adminDao)
		var data = util.JsonResult{}
		if res == 0 {
			data.Msg = "delete fail"
			_, err := w.Write(data.MapOrStructToJson())
			util.CheckErr(err)
			return
		}
		data.Code = 1
		data.Msg = "success"
		_, err = w.Write(data.MapOrStructToJson())
		util.CheckErr(err)
	}
}
