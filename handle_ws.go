package mygateapi

import (
	"strconv"

	"github.com/shopspring/decimal"
)

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

type WsOrderBook struct {
	Timestamp int64        `json:"t"`  //深度更新毫秒时间戳
	Symbol    string       `json:"s"`  //货币对
	Id        int64        `json:"id"` //深度 ID
	FirstId   int64        `json:"U"`  //自动上次深度后的第一个深度 ID
	LastId    int64        `json:"u"`  //自动上次深度后的最新深度 ID
	Bids      []PriceLevel `json:"b"`  //自上次更新以来的 bids 更新，按价格从高到低排序
	Asks      []PriceLevel `json:"a"`  //自上次更新以来的 asks 更新，按价格从地到高排序

}
type PriceLevel struct {
	Price    string `json:"p"` //档位价格
	Quantity string `json:"q"` //档位的数量
}

type WsTrade struct {
	Id           int64  `json:"id"`             //成交 ID
	CreateTime   int64  `json:"create_time"`    //成交时间，精确到秒
	CreateTimeMs int64  `json:"create_time_ms"` //成交时间，毫秒精度
	Side         string `json:"side"`           //买单或者卖单
	Symbol       string `json:"symbol"`         //交易货币对
	Amount       string `json:"amount"`         //交易数量
	Price        string `json:"price"`          //交易价
	Range        string `json:"range"`          //成交范围(格式: "开始 ID-结束 ID")
	IsInternal   bool   `json:"is_internal"`    //是否为内部成交。内部成交是指保险资金和 ADL 用户对强平指令的接管。由于不是市场深度上的正常撮合，交易价格可能会出现偏差，不会记录在 K 线上。如果不是内部交易，则该字段不会被返回。
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

func (fc *WsFutureCandles) convertToWsCandle() *WsCandles {
	return &WsCandles{
		Timestamp:   strconv.FormatInt(fc.Timestamp, 10),
		Interval:    fc.Interval,
		Open:        fc.Open,
		High:        fc.High,
		Low:         fc.Low,
		Close:       fc.Close,
		Volume:      strconv.FormatInt(fc.Volume, 10),
		Amount:      fc.Amount,
		WindowClose: fc.WindowClose,
	}
}

type WsSpotOrder struct {
	Id                 string `json:"id"`                   // 订单 ID
	User               int64  `json:"user"`                 // 用户 ID
	Text               string `json:"text"`                 // 用户自定义订单信息
	CreateTime         string `json:"create_time"`          // 订单创建时间，精确到秒
	CreateTimeMs       string `json:"create_time_ms"`       // 订单创建时间，精确到毫秒
	UpdateTime         string `json:"update_time"`          // 订单最新更新时间，精确到秒
	UpdateTimeMs       string `json:"update_time_ms"`       // 订单最新更新时间，精确到毫秒
	Event              string `json:"event"`                // 订单事件
	CurrencyPair       string `json:"currency_pair"`        // 交易货币对
	Type               string `json:"type"`                 // 订单类型
	Account            string `json:"account"`              // 账户类型
	Side               string `json:"side"`                 // 买单或者卖单
	Amount             string `json:"amount"`               // 交易数量
	Price              string `json:"price"`                // 交易价
	TimeInForce        string `json:"time_in_force"`        // Time in force 策略
	Left               string `json:"left"`                 // 交易货币未成交数量
	FilledTotal        string `json:"filled_total"`         // 已成交总金额
	FilledAmount       string `json:"filled_amount"`        // 货币成交数量
	AvgDealPrice       string `json:"avg_deal_price"`       // 订单成交均价
	Fee                string `json:"fee"`                  // 成交扣除的手续费
	FeeCurrency        string `json:"fee_currency"`         // 手续费计价单位
	PointFee           string `json:"point_fee"`            // 手续费抵扣使用的点卡数量
	GtFee              string `json:"gt_fee"`               // 手续费抵扣使用的 GT 数量
	GtDiscount         bool   `json:"gt_discount"`          // 是否开启 GT 抵扣
	RebatedFee         string `json:"rebated_fee"`          // 返还的手续费
	RebatedFeeCurrency string `json:"rebated_fee_currency"` // 返还手续费计价单位
	AutoRepay          bool   `json:"auto_repay"`           // 启用自动还款
	AutoBorrow         bool   `json:"auto_borrow"`          // 启用自动借款
	StpId              int64  `json:"stp_id"`               // stp_id
	StpAct             string `json:"stp_act"`              // stp_act
	FinishAs           string `json:"finish_as"`            // 订单的完成状态
	AmendText          string `json:"amend_text"`           // 用户在修改订单时添加的自定义数据
	BizInfo            string `json:"biz_info"`             // 业务信息（如有）
}

func handleWsData[T any](data []byte) (*WsSubscribeResult[T], error) {
	var res WsSubscribeResult[T]
	err := json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *WsCandles) HandleSubKey() string {
	return c.Interval
}

func (o *WsOrderBook) HandleSubKey() string {
	return o.Symbol
}

func (t *WsTrade) HandleSubKey() string {
	return t.Symbol
}
func convertToWsData[T any, R any](originData *WsSubscribeResult[T], targetResult *R) *WsSubscribeResult[R] {
	return &WsSubscribeResult[R]{
		Time:    originData.Time,
		TimeMs:  originData.TimeMs,
		ConnId:  originData.ConnId,
		Id:      originData.Id,
		Channel: originData.Channel,
		Event:   originData.Event,
		Error:   originData.Error,
		//Payload: originData.Payload,
		Result: targetResult,
	}
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
			//Payload: origin.Payload,
			Result: f(v),
		}
		result = append(result, r)
	}
	return result
}

