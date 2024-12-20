package mygateapi

type ContractCommon struct {
	Name              string  `json:"name"`                //合约标识
	Type              string  `json:"type"`                //合约类型, inverse - 反向合约, direct - 正向合约
	QuantoMultiplier  string  `json:"quanto_multiplier"`   //计价货币兑换为结算货币的乘数
	LeverageMin       string  `json:"leverage_min"`        //最小杠杆
	LeverageMax       string  `json:"leverage_max"`        //最大杠杆
	MaintenanceRate   string  `json:"maintenance_rate"`    //维持保证金比例
	MarkType          string  `json:"mark_type"`           //价格标记方式, internal - 内盘成交价格, index - 外部指数价格
	MarkPrice         string  `json:"mark_price"`          //当前标记价格
	IndexPrice        string  `json:"index_price"`         //当前指数价格
	LastPrice         string  `json:"last_price"`          //上一次成交价格
	MakerFeeRate      string  `json:"maker_fee_rate"`      //挂单成交的手续费率，负数代表返还后续费
	TakerFeeRate      string  `json:"taker_fee_rate"`      //吃单成交的手续费率
	OrderPriceRound   string  `json:"order_price_round"`   //委托价格最小单位
	MarkPriceRound    string  `json:"mark_price_round"`    //标记、强平等价格最小单位
	FundingRate       string  `json:"funding_rate"`        //当前资金费率
	FundingInterval   int     `json:"funding_interval"`    //资金费率应用间隔，以秒为单位
	FundingNextApply  float64 `json:"funding_next_apply"`  //下次资金费率应用时间
	RiskLimitBase     string  `json:"risk_limit_base"`     //基础风险限额,已废弃
	RiskLimitStep     string  `json:"risk_limit_step"`     //风险限额调整步长,已废弃
	RiskLimitMax      string  `json:"risk_limit_max"`      //合约允许的最大风险限额,已废弃,建议使用/futures/{settle}/risk_limit_tiers来查询风险限额
	OrderSizeMin      int64   `json:"order_size_min"`      //最小下单数量
	OrderSizeMax      int64   `json:"order_size_max"`      //最大下单数量
	OrderPriceDeviate string  `json:"order_price_deviate"` //下单价与当前标记价格允许的正负偏移量， 即下单价 order_price 需满足如下条件:
	RefDiscountRate   string  `json:"ref_discount_rate"`   //被推荐人享受交易费率折扣
	RefRebateRate     string  `json:"ref_rebate_rate"`     //推荐人享受交易费率返佣比例
	OrderbookId       int64   `json:"orderbook_id"`        //orderbook更新ID
	TradeId           int64   `json:"trade_id"`            //当前成交ID
	TradeSize         int64   `json:"trade_size"`          //历史累计成交
	PositionSize      int64   `json:"position_size"`       //当前做多用户持有仓位总和
	ConfigChangeTime  float64 `json:"config_change_time"`  //配置最后更新时间
	InDelisting       bool    `json:"in_delisting"`        //in_delisting=true 并且position_size>0时候 表示该合约处于下线过渡期 `in_delisting=true` 并且position_size=0时候 表示该合约处于下线状态
	OrdersLimit       int     `json:"orders_limit"`        //最多挂单数量
	EnableBonus       bool    `json:"enable_bonus"`        //是否支持体验金
	EnableCredit      bool    `json:"enable_credit"`       //是否支持统一账户
	CreateTime        float64 `json:"create_time"`         //表示合约的创建时间
	FundingCapRatio   string  `json:"funding_cap_ratio"`   //资金费率上限的系数。资金费率上限 = (1/市场最大杠杆 - 维持保证金率) * funding_cap_ratio
}

type PublicRestFuturesSettleContractsRes []ContractCommon
type PublicRestFuturesSettleContractsContractRes ContractCommon

type PublicRestFuturesSettleOrderBookRes struct {
	Id      int64   `json:"id"`      //深度更新 ID，深度每发生一次变化，该 ID 加 1，只有设置 with_id=true 时才返回
	Current float64 `json:"current"` //接口数据返回时间戳
	Update  float64 `json:"update"`  //深度变化时间戳
	Asks    []struct {
		P string `json:"p"` //价格 (计价货币)
		S int64  `json:"s"` //数量
	} `json:"asks"` //卖方深度列表
	Bids []struct {
		P string `json:"p"` //价格 (计价货币)
		S int64  `json:"s"` //数量
	} `json:"bids"` //买方深度列表
}

type PublicRestFuturesSettleTradesResRow struct {
	Id           int64   `json:"id"`             //成交记录 ID
	CreateTime   float64 `json:"create_time"`    //成交时间
	CreateTimeMs float64 `json:"create_time_ms"` //成交时间，保留 3 位小数的毫秒精度
	Contract     string  `json:"contract"`       //合约标识
	Size         int64   `json:"size"`           //成交数量
	Price        string  `json:"price"`          //成交价格 (计价货币)
	IsInternal   bool    `json:"is_internal"`    //是否为内部成交。内部成交是指保险基金和ADL用户对强平委托的接盘，由于不是在市场深度上的正常撮合���所以成交价格和市场有偏差，也不会计入K线。如果不是内部成交，则不返回这个字段。
}

type PublicRestFuturesSettleTradesRes []PublicRestFuturesSettleTradesResRow

type PublicRestFuturesSettleCandlesticksResRow struct {
	Ts          int64  `json:"t"`   //秒(s)精度的 Unix 时间戳
	Vol         int64  `json:"v"`   //交易量，只有市场行情的 K 线数据里有该值 (合约张数)
	C           string `json:"c"`   //收盘价 (计价货币)
	H           string `json:"h"`   //最高价 (计价货币)
	L           string `json:"l"`   //最低价 (计价货币)
	O           string `json:"o"`   //开盘价 (计价货币)
	VolCcyQuote string `json:"sum"` //交易额，单位是计价货币
}

type PublicRestFuturesSettleCandlesticksRes []PublicRestFuturesSettleCandlesticksResRow
