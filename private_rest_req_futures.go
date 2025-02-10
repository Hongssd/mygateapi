package mygateapi

type PrivateRestFuturesSettleAccountsReq struct {
	Settle *string `json:"settle"` // settle	URL	string	是	结算货币
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleAccountsAPI) Settle(settle string) *PrivateRestFuturesSettleAccountsAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

type PrivateRestFuturesSettleAccountsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleAccountsReq
}

type PrivateRestFuturesSettleAccountBookReq struct {
	Settle   *string `json:"settle"`   // settle	URL	string	是	结算货币
	Contract *string `json:"contract"` // contract	请求参数	string	否	合约标识，如果指定则只返回该合约相关数据
	Limit    *int    `json:"limit"`    // limit	请求参数	integer	否	列表返回的最大数量
	Offset   *int    `json:"offset"`   // offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
	From     *int64  `json:"from"`     // from	请求参数	integer(int64)	否	起始时间戳
	To       *int64  `json:"to"`       // to	请求参数	integer(int64)	否	终止时间戳
	Type     *string `json:"type"`     // type	请求参数	string	否	变更类型
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleAccountBookAPI) Settle(settle string) *PrivateRestFuturesSettleAccountBookAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	否	合约标识，如果指定则只返回该合约相关数据
func (api *PrivateRestFuturesSettleAccountBookAPI) Contract(contract string) *PrivateRestFuturesSettleAccountBookAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestFuturesSettleAccountBookAPI) Limit(limit int) *PrivateRestFuturesSettleAccountBookAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
func (api *PrivateRestFuturesSettleAccountBookAPI) Offset(offset int) *PrivateRestFuturesSettleAccountBookAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// from	请求参数	integer(int64)	否	起始时间戳
func (api *PrivateRestFuturesSettleAccountBookAPI) From(from int64) *PrivateRestFuturesSettleAccountBookAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	integer(int64)	否	终止时间戳
func (api *PrivateRestFuturesSettleAccountBookAPI) To(to int64) *PrivateRestFuturesSettleAccountBookAPI {
	api.req.To = GetPointer(to)
	return api
}

// type	请求参数	string	否	变更类型 可选值：dnw: 转入转出, pnl: 减仓盈亏, fee: 交易手续费, refr: 推荐人返佣, fund: 资金费用, point_dnw: 点卡转入转出, point_fee: 点卡交易手续费, point_refr: 点卡推荐人返佣, bonus_offset: 体验金抵扣
func (api *PrivateRestFuturesSettleAccountBookAPI) Type(type_ string) *PrivateRestFuturesSettleAccountBookAPI {
	api.req.Type = GetPointer(type_)
	return api
}

type PrivateRestFuturesSettleAccountBookAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleAccountBookReq
}

type PrivateRestFuturesSettlePositionsReq struct {
	Settle  *string `json:"settle"`  // settle	URL	string	是	结算货币
	Holding *bool   `json:"holding"` // holding	请求参数	boolean	否	只返回真实持仓-true,全部返回-false
	Limit   *int    `json:"limit"`   // limit	请求参数	integer	否	列表返回的最大数量
	Offset  *int    `json:"offset"`  // offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
}

// settle	URL	string	是	结算货币
func (api *PrivateRestFuturesSettlePositionsAPI) Settle(settle string) *PrivateRestFuturesSettlePositionsAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// holding	请求参数	boolean	否	只返回真实持仓-true,全部返回-false
func (api *PrivateRestFuturesSettlePositionsAPI) Holding(holding bool) *PrivateRestFuturesSettlePositionsAPI {
	api.req.Holding = GetPointer(holding)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestFuturesSettlePositionsAPI) Limit(limit int) *PrivateRestFuturesSettlePositionsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
func (api *PrivateRestFuturesSettlePositionsAPI) Offset(offset int) *PrivateRestFuturesSettlePositionsAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

type PrivateRestFuturesSettlePositionsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettlePositionsReq
}

type PrivateRestFuturesSettlePositionsContractReq struct {
	Settle   *string `json:"settle"`   // settle	URL	string	是	结算货币
	Contract *string `json:"contract"` // contract	URL	string	是	合约标识
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettlePositionsContractAPI) Settle(settle string) *PrivateRestFuturesSettlePositionsContractAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PrivateRestFuturesSettlePositionsContractAPI) Contract(contract string) *PrivateRestFuturesSettlePositionsContractAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

type PrivateRestFuturesSettlePositionsContractAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettlePositionsContractReq
}

type PrivateRestFuturesSettlePositionsContractMarginReq struct {
	Settle   *string `json:"settle"`   // settle	URL	string	是	结算货币
	Contract *string `json:"contract"` // contract	URL	string	是	合约标识
	Change   *string `json:"change"`   // change	请求参数	string	是	保证金变化数额，正数增加，负数减少
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettlePositionsContractMarginAPI) Settle(settle string) *PrivateRestFuturesSettlePositionsContractMarginAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PrivateRestFuturesSettlePositionsContractMarginAPI) Contract(contract string) *PrivateRestFuturesSettlePositionsContractMarginAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// change	请求参数	string	是	保证金变化数额，正数增加，负数减少
func (api *PrivateRestFuturesSettlePositionsContractMarginAPI) Change(change string) *PrivateRestFuturesSettlePositionsContractMarginAPI {
	api.req.Change = GetPointer(change)
	return api
}

type PrivateRestFuturesSettlePositionsContractMarginAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettlePositionsContractMarginReq
}

type PrivateRestFuturesSettlePositionsContractLeverageReq struct {
	Settle             *string `json:"settle"`               // settle	URL	string	是	结算货币
	Contract           *string `json:"contract"`             // contract	URL	string	是	合约标识
	Leverage           *string `json:"leverage"`             // leverage	请求参数	string	是	新的杠杆倍数
	CrossLeverageLimit *string `json:"cross_leverage_limit"` // cross_leverage_limit	请求参数	string	否	全仓模式下的杠杆倍数（即 leverage 为 0 时）
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettlePositionsContractLeverageAPI) Settle(settle string) *PrivateRestFuturesSettlePositionsContractLeverageAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PrivateRestFuturesSettlePositionsContractLeverageAPI) Contract(contract string) *PrivateRestFuturesSettlePositionsContractLeverageAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// leverage	请求参数	string	是	新的杠杆倍数
func (api *PrivateRestFuturesSettlePositionsContractLeverageAPI) Leverage(leverage string) *PrivateRestFuturesSettlePositionsContractLeverageAPI {
	api.req.Leverage = GetPointer(leverage)
	return api
}

// cross_leverage_limit	请求参数	string	否	全仓模式下的杠杆倍数（即 leverage 为 0 时）
func (api *PrivateRestFuturesSettlePositionsContractLeverageAPI) CrossLeverageLimit(crossLeverageLimit string) *PrivateRestFuturesSettlePositionsContractLeverageAPI {
	api.req.CrossLeverageLimit = GetPointer(crossLeverageLimit)
	return api
}

type PrivateRestFuturesSettlePositionsContractLeverageAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettlePositionsContractLeverageReq
}

type PrivateRestFuturesSettlePositionsContractRiskLimitReq struct {
	Settle    *string `json:"settle"`     // settle	URL	string	是	结算货币
	Contract  *string `json:"contract"`   // contract	URL	string	是	合约标识
	RiskLimit *string `json:"risk_limit"` // risk_limit	请求参数	string	是	新的风险限额价值
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettlePositionsContractRiskLimitAPI) Settle(settle string) *PrivateRestFuturesSettlePositionsContractRiskLimitAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PrivateRestFuturesSettlePositionsContractRiskLimitAPI) Contract(contract string) *PrivateRestFuturesSettlePositionsContractRiskLimitAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// risk_limit	请求参数	string	是	新的风险限额价值
func (api *PrivateRestFuturesSettlePositionsContractRiskLimitAPI) RiskLimit(riskLimit string) *PrivateRestFuturesSettlePositionsContractRiskLimitAPI {
	api.req.RiskLimit = GetPointer(riskLimit)
	return api
}

type PrivateRestFuturesSettlePositionsContractRiskLimitAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettlePositionsContractRiskLimitReq
}

type PrivateRestFuturesSettleDualModeReq struct {
	Settle   *string `json:"settle"`    // settle	URL	string	是	结算货币
	DualMode *bool   `json:"dual_mode"` // dual_mode	请求参数	boolean	是	是否设置为双向持仓
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleDualModeAPI) Settle(settle string) *PrivateRestFuturesSettleDualModeAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// dual_mode	请求参数	boolean	是	是否设置为双向持仓
func (api *PrivateRestFuturesSettleDualModeAPI) DualMode(dualMode bool) *PrivateRestFuturesSettleDualModeAPI {
	api.req.DualMode = GetPointer(dualMode)
	return api
}

type PrivateRestFuturesSettleDualModeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleDualModeReq
}

type PrivateRestFuturesSettleDualCompPositionsContractReq struct {
	Settle   *string `json:"settle"`   // settle	URL	string	是	结算货币
	Contract *string `json:"contract"` // contract	URL	string	是	合约标识
}

type PrivateRestFuturesSettleDualCompPositionsContractAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleDualCompPositionsContractReq
}

type PrivateRestFuturesSettleDualCompPositionsContractMarginReq struct {
	Settle   *string `json:"settle"`    // settle	URL	string	是	结算货币
	Contract *string `json:"contract"`  // contract	URL	string	是	合约标识
	Change   *string `json:"change"`    // change	请求参数	string	是	保证金变化数额，正数增加，负数减少
	DualSide *string `json:"dual_side"` // dual_side	请求参数	string	是	多头或空头仓位
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleDualCompPositionsContractMarginAPI) Settle(settle string) *PrivateRestFuturesSettleDualCompPositionsContractMarginAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PrivateRestFuturesSettleDualCompPositionsContractMarginAPI) Contract(contract string) *PrivateRestFuturesSettleDualCompPositionsContractMarginAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// change	请求参数	string	是	保证金变化数额，正数增加，负数减少
func (api *PrivateRestFuturesSettleDualCompPositionsContractMarginAPI) Change(change string) *PrivateRestFuturesSettleDualCompPositionsContractMarginAPI {
	api.req.Change = GetPointer(change)
	return api
}

// dual_side	请求参数	string	是	多头或空头仓位
func (api *PrivateRestFuturesSettleDualCompPositionsContractMarginAPI) DualSide(dualSide string) *PrivateRestFuturesSettleDualCompPositionsContractMarginAPI {
	api.req.DualSide = GetPointer(dualSide)
	return api
}

type PrivateRestFuturesSettleDualCompPositionsContractMarginAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleDualCompPositionsContractMarginReq
}

type PrivateRestFuturesSettleDualCompPositionsContractLeverageReq struct {
	Settle             *string `json:"settle"`               // settle	URL	string	是	结算货币
	Contract           *string `json:"contract"`             // contract	URL	string	是	合约标识
	Leverage           *string `json:"leverage"`             // leverage	请求参数	string	是	新的杠杆倍数
	CrossLeverageLimit *string `json:"cross_leverage_limit"` // cross_leverage_limit	请求参数	string	否	全仓模式下的杠杆倍数（即 leverage 为 0 时）
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI) Settle(settle string) *PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI) Contract(contract string) *PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// leverage	请求参数	string	是	新的杠杆倍数
func (api *PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI) Leverage(leverage string) *PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI {
	api.req.Leverage = GetPointer(leverage)
	return api
}

// cross_leverage_limit	请求参数	string	否	全仓模式下的杠杆倍数（即 leverage 为 0 时）
func (api *PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI) CrossLeverageLimit(crossLeverageLimit string) *PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI {
	api.req.CrossLeverageLimit = GetPointer(crossLeverageLimit)
	return api
}

type PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleDualCompPositionsContractLeverageReq
}

type PrivateRestFuturesSettleDualCompPositionsContractRiskLimitReq struct {
	Settle    *string `json:"settle"`     // settle	URL	string	是	结算货币
	Contract  *string `json:"contract"`   // contract	URL	string	是	合约标识
	RiskLimit *string `json:"risk_limit"` // risk_limit	请求参数	string	是	新的风险限额价值
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleDualCompPositionsContractRiskLimitAPI) Settle(settle string) *PrivateRestFuturesSettleDualCompPositionsContractRiskLimitAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PrivateRestFuturesSettleDualCompPositionsContractRiskLimitAPI) Contract(contract string) *PrivateRestFuturesSettleDualCompPositionsContractRiskLimitAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// risk_limit	请求参数	string	是	新的风险限额价值
func (api *PrivateRestFuturesSettleDualCompPositionsContractRiskLimitAPI) RiskLimit(riskLimit string) *PrivateRestFuturesSettleDualCompPositionsContractRiskLimitAPI {
	api.req.RiskLimit = GetPointer(riskLimit)
	return api
}

