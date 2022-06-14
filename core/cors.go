package core

import "net/http"

// AllowOrigin Scopes allowed across domains
var allowOrigin string

// AllowMethods Methods that allow cross-domain
var allowMethods string

// AllowHeaders Allow cross-domain headers
var allowHeaders string

// AllowOriginBool Whether a scope has been configured to allow cross-domain
var allowOriginBool bool

// AllowMethodsBool Has the method configured to allow cross-domain
var allowMethodsBool bool

// AllowHeadersBool Whether the header information that allows cross-domain has been configured
var allowHeadersBool bool

// Set up cross-domain utility functions
func setCorsFunc(w *http.ResponseWriter) {
	if allowOriginBool {
		(*w).Header().Set("Access-Control-Allow-Origin", allowOrigin)
	}
	if allowMethodsBool {
		(*w).Header().Set("Access-Control-Allow-Methods", allowMethods)
	}
	if allowHeadersBool {
		(*w).Header().Set("Access-Control-Allow-Headers", allowHeaders)
	}
}
