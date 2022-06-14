package core

import (
	"encoding/json"
	"net/http"
	"reflect"
	"venkin/logger"
)

// http 请求参数 JSON 检查以及序列化

func (cI *ControllerImpl) jsonCheckAndSerialize(c Controller) bool {
	// 反射获取传入的 Controller 接口类型数据的真实结构体类型
	rV := reflect.TypeOf(c).Elem()
	// 判断是否需要序列化请求中的数据
	// 按照约定大于配置的原则, 框架使用者声明的控制层结构体需要至少组合 ControllerImpl 结构体
	// 如果需要拿到 http 请求中的数据, 则需要额外组合一个使用者自定义的结构体用来接收 JSON 数据
	// 且使用者的控制层结构体必须先引入 ControllerImpl  结构体, 后引入需要的 JSON 序列化结构体
	// 所以如果 rV 反射参数的 NumField 结果大于 1, 则框架默认为使用者需要接收请求中的数据并需要 JSON 序列化
	if rV.NumField() > 1 {
		model := reflect.New(rV.Field(1).Type).Interface()
		var buf [512]byte
		n, _ := cI.r.Body.Read(buf[:])
		if n > 0 {
			err := json.Unmarshal(buf[:n], &model)
			// JSON 序列化发生错误则返回 500 错误
			if err != nil {
				cI.w.WriteHeader(http.StatusInternalServerError)
				cI.w.Write([]byte("JSON 序列化发生异常, 请确认请求的 Content-Type 类型以及数据结构"))
				go func() {
					logger.LogJsonSerialize(err)
				}()
				return false
			}
		}
		// 将序列化成功后的数据挂载回 ControllerImpl
		cI.data = model
		return true
	}
	return true
}
