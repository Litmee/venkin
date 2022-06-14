package conf

import "net/http"

// IsGlobalInterceptor 全局拦截器参数
var IsGlobalInterceptor GlobalInterceptor

// GlobalInterceptor 全局拦截器接口
// 自定义拦截器需要实现该接口
type GlobalInterceptor interface {
	// Interceptor 返回参数1: true 代表放行, false 代表拒绝
	// Interceptor 返回参数2: true 代表走框架默认被拦截后的返回信息, false 代表不走框架内置返回信息此时需要开发者自定义返回信息不然调用者接收不到任何提示信息
	Interceptor(w *http.ResponseWriter, r *http.Request) (bool, bool)
}

// GlobalInterceptorFunc 全局拦截器校验方法
func GlobalInterceptorFunc(w *http.ResponseWriter, r *http.Request) (bool, bool) {
	if IsGlobalInterceptor != nil {
		return IsGlobalInterceptor.Interceptor(w, r)
	}
	return true, false
}
