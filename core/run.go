package core

import (
	"net/http"
	"venkin/conf"
	"venkin/logger"
)

// Run 引擎启动
func Run(c conf.WebConf) {
	logger.LogRun("HTTP 引擎启动中")
	// 配置全局拦截器
	conf.IsGlobalInterceptor = c.IsGlobalInterceptor
	logger.LogRun("拦截器初始化完毕")

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
	logger.LogRun("跨域参数初始化完毕")
	// 启动检查函数
	go startCheck(c.Ip + ":" + c.Port)
	// 启动服务监听
	err := http.ListenAndServe(c.Ip+":"+c.Port, &GlobalHandler{})
	if err != nil {
		panic("HTTP 服务监听启动失败")
	}
}
