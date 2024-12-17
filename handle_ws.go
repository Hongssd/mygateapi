package mygateapi

type WsCandles struct {
	Timestamp   string `json:"t"` //开盘时间
	Interval    string `json:"n"` //K线周期 格式为 1m_BTC_USDT
	Open        string `json:"o"` //开盘价
	High        string `json:"h"` //最高价
	Low         string `json:"l"` //最低价
	Close       string `json:"c"` //收盘价
	Volume      string `json:"v"` //成交量
	Amount      string `json:"a"` //成交额
	WindowClose bool   `json:"w"` //是否关闭窗口，关闭表示此K线已结束
}

type WsFutureCandles struct {
	Timestamp   int64  `json:"t"` //开盘时间
	Interval    string `json:"n"` //K线周期 格式为 1m_BTC_USDT
	Open        string `json:"o"` //开盘价
	High        string `json:"h"` //最高价
	Low         string `json:"l"` //最低价
	Close       string `json:"c"` //收盘价
	Volume      int64  `json:"v"` //成交量
	Amount      string `json:"a"` //成交额
	WindowClose bool   `json:"w"` //是否关闭窗口，关闭表示此K线已结束
}

func handleWsData[T any](data []byte) (*WsSubscribeResult[T], error) {
	var res WsSubscribeResult[T]
	err := json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func splitSlice[T any, R any](origin *WsSubscribeResult[[]T], f func(o T) *R) []*WsSubscribeResult[R] {
	var result []*WsSubscribeResult[R]
	for _, v := range *origin.Result {
		r := &WsSubscribeResult[R]{
			Time:    origin.Time,
			TimeMs:  origin.TimeMs,
			ConnId:  origin.ConnId,
			Id:      origin.Id,
			Channel: origin.Channel,
			Event:   origin.Event,
			Error:   origin.Error,
			Payload: origin.Payload,
			Result:  f(v),
		}
		result = append(result, r)
	}
	return result
}
