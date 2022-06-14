package core

import (
	"time"
	"venkin/logger"
)

// GetReqBody 获取 http 请求携带的数据(已 JSON 序列化)
func (cI *ControllerImpl) GetReqBody() interface{} {
	return cI.data
}

// GetReqBodyFunc 获取 http 请求携带的数据工具方法
// 相比上面的使用者拿到数据后不需要自己做类型转换, 可以直接赋给控制层结构体内的 JSON 序列化结构体类型
func GetReqBodyFunc[T interface{}](cI *ControllerImpl) *T {
	return cI.data.(*T)
}

// SetRspBody 返回 http 请求数据
func (cI *ControllerImpl) SetRspBody(data []byte) {
	_, err := cI.w.Write(data)
	if err != nil {
		go func() {
			logger.LogHttpWriteErr(err)
		}()
		// 失败后的补偿机制
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