//type WsSpotBookTicker struct {
//	Timestamp    int64  `json:"t"` //深度更新毫秒时间戳
//	UpdateId     int64  `json:"u"` //深度更新 ID
//	CurrencyPair string `json:"s"` //货币对
//	BidPrice     string `json:"b"` //最优买单价
//	BidAmount    string `json:"B"` //最优买单数量
//	AskPrice     string `json:"a"` //最优卖单价
//	AskAmount    string `json:"A"` //最优卖单数量
//}

type WsSpotOrderBookUpdate struct {
	Timestamp    int64      `json:"t"` //深度更新毫秒时间戳
	CurrencyPair string     `json:"s"` //货币对
	FirstId      int64      `json:"U"` //自动上次深度后的第一个深度 ID
	LastId       int64      `json:"u"` //自动上次深度后的最新深度 ID
	Bids         [][]string `json:"b"` //自上次更新以来的 bids 更新，按价格从高到低排序
	Asks         [][]string `json:"a"` //自上次更新以来的 asks 更新，按价格从地到高排序
}

type WsSpotOrderBook struct {
	Timestamp    int64      `json:"t"`            //深度更新毫秒时间戳
	LastUpdateId int64      `json:"lastUpdateId"` //深度副本当前的深度 ID
	CurrencyPair string     `json:"s"`            //货币对
	Bids         [][]string `json:"bids"`         //在当前快照的 顶层 bids（买单） 数据，按价格从高到低排序
	Asks         [][]string `json:"asks"`         //在当前快照的 顶层 asks（卖单） 数据，按价格从低到高排序
}

type WsFuturesOrderBook struct {
	Timestamp int64  `json:"t"`        //深度生成时间戳（以毫秒为单位）
	Contract  string `json:"contract"` //合约名称
	Id        int64  `json:"id"`       //深度 ID
	Asks      []struct {
		Price    string `json:"p"` //档位价格
		Quantity int    `json:"s"` //档位的数量
	} `json:"asks"` //深度卖方的档位列表
	Bids []struct {
		Price    string `json:"p"` //档位价格
		Quantity int    `json:"s"` //档位的数量
	} `json:"bids"` //深度买方的档位列表
}

type WsFuturesOrderBookUpdate struct {
	Timestamp int64  `json:"t"` //深度生成时间戳（以毫秒为单位）
	Contract  string `json:"s"` //合约名称
	FirstId   int64  `json:"U"` //本次更新开始的订单簿更新 ID
	LastId    int64  `json:"u"` //本次更新结束的订单簿更新 ID
	Bids      []struct {
		Price    string `json:"p"` //变更的档位价格
		Quantity int    `json:"s"` //档位的待成交数量。如果为 0，则从订单簿中删除该价格
	} `json:"b"` //买方变动列表
	Asks []struct {
		Price    string `json:"p"` //变更的档位价格
		Quantity int    `json:"s"` //档位的待成交数量。如果为 0，则从订单簿中删除该价格
	} `json:"a"` //卖方变动列表
}

