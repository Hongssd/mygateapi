package mygateapi

type PrivateRestFuturesSettleAccountsRes struct {
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
	MarginMode             int64  `json:"margin_mode"`              // 保证金模式，0-经典保证金模式，1-跨币种保证金模式，2-组合保证金模式
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

type PrivateRestFuturesSettleAccountBookResRow struct {
	Time     float64 `json:"time"`     // 时间
	Change   string  `json:"change"`   // 变更金额
	Balance  string  `json:"balance"`  // 变更后账户余额
	Type     string  `json:"type"`     // 变更类型 dnw: 转入转出 pnl: 减仓盈亏 fee: 交易手续费 refr: 推荐人返佣 fund: 资金费用 point_dnw: 点卡转入转出 point_fee: 点卡交易手续费 point_refr: 点卡推荐人返佣 bonus_offset: 体验金抵扣
	Text     string  `json:"text"`     // 注释
	Contract string  `json:"contract"` // 合约标识，只有2023-10-30后的数据才有该字段
	TradeId  string  `json:"trade_id"` // 成交 id
	Id       string  `json:"id"`       // 账户变更记录 id
}

type PrivateRestFuturesSettleAccountBookRes []PrivateRestFuturesSettleAccountBookResRow

type PrivateRestFuturesSettlePositionsRes []GatePositionResCommon
type PrivateRestFuturesSettlePositionsContractRes GatePositionResCommon

type PrivateRestFuturesSettlePositionsContractMarginRes GatePositionResCommon
type PrivateRestFuturesSettlePositionsContractLeverageRes []GatePositionResCommon
type PrivateRestFuturesSettlePositionsContractRiskLimitRes GatePositionResCommon

type PrivateRestFuturesSettleDualModeRes struct {
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
	MarginMode             int64  `json:"margin_mode"`              // 保证金模式，0-经典保证金模式，1-跨币种保证金模式，2-组合保证金模式
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

type PrivateRestFuturesSettleDualCompPositionsContractRes []GatePositionResCommon
type PrivateRestFuturesSettleDualCompPositionsContractMarginRes []GatePositionResCommon
type PrivateRestFuturesSettleDualCompPositionsContractLeverageRes []GatePositionResCommon
type PrivateRestFuturesSettleDualCompPositionsContractRiskLimitRes []GatePositionResCommon

type PrivateRestFuturesSettleOrdersPostRes GateFuturesOrderResCommon
type PrivateRestFuturesSettleOrdersGetRes []GateFuturesOrderResCommon
type PrivateRestFuturesSettleOrdersOrderIdGetRes GateFuturesOrderResCommon
type PrivateRestFuturesSettleOrdersOrderIdDeleteRes GateFuturesOrderResCommon
type PrivateRestFuturesSettleOrdersOrderIdPutRes GateFuturesOrderResCommon

type PrivateRestFuturesSettleMyTradesResRow struct {
	Id         int64   `json:"id"`          //成交记录 ID
	CreateTime float64 `json:"create_time"` //成交时间
	Contract   string  `json:"contract"`    //合约标识
	OrderId    string  `json:"order_id"`    //成交记录关联订单 ID
	Size       int64   `json:"size"`        //成交数量
	CloseSize  int64   `json:"close_size"`  //平仓数量 close_size=0 && size＞0 开多 close_size=0 && size＜0 开空 close_size>0 && size>0 && size <= close_size 平空 close_size>0 && size>0 && size > close_size 平空且开多 close_size<0 && size<0 && size >= close_size 平多 close_size<0 && size<0 && size < close_size 平多且开空
	Price      string  `json:"price"`       //成交价格
	Role       string  `json:"role"`        //成交角色， taker - 吃单, maker - 做单
	Text       string  `json:"text"`        //订单的自定义信息
	Fee        string  `json:"fee"`         //成交手续费
	PointFee   string  `json:"point_fee"`   //成交点卡手续费
}
type PrivateRestFuturesSettleMyTradesRes []PrivateRestFuturesSettleMyTradesResRow

type PrivateRestFuturesSettleOrdersTimeRangeRes []struct {
	Id           int64   `json:"id"`             // 合约订单 ID
	User         int64   `json:"user"`           // 用户 ID
	CreateTime   float64 `json:"create_time"`    // 订单创建时间
	FinishTime   string  `json:"finish_time"`    // 订单结束时间，未结束订单无此字段返回
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
