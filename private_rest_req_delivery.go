package mygateapi

type PrivateRestDeliverySettleAccountsReq struct {
	Settle *string `json:"settle"` //结算货币
}

// 结算货币
func (api *PrivateRestDeliverySettleAccountsAPI) Settle(settle string) *PrivateRestDeliverySettleAccountsAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

type PrivateRestDeliverySettleAccountsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettleAccountsReq
}

type PrivateRestDeliverySettleAccountBookReq struct {
	Settle *string `json:"settle"` //结算货币
	Limit  *int    `json:"limit"`  //列表返回的最大数量
	From   *int64  `json:"from"`   //起始时间戳
	To     *int64  `json:"to"`     //终止时间戳
	Type   *string `json:"type"`   //类型
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettleAccountBookAPI) Settle(settle string) *PrivateRestDeliverySettleAccountBookAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestDeliverySettleAccountBookAPI) Limit(limit int) *PrivateRestDeliverySettleAccountBookAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// from	请求参数	integer(int64)	否	起始时间戳
func (api *PrivateRestDeliverySettleAccountBookAPI) From(from int64) *PrivateRestDeliverySettleAccountBookAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	integer(int64)	否	终止时间戳
func (api *PrivateRestDeliverySettleAccountBookAPI) To(to int64) *PrivateRestDeliverySettleAccountBookAPI {
	api.req.To = GetPointer(to)
	return api
}

// type	请求参数	string	否
func (api *PrivateRestDeliverySettleAccountBookAPI) Type(type_ string) *PrivateRestDeliverySettleAccountBookAPI {
	api.req.Type = GetPointer(type_)
	return api
}

type PrivateRestDeliverySettleAccountBookAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettleAccountBookReq
}

type PrivateRestDeliverySettlePositionsReq struct {
	Settle *string `json:"settle"` //结算货币
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePositionsAPI) Settle(settle string) *PrivateRestDeliverySettlePositionsAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

type PrivateRestDeliverySettlePositionsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePositionsReq
}

type PrivateRestDeliverySettlePositionsContractReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePositionsContractAPI) Settle(settle string) *PrivateRestDeliverySettlePositionsContractAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PrivateRestDeliverySettlePositionsContractAPI) Contract(contract string) *PrivateRestDeliverySettlePositionsContractAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

type PrivateRestDeliverySettlePositionsContractAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePositionsContractReq
}

type PrivateRestDeliverySettlePositionsContractMarginReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
	Change   *string `json:"change"`   //保证金变化数额，正数增加，负数减少
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePositionsContractMarginAPI) Settle(settle string) *PrivateRestDeliverySettlePositionsContractMarginAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PrivateRestDeliverySettlePositionsContractMarginAPI) Contract(contract string) *PrivateRestDeliverySettlePositionsContractMarginAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// change	请求参数	string	是	保证金变化数额，正数增加，负数减少
func (api *PrivateRestDeliverySettlePositionsContractMarginAPI) Change(change string) *PrivateRestDeliverySettlePositionsContractMarginAPI {
	api.req.Change = GetPointer(change)
	return api
}

type PrivateRestDeliverySettlePositionsContractMarginAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePositionsContractMarginReq
}

type PrivateRestDeliverySettlePositionsContractLeverageReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
	Leverage *string `json:"leverage"` //新的杠杆倍数
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePositionsContractLeverageAPI) Settle(settle string) *PrivateRestDeliverySettlePositionsContractLeverageAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PrivateRestDeliverySettlePositionsContractLeverageAPI) Contract(contract string) *PrivateRestDeliverySettlePositionsContractLeverageAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// leverage	请求参数	string	是	新的杠杆倍数
func (api *PrivateRestDeliverySettlePositionsContractLeverageAPI) Leverage(leverage string) *PrivateRestDeliverySettlePositionsContractLeverageAPI {
	api.req.Leverage = GetPointer(leverage)
	return api
}

type PrivateRestDeliverySettlePositionsContractLeverageAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePositionsContractLeverageReq
}

