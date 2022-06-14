package core

import (
	"net/http"
	"venkin/conf"
	"venkin/logger"
)

// Run 引擎启动
func Run(c conf.WebConf) {
	logger.LogRun()
	// 配置全局拦截器
	conf.IsGlobalInterceptor = c.IsGlobalInterceptor

	// 跨域参数的预处理
	if c.AllowOrigin != "" {
		allowOrigin = c.AllowOrigin
		allowOriginBool = true
	}
	if c.AllowMethods != "" {
		allowMethods = c.AllowMethods
		allowMethodsBool = true
	}
	if c.AllowHeaders != "" {
		allowHeaders = c.AllowHeaders
		allowHeadersBool = true
	}
	// 启动服务监听
	err := http.ListenAndServe(c.Ip+":"+c.Port, &GlobalHandler{})
	if err != nil {
		panic("http 服务监听启动失败")
	}
}