func (o *WsSpotOrderBookUpdate) convertToWsOrderBook() *WsOrderBook {
	var bids, asks []PriceLevel
	for _, v := range o.Bids {
		bids = append(bids, PriceLevel{
			Price:    v[0],
			Quantity: v[1],
		})
	}
	for _, v := range o.Asks {
		asks = append(asks, PriceLevel{
			Price:    v[0],
			Quantity: v[1],
		})
	}
	return &WsOrderBook{
		Timestamp: o.Timestamp,
		Symbol:    o.CurrencyPair,
		Id:        o.LastId,
		FirstId:   o.FirstId,
		LastId:    o.LastId,
		Bids:      bids,
		Asks:      asks,
	}
}
func (o *WsSpotOrderBook) convertToWsOrderBook() *WsOrderBook {
	var bids, asks []PriceLevel
	for _, v := range o.Bids {
		bids = append(bids, PriceLevel{
			Price:    v[0],
			Quantity: v[1],
		})
	}
	for _, v := range o.Asks {
		asks = append(asks, PriceLevel{
			Price:    v[0],
			Quantity: v[1],
		})
	}

	return &WsOrderBook{
		Timestamp: o.Timestamp,
		Symbol:    o.CurrencyPair,
		Id:        o.LastUpdateId,
		FirstId:   o.LastUpdateId,
		LastId:    o.LastUpdateId,
		Bids:      bids,
		Asks:      asks,
	}
}
func (o *WsFuturesOrderBookUpdate) convertToWsOrderBook() *WsOrderBook {
	var bids, asks []PriceLevel
	for _, v := range o.Bids {
		bids = append(bids, PriceLevel{
			Price:    v.Price,
			Quantity: strconv.Itoa(v.Quantity),
		})
	}
	for _, v := range o.Asks {
		asks = append(asks, PriceLevel{
			Price:    v.Price,
			Quantity: strconv.Itoa(v.Quantity),
		})
	}

	return &WsOrderBook{
		Timestamp: o.Timestamp,
		Symbol:    o.Contract,
		Id:        o.LastId,
		FirstId:   o.FirstId,
		LastId:    o.LastId,
		Bids:      bids,
		Asks:      asks,
	}
}
func (o *WsFuturesOrderBook) convertToWsOrderBook() *WsOrderBook {
	var bids, asks []PriceLevel
	for _, v := range o.Bids {
		bids = append(bids, PriceLevel{
			Price:    v.Price,
			Quantity: strconv.Itoa(v.Quantity),
		})
	}
	for _, v := range o.Asks {
		asks = append(asks, PriceLevel{
			Price:    v.Price,
			Quantity: strconv.Itoa(v.Quantity),
		})
	}

	return &WsOrderBook{
		Timestamp: o.Timestamp,
		Symbol:    o.Contract,
		Id:        o.Id,
		FirstId:   o.Id,
		LastId:    o.Id,
		Bids:      bids,
		Asks:      asks,
	}
}

type WsSpotTrade struct {
	Id           int64  `json:"id"`             //成交 ID
	CreateTime   int64  `json:"create_time"`    //成交时间，精确到秒
	CreateTimeMs string `json:"create_time_ms"` //成交时间，毫秒精度
	Side         string `json:"side"`           //买单或者卖单
	CurrencyPair string `json:"currency_pair"`  //交易货币对
	Amount       string `json:"amount"`         //交易数量
	Price        string `json:"price"`          //交易价
	Range        string `json:"range"`          //成交范围(格式: "开始 ID-结束 ID")
}
type WsFuturesTrade struct {
	Contract     string `json:"contract"`       //合约名称
	Size         int64  `json:"size"`           //交易数量
	Id           int64  `json:"id"`             //交易 ID
	CreateTime   int64  `json:"create_time"`    //交易消息创建时间
	CreateTimeMs int64  `json:"create_time_ms"` //交易消息创建时间（以毫秒为单位）
	Price        string `json:"price"`          //交易价格
	IsInternal   bool   `json:"is_internal"`    //是否为内部成交。内部成交是指保险资金和 ADL 用户对强平指令的接管。由于不是市场深度上的正常撮合，交易价格可能会出现偏差，不会记录在 K 线上。如果不是内部交易，则该字段不会被返回。
}

func (t *WsSpotTrade) convertToWsTrade() *WsTrade {
	ms, _ := decimal.NewFromString(t.CreateTimeMs)
	return &WsTrade{
		Id:           t.Id,
		CreateTime:   t.CreateTime,
		CreateTimeMs: ms.IntPart(),
		Side:         t.Side,
		Symbol:       t.CurrencyPair,
		Amount:       t.Amount,
		Price:        t.Price,
		Range:        t.Range,
	}
}
func (t *WsFuturesTrade) convertToWsTrade() *WsTrade {
	size := decimal.NewFromInt(t.Size)
	Side := ""
	if size.GreaterThan(decimal.Zero) {
		Side = "buy"
	} else {
		Side = "sell"
	}
	return &WsTrade{
		Id:           t.Id,
		CreateTime:   t.CreateTime,
		CreateTimeMs: t.CreateTimeMs,
		Side:         Side,
		Symbol:       t.Contract,
		Amount:       size.Abs().String(),
		Price:        t.Price,
		Range:        "",
		IsInternal:   false,
	}

}

