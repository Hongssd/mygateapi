package mygateapi

type GatePositionResCommon struct {
	User              int64  `json:"user"`               // 用户ID
	Contract          string `json:"contract"`           // 合约标识
	Size              int64  `json:"size"`               // 头寸大小
	Leverage          string `json:"leverage"`           // 杠杆倍数，0代表全仓，正数代表逐仓
	RiskLimit         string `json:"risk_limit"`         // 风险限额
	LeverageMax       string `json:"leverage_max"`       // 当前风险限额下，允许的最大杠杆倍数
	MaintenanceRate   string `json:"maintenance_rate"`   // 当前风险限额下，维持保证金比例
	Value             string `json:"value"`              // 按结算币种标记价格计算的合约价值
	Margin            string `json:"margin"`             // 保证金
	EntryPrice        string `json:"entry_price"`        // 开仓价格
	LiqPrice          string `json:"liq_price"`          // 爆仓价格
	MarkPrice         string `json:"mark_price"`         // 合约当前标记价格
	InitialMargin     string `json:"initial_margin"`     // 仓位占用的起始保证金，适用于统一账户
	MaintenanceMargin string `json:"maintenance_margin"` // 仓位所需的维持保证金，适用于统一账户
	UnrealisedPnl     string `json:"unrealised_pnl"`     // 未实现盈亏
	RealisedPnl       string `json:"realised_pnl"`       // 已实现盈亏
	PnlPnl            string `json:"pnl_pnl"`            // 已实现盈亏-仓位盈亏
	PnlFund           string `json:"pnl_fund"`           // 已实现盈亏-资金费用
	PnlFee            string `json:"pnl_fee"`            // 已实现盈亏-手续费
	HistoryPnl        string `json:"history_pnl"`        // 已平仓的仓位总盈亏
	LastClosePnl      string `json:"last_close_pnl"`     // 最近一次平仓的盈亏
	RealisedPoint     string `json:"realised_point"`     // 点卡已实现盈亏
	HistoryPoint      string `json:"history_point"`      // 已平仓的点卡总盈亏
	AdlRanking        int    `json:"adl_ranking"`        // 自动减仓排名，共1-5个等级，1 最高，5 最低，特殊情况 6 是没有持仓或在爆仓中
	PendingOrders     int    `json:"pending_orders"`     // 当前未完成委托数量
	CloseOrder        struct {
		Id    int64  `json:"id"`     // 委托ID
		Price string `json:"price"`  // 委托价格
		IsLiq bool   `json:"is_liq"` // 是否为强制平仓
	} `json:"close_order"` // 当前平仓委托信息，如果没有平仓则为null
	Mode               string `json:"mode"`                 // 持仓模式
	CrossLeverageLimit string `json:"cross_leverage_limit"` // 全仓模式下的杠杆倍数（即 leverage 为 0 时）
	UpdateTime         int64  `json:"update_time"`          // 最后更新时间
	UpdateId           int64  `json:"update_id"`            // 更新id，仓位每更新一次，数值会+1
	OpenTime           int64  `json:"open_time"`            // 开仓时间
}
type GateFuturesOrderResCommon struct {
	Id           int64   `json:"id"`             // 合约订单 ID
	User         int64   `json:"user"`           // 用户 ID
	CreateTime   float64 `json:"create_time"`    // 订单创建时间
	FinishTime   float64 `json:"finish_time"`    // 订单结束时间，未结束订单无此字段返回
	FinishAs     string  `json:"finish_as"`      // 结束方式，包括：filled: 完全成交 cancelled: 用户撤销 liquidated: 强制平仓撤销 ioc: 未立即完全成交，因为tif设置为ioc auto_deleveraged: 自动减仓撤销 reduce_only: 增持仓位撤销，因为设置reduce_only或平仓 position_closed: 因为仓位平掉了，所以挂单被撤掉 reduce_out: 只减仓被排除的不容易成交的挂单 stp: 订单发生自成交限制而被撤销
	Status       string  `json:"status"`         // 订单状态。open: 等待处理 finished: 已结束的订单
	Contract     string  `json:"contract"`       // 合约标识
	Size         int64   `json:"size"`           // 交易数量，正数为买入，负数为卖出。平仓委托则设置为0。
	Iceberg      int64   `json:"iceberg"`        // 冰山委托显示数量。0为完全不隐藏。注意，隐藏部分成交按照taker收取手续费。
	Price        string  `json:"price"`          // 委托价。价格为0并且tif为ioc，代表市价委托。
	Close        bool    `json:"close"`          // 设置为 true 的时候执行平仓操作，并且size应设置为0
	IsClose      bool    `json:"is_close"`       // 是否为平仓委托。对应请求中的close。
	ReduceOnly   bool    `json:"reduce_only"`    // 设置为 true 的时候，为只减仓委托
	IsReduceOnly bool    `json:"is_reduce_only"` // 是否为只减仓委托。对应请求中的reduce_only。
	IsLiq        bool    `json:"is_liq"`         // 是否为强制平仓委托
	Tif          string  `json:"tif"`            // Time in force 策略，市价单当前只支持 ioc 模式
	Left         int64   `json:"left"`           // 未成交数量
	FillPrice    string  `json:"fill_price"`     // 成交价
	Text         string  `json:"text"`           // 订单自定义信息，用户可以用该字段设置自定义 ID，用户自定义字段必须满足以下条件： 1. 必须以 t- 开头 2. 不计算 t- ，长度不能超��� 28 字节 3. 输入内容只能包含数字、字母、下划线(_)、中划线(-) 或者点(.) 除用户自定义信息以外，以下为内部保留字段，标识订单来源: - web: 网页 - api: API 调用 - app: 移动端 - auto_deleveraging: 自动减仓 - liquidation: 强制平仓 - insurance: 保险
	Tkfr         string  `json:"tkfr"`           // 吃单费率
	Mkfr         string  `json:"mkfr"`           // 做单费率
	Refu         int64   `json:"refu"`           // 推荐人用户 ID
	AutoSize     string  `json:"auto_size"`      // 双仓模式下用于设置平仓的方向，close_long 平多头， close_short 平空头，需要同时设置 size 为 0
	StpId        int64   `json:"stp_id"`         // 订单所属的STP用户组id，同一个STP用户组内用户之间的订单不允许发生自成交。 1. 如果撮合时两个订单的 stp_id 非 0 且相等，则不成交，而是根据 taker 的 stp_act 执行相应策略。 2. 没有设置STP用户组成交的订单，stp_id 默认返回 0。
	StpAct       string  `json:"stp_act"`        // Self-Trading Prevention Action,用户可以用该字段设置自定义限制自成交策略。 1. 用户在设置加入STP用户组后，可以通过传递 stp_act 来限制用户发生自成交的策略，没有传递 stp_act 默认按照 cn 的策略。 2. 用户在没有设置加入STP用户组时，传递 stp_act 参数会报错。 3. 用户没有使用 stp_act 发生成交的订单，stp_act 返回 -。 - cn: Cancel newest,取消新订单，保留老订单 - co: Cancel oldest,取消⽼订单，保留新订单 - cb: Cancel both,新旧订单都取消
	AmendText    string  `json:"amend_text"`     // 用户修改订单时备注的信息
	BizInfo      string  `json:"biz_info"`       // 附加信息
}

