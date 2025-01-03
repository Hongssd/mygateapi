package mygateapi

type PrivateRestDeliverySettleAccountsRes struct {
	Total                  string `json:"total"`                    // 钱包余额是用户累计充值提现和盈亏结算(包括已实现盈亏, 资金费用,手续费及推荐返佣)之后的余额, 不包含未实现盈亏. total = SUM(history_dnw, history_pnl, history_fee, history_refr, history_fund)
	UnrealisedPnl          string `json:"unrealised_pnl"`           // 未实现盈亏
	PositionMargin         string `json:"position_margin"`          // 头寸保证金
	OrderMargin            string `json:"order_margin"`             // 未完成订单的保证金
	Available              string `json:"available"`                // 可用的转出或交易的额度，统一账户下包含授信额度的可用额度(有包含体验金,体验金无法转出,所以要转出,转出金额需要扣除体验金)
	Point                  string `json:"point"`                    // 点卡数额
	Currency               string `json:"currency"`                 // 结算币种
	InDualMode             bool   `json:"in_dual_mode"`             // 是否为双向持仓模式
	EnableCredit           bool   `json:"enable_credit"`            // 是否开启统一账户模式
	PositionInitialMargin  string `json:"position_initial_margin"`  // 头寸占用的起始保证金，适用于统一账户模式
	MaintenanceMargin      string `json:"maintenance_margin"`       // 头寸占用的维持保证金，适用于新经典账户保证金模式和统一账户模式
	Bonus                  string `json:"bonus"`                    // 体验金
	EnableEvolvedClassic   bool   `json:"enable_evolved_classic"`   // 经典账户保证金模式,true-新模式,false-老模式
	CrossOrderMargin       string `json:"cross_order_margin"`       // 全仓挂单保证金，适用于新经典账户保证金模式
	CrossInitialMargin     string `json:"cross_initial_margin"`     // 全仓初始保证金，适用于新经典账户保证金模式
	CrossMaintenanceMargin string `json:"cross_maintenance_margin"` // 全仓维持保证金，适用于新经典账户保证金模式
	CrossUnrealisedPnl     string `json:"cross_unrealised_pnl"`     // 全仓未实现盈亏，适用于新经典账户保证金模式
	CrossAvailable         string `json:"cross_available"`          // 全仓可用额度，适用于新经典账户保证金模式
	IsolatedPositionMargin string `json:"isolated_position_margin"` // 逐仓仓位保证金，适用于新经典账户保证金模式
	EnableNewDualMode      bool   `json:"enable_new_dual_mode"`     // 是否开启新的双向持仓模式
	MarginMode             int    `json:"margin_mode"`              // 保证金模式，0-经典保证金模式，1-跨币种保证金模式，2-组合保证金模式
	History                struct {
		Dnw         string `json:"dnw"`          // 累计转入转出
		Pnl         string `json:"pnl"`          // 累计交易盈亏
		Fee         string `json:"fee"`          // 累计手续费
		Refr        string `json:"refr"`         // 累计获取的推荐人返佣
		Fund        string `json:"fund"`         // 累计资金费用
		PointDnw    string `json:"point_dnw"`    // 累计点卡转入转出
		PointFee    string `json:"point_fee"`    // 累计点卡抵扣手续费
		PointRefr   string `json:"point_refr"`   // 累计获取的点卡推荐人返佣
		BonusDnw    string `json:"bonus_dnw"`    // 累计体验金转入转出
		BonusOffset string `json:"bonus_offset"` // 累计体验金抵扣
	} `json:"history"` // 累计统计数据
}

type PrivateRestDeliverySettleAccountBookRes []struct {
	Time     float64 `json:"time"`     // 时间
	Change   string  `json:"change"`   // 变更金额
	Balance  string  `json:"balance"`  // 变更后账户余额
	Type     string  `json:"type"`     // 变更类型
	Text     string  `json:"text"`     // 注释
	Contract string  `json:"contract"` // 合约标识，只有2023-10-30后的数据才有该字段
	TradeID  string  `json:"trade_id"` // 成交 id
	ID       string  `json:"id"`       // 账户变更记录 id
}

type PrivateRestDeliverySettlePositionsRes []GatePositionResCommon
type PrivateRestDeliverySettlePositionsContractRes GatePositionResCommon
type PrivateRestDeliverySettlePositionsContractMarginRes GatePositionResCommon
type PrivateRestDeliverySettlePositionsContractLeverageRes GatePositionResCommon
type PrivateRestDeliverySettlePositionsContractRiskLimitRes GatePositionResCommon

type PrivateRestDeliverySettleOrdersPostRes GateFuturesOrderResCommon
type PrivateRestDeliverySettleOrdersGetRes []GateFuturesOrderResCommon
type PrivateRestDeliverySettleOrdersDeleteRes []GateFuturesOrderResCommon
type PrivateRestDeliverySettleOrdersOrderIdGetRes GateFuturesOrderResCommon
type PrivateRestDeliverySettleOrdersOrderIdDeleteRes GateFuturesOrderResCommon

