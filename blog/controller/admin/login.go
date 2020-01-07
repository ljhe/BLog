package admin

import (
	"Blog/blog/dao"
	"Blog/blog/service"
	"Blog/helper"
	"Blog/util"
	"html/template"
	"net/http"
)

var adminService service.AdminService
var adminDao dao.Admin
var SessionID string

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		f, _ := template.ParseFiles("view/admin/login.html")
		_ = f.Execute(w, nil)
	} else if r.Method == "POST" {
		adminDao.User = r.PostFormValue("username")
		adminDao.Password = util.GetMD5String(r.PostFormValue("pass"))
		var data = util.JsonResult{}
		// 检查账户密码是否为空
		if adminDao.User == "" || adminDao.Password == "" {
			data.Msg = "user and password must be not null"
			_, err := w.Write(data.MapOrStructToJson())
			util.CheckErr(err)
			return
		}
		res := adminService.SelectByNameAndPassword(adminDao)
		// 判断账户密码是否正确
		if res == nil {
			data.Msg = "user or password is error"
			_, err := w.Write(data.MapOrStructToJson())
			util.CheckErr(err)
			return
		}
		// 存储 session
		SessionID = helper.MySessionMgr.StartSession(w, r)
		helper.MySessionMgr.SetSessionVal(SessionID, "userInfo", adminDao.User)

		data.Code = 1
		_, err := w.Write(data.MapOrStructToJson())
		util.CheckErr(err)
	}
}

func Logout(w http.ResponseWriter, r *http.Request)  {
	// 删除 session 处理登出
	helper.MySessionMgr.EndSessionBy(SessionID)
	http.Redirect(w, r, "/admin/login", http.StatusFound)
}

func CheckLogin(w http.ResponseWriter, r *http.Request) {
	// 判断是否登陆 如果未登陆 跳转到登陆页面
	_, t := helper.MySessionMgr.GetSessionVal(SessionID, "userInfo")
	if t == false {
		http.Redirect(w, r, "/admin/login", http.StatusFound)
	}
}
