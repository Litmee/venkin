package conf

// WebConf Web 配置中心结构体
type WebConf struct {
	Ip                  string
	Port                string
	IsGlobalInterceptor GlobalInterceptor
	AllowOrigin         string
	AllowMethods        string
	AllowHeaders        string
	MysqlUrl            string
}