type PrivateRestDeliverySettlePositionsContractRiskLimitReq struct {
	Settle    *string `json:"settle"`     //结算货币
	Contract  *string `json:"contract"`   //合约标识
	RiskLimit *string `json:"risk_limit"` //新的风险限额
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePositionsContractRiskLimitAPI) Settle(settle string) *PrivateRestDeliverySettlePositionsContractRiskLimitAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PrivateRestDeliverySettlePositionsContractRiskLimitAPI) Contract(contract string) *PrivateRestDeliverySettlePositionsContractRiskLimitAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// risk_limit	请求参数	string	是	新的风险限额
func (api *PrivateRestDeliverySettlePositionsContractRiskLimitAPI) RiskLimit(riskLimit string) *PrivateRestDeliverySettlePositionsContractRiskLimitAPI {
	api.req.RiskLimit = GetPointer(riskLimit)
	return api
}

type PrivateRestDeliverySettlePositionsContractRiskLimitAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePositionsContractRiskLimitReq
}

type PrivateRestDeliverySettleOrdersPostReq GateFuturesOrderReqCommon

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettleOrdersPostAPI) Settle(settle string) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PrivateRestDeliverySettleOrdersPostAPI) Contract(contract string) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// size	请求参数	integer	是	交易数量，正数为买入，负数为卖出。平仓委托则设置为0。
func (api *PrivateRestDeliverySettleOrdersPostAPI) Size(size int64) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.Size = GetPointer(size)
	return api
}

// iceberg	请求参数	integer	否	冰山委托显示数量。0为完全不隐藏。注意，隐藏部分成交按照taker收取手续费。
func (api *PrivateRestDeliverySettleOrdersPostAPI) Iceberg(iceberg int64) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.Iceberg = GetPointer(iceberg)
	return api
}

// price	请求参数	string	是	委托价。价格为0并且tif为ioc，代表市价委托。
func (api *PrivateRestDeliverySettleOrdersPostAPI) Price(price string) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.Price = GetPointer(price)
	return api
}

// close	请求参数	boolean	否	设置为 true 的时候执行平仓操作，并且size应设置为0
func (api *PrivateRestDeliverySettleOrdersPostAPI) Close(close bool) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.Close = GetPointer(close)
	return api
}

// reduce_only	请求参数	boolean	否	设置为 true 的时候，为只减仓委托
func (api *PrivateRestDeliverySettleOrdersPostAPI) ReduceOnly(reduceOnly bool) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// tif	请求参数	string	是	Time in force 策略，市价单当前只支持 ioc 模式
func (api *PrivateRestDeliverySettleOrdersPostAPI) Tif(tif string) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.Tif = GetPointer(tif)
	return api
}

// text	请求参数	string	否	订单自定义信息，用户可以用该字段设置自定义 ID，用户自定义字段必须满足以下条件：
func (api *PrivateRestDeliverySettleOrdersPostAPI) Text(text string) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.Text = GetPointer(text)
	return api
}

// auto_size	请求参数	string	否	双仓模式下用于设置平仓的方向，close_long 平多头， close_short 平空头，需要同时设置 size 为 0
func (api *PrivateRestDeliverySettleOrdersPostAPI) AutoSize(autoSize string) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.AutoSize = GetPointer(autoSize)
	return api
}

// stp_act	请求参数	string	否	Self-Trading Prevention Action,用户可以用该字段设置自定义限制自成交策略。
func (api *PrivateRestDeliverySettleOrdersPostAPI) StpAct(stpAct string) *PrivateRestDeliverySettleOrdersPostAPI {
	api.req.StpAct = GetPointer(stpAct)
	return api
}

type PrivateRestDeliverySettleOrdersPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettleOrdersPostReq
}

