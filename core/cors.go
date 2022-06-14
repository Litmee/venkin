package core

import "net/http"

// AllowOrigin 允许跨域的范围
var allowOrigin string

// AllowMethods 允许跨域的方法
var allowMethods string

// AllowHeaders 允许跨域的头信息
var allowHeaders string

// AllowOriginBool 是否已配置允许跨域的范围
var allowOriginBool bool

// AllowMethodsBool 是否已配置允许跨域的方法
var allowMethodsBool bool

// AllowHeadersBool 是否已配置允许跨域的头信息
var allowHeadersBool bool

// 设置跨域工具函数
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
