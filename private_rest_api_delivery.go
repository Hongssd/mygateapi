package mygateapi

// gate PrivateRestDeliverySettleAccounts PrivateRest接口 GET 获取合约账户
func (client *PrivateRestClient) NewPrivateRestDeliverySettleAccounts() *PrivateRestDeliverySettleAccountsAPI {
	return &PrivateRestDeliverySettleAccountsAPI{
		client: client,
		req:    &PrivateRestDeliverySettleAccountsReq{},
	}
}

func (api *PrivateRestDeliverySettleAccountsAPI) Do() (*GateRestRes[PrivateRestDeliverySettleAccountsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettleAccounts])
	return gateCallApiWithSecret[PrivateRestDeliverySettleAccountsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettleAccountBook PrivateRest接口 GET 获取合约账户交易记录
func (client *PrivateRestClient) NewPrivateRestDeliverySettleAccountBook() *PrivateRestDeliverySettleAccountBookAPI {
	return &PrivateRestDeliverySettleAccountBookAPI{
		client: client,
		req:    &PrivateRestDeliverySettleAccountBookReq{},
	}
}

func (api *PrivateRestDeliverySettleAccountBookAPI) Do() (*GateRestRes[PrivateRestDeliverySettleAccountBookRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettleAccountBook])
	return gateCallApiWithSecret[PrivateRestDeliverySettleAccountBookRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettlePositions PrivateRest接口 GET 获取合约持仓
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePositions() *PrivateRestDeliverySettlePositionsAPI {
	return &PrivateRestDeliverySettlePositionsAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePositionsReq{},
	}
}

func (api *PrivateRestDeliverySettlePositionsAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePositionsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePositions])
	return gateCallApiWithSecret[PrivateRestDeliverySettlePositionsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettlePositionsContract PrivateRest接口 GET 获取单个仓位信息
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePositionsContract() *PrivateRestDeliverySettlePositionsContractAPI {
	return &PrivateRestDeliverySettlePositionsContractAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePositionsContractReq{},
	}
}

func (api *PrivateRestDeliverySettlePositionsContractAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePositionsContractRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePositionsContract])
	return gateCallApiWithSecret[PrivateRestDeliverySettlePositionsContractRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettlePositionsContractMargin PrivateRest接口 POST 更新仓位保证金
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePositionsContractMargin() *PrivateRestDeliverySettlePositionsContractMarginAPI {
	return &PrivateRestDeliverySettlePositionsContractMarginAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePositionsContractMarginReq{},
	}
}

func (api *PrivateRestDeliverySettlePositionsContractMarginAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePositionsContractMarginRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePositionsContractMargin])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestDeliverySettlePositionsContractMarginRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestDeliverySettlePositionsContractLeverage PrivateRest接口 POST 更新仓位杠杆
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePositionsContractLeverage() *PrivateRestDeliverySettlePositionsContractLeverageAPI {
	return &PrivateRestDeliverySettlePositionsContractLeverageAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePositionsContractLeverageReq{},
	}
}

func (api *PrivateRestDeliverySettlePositionsContractLeverageAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePositionsContractLeverageRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePositionsContractLeverage])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestDeliverySettlePositionsContractLeverageRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestDeliverySettlePositionsContractRiskLimit PrivateRest接口 POST 更新仓位风险限额
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePositionsContractRiskLimit() *PrivateRestDeliverySettlePositionsContractRiskLimitAPI {
	return &PrivateRestDeliverySettlePositionsContractRiskLimitAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePositionsContractRiskLimitReq{},
	}
}

func (api *PrivateRestDeliverySettlePositionsContractRiskLimitAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePositionsContractRiskLimitRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePositionsContractRiskLimit])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestDeliverySettlePositionsContractRiskLimitRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestDeliverySettleOrdersPost PrivateRest接口 POST 合约交易下单
func (client *PrivateRestClient) NewPrivateRestDeliverySettleOrdersPost() *PrivateRestDeliverySettleOrdersPostAPI {
	return &PrivateRestDeliverySettleOrdersPostAPI{
		client: client,
		req:    &PrivateRestDeliverySettleOrdersPostReq{},
	}
}