type PrivateRestFuturesSettleDualCompPositionsContractRiskLimitAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleDualCompPositionsContractRiskLimitReq
}
type PrivateRestFuturesSettleOrdersPostReq GateFuturesOrderReqCommon

// settle	URL	string	是	结算货币
func (api *PrivateRestFuturesSettleOrdersPostAPI) Settle(settle string) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	body	string	是	合约标识
func (api *PrivateRestFuturesSettleOrdersPostAPI) Contract(contract string) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// size	body	integer(int64)	是	必选。交易数量，正数为买入，负数为卖出。平仓委托则设置为0。
func (api *PrivateRestFuturesSettleOrdersPostAPI) Size(size int64) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.Size = GetPointer(size)
	return api
}

// iceberg	body	integer(int64)	否	冰山委托显示数量。0为完全不隐藏。注意，隐藏部分成交按照taker收取手续费。
func (api *PrivateRestFuturesSettleOrdersPostAPI) Iceberg(iceberg int64) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.Iceberg = GetPointer(iceberg)
	return api
}

// price	body	string	否	委托价。价格为0并且tif为ioc，代表市价委托。
func (api *PrivateRestFuturesSettleOrdersPostAPI) Price(price string) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.Price = GetPointer(price)
	return api
}

// close	body	boolean	否	设置为 true 的时候执行平仓操作，并且size应设置为0
func (api *PrivateRestFuturesSettleOrdersPostAPI) Close(close bool) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.Close = GetPointer(close)
	return api
}

// reduce_only	body	boolean	否	设置为 true 的时候，为只减仓委托
func (api *PrivateRestFuturesSettleOrdersPostAPI) ReduceOnly(reduceOnly bool) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// tif	body	string	否	Time in force 策略，市价单当前只支持 ioc 模式
func (api *PrivateRestFuturesSettleOrdersPostAPI) Tif(tif string) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.Tif = GetPointer(tif)
	return api
}

// text	body	string	否	订单自定义信息，用户可以用该字段设置自定义 ID，用户自定义字段必须满足以下条件：
func (api *PrivateRestFuturesSettleOrdersPostAPI) Text(text string) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.Text = GetPointer(text)
	return api
}

// auto_size	body	string	否	双仓模式下用于设置平仓的方向，close_long 平多头， close_short 平空头，需要同时设置 size 为 0
func (api *PrivateRestFuturesSettleOrdersPostAPI) AutoSize(autoSize string) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.AutoSize = GetPointer(autoSize)
	return api
}

// stp_act	body	string	否	Self-Trading Prevention Action,用户可以用该字段设置自定义限制自成交策略。
func (api *PrivateRestFuturesSettleOrdersPostAPI) StpAct(stpAct string) *PrivateRestFuturesSettleOrdersPostAPI {
	api.req.StpAct = GetPointer(stpAct)
	return api
}

type PrivateRestFuturesSettleOrdersPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleOrdersPostReq
}

type PrivateRestFuturesSettleOrdersGetReq struct {
	Settle *string `json:"settle"`  // settle	URL	string	是	结算货币
	Status *string `json:"status"`  // status	请求参数	string	是	基于状态查询订单列表
	Limit  *int    `json:"limit"`   // limit	请求参数	integer	否	列表返回的最大数量
	Offset *int    `json:"offset"`  // offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
	LastId *string `json:"last_id"` // last_id	请求参数	string	否	以上个列表的最后一条记录的 ID 作为下个列表的起点
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleOrdersGetAPI) Settle(settle string) *PrivateRestFuturesSettleOrdersGetAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// status	请求参数	string	是	基于状态查询订单列表
func (api *PrivateRestFuturesSettleOrdersGetAPI) Status(status string) *PrivateRestFuturesSettleOrdersGetAPI {
	api.req.Status = GetPointer(status)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestFuturesSettleOrdersGetAPI) Limit(limit int) *PrivateRestFuturesSettleOrdersGetAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
func (api *PrivateRestFuturesSettleOrdersGetAPI) Offset(offset int) *PrivateRestFuturesSettleOrdersGetAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// last_id	请求参数	string	否	以上个列表的最后一条记录的 ID 作为下个列表的起点
func (api *PrivateRestFuturesSettleOrdersGetAPI) LastId(lastId string) *PrivateRestFuturesSettleOrdersGetAPI {
	api.req.LastId = GetPointer(lastId)
	return api
}

type PrivateRestFuturesSettleOrdersGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleOrdersGetReq
}

type PrivateRestFuturesSettleOrdersOrderIdGetReq struct {
	Settle  *string `json:"settle"`   // settle	URL	string	是	结算货币
	OrderId *string `json:"order_id"` // order_id	URL	string	是	订单 ID
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleOrdersOrderIdGetAPI) Settle(settle string) *PrivateRestFuturesSettleOrdersOrderIdGetAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// order_id	URL	string	是	订单 ID
func (api *PrivateRestFuturesSettleOrdersOrderIdGetAPI) OrderId(orderId string) *PrivateRestFuturesSettleOrdersOrderIdGetAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

type PrivateRestFuturesSettleOrdersOrderIdGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleOrdersOrderIdGetReq
}

type PrivateRestFuturesSettleOrdersOrderIdDeleteReq struct {
	Settle  *string `json:"settle"`   // settle	URL	string	是	结算货币
	OrderId *string `json:"order_id"` // order_id	URL	string	是	订单 ID
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleOrdersOrderIdDeleteAPI) Settle(settle string) *PrivateRestFuturesSettleOrdersOrderIdDeleteAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// order_id	URL	string	是	订单 ID
func (api *PrivateRestFuturesSettleOrdersOrderIdDeleteAPI) OrderId(orderId string) *PrivateRestFuturesSettleOrdersOrderIdDeleteAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

type PrivateRestFuturesSettleOrdersOrderIdDeleteAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleOrdersOrderIdDeleteReq
}

type PrivateRestFuturesSettleOrdersOrderIdPutReq struct {
	Settle    *string `json:"settle"`     // settle	URL	string	是	结算货币
	OrderId   *string `json:"order_id"`   // order_id	URL	string	是	成功创建订单时返回的订单 ID 或者用户创建时指定的自定义 ID（即 text 字段）。
	Size      *int64  `json:"size"`       // size	body	integer(int64)	否	新的委托大小。包括已成交委托的部分。
	Price     *string `json:"price"`      // price	body	string	否	新的委托价格。
	AmendText *string `json:"amend_text"` // amend_text	body	string	否	用户可以备注这次修改的信息。
	BizInfo   *string `json:"biz_info"`   // biz_info	body	string	否	用户可以备注这次修改的信息，比如ao。
	Bbo       *string `json:"bbo"`        // bbo	body	string	否	用户可以对手价进行修改。
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleOrdersOrderIdPutAPI) Settle(settle string) *PrivateRestFuturesSettleOrdersOrderIdPutAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// order_id	URL	string	是	成功创建订单时返回的订单 ID 或者用户创建时指定的自定义 ID（即 text 字段）。
func (api *PrivateRestFuturesSettleOrdersOrderIdPutAPI) OrderId(orderId string) *PrivateRestFuturesSettleOrdersOrderIdPutAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// size	body	integer(int64)	否	新的委托大小。包括已成交委托的部分。
func (api *PrivateRestFuturesSettleOrdersOrderIdPutAPI) Size(size int64) *PrivateRestFuturesSettleOrdersOrderIdPutAPI {
	api.req.Size = GetPointer(size)
	return api
}

// price	body	string	否	新的委托价格。
func (api *PrivateRestFuturesSettleOrdersOrderIdPutAPI) Price(price string) *PrivateRestFuturesSettleOrdersOrderIdPutAPI {
	api.req.Price = GetPointer(price)
	return api
}

// amend_text	body	string	否	用户可以备注这次修改的信息。
func (api *PrivateRestFuturesSettleOrdersOrderIdPutAPI) AmendText(amendText string) *PrivateRestFuturesSettleOrdersOrderIdPutAPI {
	api.req.AmendText = GetPointer(amendText)
	return api
}

// biz_info	body	string	否	用户可以备注这次修改的信息，比如ao。
func (api *PrivateRestFuturesSettleOrdersOrderIdPutAPI) BizInfo(bizInfo string) *PrivateRestFuturesSettleOrdersOrderIdPutAPI {
	api.req.BizInfo = GetPointer(bizInfo)
	return api
}