type GateFuturesPriceTriggeredOrderResCommon struct {
	Initial struct {
		Contract     string `json:"contract,omitempty"`       // 合约标识
		Size         int64  `json:"size,omitempty"`           // 交易数量，正数为买入，负数为卖出，平仓操作必须为0
		Price        string `json:"price,omitempty"`          // 交易价，当价格为 0 时，表示通过市价方式来下单
		Close        bool   `json:"close,omitempty"`          // 设置为 true 的时候执行平仓操作
		Tif          string `json:"tif,omitempty"`            // Time in force 策略，市价单当前只支持 ioc 模式
		Text         string `json:"text,omitempty"`           // 订单的来源，包括：web: 网页 api: API 调用 app: 移动端
		ReduceOnly   bool   `json:"reduce_only,omitempty"`    // 设置为 true 的时候执行自动减仓操作
		AutoSize     string `json:"auto_size,omitempty"`      // 双仓模式下用于设置平仓的方向，close_long 平多头， close_short 平空头，需要同时设置 size 为 0
		IsReduceOnly bool   `json:"is_reduce_only,omitempty"` // 是否为只减仓委托。对应请求中的reduce_only。
		IsClose      bool   `json:"is_close,omitempty"`       // 是否为平仓委托。对应请求中的close。
	} `json:"initial"` // 初始订单信息
	Trigger struct {
		StrategyType int    `json:"strategy_type,omitempty"` // 触发策略
		PriceType    int    `json:"price_type,omitempty"`    // 参考价格类型。 0 - 最新成交价，1 - 标记价格，2 - 指数价格
		Price        string `json:"price,omitempty"`         // 价格触发时为价格，价差触发时为价差
		Rule         int    `json:"rule,omitempty"`          // 价格条件类型
		Expiration   int    `json:"expiration,omitempty"`    // 最长等待触发时间，超时则取消该订单，单位是秒 s
	} `json:"trigger"` // 触发条件
	Id         int64   `json:"id,omitempty"`          // 自动订单 ID
	User       int64   `json:"user,omitempty"`        // 用户 ID
	CreateTime float64 `json:"create_time,omitempty"` // 创建时间
	FinishTime float64 `json:"finish_time,omitempty"` // 结束时间
	TradeId    int64   `json:"trade_id,omitempty"`    // 触发后委托单ID
	Status     string  `json:"status,omitempty"`      // 订单状态
	FinishAs   string  `json:"finish_as,omitempty"`   // 结束状态，cancelled - 被取消；succeeded - 成功；failed - 失败；expired - 过期
	Reason     string  `json:"reason,omitempty"`      // 订单结束的附加描述信息
	OrderType  string  `json:"order_type,omitempty"`  // 止盈止损的类型，包括：close-long-order: 委托单止盈止损，平做多仓 close-short-order: 委托单止盈止损，平做空仓 close-long-position: 仓位止盈止损，平多仓 close-short-position: 仓位止盈止损，平空仓 plan-close-long-position: 仓位计划止盈止损，平多仓 plan-close-short-position: 仓位计划止盈止损，平空仓 其中委托单止盈止损的两种类型只读，不能通过请求传入
	MeOrderId  int64   `json:"me_order_id,omitempty"` // 委托单止盈止损对应的委托 ID
}