func (api *PrivateRestDeliverySettleOrdersPostAPI) Do() (*GateRestRes[PrivateRestDeliverySettleOrdersPostRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettleOrdersPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestDeliverySettleOrdersPostRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestDeliverySettleOrdersGet PrivateRest接口 GET 查询合约订单列表
func (client *PrivateRestClient) NewPrivateRestDeliverySettleOrdersGet() *PrivateRestDeliverySettleOrdersGetAPI {
	return &PrivateRestDeliverySettleOrdersGetAPI{
		client: client,
		req:    &PrivateRestDeliverySettleOrdersGetReq{},
	}
}

func (api *PrivateRestDeliverySettleOrdersGetAPI) Do() (*GateRestRes[PrivateRestDeliverySettleOrdersGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettleOrdersGet])
	return gateCallApiWithSecret[PrivateRestDeliverySettleOrdersGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettleOrdersDelete PrivateRest接口 DELETE 批量取消状态为open的订单
func (client *PrivateRestClient) NewPrivateRestDeliverySettleOrdersDelete() *PrivateRestDeliverySettleOrdersDeleteAPI {
	return &PrivateRestDeliverySettleOrdersDeleteAPI{
		client: client,
		req:    &PrivateRestDeliverySettleOrdersDeleteReq{},
	}
}

func (api *PrivateRestDeliverySettleOrdersDeleteAPI) Do() (*GateRestRes[PrivateRestDeliverySettleOrdersDeleteRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettleOrdersDelete])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestDeliverySettleOrdersDeleteRes](api.client.c, url, reqBody, DELETE)
}

// gate PrivateRestDeliverySettleOrdersOrderIdGet PrivateRest接口 GET 查询单个订单详情
func (client *PrivateRestClient) NewPrivateRestDeliverySettleOrdersOrderIdGet() *PrivateRestDeliverySettleOrdersOrderIdGetAPI {
	return &PrivateRestDeliverySettleOrdersOrderIdGetAPI{
		client: client,
		req:    &PrivateRestDeliverySettleOrdersOrderIdGetReq{},
	}
}

func (api *PrivateRestDeliverySettleOrdersOrderIdGetAPI) Do() (*GateRestRes[PrivateRestDeliverySettleOrdersOrderIdGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettleOrdersOrderIdGet])
	return gateCallApiWithSecret[PrivateRestDeliverySettleOrdersOrderIdGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettleOrdersOrderIdDelete PrivateRest接口 DELETE 撤销单个订单
func (client *PrivateRestClient) NewPrivateRestDeliverySettleOrdersOrderIdDelete() *PrivateRestDeliverySettleOrdersOrderIdDeleteAPI {
	return &PrivateRestDeliverySettleOrdersOrderIdDeleteAPI{
		client: client,
		req:    &PrivateRestDeliverySettleOrdersOrderIdDeleteReq{},
	}
}

func (api *PrivateRestDeliverySettleOrdersOrderIdDeleteAPI) Do() (*GateRestRes[PrivateRestDeliverySettleOrdersOrderIdDeleteRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettleOrdersOrderIdDelete])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestDeliverySettleOrdersOrderIdDeleteRes](api.client.c, url, reqBody, DELETE)
}

// gate PrivateRestDeliverySettleMyTrades PrivateRest接口 GET 查询个人成交记录
func (client *PrivateRestClient) NewPrivateRestDeliverySettleMyTrades() *PrivateRestDeliverySettleMyTradesAPI {
	return &PrivateRestDeliverySettleMyTradesAPI{
		client: client,
		req:    &PrivateRestDeliverySettleMyTradesReq{},
	}
}

func (api *PrivateRestDeliverySettleMyTradesAPI) Do() (*GateRestRes[PrivateRestDeliverySettleMyTradesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettleMyTrades])
	return gateCallApiWithSecret[PrivateRestDeliverySettleMyTradesRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettlePositionClose PrivateRest接口 GET 查询平仓历史
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePositionClose() *PrivateRestDeliverySettlePositionCloseAPI {
	return &PrivateRestDeliverySettlePositionCloseAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePositionCloseReq{},
	}
}