type WsSpotCrossBalance struct {
	Timestamp    int    `json:"timestamp"`     //余额更新时间戳戳，精确到秒
	TimestampMs  int    `json:"timestamp_ms"`  //余额更新时间戳戳，精确到毫秒
	User         string `json:"user"`          //用户 ID
	Currency     string `json:"currency"`      //变更的货币
	Change       string `json:"change"`        //变更数量
	Total        string `json:"total"`         //总现货余额
	Available    string `json:"available"`     //可用现货余额
	Freeze       string `json:"freeze"`        //冻结余额数量
	FreezeChange string `json:"freeze_change"` //余额锁定的变动金额
	ChangeType   string `json:"change_type"`   //余额变动类型 withdraw deposit trade-fee-deduct order-create: 订单创建 order-match: 订单成交更新 order-update: 取消订单或修改订单 margin-transfer future-transfer cross-margin-transfer other
}

type WsSpotTicker struct {
	CurrencyPair     string `json:"currency_pair"`     //货币对
	Last             string `json:"last"`              //最新成交价
	LowestAsk        string `json:"lowest_ask"`        //卖方最低价
	HighestBid       string `json:"highest_bid"`       //买方最高价
	ChangePercentage string `json:"change_percentage"` //涨跌百分比，跌用负数标识，如 -7.45
	BaseVolume       string `json:"base_volume"`       //交易货币成交量
	QuoteVolume      string `json:"quote_volume"`      //计价货币成交量
	High24h          string `json:"high_24h"`          //24 小时最高价
	Low24h           string `json:"low_24h"`           //24 小时最低价
}

type WsSpotBalance struct {
	Timestamp    int    `json:"timestamp"`     //余额更新时间戳戳，精确到秒
	TimestampMs  int    `json:"timestamp_ms"`  //余额更新时间戳戳，精确到毫秒
	User         string `json:"user"`          //用户 ID
	Currency     string `json:"currency"`      //变更的货币
	Change       string `json:"change"`        //变更数量
	Total        string `json:"total"`         //总现货余额
	Available    string `json:"available"`     //可用现货余额
	Freeze       string `json:"freeze"`        //冻结余额数量
	FreezeChange string `json:"freeze_change"` //余额锁定的变动金额
	ChangeType   string `json:"change_type"`   //余额变动类型
}

type WsFuturesTicker struct {
	Contract              string `json:"contract"`                //合约名称
	Last                  string `json:"last"`                    //最新成交价
	ChangePercentage      string `json:"change_percentage"`       //涨跌幅
	FundingRate           string `json:"funding_rate"`            //资金费率
	FundingRateIndicative string `json:"funding_rate_indicative"` //下一期参考资金费率
	MarkPrice             string `json:"mark_price"`              //标记价格
	IndexPrice            string `json:"index_price"`             //指数价格
	TotalSize             string `json:"total_size"`              //总数量
	Volume24h             string `json:"volume_24h"`              //24 小时成交量
	QuantoBaseRate        string `json:"quanto_base_rate"`        //双币种合约中基础货币与结算货币的汇率。不存在于其他类型的合同中
	Volume24hBtc          string `json:"volume_24h_btc"`          //近 24 小时 BTC 交易量（已弃用，请使用volume_24h_base、volume_24h_quote、volume_24h_settle代替）
	Volume24hUsd          string `json:"volume_24h_usd"`          //近 24 小时美元交易量（已弃用，请使用volume_24h_base、volume_24h_quote、volume_24h_settle 代替）
	Volume24hQuote        string `json:"volume_24h_quote"`        //近 24 小时交易量，以计价货币计
	Volume24hSettle       string `json:"volume_24h_settle"`       //近 24 小时交易量，以结算货币计
	Volume24hBase         string `json:"volume_24h_base"`         //近 24 小时交易量，以基础货币计
	Low24h                string `json:"low_24h"`                 //近 24 小时最低交易价
	High24h               string `json:"high_24h"`                //近 24 小时最高交易价
}

