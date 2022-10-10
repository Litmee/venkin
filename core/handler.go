package core

import (
	"net/http"
	"reflect"
	"venkin/conf"
	"venkin/logger"
)

// GlobalHandler Global route callback structure
type GlobalHandler struct {
}

// Global routing callback distribution
func (g *GlobalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Configure cross-domain information for requests
	setCorsFunc(&w)
	// Cross-domain preflight processing
	if r.Method == "OPTIONS" {
		return
	}
	// Call the request dispatch function
	handOutCore(&w, r)
}

// Request distribution center function
func handOutCore(w *http.ResponseWriter, r *http.Request) {
	// Global interceptor judgment
	a, b := conf.IsGlobalInterceptor.Interceptor(w, r)
	if a {
		// get route map
		c, ok := RouterMap[r.URL.Path]
		if ok {
			// Obtain the business structure of the user-built routing control layer through reflection
			copyC := reflect.New(reflect.TypeOf(c).Elem()).Interface().(Controller)
			// initialization ResponseWriter & Request
			copyC.initRsp(w)
			copyC.initReq(r)
			// http method speculative execution
			copyC.judgeMethod(r.Method, copyC)
		} else {
			(*w).WriteHeader(http.StatusNotFound)
			_, err := (*w).Write([]byte("404 url not found"))
			logger.LogHttpWriteErr(err)
		}
		return
	}
	if b {
		// Built-in return information
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		_, err := (*w).Write([]byte("405 is interceptor"))
		if err != nil {
			// logger.LogHttpWriteErr(err)
			return
		}
	}
}
