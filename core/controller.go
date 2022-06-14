package core

import "net/http"

// Controller Control Layer Interface Model
// Used to define the type of request
// Get()      Http GET Request
// Post()     Http POST Request
// Put()      Http PUT Request
// Delete()   Http DELETE Request
// Other()    Http Other Requests
type Controller interface {
	Get()
	Post()
	Put()
	Delete()
	Other()
	initReq(r *http.Request)
	initRsp(w http.ResponseWriter)
	judgeMethod(m string, c Controller)
}
