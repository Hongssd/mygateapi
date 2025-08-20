package mygateapi

type PublicRestSpotCurrenciesResRow struct {
	Currency         string `json:"currency"`          // string 币种名称
	Delisted         bool   `json:"delisted"`          // boolean 是否下架
	WithdrawDisabled bool   `json:"withdraw_disabled"` // boolean 是否暂停提现
	WithdrawDelayed  bool   `json:"withdraw_delayed"`  // boolean 提现是否存在延迟
	DepositDisabled  bool   `json:"deposit_disabled"`  // boolean 是否暂停充值
	TradeDisabled    bool   `json:"trade_disabled"`    // boolean 是否暂停交易
	FixedRate        string `json:"fixed_rate"`        // string 固定交易手续费率。仅限固定交易费率的币种，普通币种该字段无效
	Chain            string `json:"chain"`             // string 币对应的链
}
type PublicRestSpotCurrenciesRes []PublicRestSpotCurrenciesResRow
type PublicRestSpotCurrenciesCurrencyRes PublicRestSpotCurrenciesResRow

type PublicRestSpotCurrencyPairCommon struct {
	ID              string `json:"id"`               // string 交易对名称
	Base            string `json:"base"`             // string 交易对中的基础币种
	Quote           string `json:"quote"`            // string 交易对中的报价币种
	Fee             string `json:"fee"`              // string 交易手续费率
	MinBaseAmount   string `json:"min_base_amount"`  // string 最小基础币种下单量
	MinQuoteAmount  string `json:"min_quote_amount"` // string 最小报价币种下单量
	MaxBaseAmount   string `json:"max_base_amount"`  // string 最大基础币种下单量
	MaxQuoteAmount  string `json:"max_quote_amount"` // string 最大报价币种下单量
	AmountPrecision int    `json:"amount_precision"` // integer 交易对中基础币种的交易精度
	Precision       int    `json:"precision"`        // integer 交易对中报价币种的交易精度
	TradeStatus     string `json:"trade_status"`     // string 交易对状态
	SellStart       int    `json:"sell_start"`       // integer 卖单开放时间
	BuyStart        int    `json:"buy_start"`        // integer 买单开放时间
}

type PublicRestSpotCurrencyPairsAllRes []PublicRestSpotCurrencyPairCommon

type PublicRestSpotCurrencyPairsCurrencyPairRes PublicRestSpotCurrencyPairCommon

type PublicRestSpotTickerResRow struct {
	CurrencyPair     string `json:"currency_pair"`     // string 交易对
	Last             string `json:"last"`              // string 最新成交价
	LowestAsk        string `json:"lowest_ask"`        // string 最新卖方最低价
	LowestSize       string `json:"lowest_size"`       // string 最新卖方最低价数量；批量查询时不存在；单个查询时存在,如果没有数据时为空
	HighestBid       string `json:"highest_bid"`       // string 最新买方最高价
	HighestSize      string `json:"highest_size"`      // string 最新买方最高价数量；批量查询时不存在；单个查询时存在,如果没有数据时为空
	ChangePercentage string `json:"change_percentage"` // string 最近24h涨跌百分比，跌用负数标识，如 -7.45
	ChangeUtc0       string `json:"change_utc0"`       // string utc0时区，最近24h涨跌百分比，跌用负数标识，如 -7.45
	ChangeUtc8       string `json:"change_utc8"`       // string utc8时区，最近24h涨跌百分比，跌用负数标识，如 -7.45
	BaseVolume       string `json:"base_volume"`       // string 最近24h交易货币成交量
	QuoteVolume      string `json:"quote_volume"`      // string 最近24h计价货币成交量
	High24h          string `json:"high_24h"`          // string 24小时最高价
	Low24h           string `json:"low_24h"`           // string 24小时最低价
	EtfNetValue      string `json:"etf_net_value"`     // string ETF 净值
	EtfPreNetValue   string `json:"etf_pre_net_value"` // string|null ETF 前一再平衡点净值
	EtfPreTimestamp  int    `json:"etf_pre_timestamp"` // integer|null ETF 前一再平衡时间
	EtfLeverage      string `json:"etf_leverage"`      // string|null ETF 当前杠杆率
}