type PrivateRestDeliverySettleOrdersGetReq struct {
	Contract   *string `json:"contract"`    //合约标识
	Status     *string `json:"status"`      //基于状态查询订单列表
	Limit      *int    `json:"limit"`       //列表返回的最大数量
	Offset     *int    `json:"offset"`      //列表返回的偏移量，从 0 开始
	LastId     *string `json:"last_id"`     //以上个列表的最后一条记录的 ID 作为下个列表的起点
	CountTotal *int    `json:"count_total"` //是否需要返回列表总数，默认为 0 不返回
	Settle     *string `json:"settle"`      //结算货币
}

// contract	URL	string	是	合约标识
func (api *PrivateRestDeliverySettleOrdersGetAPI) Contract(contract string) *PrivateRestDeliverySettleOrdersGetAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// status	请求参数	string	否	基于状态查询订单列表
func (api *PrivateRestDeliverySettleOrdersGetAPI) Status(status string) *PrivateRestDeliverySettleOrdersGetAPI {
	api.req.Status = GetPointer(status)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestDeliverySettleOrdersGetAPI) Limit(limit int) *PrivateRestDeliverySettleOrdersGetAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
func (api *PrivateRestDeliverySettleOrdersGetAPI) Offset(offset int) *PrivateRestDeliverySettleOrdersGetAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// last_id	请求参数	string	否	以上个列表的最后一条记录的 ID 作为下个列表的起点
func (api *PrivateRestDeliverySettleOrdersGetAPI) LastId(lastId string) *PrivateRestDeliverySettleOrdersGetAPI {
	api.req.LastId = GetPointer(lastId)
	return api
}

// count_total	请求参数	integer	否	是否需要返回列表总数，默认为 0 不返回
func (api *PrivateRestDeliverySettleOrdersGetAPI) CountTotal(countTotal int) *PrivateRestDeliverySettleOrdersGetAPI {
	api.req.CountTotal = GetPointer(countTotal)
	return api
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettleOrdersGetAPI) Settle(settle string) *PrivateRestDeliverySettleOrdersGetAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

type PrivateRestDeliverySettleOrdersGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettleOrdersGetReq
}

type PrivateRestDeliverySettleOrdersDeleteReq struct {
	Contract *string `json:"contract"` //合约标识
	Side     *string `json:"side"`     //指定全部买单或全部卖单，不指定则两者都包括
	Settle   *string `json:"settle"`   //结算货币
}

// contract	URL	string	是	合约标识
func (api *PrivateRestDeliverySettleOrdersDeleteAPI) Contract(contract string) *PrivateRestDeliverySettleOrdersDeleteAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// side	请求参数	string	否	指定全部买单或全部卖单，不指定则两者都包括
func (api *PrivateRestDeliverySettleOrdersDeleteAPI) Side(side string) *PrivateRestDeliverySettleOrdersDeleteAPI {
	api.req.Side = GetPointer(side)
	return api
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettleOrdersDeleteAPI) Settle(settle string) *PrivateRestDeliverySettleOrdersDeleteAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

type PrivateRestDeliverySettleOrdersDeleteAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettleOrdersDeleteReq
}

type PrivateRestDeliverySettleOrdersOrderIdGetReq struct {
	Settle  *string `json:"settle"`   //结算货币
	OrderId *string `json:"order_id"` //成功创建订单时返回的 ID
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettleOrdersOrderIdGetAPI) Settle(settle string) *PrivateRestDeliverySettleOrdersOrderIdGetAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// order_id	URL	string	是	成功创建订单时返回的 ID
func (api *PrivateRestDeliverySettleOrdersOrderIdGetAPI) OrderId(orderId string) *PrivateRestDeliverySettleOrdersOrderIdGetAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

type PrivateRestDeliverySettleOrdersOrderIdGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettleOrdersOrderIdGetReq
}

type PrivateRestDeliverySettleOrdersOrderIdDeleteReq struct {
	Settle  *string `json:"settle"`   //结算货币
	OrderId *string `json:"order_id"` //成功创建订单时返回的 ID
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettleOrdersOrderIdDeleteAPI) Settle(settle string) *PrivateRestDeliverySettleOrdersOrderIdDeleteAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// order_id	URL	string	是	成功创建订单时返回的 ID
func (api *PrivateRestDeliverySettleOrdersOrderIdDeleteAPI) OrderId(orderId string) *PrivateRestDeliverySettleOrdersOrderIdDeleteAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

