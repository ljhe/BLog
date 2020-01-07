package server

import (
	"Blog/util"
	"net/http"
	"strings"
)

func Web() *MyMux {
	return &MyMux{make(map[string][]*MyHandler)}
}

type MyHandler struct {
	path string
	http.Handler
}

type MyMux struct {
	handlers map[string][]*MyHandler
}

// 进行路由分配
func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handler := range m.handlers[strings.ToLower(r.Method)] {
		if handler.path == r.URL.Path {
			// 调用的是func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) { f(w, r) }
			handler.ServeHTTP(w, r)
			return
		} else if strings.HasPrefix(r.URL.Path, handler.path) {
			// 处理静态文件
			handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
	return
}

func (m *MyMux) Listen(port string) {
	err := http.ListenAndServe(port, m)
	util.CheckErr(err)
}

// 添加路由
func (m *MyMux) AddRoute(mode string, path string, handlerFunc http.HandlerFunc, handler http.Handler) {
	m.add(mode, path, handlerFunc, handler)
}

/*
 *mode         Post|Get|Put|Delete
 *path  	   前缀
 *handlerFun   http.HandlerFunc
 *handler      http.Handler
 */
func (m *MyMux) add(mode, path string, handlerFun http.HandlerFunc, handler http.Handler) {
	var h *MyHandler
	if handlerFun != nil {
		h = &MyHandler{
			strings.ToLower(path),
			handlerFun,
		}
	} else if handler != nil {
		h = &MyHandler{
			strings.ToLower(path),
			handler,
		}
	}

	// 下面是存储路由的核心代码，这里的路由m.handlers存储的格式是Get|Post|Put|Delete:String:Handler
	m.handlers[strings.ToLower(mode)] = append(
		m.handlers[strings.ToLower(mode)],
		h,
	)
}