type WsFuturesOrder struct {
	CreateTime      int64   `json:"create_time"`       // 订单创建时间（已弃用）
	CreateTimeMs    int64   `json:"create_time_ms"`    // 订单创建时间戳（以毫秒为单位）
	FillPrice       float64 `json:"fill_price"`        // 订单成交价格 Float
	FinishAs        string  `json:"finish_as"`         // 订单是如何完成的 String
	Iceberg         int64   `json:"iceberg"`           // 冰山下单显示的数量 Integer
	Id              int64   `json:"id"`                // 订单 ID Integer
	IsClose         bool    `json:"is_close"`          // 是否为 close position Boolean
	IsLiq           bool    `json:"is_liq"`            // 是否为 liquidation Boolean
	Left            int64   `json:"left"`              // 剩余可交易数量 Integer
	Mkfr            float64 `json:"mkfr"`              // Maker 费用 Float
	IsReduceOnly    bool    `json:"is_reduce_only"`    // 是否为 reduce-only Boolean
	Status          string  `json:"status"`            // 订单状态 String
	Tkfr            float64 `json:"tkfr"`              // taker 费用 Float
	Price           float64 `json:"price"`             // 订单价格 Float
	Refu            int64   `json:"refu"`              // 推荐用户 ID Integer
	Refr            float64 `json:"refr"`              // 推荐用户费用 Float
	Size            int64   `json:"size"`              // 订单大小 Integer
	Text            string  `json:"text"`              // 用户定义的信息 String
	Tif             string  `json:"tif"`               // 有效时间 String
	FinishTime      int64   `json:"finish_time"`       // 订单结束 unix 时间戳（以秒为单位） Integer
	FinishTimeMs    int64   `json:"finish_time_ms"`    // 订单结束 unix 时间戳（以毫秒为单位） Integer
	User            string  `json:"user"`              // 用户 ID String
	Contract        string  `json:"contract"`          // 合约名称 String
	StpId           string  `json:"stp_id"`            // stp_id String
	StpAct          string  `json:"stp_act"`           // stp_act String
	AmendText       string  `json:"amend_text"`        // 用户修改订单时备注的自定义数据 String
	UpdateId        int64   `json:"update_id"`         // 更新id Integer
	UpdateTime      int64   `json:"update_time"`       // 更新时间 (毫秒时间戳) Integer
	BizInfo         string  `json:"biz_info"`          // 业务信息 String
	StopProfitPrice string  `json:"stop_profit_price"` // 止盈价格
	StopLossPrice   string  `json:"stop_loss_price"`   // 止损价格
}

type WsFuturesPosition struct {
	Contract        string  `json:"contract"`         //合约名称
	EntryPrice      float64 `json:"entry_price"`      //开仓价格
	HistoryPnl      float64 `json:"history_pnl"`      //已平仓的仓位总盈亏
	HistoryPoint    float64 `json:"history_point"`    //已平仓的点卡总盈亏
	LastClosePnl    float64 `json:"last_close_pnl"`   //最近一次平仓的盈亏
	Leverage        int     `json:"leverage"`         //杠杆倍数，0代表全仓，正数代表逐仓
	LeverageMax     int     `json:"leverage_max"`     //当前风险限额下，允许的最大杠杆倍数
	LiqPrice        float64 `json:"liq_price"`        //爆仓价格
	MaintenanceRate float64 `json:"maintenance_rate"` //当前风险限额下，维持保证金比例
	Margin          float64 `json:"margin"`           //保证金
	Mode            string  `json:"mode"`             //模式
	RealisedPnl     float64 `json:"realised_pnl"`     //已实现盈亏
	RealisedPoint   float64 `json:"realised_point"`   //点卡已实现盈亏
	RiskLimit       int     `json:"risk_limit"`       //风险限额
	Size            int     `json:"size"`             //合约 size
	Time            int     `json:"time"`             //更新 unix 时间戳
	TimeMs          int     `json:"time_ms"`          //更新 unix 时间戳（以毫秒为单位）
	User            string  `json:"user"`             //用户 ID
	UpdateId        int     `json:"update_id"`        //消息序列号，每次推送 order 之后会自增 1
}

type WsFuturesBalance struct {
	Balance  float64 `json:"balance"`  //余额最终数量
	Change   float64 `json:"change"`   //余额变化数量
	Text     string  `json:"text"`     //附带信息
	Time     int     `json:"time"`     //时间
	TimeMs   int64   `json:"time_ms"`  //时间（以毫秒为单位）
	Type     string  `json:"type"`     //类型
	User     string  `json:"user"`     //用户 ID
	Currency string  `json:"currency"` //币种
}
