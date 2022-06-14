package core

import "net/http"

// Controller 控制层接口模型
// 用来定义控制层结构
// Get()      Http GET 请求
// Post()     Http POST 请求
// Put()      Http PUT 请求
// Delete()   Http DELETE 请求
// Other()    Http 其他请求
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