type PrivateRestDeliverySettleOrdersOrderIdDeleteAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettleOrdersOrderIdDeleteReq
}

type PrivateRestDeliverySettleMyTradesReq struct {
	Settle     *string `json:"settle"`      //结算货币
	Contract   *string `json:"contract"`    //合约标识
	Order      *int64  `json:"order"`       //委托ID，如果指定则返回该委托相关数据
	Limit      *int    `json:"limit"`       //列表返回的最大数量
	Offset     *int    `json:"offset"`      //列表返回的偏移量，从 0 开始
	LastId     *string `json:"last_id"`     //以上个列表的最后一条记录的 ID 作为下个列表的起点
	CountTotal *int    `json:"count_total"` //是否需要返回列表总数，默认为 0 不返回
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettleMyTradesAPI) Settle(settle string) *PrivateRestDeliverySettleMyTradesAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	否	合约标识
func (api *PrivateRestDeliverySettleMyTradesAPI) Contract(contract string) *PrivateRestDeliverySettleMyTradesAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// order	请求参数	integer	否	委托ID，如果指定则返回该委托相关数据
func (api *PrivateRestDeliverySettleMyTradesAPI) Order(order int64) *PrivateRestDeliverySettleMyTradesAPI {
	api.req.Order = GetPointer(order)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestDeliverySettleMyTradesAPI) Limit(limit int) *PrivateRestDeliverySettleMyTradesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
func (api *PrivateRestDeliverySettleMyTradesAPI) Offset(offset int) *PrivateRestDeliverySettleMyTradesAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// last_id	请求参数	string	否	以上个列表的最后一条记录的 ID 作为下个列表的起点
func (api *PrivateRestDeliverySettleMyTradesAPI) LastId(lastId string) *PrivateRestDeliverySettleMyTradesAPI {
	api.req.LastId = GetPointer(lastId)
	return api
}

// count_total	请求参数	integer	否	是否需要返回列表总数，默认为 0 不返回
func (api *PrivateRestDeliverySettleMyTradesAPI) CountTotal(countTotal int) *PrivateRestDeliverySettleMyTradesAPI {
	api.req.CountTotal = GetPointer(countTotal)
	return api
}

type PrivateRestDeliverySettleMyTradesAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettleMyTradesReq
}

type PrivateRestDeliverySettlePositionCloseReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
	Limit    *int    `json:"limit"`    //列表返回的最大数量
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePositionCloseAPI) Settle(settle string) *PrivateRestDeliverySettlePositionCloseAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PrivateRestDeliverySettlePositionCloseAPI) Contract(contract string) *PrivateRestDeliverySettlePositionCloseAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestDeliverySettlePositionCloseAPI) Limit(limit int) *PrivateRestDeliverySettlePositionCloseAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PrivateRestDeliverySettlePositionCloseAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePositionCloseReq
}

type PrivateRestDeliverySettleLiquidatesReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
	Limit    *int    `json:"limit"`    //列表返回的最大数量
	At       *int    `json:"at"`       //指定时间戳的强平历史
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettleLiquidatesAPI) Settle(settle string) *PrivateRestDeliverySettleLiquidatesAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PrivateRestDeliverySettleLiquidatesAPI) Contract(contract string) *PrivateRestDeliverySettleLiquidatesAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestDeliverySettleLiquidatesAPI) Limit(limit int) *PrivateRestDeliverySettleLiquidatesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// at	请求参数	integer	否	指定时间戳的强平历史
func (api *PrivateRestDeliverySettleLiquidatesAPI) At(at int) *PrivateRestDeliverySettleLiquidatesAPI {
	api.req.At = GetPointer(at)
	return api
}

type PrivateRestDeliverySettleLiquidatesAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettleLiquidatesReq
}

type PrivateRestDeliverySettleSettlementsReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
	Limit    *int    `json:"limit"`    //列表返回的最大数量
	At       *int    `json:"at"`       //指定时间戳的结算历史
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettleSettlementsAPI) Settle(settle string) *PrivateRestDeliverySettleSettlementsAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PrivateRestDeliverySettleSettlementsAPI) Contract(contract string) *PrivateRestDeliverySettleSettlementsAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestDeliverySettleSettlementsAPI) Limit(limit int) *PrivateRestDeliverySettleSettlementsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// at	请求参数	integer	否	指定时间戳的结算历史
func (api *PrivateRestDeliverySettleSettlementsAPI) At(at int) *PrivateRestDeliverySettleSettlementsAPI {
	api.req.At = GetPointer(at)
	return api
}

type PrivateRestDeliverySettleSettlementsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettleSettlementsReq
}

type PrivateRestDeliverySettlePriceOrdersPostReq struct {
	Settle  *string                    `json:"settle"`  //结算货币
	Initial *GateFuturesOrderReqCommon `json:"initial"` //body	是
	Trigger *struct {
		StrategyType *int    `json:"strategy_type"` //body	integer(int32)	否	触发策略
		PriceType    *int    `json:"price_type"`    //body	integer(int32)	否	参考价格类型。 0 - 最新成交价，1 - 标记价格，2 - 指数价格
		Price        *string `json:"price"`         //body	string	否	价格触发时为价格，价差触发时为价差
		Rule         *int    `json:"rule"`          //body	integer(int32)	否	价格条件类型
		Expiration   *int    `json:"expiration"`    //body	integer	否	最长等待触发时间，超时则取消该订单，单位是秒 s
	} `json:"trigger"` //body	是
	OrderType *string `json:"order_type"` //body	string	否	止盈止损的类型
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) Settle(settle string) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// initial	body	是
// contract 请求参数	string	是	合约标识
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) Contract(contract string) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Initial.Contract = GetPointer(contract)
	return api
}

// size	请求参数	integer	否	交易数量，正数为买入，负数为卖出。平仓委托则设置为0。
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) Size(size int64) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Initial.Size = GetPointer(size)
	return api
}

// price 请求参数	string	是	委托价。价格为0并且tif为ioc，代表市价委托。
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) InitialPrice(price string) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Initial.Price = GetPointer(price)
	return api
}

// close	请求参数	boolean	否	设置为 true 的时候执行平仓操作，并且size应设置为0
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) Close(close bool) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Initial.Close = GetPointer(close)
	return api
}

// tif 请求参数	string	否	Time in force 策略，市价单当前只支持 ioc 模式
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) Tif(tif string) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Initial.Tif = GetPointer(tif)
	return api
}

// text	请求参数	string	否	订单自定义信息，用户可以用该字段设置自定义 ID，用户自定义字段必须满足以下条件：
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) Text(text string) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Initial.Text = GetPointer(text)
	return api
}

// reduce_only	请求参数	boolean	否	设置为 true 的时候执行自动减仓操作
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) ReduceOnly(reduceOnly bool) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Initial.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// auto_size	请求参数	string	否	双仓模式下用于设置平仓的方向，close_long 平多头， close_short 平空头，需要同时设置 size 为 0
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) AutoSize(autoSize string) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Initial.AutoSize = GetPointer(autoSize)
	return api
}

// trigger body 是
// strategy_type	请求参数	integer	否	触发策略
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) StrategyType(strategyType int) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Trigger.StrategyType = GetPointer(strategyType)
	return api
}

// price_type	请求参数	integer	否	参考价格类型。 0 - 最新成交价，1 - 标记价格，2 - 指数价格
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) PriceType(priceType int) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Trigger.PriceType = GetPointer(priceType)
	return api
}

// price	请求参数	string	否	价格触发时为价格，价差触发时为价差
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) TriggerPrice(price string) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Trigger.Price = GetPointer(price)
	return api
}

// rule	请求参数	integer	否	价格条件类型
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) Rule(rule int) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Trigger.Rule = GetPointer(rule)
	return api
}

