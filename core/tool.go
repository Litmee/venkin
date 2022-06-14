package core

import (
	"net"
	"time"
	"venkin/logger"
)

// GetReqBody Get The Data Carried By The Http Request
// The Data Has Been Serialized To The Type Assigned By The User At The Beginning
// But The Corresponding Type Conversion Is Still Required After Getting It
func (cI *ControllerImpl) GetReqBody() interface{} {
	return cI.data
}

// GetReqBodyFunc Get The Data Tool Method Carried By The Http Request
// Compared With The GetReqBody Method
// Users Don't Need To Do Type Conversion After They Get The Data
// They Can Directly Assign The Required Type After Getting The Return Value
// This Function Is Recommended To Avoid Unnecessary Type Conversion Errors During Development
func GetReqBodyFunc[T interface{}](cI *ControllerImpl) *T {
	return cI.data.(*T)
}

// SetRspBody Return The Data Required By The Http Request
func (cI *ControllerImpl) SetRspBody(data []byte) {
	_, err := cI.w.Write(data)
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
		// Compensation mechanism after failure
		i := 0
		for i < 3 {
			i += 1
			_, err = cI.w.Write(data)
			if err != nil {
				go func() {
					logger.LogHttpWriteErr(err)
				}()
				time.Sleep(time.Millisecond * 500)
				continue
			}
			break
		}
	}
}

// Engine Start Check Function
func startCheck(addr string) {
	var conn net.Conn
	var err error
	defer func(conn *net.Conn) {
		err := (*conn).Close()
		if err != nil {
			return
		}
	}(&conn)
	for {
		conn, err = net.Dial("tcp", addr)
		if err == nil {
			logger.LogRun("HTTP Engine Started Successfully")
			break
		}
	}
}
