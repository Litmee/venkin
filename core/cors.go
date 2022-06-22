package core

import "net/http"

var (
	// AllowOrigin Scopes allowed across domains
	allowOrigin string
	// AllowMethods Methods that allow cross-domain
	allowMethods string
	// AllowHeaders Allow cross-domain headers
	allowHeaders string
	// AllowOriginBool Whether a scope has been configured to allow cross-domain
	allowOriginBool bool
	// AllowMethodsBool Has the method configured to allow cross-domain
	allowMethodsBool bool
	// AllowHeadersBool Whether the header information that allows cross-domain has been configured
	allowHeadersBool bool
)

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