type PublicRestSpotTickersRes []PublicRestSpotTickerResRow

type PublicRestSpotOrderBookRes struct {
	ID      int64      `json:"id"`      // 深度更新ID。深度每发生一次变化，ID 就会更新一次。仅在 with_id 设置为 true 该值有效
	Current int64      `json:"current"` // 接口数据返回 ms 时间戳
	Update  int64      `json:"update"`  // 深度变化 ms 时间戳
	Asks    [][]string `json:"asks"`    // 卖方深度列表
	Bids    [][]string `json:"bids"`    // 买方深度列表
}

type PublicRestSpotTradesResRow struct {
	ID           string `json:"id"`             // 成交记录 ID
	CreateTime   string `json:"create_time"`    // 成交时间
	CreateTimeMs string `json:"create_time_ms"` // 成交时间，毫秒精度
	CurrencyPair string `json:"currency_pair"`  // 交易货币对
	Side         string `json:"side"`           // 买单或者卖单
	Role         string `json:"role"`           // 交易角色，公共接口无此字段返回
	Amount       string `json:"amount"`         // 交易数量
	Price        string `json:"price"`          // 交易价
	OrderID      string `json:"order_id"`       // 关联的订单 ID，公共接口无此字段返回
	Fee          string `json:"fee"`            // 成交扣除的手续费，公共接口无此字段返回
	FeeCurrency  string `json:"fee_currency"`   // 手续费计价单位，公共接口无此字段返回
	PointFee     string `json:"point_fee"`      // 手续费抵扣使用的点卡数量，公共接口无此字段返回
	GtFee        string `json:"gt_fee"`         // 手续费抵扣使用的 GT 数量，公共接口无此字段返回
	AmendText    string `json:"amend_text"`     // 用户修改订单时备注的信息
	SequenceID   string `json:"sequence_id"`    // 单市场连续成交ID
	Text         string `json:"text"`           // 订单的自定义信息，公共接口无此字段返回
}

type PublicRestSpotTradesRes []PublicRestSpotTradesResRow

type PublicRestSpotCandlesticksMiddle []PublicRestSpotCandlesticksMiddleRow
type PublicRestSpotCandlesticksMiddleRow [9]interface{}

func (middle *PublicRestSpotCandlesticksMiddle) ConvertToRes() *PublicRestSpotCandlesticksRes {
	resList := PublicRestSpotCandlesticksRes{}
	for _, v := range *middle {
		resList = append(resList, *v.ConvertToRes())
	}
	return &resList
}

func (middle *PublicRestSpotCandlesticksMiddleRow) ConvertToRes() *PublicRestSpotCandlesticksResRow {
	res := PublicRestSpotCandlesticksResRow{
		Ts:          middle[0].(string),
		VolCcyQuote: middle[1].(string),
		C:           middle[2].(string),
		H:           middle[3].(string),
		L:           middle[4].(string),
		O:           middle[5].(string),
		VolCcy:      middle[6].(string),
		Confirm:     middle[7].(string),
	}
	return &res
}

type PublicRestSpotCandlesticksResRow struct {
	Ts          string `json:"ts"`            // 秒(s)精度的 Unix 时间戳
	VolCcyQuote string `json:"vol_ccy_quote"` // 计价货币交易额
	O           string `json:"o"`             // 开盘价
	H           string `json:"h"`             // 最高价
	L           string `json:"l"`             // 最低价
	C           string `json:"c"`             // 收盘价
	VolCcy      string `json:"vol_ccy"`       // 基础货币交易量
	Confirm     string `json:"confirm"`       // 窗口是否关闭，true 代表此段K线蜡烛图数据结束，false 代表此段K线蜡烛图数据尚未结束
}

type PublicRestSpotCandlesticksRes []PublicRestSpotCandlesticksResRow

type PublicRestSpotTimeRes struct {
	ServerTime int64 `json:"server_time"` // 服务器时间
}
