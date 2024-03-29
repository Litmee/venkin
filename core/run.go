package core

import (
	"database/sql"
	"net/http"
	"venkin/conf"
	"venkin/logger"
	"venkin/orm"
)

// Run Engine Start
func Run(c *conf.WebConf) {
	logger.LogRun("HTTP Engine Starting")
	if c == nil {
		// Enable Default Configuration
		c = conf.DefaultWebConf
		logger.LogRun("####################### Default Configuration Enabled #######################")
		logger.LogRun("Port: " + c.Port)
		logger.LogRun("GlobalInterceptor: nil")
		logger.LogRun("Access-Control-Allow-Origin: " + c.AllowOrigin)
		logger.LogRun("Access-Control-Allow-Headers: " + c.AllowHeaders)
		logger.LogRun("####################### Default Configuration Enabled #######################")
	}
	// Configure Global Interceptors
	if c.IsGlobalInterceptor != nil {
		conf.IsGlobalInterceptor = c.IsGlobalInterceptor
		logger.LogRun("The Interceptor Is Initialized")
	}

	// Preprocessing Of Cross-Domain Parameters
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
	logger.LogRun("Cross-Domain Parameter Initialization Is Complete")

	if c.MySqlAddr != "" {
		// Connect to the database
		db, err := sql.Open("mysql", c.MySqlAddr)
		db.SetMaxIdleConns(3)
		// Maximum number of connections
		db.SetMaxOpenConns(20)
		if err != nil {
			panic("Database connection failed")
		}
		orm.GlobalDB = db
		orm.GlobalDB.Ping()
		orm.InitGlobalSqlMap()
		defer orm.GlobalDB.Close()
		logger.LogRun("Database connection succeeded")
	}

	// Call Check Function
	go startCheck(conf.Ip + c.Port)

	// Start Service Monitoring
	err := http.ListenAndServe(conf.Ip+c.Port, &GlobalHandler{})
	if err != nil {
		panic("HTTP Service Monitoring Failed To Start")
	}
}
