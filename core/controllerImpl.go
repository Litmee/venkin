package core

import (
	"net/http"
	"venkin/logger"
)

// ControllerImpl Controller 接口实现体
type ControllerImpl struct {
	w    http.ResponseWriter
	r    *http.Request
	data interface{}
}

// 结构体 *http.Request 初始化方法
func (cI *ControllerImpl) initReq(r *http.Request) {
	cI.r = r
}

// 结构体 http.ResponseWriter 初始化方法
func (cI *ControllerImpl) initRsp(w http.ResponseWriter) {
	cI.w = w
}

// http method 推断方法, 决定去调用哪个请求
func (cI *ControllerImpl) judgeMethod(m string, c Controller) {
	// 进行 JSON 检查以及序列化操作
	if !cI.jsonCheckAndSerialize(c) {
		return
	}
	// JSON 序列化成功后根据 http 请求的 method 类型调用控制层模型对应实现的方法
	if m == "GET" {
		c.Get()
		return
	}
	if m == "POST" {
		c.Post()
		return
	}
	if m == "PUT" {
		c.Put()
		return
	}
	if m == "DELETE" {
		c.Delete()
		return
	}
	c.Other()
}

// Get 路由默认的 GET 请求处理方法
func (cI *ControllerImpl) Get() {
	// 开启新协程记录日志
	go func() {
		logger.LogHttpMethodErr(cI.r.URL.String(), cI.r.Method)
	}()
	// 回写 http 请求状态
	cI.w.WriteHeader(http.StatusNotFound)
	_, err := cI.w.Write([]byte("404 Method Not Found"))
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
	}
}

// Post 路由默认的 POST 请求处理方法
func (cI *ControllerImpl) Post() {
	// 开启新协程记录日志
	go func() {
		logger.LogHttpMethodErr(cI.r.URL.String(), cI.r.Method)
	}()
	// 回写 http 请求状态
	cI.w.WriteHeader(http.StatusNotFound)
	_, err := cI.w.Write([]byte("404 Method Not Found"))
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
	}
}

// Put 路由默认的 POST 请求处理方法
func (cI *ControllerImpl) Put() {
	// 开启新协程记录日志
	go func() {
		logger.LogHttpMethodErr(cI.r.URL.String(), cI.r.Method)
	}()
	// 回写 http 请求状态
	cI.w.WriteHeader(http.StatusNotFound)
	_, err := cI.w.Write([]byte("404 Method Not Found"))
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
	}
}

// Delete 路由默认的 DELETE 请求处理方法
func (cI *ControllerImpl) Delete() {
	// 开启新协程记录日志
	go func() {
		logger.LogHttpMethodErr(cI.r.URL.String(), cI.r.Method)
	}()
	// 回写 http 请求状态
	cI.w.WriteHeader(http.StatusNotFound)
	_, err := cI.w.Write([]byte("404 Method Not Found"))
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
	}
}

// Other 路由默认的其他类型请求处理方法
func (cI *ControllerImpl) Other() {
	// 开启新协程记录日志
	go func() {
		logger.LogHttpMethodErr(cI.r.URL.String(), cI.r.Method)
	}()
	// 回写 http 请求状态
	cI.w.WriteHeader(http.StatusNotFound)
	_, err := cI.w.Write([]byte("404 Method Not Found"))
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
	}
}
