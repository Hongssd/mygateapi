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

type WsSpotOrder struct {
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

type WsSpotBookTicker struct {
	Timestamp    int64  `json:"t"` //深度更新毫秒时间戳
	UpdateId     int64  `json:"u"` //深度更新 ID
	CurrencyPair string `json:"s"` //货币对
	BidPrice     string `json:"b"` //最优买单价
	BidAmount    string `json:"B"` //最优买单数量
	AskPrice     string `json:"a"` //最优卖单价
	AskAmount    string `json:"A"` //最优卖单数量
}

type WsSpotOrderBookUdpate struct {
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
	CreateTime   int    `json:"create_time"`    //订单创建时间（已弃用）
	CreateTimeMs int    `json:"create_time_ms"` //订单创建时间戳（以毫秒为单位）
	FillPrice    string `json:"fill_price"`     //订单成交价格
	FinishAs     string `json:"finish_as"`      //订单是如何完成的。 filled：全部成交 cancelled：手动取消 liquidated：因清算而取消 ioc：生效时间为 IOC，立即完成 auto_deleveraging：ADL 完成 reduce_only：因减仓设置而增仓而取消 position_close：因平仓而取消 stp：订单发生自成交限制而被撤销 _new：新建 _update：成交或部分成交或更新订单 reduce_out: 只减仓被排除的不容易成交的挂单
	Iceberg      int    `json:"iceberg"`        //冰山下单显示的数量，不指定或传 0 都默认为普通下单。目前不支持全部冰山。
	Id           int    `json:"id"`             //订单 ID
	IsClose      bool   `json:"is_close"`       //该订单是否为 close position
	IsLiq        bool   `json:"is_liq"`         //该订单是否为 liquidation
	Left         int    `json:"left"`           //剩余可交易数量
	Mkfr         string `json:"mkfr"`           //Maker 费用
	IsReduceOnly bool   `json:"is_reduce_only"` //该订单是否为 reduce-only
	Status       string `json:"status"`         //订单状态 open: 等待交易 finished: 完成
	Tkfr         string `json:"tkfr"`           //taker 费用
	Price        string `json:"price"`          //订单价格。 0 表示市价订单，tif 设置为 ioc
	Refu         int    `json:"refu"`           //推荐用户 ID
	Refr         string `json:"refr"`           //推荐用户费用
	Size         int    `json:"size"`           //订单大小。指定正数进行出价，指定负数进行询问
	Text         string `json:"text"`           //用户定义的信息
	Tif          string `json:"tif"`            //有效时间 gtc：GoodTillCancelled ioc：ImmediateOrCancelled，仅接受者 poc：PendingOrCancelled，只进行后订单，始终享受挂单费用 fok： FillOrKill，完全填充或不填充 type=market 时仅支持 ioc 和 fok
	FinishTime   int    `json:"finish_time"`    //订单更新 unix 时间戳（以秒为单位）
	FinishTimeMs int    `json:"finish_time_ms"` //订单更新 unix 时间戳（以毫秒为单位）
	User         string `json:"user"`           //用户 ID
	Contract     string `json:"contract"`       //合约名称
	StpId        string `json:"stp_id"`         //同一 stp_id 组内的用户之间的订单不允许自交易 1.如果匹配的两个订单的 stp_id 非零且相等，则不会被执行。而是根据 taker 的 stp_act 执行相应的策略。 2.对于未设置 STP 组的订单，stp_id 默认返回 0
	StpAct       string `json:"stp_act"`        //自我交易预防行动。用户可以通过该字段设置自我交易防范策略 1.用户加入 STP Group 后，可以通过 stp_act 来限制用户的自我交易防范策略。如果不传 stp_act，则默认为 cn 策略。 2.当用户没有加入 STP 组时，传递 stp_act 参数时会返回错误。 3.如果用户下单时没有使用'stp_act'，'stp_act'将返回'-' cn: 取消最新订单，取消新订单并保留旧订单 co: 取消最旧订单，取消旧订单并保留新订单 cb：取消两者，新旧订单都会被取消
	AmendText    string `json:"amend_text"`     //用户修改订单时备注的自定义数据
}
type WsFuturesTrade struct {
	Contract     string `json:"contract"`       //合约名称
	Size         int    `json:"size"`           //交易数量
	Id           int    `json:"id"`             //交易 ID
	CreateTime   int    `json:"create_time"`    //交易消息创建时间
	CreateTimeMs int    `json:"create_time_ms"` //交易消息创建时间（以毫秒为单位）
	Price        string `json:"price"`          //交易价格
	IsInternal   bool   `json:"is_internal"`    //是否为内部成交。内部成交是指保险资金和 ADL 用户对强平指令的接管。由于不是市场深度上的正常撮合，交易价格可能会出现偏差，不会记录在 K 线上。如果不是内部交易，则该字段不会被返回。
}

type WsFuturesPosition struct {
	Contract        string  `json:"contract"`         //合约名称
	EntryPrice      string  `json:"entry_price"`      //开仓价格
	HistoryPnl      string  `json:"history_pnl"`      //已平仓的仓位总盈亏
	HistoryPoint    string  `json:"history_point"`    //已平仓的点卡总盈亏
	LastClosePnl    string  `json:"last_close_pnl"`   //最近一次平仓的盈亏
	Leverage        int     `json:"leverage"`         //杠杆倍数，0代表全仓，正数代表逐仓
	LeverageMax     int     `json:"leverage_max"`     //当前风险限额下，允许的最大杠杆倍数
	LiqPrice        float64 `json:"liq_price"`        //爆仓价格
	MaintenanceRate float64 `json:"maintenance_rate"` //当前风险限额下，维持保证金比例
	Margin          string  `json:"margin"`           //保证金
	RealisedPnl     string  `json:"realised_pnl"`     //已实现盈亏
	RealisedPoint   string  `json:"realised_point"`   //点卡已实现盈亏
	RiskLimit       int     `json:"risk_limit"`       //风险限额
	Size            int     `json:"size"`             //合约 size
	Time            int     `json:"time"`             //更新 unix 时间戳
	TimeMs          int     `json:"time_ms"`          //更新 unix 时间戳（以毫秒为单位）
	User            string  `json:"user"`             //用户 ID
	UpdateId        int     `json:"update_id"`        //消息序列号，每次推送 order 之后会自增 1
}

type WsFuturesBalance struct {
	Balance  string `json:"balance"`  //余额最终数量
	Change   string `json:"change"`   //余额变化数量
	Text     string `json:"text"`     //附带信息
	Time     int    `json:"time"`     //时间
	TimeMs   int    `json:"time_ms"`  //时间（以毫秒为单位）
	Type     string `json:"type"`     //类型
	User     string `json:"user"`     //用户 ID
	Currency string `json:"currency"` //币种
}
