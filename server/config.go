package server

import (
	"Blog/blog/controller/admin"
	"net/http"
)

// 控制器各方法的接口
type Controller interface {
	Login(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	Welcome(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	AdminList(w http.ResponseWriter, r *http.Request)
	AdminAdd(w http.ResponseWriter, r *http.Request)
	ChangeStatus(w http.ResponseWriter, r *http.Request)
	AdminDelete(w http.ResponseWriter, r *http.Request)
}

// 实现控制器各方法接口
func (router Router) Login(w http.ResponseWriter, r *http.Request) {
	admin.Login(w, r)
}
func (router Router) Index(w http.ResponseWriter, r *http.Request) {
	admin.Index(w, r)
}
func (router Router) Welcome(w http.ResponseWriter, r *http.Request) {
	admin.Welcome(w, r)
}
func (router Router) Logout(w http.ResponseWriter, r *http.Request) {
	admin.Logout(w, r)
}
func (router Router) AdminList(w http.ResponseWriter, r *http.Request) {
	admin.List(w, r)
}
func (router Router) AdminAdd(w http.ResponseWriter, r *http.Request) {
	admin.Add(w, r)
}
func (router Router) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	admin.ChangeStatus(w, r)
}
func (router Router) AdminDelete(w http.ResponseWriter, r *http.Request) {
	admin.Delete(w, r)
}

// 路由配置
func Config() Router {
	var r Router
	// 加载静态文件
	r.Static = map[int][]string{
		0: {"GET", "/public/static/css", "./public/static/css"},
		1: {"GET", "/public/static/images", "./public/static/images"},
		2: {"GET", "/public/static/lib", "./public/static/lib"},
		3: {"GET", "/public/static/js", "./public/static/js"},
	}
	// 加载动态路由
	r.Config = map[int][]interface{}{
		0: {"GET", "/admin/login", r.Login},
		1: {"POST", "/admin/login", r.Login},
		2: {"GET", "/admin/index", r.Index},
		3: {"GET", "/admin/welcome", r.Welcome},
		4: {"GET", "/admin/logout", r.Logout},
		5: {"GET", "/admin/admin-list", r.AdminList},
		6: {"POST", "/admin/admin-list", r.AdminList},
		7: {"GET", "/admin/admin-add", r.AdminAdd},
		8: {"POST", "/admin/admin-add", r.AdminAdd},
		9: {"POST", "/admin/change-status", r.ChangeStatus},
		10: {"POST", "/admin/admin-delete", r.AdminDelete},
	}
	return r
}
