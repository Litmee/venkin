package conf

import "net/http"

// IsGlobalInterceptor Global interceptor parameters
var IsGlobalInterceptor GlobalInterceptor = &DefaultInterceptor{}

// GlobalInterceptor Global interceptor interface
// Custom interceptors need to implement this interface
type GlobalInterceptor interface {
	// Interceptor Return parameter 1: true means release, false means reject
	// Interceptor Return parameter 2: true represents the return information after the framework is intercepted by default
	// false represents the return information built-in without the framework
	// At this time, the developer needs to customize the return information, otherwise the caller will not receive any prompt information
	Interceptor(w *http.ResponseWriter, r *http.Request) (bool, bool)
}

// DefaultInterceptor The interceptor interface that comes with the framework implements the structure
// Of course it does not have the actual interception function
type DefaultInterceptor struct{}

// Interceptor Global interceptor verification method
func (di *DefaultInterceptor) Interceptor(w *http.ResponseWriter, r *http.Request) (bool, bool) {
	return true, false
}
