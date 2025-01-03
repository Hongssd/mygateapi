package mygateapi

type DeliveryContractCommon struct {
	Name                string  `json:"name,omitempty"`                  //合约标识
	Underlying          string  `json:"underlying,omitempty"`            //标的物
	Cycle               string  `json:"cycle,omitempty"`                 //周期类型, 季度合约, 周合约等
	Type                string  `json:"type,omitempty"`                  //合约类型, inverse - 反向合约, direct - 正向合约
	QuantoMultiplier    string  `json:"quanto_multiplier,omitempty"`     //计价货币兑换为结算货币的乘数
	LeverageMin         string  `json:"leverage_min,omitempty"`          //最小杠杆
	LeverageMax         string  `json:"leverage_max,omitempty"`          //最大杠杆
	MaintenanceRate     string  `json:"maintenance_rate,omitempty"`      //维持保证金比例
	MarkType            string  `json:"mark_type,omitempty"`             //价格标记方式, internal - 内盘成交价格, index - 外部指数价格
	MarkPrice           string  `json:"mark_price,omitempty"`            //当前标记价格
	IndexPrice          string  `json:"index_price,omitempty"`           //当前指数价格
	LastPrice           string  `json:"last_price,omitempty"`            //上一次成交价格
	MakerFeeRate        string  `json:"maker_fee_rate,omitempty"`        //挂单成交的手续费率，负数代表返还后续费
	TakerFeeRate        string  `json:"taker_fee_rate,omitempty"`        //吃单成交的手续费率
	OrderPriceRound     string  `json:"order_price_round,omitempty"`     //委托价格最小单位
	MarkPriceRound      string  `json:"mark_price_round,omitempty"`      //标记、强平等价格最小单位
	BasisRate           string  `json:"basis_rate,omitempty"`            //当前合理基差率
	BasisValue          string  `json:"basis_value,omitempty"`           //当前合理基差值
	BasisImpactValue    string  `json:"basis_impact_value,omitempty"`    //计算合理基差率时加权深度影响额
	SettlePrice         string  `json:"settle_price,omitempty"`          //预计结算价格
	SettlePriceInterval int     `json:"settle_price_interval,omitempty"` //结算价格更新间隔
	SettlePriceDuration int     `json:"settle_price_duration,omitempty"` //加权平均计算结算价格时长, 单位秒
	ExpireTime          int64   `json:"expire_time,omitempty"`           //合约到期时间戳
	RiskLimitBase       string  `json:"risk_limit_base,omitempty"`       //基础风险限额
	RiskLimitStep       string  `json:"risk_limit_step,omitempty"`       //风险限额调整步长
	RiskLimitMax        string  `json:"risk_limit_max,omitempty"`        //合约允许的最大风险限额
	OrderSizeMin        int64   `json:"order_size_min,omitempty"`        //最小下单数量
	OrderSizeMax        int64   `json:"order_size_max,omitempty"`        //最大下单数量
	OrderPriceDeviate   string  `json:"order_price_deviate,omitempty"`   //下单价与当前标记价格允许的正负偏移量， 即下单价 order_price 需满足如下条件:
	RefDiscountRate     string  `json:"ref_discount_rate,omitempty"`     //被推荐人享受交易费率折扣
	RefRebateRate       string  `json:"ref_rebate_rate,omitempty"`       //推荐人享受交易费率返佣比例
	OrderbookID         int64   `json:"orderbook_id,omitempty"`          //orderbook更新ID
	TradeID             int64   `json:"trade_id,omitempty"`              //当前成交ID
	TradeSize           int64   `json:"trade_size,omitempty"`            //历史累计成交
	PositionSize        int64   `json:"position_size,omitempty"`         //当前做多用户持有仓位总和
	ConfigChangeTime    float64 `json:"config_change_time,omitempty"`    //配置最后更新时间
	InDelisting         bool    `json:"in_delisting,omitempty"`          //合约下线中
	OrdersLimit         int     `json:"orders_limit,omitempty"`          //最多挂单数量
}
type PublicRestDeliverySettleContractsRes []DeliveryContractCommon
type PublicRestDeliverySettleContractsContractRes DeliveryContractCommon

