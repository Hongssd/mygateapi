package mygateapi

// gate PrivateRestFuturesSettleAccounts PrivateRest接口 GET 获取合约账户
func (client *PrivateRestClient) NewPrivateRestFuturesSettleAccounts() *PrivateRestFuturesSettleAccountsAPI {
	return &PrivateRestFuturesSettleAccountsAPI{
		client: client,
		req:    &PrivateRestFuturesSettleAccountsReq{},
	}
}

func (api *PrivateRestFuturesSettleAccountsAPI) Do() (*GateRestRes[PrivateRestFuturesSettleAccountsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleAccounts])
	return gateCallApiWithSecret[PrivateRestFuturesSettleAccountsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestFuturesSettleAccountBook PrivateRest接口 GET 获取合约账户交易记录
func (client *PrivateRestClient) NewPrivateRestFuturesSettleAccountBook() *PrivateRestFuturesSettleAccountBookAPI {
	return &PrivateRestFuturesSettleAccountBookAPI{
		client: client,
		req:    &PrivateRestFuturesSettleAccountBookReq{},
	}
}

func (api *PrivateRestFuturesSettleAccountBookAPI) Do() (*GateRestRes[PrivateRestFuturesSettleAccountBookRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleAccountBook])
	return gateCallApiWithSecret[PrivateRestFuturesSettleAccountBookRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestFuturesSettlePositions PrivateRest接口 GET 获取用户仓位列表
func (client *PrivateRestClient) NewPrivateRestFuturesSettlePositions() *PrivateRestFuturesSettlePositionsAPI {
	return &PrivateRestFuturesSettlePositionsAPI{
		client: client,
		req:    &PrivateRestFuturesSettlePositionsReq{},
	}
}

func (api *PrivateRestFuturesSettlePositionsAPI) Do() (*GateRestRes[PrivateRestFuturesSettlePositionsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettlePositions])
	return gateCallApiWithSecret[PrivateRestFuturesSettlePositionsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestFuturesSettlePositionsContract PrivateRest接口 GET 获取用户单个仓位信息
func (client *PrivateRestClient) NewPrivateRestFuturesSettlePositionsContract() *PrivateRestFuturesSettlePositionsContractAPI {
	return &PrivateRestFuturesSettlePositionsContractAPI{
		client: client,
		req:    &PrivateRestFuturesSettlePositionsContractReq{},
	}
}

func (api *PrivateRestFuturesSettlePositionsContractAPI) Do() (*GateRestRes[PrivateRestFuturesSettlePositionsContractRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettlePositionsContract])
	return gateCallApiWithSecret[PrivateRestFuturesSettlePositionsContractRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestFuturesSettlePositionsContractMargin PrivateRest接口 POST 更新仓位保证金
func (client *PrivateRestClient) NewPrivateRestFuturesSettlePositionsContractMargin() *PrivateRestFuturesSettlePositionsContractMarginAPI {
	return &PrivateRestFuturesSettlePositionsContractMarginAPI{
		client: client,
		req:    &PrivateRestFuturesSettlePositionsContractMarginReq{},
	}
}

func (api *PrivateRestFuturesSettlePositionsContractMarginAPI) Do() (*GateRestRes[PrivateRestFuturesSettlePositionsContractMarginRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettlePositionsContractMargin])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestFuturesSettlePositionsContractMarginRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestFuturesSettlePositionsContractLeverage PrivateRest接口 POST 更新仓位杠杆
func (client *PrivateRestClient) NewPrivateRestFuturesSettlePositionsContractLeverage() *PrivateRestFuturesSettlePositionsContractLeverageAPI {
	return &PrivateRestFuturesSettlePositionsContractLeverageAPI{
		client: client,
		req:    &PrivateRestFuturesSettlePositionsContractLeverageReq{},
	}
}

func (api *PrivateRestFuturesSettlePositionsContractLeverageAPI) Do() (*GateRestRes[PrivateRestFuturesSettlePositionsContractLeverageRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettlePositionsContractLeverage])
	return gateCallApiWithSecret[PrivateRestFuturesSettlePositionsContractLeverageRes](api.client.c, url, NIL_REQBODY, POST)
}

// gate PrivateRestFuturesSettlePositionsContractRiskLimit PrivateRest接口 POST 更新仓位风险限额
func (client *PrivateRestClient) NewPrivateRestFuturesSettlePositionsContractRiskLimit() *PrivateRestFuturesSettlePositionsContractRiskLimitAPI {
	return &PrivateRestFuturesSettlePositionsContractRiskLimitAPI{
		client: client,
		req:    &PrivateRestFuturesSettlePositionsContractRiskLimitReq{},
	}
}

func (api *PrivateRestFuturesSettlePositionsContractRiskLimitAPI) Do() (*GateRestRes[PrivateRestFuturesSettlePositionsContractRiskLimitRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettlePositionsContractRiskLimit])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestFuturesSettlePositionsContractRiskLimitRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestFuturesSettleDualMode PrivateRest接口 POST 设置持仓模式
func (client *PrivateRestClient) NewPrivateRestFuturesSettleDualMode() *PrivateRestFuturesSettleDualModeAPI {
	return &PrivateRestFuturesSettleDualModeAPI{
		client: client,
		req:    &PrivateRestFuturesSettleDualModeReq{},
	}
}

func (api *PrivateRestFuturesSettleDualModeAPI) Do() (*GateRestRes[PrivateRestFuturesSettleDualModeRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleDualMode])
	return gateCallApiWithSecret[PrivateRestFuturesSettleDualModeRes](api.client.c, url, NIL_REQBODY, POST)
}

// gate PrivateRestFuturesSettleDualCompPositionsContract PrivateRest接口 POST 获取双仓模式下的持仓信息
func (client *PrivateRestClient) NewPrivateRestFuturesSettleDualCompPositionsContract() *PrivateRestFuturesSettleDualCompPositionsContractAPI {
	return &PrivateRestFuturesSettleDualCompPositionsContractAPI{
		client: client,
		req:    &PrivateRestFuturesSettleDualCompPositionsContractReq{},
	}
}

func (api *PrivateRestFuturesSettleDualCompPositionsContractAPI) Do() (*GateRestRes[PrivateRestFuturesSettleDualCompPositionsContractRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleDualCompPositionsContract])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestFuturesSettleDualCompPositionsContractRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestFuturesSettleDualCompPositionsContractMargin PrivateRest接口 POST 更新双仓模式下的持仓保证金
func (client *PrivateRestClient) NewPrivateRestFuturesSettleDualCompPositionsContractMargin() *PrivateRestFuturesSettleDualCompPositionsContractMarginAPI {
	return &PrivateRestFuturesSettleDualCompPositionsContractMarginAPI{
		client: client,
		req:    &PrivateRestFuturesSettleDualCompPositionsContractMarginReq{},
	}
}

func (api *PrivateRestFuturesSettleDualCompPositionsContractMarginAPI) Do() (*GateRestRes[PrivateRestFuturesSettleDualCompPositionsContractMarginRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleDualCompPositionsContractMargin])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestFuturesSettleDualCompPositionsContractMarginRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestFuturesSettleDualCompPositionsContractLeverage PrivateRest接口 POST 更新双仓模式下的持仓杠杆
func (client *PrivateRestClient) NewPrivateRestFuturesSettleDualCompPositionsContractLeverage() *PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI {
	return &PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI{
		client: client,
		req:    &PrivateRestFuturesSettleDualCompPositionsContractLeverageReq{},
	}
}

func (api *PrivateRestFuturesSettleDualCompPositionsContractLeverageAPI) Do() (*GateRestRes[PrivateRestFuturesSettleDualCompPositionsContractLeverageRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleDualCompPositionsContractLeverage])
	return gateCallApiWithSecret[PrivateRestFuturesSettleDualCompPositionsContractLeverageRes](api.client.c, url, NIL_REQBODY, POST)
}

// gate PrivateRestFuturesSettleDualCompPositionsContractRiskLimit PrivateRest接口 POST 更新双仓模式下的持仓风险限额
func (client *PrivateRestClient) NewPrivateRestFuturesSettleDualCompPositionsContractRiskLimit() *PrivateRestFuturesSettleDualCompPositionsContractRiskLimitAPI {
	return &PrivateRestFuturesSettleDualCompPositionsContractRiskLimitAPI{
		client: client,
		req:    &PrivateRestFuturesSettleDualCompPositionsContractRiskLimitReq{},
	}
}

func (api *PrivateRestFuturesSettleDualCompPositionsContractRiskLimitAPI) Do() (*GateRestRes[PrivateRestFuturesSettleDualCompPositionsContractRiskLimitRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleDualCompPositionsContractRiskLimit])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestFuturesSettleDualCompPositionsContractRiskLimitRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestFuturesSettleOrdersPost PrivateRest接口 POST 合约交易下单
func (client *PrivateRestClient) NewPrivateRestFuturesSettleOrdersPost() *PrivateRestFuturesSettleOrdersPostAPI {
	return &PrivateRestFuturesSettleOrdersPostAPI{
		client: client,
		req:    &PrivateRestFuturesSettleOrdersPostReq{},
	}
}

func (api *PrivateRestFuturesSettleOrdersPostAPI) Do() (*GateRestRes[PrivateRestFuturesSettleOrdersPostRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleOrdersPost])
	// 删除settle字段
	reqBody, err := removeSettleFromReqBody(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestFuturesSettleOrdersPostRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestFuturesSettleOrdersGet PrivateRest接口 GET 查询合约订单列表
func (client *PrivateRestClient) NewPrivateRestFuturesSettleOrdersGet() *PrivateRestFuturesSettleOrdersGetAPI {
	return &PrivateRestFuturesSettleOrdersGetAPI{
		client: client,
		req:    &PrivateRestFuturesSettleOrdersGetReq{},
	}
}

func (api *PrivateRestFuturesSettleOrdersGetAPI) Do() (*GateRestRes[PrivateRestFuturesSettleOrdersGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleOrdersGet])
	return gateCallApiWithSecret[PrivateRestFuturesSettleOrdersGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestFuturesSettleOrdersTimeRange PrivateRest接口 GET 查询合约订单列表(时间区间)
func (client *PrivateRestClient) NewPrivateRestFuturesSettleOrdersTimeRange() *PrivateRestFuturesSettleOrdersTimeRangeAPI {
	return &PrivateRestFuturesSettleOrdersTimeRangeAPI{
		client: client,
		req:    &PrivateRestFuturesSettleOrdersTimeRangeReq{},
	}
}

func (api *PrivateRestFuturesSettleOrdersTimeRangeAPI) Do() (*GateRestRes[PrivateRestFuturesSettleOrdersTimeRangeRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleOrdersTimeRange])
	return gateCallApiWithSecret[PrivateRestFuturesSettleOrdersTimeRangeRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestFuturesSettleOrdersOrderIdGet PrivateRest接口 GET 查询单个订单详情
func (client *PrivateRestClient) NewPrivateRestFuturesSettleOrdersOrderIdGet() *PrivateRestFuturesSettleOrdersOrderIdGetAPI {
	return &PrivateRestFuturesSettleOrdersOrderIdGetAPI{
		client: client,
		req:    &PrivateRestFuturesSettleOrdersOrderIdGetReq{},
	}
}

func (api *PrivateRestFuturesSettleOrdersOrderIdGetAPI) Do() (*GateRestRes[PrivateRestFuturesSettleOrdersOrderIdGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleOrdersOrderIdGet])
	return gateCallApiWithSecret[PrivateRestFuturesSettleOrdersOrderIdGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestFuturesSettleOrdersOrderIdDelete PrivateRest接口 DELETE 撤销单个订单
func (client *PrivateRestClient) NewPrivateRestFuturesSettleOrdersOrderIdDelete() *PrivateRestFuturesSettleOrdersOrderIdDeleteAPI {
	return &PrivateRestFuturesSettleOrdersOrderIdDeleteAPI{
		client: client,
		req:    &PrivateRestFuturesSettleOrdersOrderIdDeleteReq{},
	}
}

func (api *PrivateRestFuturesSettleOrdersOrderIdDeleteAPI) Do() (*GateRestRes[PrivateRestFuturesSettleOrdersOrderIdDeleteRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleOrdersOrderIdDelete])
	return gateCallApiWithSecret[PrivateRestFuturesSettleOrdersOrderIdDeleteRes](api.client.c, url, NIL_REQBODY, DELETE)
}

// gate PrivateRestFuturesSettleOrdersOrderIdPut PrivateRest接口 PUT 修改单个订单
func (client *PrivateRestClient) NewPrivateRestFuturesSettleOrdersOrderIdPut() *PrivateRestFuturesSettleOrdersOrderIdPutAPI {
	return &PrivateRestFuturesSettleOrdersOrderIdPutAPI{
		client: client,
		req:    &PrivateRestFuturesSettleOrdersOrderIdPutReq{},
	}
}

func (api *PrivateRestFuturesSettleOrdersOrderIdPutAPI) Do() (*GateRestRes[PrivateRestFuturesSettleOrdersOrderIdPutRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleOrdersOrderIdPut])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestFuturesSettleOrdersOrderIdPutRes](api.client.c, url, reqBody, PUT)
}

// gate PrivateRestFuturesSettleMyTrades PrivateRest接口 GET 查询个人成交记录
func (client *PrivateRestClient) NewPrivateRestFuturesSettleMyTrades() *PrivateRestFuturesSettleMyTradesAPI {
	return &PrivateRestFuturesSettleMyTradesAPI{
		client: client,
		req:    &PrivateRestFuturesSettleMyTradesReq{},
	}
}

func (api *PrivateRestFuturesSettleMyTradesAPI) Do() (*GateRestRes[PrivateRestFuturesSettleMyTradesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettleMyTrades])
	return gateCallApiWithSecret[PrivateRestFuturesSettleMyTradesRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestFuturesSettlePriceOrdersPost PrivateRest接口 POST 创建价格触发订单
func (client *PrivateRestClient) NewPrivateRestFuturesSettlePriceOrdersPost() *PrivateRestFuturesSettlePriceOrdersPostAPI {
	return &PrivateRestFuturesSettlePriceOrdersPostAPI{
		client: client,
		req:    &PrivateRestFuturesSettlePriceOrdersPostReq{},
	}
}

func (api *PrivateRestFuturesSettlePriceOrdersPostAPI) Do() (*GateRestRes[PrivateRestFuturesSettlePriceOrdersPostRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettlePriceOrdersPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestFuturesSettlePriceOrdersPostRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestFuturesSettlePriceOrdersGet PrivateRest接口 GET 查询进行中自动订单列表
func (client *PrivateRestClient) NewPrivateRestFuturesSettlePriceOrdersGet() *PrivateRestFuturesSettlePriceOrdersGetAPI {
	return &PrivateRestFuturesSettlePriceOrdersGetAPI{
		client: client,
		req:    &PrivateRestFuturesSettlePriceOrdersGetReq{},
	}
}

func (api *PrivateRestFuturesSettlePriceOrdersGetAPI) Do() (*GateRestRes[PrivateRestFuturesSettlePriceOrdersGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettlePriceOrdersGet])
	return gateCallApiWithSecret[PrivateRestFuturesSettlePriceOrdersGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestFuturesSettlePriceOrdersDelete PrivateRest接口 DELETE 批量取消自动订单
func (client *PrivateRestClient) NewPrivateRestFuturesSettlePriceOrdersDelete() *PrivateRestFuturesSettlePriceOrdersDeleteAPI {
	return &PrivateRestFuturesSettlePriceOrdersDeleteAPI{
		client: client,
		req:    &PrivateRestFuturesSettlePriceOrdersDeleteReq{},
	}
}

func (api *PrivateRestFuturesSettlePriceOrdersDeleteAPI) Do() (*GateRestRes[PrivateRestFuturesSettlePriceOrdersDeleteRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettlePriceOrdersDelete])
	return gateCallApiWithSecret[PrivateRestFuturesSettlePriceOrdersDeleteRes](api.client.c, url, NIL_REQBODY, DELETE)
}


// gate PrivateRestFuturesSettlePriceOrdersOrderIdGet PrivateRest接口 GET 查询单个自动订单详情
func (client *PrivateRestClient) NewPrivateRestFuturesSettlePriceOrdersOrderIdGet() *PrivateRestFuturesSettlePriceOrdersOrderIdGetAPI {
	return &PrivateRestFuturesSettlePriceOrdersOrderIdGetAPI{
		client: client,
		req:    &PrivateRestFuturesSettlePriceOrdersOrderIdGetReq{},
	}
}

func (api *PrivateRestFuturesSettlePriceOrdersOrderIdGetAPI) Do() (*GateRestRes[PrivateRestFuturesSettlePriceOrdersOrderIdGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettlePriceOrdersOrderIdGet])
	return gateCallApiWithSecret[PrivateRestFuturesSettlePriceOrdersOrderIdGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestFuturesSettlePriceOrdersOrderIdDelete PrivateRest接口 DELETE 撤销单个自动订单
func (client *PrivateRestClient) NewPrivateRestFuturesSettlePriceOrdersOrderIdDelete() *PrivateRestFuturesSettlePriceOrdersOrderIdDeleteAPI {
	return &PrivateRestFuturesSettlePriceOrdersOrderIdDeleteAPI{
		client: client,
		req:    &PrivateRestFuturesSettlePriceOrdersOrderIdDeleteReq{},
	}
}

func (api *PrivateRestFuturesSettlePriceOrdersOrderIdDeleteAPI) Do() (*GateRestRes[PrivateRestFuturesSettlePriceOrdersOrderIdDeleteRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestFuturesSettlePriceOrdersOrderIdDelete])
	return gateCallApiWithSecret[PrivateRestFuturesSettlePriceOrdersOrderIdDeleteRes](api.client.c, url, NIL_REQBODY, DELETE)
}
