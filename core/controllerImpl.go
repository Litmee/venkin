package core

import (
	"net/http"
	"venkin/logger"
)

// ControllerImpl Controller interface implementation
type ControllerImpl struct {
	w    http.ResponseWriter
	r    *http.Request
	data interface{}
}

// structure *http.Request parameter initialization method
func (cI *ControllerImpl) initReq(r *http.Request) {
	cI.r = r
}

// structure http.ResponseWriter parameter initialization method
func (cI *ControllerImpl) initRsp(w http.ResponseWriter) {
	cI.w = w
}

// Infer method, decide which request to call
func (cI *ControllerImpl) judgeMethod(m string, c Controller) {
	// JSON inspection and serialization
	if !cI.jsonCheckAndSerialize(c) {
		return
	}
	// After the JSON serialization is successful
	// the method corresponding to the implementation of the control layer model is called according to the method type of the http request
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

// Get Route default GET request handling method
func (cI *ControllerImpl) Get() {
	// Start new coroutine logging
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

// Post Route default POST request handling method
func (cI *ControllerImpl) Post() {
	// Start new coroutine logging
	go func() {
		logger.LogHttpMethodErr(cI.r.URL.String(), cI.r.Method)
	}()
	// write-back http request status
	cI.w.WriteHeader(http.StatusNotFound)
	_, err := cI.w.Write([]byte("404 Method Not Found"))
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
	}
}

// Put Route default POST request handling method
func (cI *ControllerImpl) Put() {
	// Start new coroutine logging
	go func() {
		logger.LogHttpMethodErr(cI.r.URL.String(), cI.r.Method)
	}()
	// write-back http request status
	cI.w.WriteHeader(http.StatusNotFound)
	_, err := cI.w.Write([]byte("404 Method Not Found"))
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
	}
}

// Delete Route the default DELETE request handler
func (cI *ControllerImpl) Delete() {
	// Start new coroutine logging
	go func() {
		logger.LogHttpMethodErr(cI.r.URL.String(), cI.r.Method)
	}()
	// write-back http request status
	cI.w.WriteHeader(http.StatusNotFound)
	_, err := cI.w.Write([]byte("404 Method Not Found"))
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
	}
}

// Other Route default other types of request processing methods
func (cI *ControllerImpl) Other() {
	// Start new coroutine logging
	go func() {
		logger.LogHttpMethodErr(cI.r.URL.String(), cI.r.Method)
	}()
	// write-back http request status
	cI.w.WriteHeader(http.StatusNotFound)
	_, err := cI.w.Write([]byte("404 Method Not Found"))
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
	}
}