type PrivateRestDeliverySettleMyTradesResRow struct {
	ID         int64   `json:"id"`          // 成交记录 ID
	CreateTime float64 `json:"create_time"` // 成交时间
	Contract   string  `json:"contract"`    // 合约标识
	OrderID    string  `json:"order_id"`    // 成交记录关联订单 ID
	Size       int64   `json:"size"`        // 成交数量
	CloseSize  int64   `json:"close_size"`  // 平仓数量 close_size=0 && size＞0 开多 close_size=0 && size＜0 开空 close_size>0 && size>0 && size <= close_size 平空 close_size>0 && size>0 && size > close_size 平空且开多 close_size<0 && size<0 && size >= close_size 平多 close_size<0 && size<0 && size < close_size 平多且开空
	Price      string  `json:"price"`       // 成交价格
	Role       string  `json:"role"`        // 成交角色， taker - 吃单, maker - 做单
	Text       string  `json:"text"`        // 订单的自定义信息
	Fee        string  `json:"fee"`         // 成交手续费
	PointFee   string  `json:"point_fee"`   // 成交点卡手续费
}
type PrivateRestDeliverySettleMyTradesRes []PrivateRestDeliverySettleMyTradesResRow

type PrivateRestDeliverySettlePositionCloseResRow struct {
	Time          float64 `json:"time"`            // 平仓时间
	Contract      string  `json:"contract"`        // 合约标识
	Side          string  `json:"side"`            // 多空方向
	Pnl           string  `json:"pnl"`             // 盈亏
	PnlPnl        string  `json:"pnl_pnl"`         // 盈亏-仓位盈亏
	PnlFund       string  `json:"pnl_fund"`        // 盈亏-资金费用
	PnlFee        string  `json:"pnl_fee"`         // 盈亏-手续费
	Text          string  `json:"text"`            // 平仓委托的来源，具体取值参见order.text字段
	MaxSize       string  `json:"max_size"`        // 最大持仓量
	AccumSize     string  `json:"accum_size"`      // 累计平仓量
	FirstOpenTime int64   `json:"first_open_time"` // 开仓时间
	LongPrice     string  `json:"long_price"`      // side为long时表示开仓均价，为short时表示平仓均价
	ShortPrice    string  `json:"short_price"`     // side为long时表示平仓均价，为short时表示开仓均价
}

type PrivateRestDeliverySettlePositionCloseRes []PrivateRestDeliverySettlePositionCloseResRow

type PrivateRestDeliverySettleLiquidatesResRow struct {
	Time       int64  `json:"time"`        // 强制平仓时间
	Contract   string `json:"contract"`    // 合约标识
	Leverage   string `json:"leverage"`    // 杠杆倍数，公共接口无该字段返回
	Size       int64  `json:"size"`        // 仓位大小
	Margin     string `json:"margin"`      // 保证金，公共接口无该字段返回
	EntryPrice string `json:"entry_price"` // 平均开仓价，公共接口无该字段返回
	LiqPrice   string `json:"liq_price"`   // 强制平仓价，公共接口无该字段返回
	MarkPrice  string `json:"mark_price"`  // 市场标记价，公共接口无该字段返回
	OrderID    int64  `json:"order_id"`    // 强平委托ID，公共接口无该字段返回
	OrderPrice string `json:"order_price"` // 强平委托价
	FillPrice  string `json:"fill_price"`  // 强平委托吃单平均成交价
	Left       int64  `json:"left"`        // 强平委托挂单大小
}
type PrivateRestDeliverySettleLiquidatesRes []PrivateRestDeliverySettleLiquidatesResRow

type PrivateRestDeliverySettleSettlementsResRow struct {
	Time        int64  `json:"time"`         // 强制平仓时间
	Contract    string `json:"contract"`     // 合约标识
	Leverage    string `json:"leverage"`     // 杠杆倍数
	Size        int64  `json:"size"`         // 仓位大小
	Margin      string `json:"margin"`       // 保证金
	EntryPrice  string `json:"entry_price"`  // 平均开仓价
	SettlePrice string `json:"settle_price"` // 结算价
	Profit      string `json:"profit"`       // 盈利
	Fee         string `json:"fee"`          // 结算费
}
type PrivateRestDeliverySettleSettlementsRes []PrivateRestDeliverySettleSettlementsResRow

type PrivateRestDeliverySettlePriceOrdersPostRes struct {
	ID int64 `json:"id"` // 自动订单 ID
}

type PrivateRestDeliverySettlePriceOrdersGetRes []GateFuturesPriceTriggeredOrderResCommon
type PrivateRestDeliverySettlePriceOrdersDeleteRes []GateFuturesPriceTriggeredOrderResCommon
type PrivateRestDeliverySettlePriceOrdersOrderIdGetRes GateFuturesPriceTriggeredOrderResCommon
type PrivateRestDeliverySettlePriceOrdersOrderIdDeleteRes GateFuturesPriceTriggeredOrderResCommon