// bbo	body	string	否	用户可以对手价进行修改。
func (api *PrivateRestFuturesSettleOrdersOrderIdPutAPI) Bbo(bbo string) *PrivateRestFuturesSettleOrdersOrderIdPutAPI {
	api.req.Bbo = GetPointer(bbo)
	return api
}

type PrivateRestFuturesSettleOrdersOrderIdPutAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleOrdersOrderIdPutReq
}

type PrivateRestFuturesSettleMyTradesReq struct {
	Settle   *string `json:"settle"`   // settle	URL	string	是	结算货币
	Contract *string `json:"contract"` // contract	请求参数	string	否	合约标识，如果指定则只返回该合约相关数据
	Order    *int64  `json:"order"`    // order	请求参数	integer(int64)	否	委托ID，如果指定则返回该委托相关数据
	Limit    *int    `json:"limit"`    // limit	请求参数	integer	否	列表返回的最大数量
	Offset   *int    `json:"offset"`   // offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
	LastId   *string `json:"last_id"`  // last_id	请求参数	string	否	以上个列表的最后一条记录的 ID 作为下个列表的起点。
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleMyTradesAPI) Settle(settle string) *PrivateRestFuturesSettleMyTradesAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	否	合约标识，如果指定则只返回该合约相关数据
func (api *PrivateRestFuturesSettleMyTradesAPI) Contract(contract string) *PrivateRestFuturesSettleMyTradesAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// order	请求参数	integer(int64)	否	委托ID，如果指定则返回该委托相关数据
func (api *PrivateRestFuturesSettleMyTradesAPI) Order(order int64) *PrivateRestFuturesSettleMyTradesAPI {
	api.req.Order = GetPointer(order)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestFuturesSettleMyTradesAPI) Limit(limit int) *PrivateRestFuturesSettleMyTradesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
func (api *PrivateRestFuturesSettleMyTradesAPI) Offset(offset int) *PrivateRestFuturesSettleMyTradesAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// last_id	请求参数	string	否	以上个列表的最后一条记录的 ID 作为下个列表的起点。
func (api *PrivateRestFuturesSettleMyTradesAPI) LastId(lastId string) *PrivateRestFuturesSettleMyTradesAPI {
	api.req.LastId = GetPointer(lastId)
	return api
}

type PrivateRestFuturesSettleMyTradesAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleMyTradesReq
}

type PrivateRestFuturesSettleOrdersTimeRangeReq struct {
	Settle   *string `json:"settle"`   // settle	URL	string	是	结算货币
	Contract *string `json:"contract"` // contract	请求参数	string	否	合约标识，如果指定则只返回该合约相关数据
	From     *int64  `json:"from"`     // from	请求参数	integer(int64)	否	开始时间，Unix 时间戳
	To       *int64  `json:"to"`       // to	请求参数	integer(int64)	否	终止时间戳
	Limit    *int    `json:"limit"`    // limit	请求参数	integer	否	列表返回的最大数量
	Offset   *int    `json:"offset"`   // offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
}
type PrivateRestFuturesSettleOrdersTimeRangeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestFuturesSettleOrdersTimeRangeReq
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PrivateRestFuturesSettleOrdersTimeRangeAPI) Settle(settle string) *PrivateRestFuturesSettleOrdersTimeRangeAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	否	合约标识，如果指定则只返回该合约相关数据
func (api *PrivateRestFuturesSettleOrdersTimeRangeAPI) Contract(contract string) *PrivateRestFuturesSettleOrdersTimeRangeAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// from	请求参数	integer(int64)	否	起始时间戳
func (api *PrivateRestFuturesSettleOrdersTimeRangeAPI) From(from int64) *PrivateRestFuturesSettleOrdersTimeRangeAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	integer(int64)	否	终止时间戳
func (api *PrivateRestFuturesSettleOrdersTimeRangeAPI) To(to int64) *PrivateRestFuturesSettleOrdersTimeRangeAPI {
	api.req.To = GetPointer(to)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestFuturesSettleOrdersTimeRangeAPI) Limit(limit int) *PrivateRestFuturesSettleOrdersTimeRangeAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
func (api *PrivateRestFuturesSettleOrdersTimeRangeAPI) Offset(offset int) *PrivateRestFuturesSettleOrdersTimeRangeAPI {
	api.req.Offset = GetPointer(offset)
	return api
}
