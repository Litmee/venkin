package core

import (
	"net/http"
	"reflect"
	"venkin/conf"
	"venkin/logger"
)

// GlobalHandler 全局路由回调结构体
type GlobalHandler struct {
}

// 全局路由回调分发
func (g *GlobalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 给请求配置跨域信息
	setCorsFunc(&w)
	// 跨域预检处理
	if r.Method == "OPTIONS" {
		return
	}
	// 调用请求分发函数
	handOutCore(&w, r)
}

// 请求分发中心函数
func handOutCore(w *http.ResponseWriter, r *http.Request) {
	// 全局拦截器判断
	a, b := conf.GlobalInterceptorFunc(w, r)
	if a {
		// 获取路由映射
		c, ok := RouterMap[r.URL.Path]
		if ok {
			// 通过反射获取用户自建的路由控制层的业务结构体
			copyC := reflect.New(reflect.TypeOf(c).Elem()).Interface().(Controller)
			// 初始化 ResponseWriter & Request
			copyC.initRsp(*w)
			copyC.initReq(r)
			// http method 推断执行
			copyC.judgeMethod(r.Method, copyC)
		} else {
			(*w).WriteHeader(http.StatusNotFound)
			_, err := (*w).Write([]byte("404 url not found"))
			logger.LogHttpWriteErr(err)
		}
		return
	}
	if b {
		// 内置返回信息
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		_, err := (*w).Write([]byte("405 is interceptor"))
		if err != nil {
			logger.LogHttpWriteErr(err)
			return
		}
	}
}