func (api *PrivateRestDeliverySettlePositionCloseAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePositionCloseRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePositionClose])
	return gateCallApiWithSecret[PrivateRestDeliverySettlePositionCloseRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettleLiquidates PrivateRest接口 GET 查询强平历史
func (client *PrivateRestClient) NewPrivateRestDeliverySettleLiquidates() *PrivateRestDeliverySettleLiquidatesAPI {
	return &PrivateRestDeliverySettleLiquidatesAPI{
		client: client,
		req:    &PrivateRestDeliverySettleLiquidatesReq{},
	}
}

func (api *PrivateRestDeliverySettleLiquidatesAPI) Do() (*GateRestRes[PrivateRestDeliverySettleLiquidatesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettleLiquidates])
	return gateCallApiWithSecret[PrivateRestDeliverySettleLiquidatesRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettleSettlements PrivateRest接口 GET 查询结算记录
func (client *PrivateRestClient) NewPrivateRestDeliverySettleSettlements() *PrivateRestDeliverySettleSettlementsAPI {
	return &PrivateRestDeliverySettleSettlementsAPI{
		client: client,
		req:    &PrivateRestDeliverySettleSettlementsReq{},
	}
}

func (api *PrivateRestDeliverySettleSettlementsAPI) Do() (*GateRestRes[PrivateRestDeliverySettleSettlementsRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettleSettlements])
	return gateCallApiWithSecret[PrivateRestDeliverySettleSettlementsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettlePriceOrdersPost PrivateRest接口 POST 创建价格触发订单
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePriceOrdersPost() *PrivateRestDeliverySettlePriceOrdersPostAPI {
	return &PrivateRestDeliverySettlePriceOrdersPostAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePriceOrdersPostReq{},
	}
}

func (api *PrivateRestDeliverySettlePriceOrdersPostAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePriceOrdersPostRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePriceOrdersPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestDeliverySettlePriceOrdersPostRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestDeliverySettlePriceOrdersGet PrivateRest接口 GET 查询自动订单列表
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePriceOrdersGet() *PrivateRestDeliverySettlePriceOrdersGetAPI {
	return &PrivateRestDeliverySettlePriceOrdersGetAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePriceOrdersGetReq{},
	}
}

func (api *PrivateRestDeliverySettlePriceOrdersGetAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePriceOrdersGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePriceOrdersGet])
	return gateCallApiWithSecret[PrivateRestDeliverySettlePriceOrdersGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettlePriceOrdersDelete PrivateRest接口 DELETE 批量取消自动订单
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePriceOrdersDelete() *PrivateRestDeliverySettlePriceOrdersDeleteAPI {
	return &PrivateRestDeliverySettlePriceOrdersDeleteAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePriceOrdersDeleteReq{},
	}
}

func (api *PrivateRestDeliverySettlePriceOrdersDeleteAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePriceOrdersDeleteRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePriceOrdersDelete])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestDeliverySettlePriceOrdersDeleteRes](api.client.c, url, reqBody, DELETE)
}

// gate PrivateRestDeliverySettlePriceOrdersOrderIdGet PrivateRest接口 GET 查询单个自动订单详情
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePriceOrdersOrderIdGet() *PrivateRestDeliverySettlePriceOrdersOrderIdGetAPI {
	return &PrivateRestDeliverySettlePriceOrdersOrderIdGetAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePriceOrdersOrderIdGetReq{},
	}
}

func (api *PrivateRestDeliverySettlePriceOrdersOrderIdGetAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePriceOrdersOrderIdGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePriceOrdersOrderIdGet])
	return gateCallApiWithSecret[PrivateRestDeliverySettlePriceOrdersOrderIdGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestDeliverySettlePriceOrdersOrderIdDelete PrivateRest接口 DELETE 撤销单个自动订单
func (client *PrivateRestClient) NewPrivateRestDeliverySettlePriceOrdersOrderIdDelete() *PrivateRestDeliverySettlePriceOrdersOrderIdDeleteAPI {
	return &PrivateRestDeliverySettlePriceOrdersOrderIdDeleteAPI{
		client: client,
		req:    &PrivateRestDeliverySettlePriceOrdersOrderIdDeleteReq{},
	}
}

func (api *PrivateRestDeliverySettlePriceOrdersOrderIdDeleteAPI) Do() (*GateRestRes[PrivateRestDeliverySettlePriceOrdersOrderIdDeleteRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestDeliverySettlePriceOrdersOrderIdDelete])
	return gateCallApiWithSecret[PrivateRestDeliverySettlePriceOrdersOrderIdDeleteRes](api.client.c, url, NIL_REQBODY, DELETE)
}
