package conf

const Ip = "127.0.0.1:"

// DefaultWebConf Default global configuration parameters
var DefaultWebConf *WebConf

// WebConf Configure the central structure
type WebConf struct {
	// port
	Port string
	// global interceptor
	IsGlobalInterceptor GlobalInterceptor
	// scopes allowed across domains
	AllowOrigin string
	// allow cross-origin request types
	AllowMethods string
	// request headers that are allowed to be carried across domains
	AllowHeaders string
	// Mysql Addr
	MySqlAddr string
}

func init() {
	// Default configuration properties
	DefaultWebConf = &WebConf{
		Port:                "8062",
		IsGlobalInterceptor: nil,
		AllowOrigin:         "*",
		AllowHeaders:        "Content-Type",
	}
}
