package mygateapi

import (
	"fmt"
)

type GateErrorRes struct {
	StatusCode int    `json:"status_code"`
	Label      string `json:"label"`
	Msg        string `json:"message"`
}

type GateTimeRes struct {
	InTime  string `json:"inTime"`  //REST网关接收请求时的时间戳，Unix时间戳的微秒数格式，如 1597026383085123返回的时间是请求验证后的时间
	OutTime string `json:"outTime"` //REST网关发送响应时的时间戳，Unix时间戳的微秒数格式，如 1597026383085123
}
type GateRestRes[T any] struct {
	GateErrorRes   //错误信息
	GateTimeRes    //时间戳
	Data         T `json:"data"` //请求结果
}

func is2xx(httpCode int) bool {
	return httpCode >= 200 && httpCode < 300
}
func handlerCommonRest[T any](data []byte, respHeaderMap map[string]string, statusCode int) (*GateRestRes[T], error) {
	res := &GateRestRes[T]{}
	log.Debug(string(data))

	res.StatusCode = statusCode
	//log.Debug(res.StatusCode)
	//捕获网关出入站时间信息 结果为微秒时间戳
	if inTimeStr, ok := respHeaderMap["X-In-Time"]; ok {
		res.InTime = inTimeStr
	}
	if outTimeStr, ok := respHeaderMap["X-Out-Time"]; ok {
		res.OutTime = outTimeStr
	}

	if is2xx(res.StatusCode) {
		//请求成功，捕获结构体
		var d T
		if len(data) == 0 {
			res.Data = d
		} else {
			err := json.Unmarshal(data, &d)
			if err != nil {
				log.Error("rest返回值获取失败", err)
				return res, err
			}
			res.Data = d
		}
	} else {
		var errRes GateErrorRes
		err := json.Unmarshal(data, &errRes)
		if err != nil {
			log.Error("rest错误信息获取失败", err)
			return res, err
		}
		res.GateErrorRes = errRes
	}
	return res, res.handlerError()
}
func (err *GateErrorRes) handlerError() error {
	if !is2xx(err.StatusCode) {
		//不为2xx的返回代码代表请求失败
		return fmt.Errorf("request error:[label:%v][message:%v]", err.Label, err.Msg)
	} else {
		return nil
	}

}