// expiration	请求参数	integer	否	最长等待触发时间，超时则取消该订单，单位是秒 s
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) Expiration(expiration int) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.Trigger.Expiration = GetPointer(expiration)
	return api
}

// order_type	请求参数	string	否	止盈止损的类型
func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) OrderType(orderType string) *PrivateRestDeliverySettlePriceOrdersPostAPI {
	api.req.OrderType = GetPointer(orderType)
	return api
}

type PrivateRestDeliverySettlePriceOrdersPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePriceOrdersPostReq
}

type PrivateRestDeliverySettlePriceOrdersDeleteReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePriceOrdersDeleteAPI) Settle(settle string) *PrivateRestDeliverySettlePriceOrdersDeleteAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PrivateRestDeliverySettlePriceOrdersDeleteAPI) Contract(contract string) *PrivateRestDeliverySettlePriceOrdersDeleteAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

type PrivateRestDeliverySettlePriceOrdersDeleteAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePriceOrdersDeleteReq
}

type PrivateRestDeliverySettlePriceOrdersGetReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Status   *string `json:"status"`   //基于状态查询订单列表
	Contract *string `json:"contract"` //合约标识，如果指定则只返回该合约相关数据
	Limit    *int    `json:"limit"`    //列表返回的最大数量
	Offset   *int    `json:"offset"`   //列表返回的偏移量，从 0 开始
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePriceOrdersGetAPI) Settle(settle string) *PrivateRestDeliverySettlePriceOrdersGetAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// status	请求参数	string	否	基于状态查询订单列表
func (api *PrivateRestDeliverySettlePriceOrdersGetAPI) Status(status string) *PrivateRestDeliverySettlePriceOrdersGetAPI {
	api.req.Status = GetPointer(status)
	return api
}

// contract	请求参数	string	否	合约标识，如果指定则只返回该合约相关数据
func (api *PrivateRestDeliverySettlePriceOrdersGetAPI) Contract(contract string) *PrivateRestDeliverySettlePriceOrdersGetAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestDeliverySettlePriceOrdersGetAPI) Limit(limit int) *PrivateRestDeliverySettlePriceOrdersGetAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
func (api *PrivateRestDeliverySettlePriceOrdersGetAPI) Offset(offset int) *PrivateRestDeliverySettlePriceOrdersGetAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

type PrivateRestDeliverySettlePriceOrdersGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePriceOrdersGetReq
}

type PrivateRestDeliverySettlePriceOrdersOrderIdGetReq struct {
	Settle  *string `json:"settle"`   //结算货币
	OrderId *string `json:"order_id"` //成功创建订单时返回的 ID
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePriceOrdersOrderIdGetAPI) Settle(settle string) *PrivateRestDeliverySettlePriceOrdersOrderIdGetAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// order_id	URL	string	是	成功创建订单时返回的 ID
func (api *PrivateRestDeliverySettlePriceOrdersOrderIdGetAPI) OrderId(orderId string) *PrivateRestDeliverySettlePriceOrdersOrderIdGetAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

type PrivateRestDeliverySettlePriceOrdersOrderIdGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePriceOrdersOrderIdGetReq
}

type PrivateRestDeliverySettlePriceOrdersOrderIdDeleteReq struct {
	Settle  *string `json:"settle"`   //结算货币
	OrderId *string `json:"order_id"` //成功创建订单时返回的 ID
}

// settle	URL	string	是	结算货币
func (api *PrivateRestDeliverySettlePriceOrdersOrderIdDeleteAPI) Settle(settle string) *PrivateRestDeliverySettlePriceOrdersOrderIdDeleteAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// order_id	URL	string	是	成功创建订单时返回的 ID
func (api *PrivateRestDeliverySettlePriceOrdersOrderIdDeleteAPI) OrderId(orderId string) *PrivateRestDeliverySettlePriceOrdersOrderIdDeleteAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

type PrivateRestDeliverySettlePriceOrdersOrderIdDeleteAPI struct {
	client *PrivateRestClient
	req    *PrivateRestDeliverySettlePriceOrdersOrderIdDeleteReq
}
