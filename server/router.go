package server

import (
	"Blog/db"
	"Blog/helper"
	"net/http"
)

type App struct {
	Router     Router
	Connector  *db.Connector
	SessionMgr *helper.SessionMgr
}

func New() *App {
	return &App{
		Config(),
		db.Constructor(),
		helper.NewSessionMgr("cookieName", 3600),
	}
}

type Router struct {
	Config map[int][]interface{}
	Static map[int][]string
}

func (app App) Run(port string) {
	// 连接数据库
	app.Connector.Connect()

	web := Web()
	// 添加动态路由
	for _, v := range app.Router.Config {
		web.AddRoute(v[0].(string), v[1].(string), v[2].(func(http.ResponseWriter, *http.Request)), nil)
	}
	// 添加静态文件
	for _, v := range app.Router.Static {
		web.AddRoute(v[0], v[1], nil, http.StripPrefix(v[1], http.FileServer(http.Dir(v[2]))))
	}
	web.Listen(port)
}
