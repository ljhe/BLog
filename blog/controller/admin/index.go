package admin

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	CheckLogin(w, r)
	if r.Method == "GET" {
		f, _ := template.ParseFiles("view/admin/index.html")
		_ = f.Execute(w, nil)
	}
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	CheckLogin(w, r)
	if r.Method == "GET" {
		f, _ := template.ParseFiles("view/admin/welcome.html")
		_ = f.Execute(w, nil)
	}
}