type PublicRestDeliverySettleOrderBookRes struct {
	ID      int64   `json:"id,omitempty"`      //深度更新 ID，深度每发生一次变化，该 ID 加 1，只有设置 with_id=true 时才返回
	Current float64 `json:"current,omitempty"` //接口数据返回时间戳
	Update  float64 `json:"update,omitempty"`  //深度变化时间戳
	Asks    []struct {
		Price    string `json:"p,omitempty"` //价格 (计价货币)
		Quantity int64  `json:"s,omitempty"` //数量
	} `json:"asks,omitempty"` //卖方深度列表
	Bids []struct {
		Price    string `json:"p,omitempty"` //价格 (计价货币)
		Quantity int64  `json:"s,omitempty"` //数量
	}
}

type PublicRestDeliverySettleTradesRes []struct {
	ID           int64   `json:"id,omitempty"`             //成交记录 ID
	CreateTime   float64 `json:"create_time,omitempty"`    //成交时间
	CreateTimeMs float64 `json:"create_time_ms,omitempty"` //成交时间，保留 3 位小数的毫秒精度
	Contract     string  `json:"contract,omitempty"`       //合约标识
	Size         int64   `json:"size,omitempty"`           //成交数量
	Price        string  `json:"price,omitempty"`          //成交价格 (计价货币)
	IsInternal   bool    `json:"is_internal,omitempty"`    //是否为内部成交
}

type PublicRestDeliverySettleCandlesticksRes []struct {
	Timestamp int64  `json:"t,omitempty"` //秒 s 精度的 Unix 时间戳
	Volume    int64  `json:"v,omitempty"` //交易量，只有市场行情的 K 线数据里有该值 (合约张数)
	Close     string `json:"c,omitempty"` //收盘价 (计价货币)
	High      string `json:"h,omitempty"` //最高价 (计价货币)
	Low       string `json:"l,omitempty"` //最低价 (计价货币)
	Open      string `json:"o,omitempty"` //开盘价 (计价货币)
}

type PublicRestDeliverySettleTickersResRow struct {
	Contract              string `json:"contract,omitempty"`                //合约标识
	Last                  string `json:"last,omitempty"`                    //最新成交价
	ChangePercentage      string `json:"change_percentage,omitempty"`       //涨跌百分比，跌用负数标识，如 -7.45
	TotalSize             string `json:"total_size,omitempty"`              //当前合约总持仓量
	Low24h                string `json:"low_24h,omitempty"`                 //最近24小时最低价
	High24h               string `json:"high_24h,omitempty"`                //最近24小时最高价
	Volume24h             string `json:"volume_24h,omitempty"`              //最近24小时成交总量
	Volume24hBtc          string `json:"volume_24h_btc,omitempty"`          //最近24小时成交总量，BTC单位(即将废弃，建议使用 volume_24h_base, volume_24h_quote, volume_24h_settle)
	Volume24hUsd          string `json:"volume_24h_usd,omitempty"`          //最近24小时成交总量，USD单位(即将废弃，建议使用 volume_24h_base, volume_24h_quote, volume_24h_settle)
	Volume24hBase         string `json:"volume_24h_base,omitempty"`         //最近24小时成交量，以基础货币为单位
	Volume24hQuote        string `json:"volume_24h_quote,omitempty"`        //最近24小时成交量，以计价货币为单位
	Volume24hSettle       string `json:"volume_24h_settle,omitempty"`       //最近24小时成交量，以结算货币为单位
	MarkPrice             string `json:"mark_price,omitempty"`              //最近标记价格
	FundingRate           string `json:"funding_rate,omitempty"`            //资金费率
	FundingRateIndicative string `json:"funding_rate_indicative,omitempty"` //下一周期预测资金费率
	IndexPrice            string `json:"index_price,omitempty"`             //指数价格
	QuantoBaseRate        string `json:"quanto_base_rate,omitempty"`        //双币种合约中，基础货币和结算货币的汇率。其他类型合约中此字段不存在。
	BasisRate             string `json:"basis_rate,omitempty"`              //基差率
	BasisValue            string `json:"basis_value,omitempty"`             //基差数值
	LowestAsk             string `json:"lowest_ask,omitempty"`              //最新卖
	LowestSize            string `json:"lowest_size,omitempty"`             //最新卖方最低价的挂单量
	HighestBid            string `json:"highest_bid,omitempty"`             //最新买方最高价
	HighestSize           string `json:"highest_size,omitempty"`            //最新买方最高价的挂单量
}

type PublicRestDeliverySettleTickersRes []PublicRestDeliverySettleTickersResRow
