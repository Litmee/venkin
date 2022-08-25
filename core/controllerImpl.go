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
func (cI *ControllerImpl) initRsp(w *http.ResponseWriter) {
	cI.w = *w
}

// Infer method, decide which request to call
func (cI *ControllerImpl) judgeMethod(m string, c Controller) {
	// JSON inspection and serialization
	if !cI.jsonCheckAndSerialize(c) {
		return
	}
	// After the JSON serialization is successful
	// The method corresponding to the implementation of the control layer model is called according to the method type of the http request
	if m == http.MethodGet {
		c.Get()
		return
	}
	if m == http.MethodPost {
		c.Post()
		return
	}
	if m == http.MethodPut {
		c.Put()
		return
	}
	if m == http.MethodDelete {
		c.Delete()
		return
	}
	c.Other()
}

// Get Route default GET request handling method
func (cI *ControllerImpl) Get() {
	unifiedNoMethod(cI)
}

// Post Route default POST request handling method
func (cI *ControllerImpl) Post() {
	unifiedNoMethod(cI)
}

// Put Route default POST request handling method
func (cI *ControllerImpl) Put() {
	unifiedNoMethod(cI)
}

// Delete Route the default DELETE request handler
func (cI *ControllerImpl) Delete() {
	unifiedNoMethod(cI)
}

// Other Route default other types of request processing methods
func (cI *ControllerImpl) Other() {
	unifiedNoMethod(cI)
}

// Unified processing function for non-existent request methods
func unifiedNoMethod(c *ControllerImpl) {
	// logging
	logger.LogHttpMethodErr(c.r.URL.String(), c.r.Method)
	// write-back http request status
	c.w.WriteHeader(http.StatusNotFound)
	c.w.Write([]byte("404 Method Not Found"))
}
