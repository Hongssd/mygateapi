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

type WsOrder struct {
	Id                 string `json:"id"`
	Text               string `json:"text"`
	CreateTime         string `json:"create_time"`
	UpdateTime         string `json:"update_time"`
	CurrencyPair       string `json:"currency_pair"`
	Type               string `json:"type"`
	Account            string `json:"account"`
	Side               string `json:"side"`
	Amount             string `json:"amount"`
	Price              string `json:"price"`
	TimeInForce        string `json:"time_in_force"`
	Left               string `json:"left"`
	FilledTotal        string `json:"filled_total"`
	AvgDealPrice       string `json:"avg_deal_price"`
	Fee                string `json:"fee"`
	FeeCurrency        string `json:"fee_currency"`
	PointFee           string `json:"point_fee"`
	GtFee              string `json:"gt_fee"`
	RebatedFee         string `json:"rebated_fee"`
	RebatedFeeCurrency string `json:"rebated_fee_currency"`
	CreateTimeMs       string `json:"create_time_ms"`
	UpdateTimeMs       string `json:"update_time_ms"`
	User               int    `json:"user"`
	Event              string `json:"event"`
	StpId              int    `json:"stp_id"`
	StpAct             string `json:"stp_act"`
	FinishAs           string `json:"finish_as"`
	BizInfo            string `json:"biz_info"`
	AmendText          string `json:"amend_text"`
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
