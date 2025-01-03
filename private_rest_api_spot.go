package mygateapi

// gate PrivateRestSpotAccounts PrivateRest接口 GET 获取现货交易账户列表
func (client *PrivateRestClient) NewPrivateRestSpotInstruments() *PrivateRestSpotAccountsAPI {
	return &PrivateRestSpotAccountsAPI{
		client: client,
		req:    &PrivateRestSpotAccountsReq{},
	}
}
func (api *PrivateRestSpotAccountsAPI) Do() (*GateRestRes[PrivateRestSpotAccountsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestSpotAccounts])
	return gateCallApiWithSecret[PrivateRestSpotAccountsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestSpotOrdersPost PrivateRest接口 POST 下单
func (client *PrivateRestClient) NewPrivateRestSpotOrdersPost() *PrivateRestSpotOrdersPostAPI {
	return &PrivateRestSpotOrdersPostAPI{
		client: client,
		req:    &PrivateRestSpotOrdersPostReq{},
	}
}
func (api *PrivateRestSpotOrdersPostAPI) Do() (*GateRestRes[PrivateRestSpotOrdersPostRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestSpotOrdersPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestSpotOrdersPostRes](api.client.c, url, reqBody, POST)
}

// gate PrivateRestSpotAccountBook PrivateRest接口 GET 查询现货账户变动历史
func (client *PrivateRestClient) NewPrivateRestSpotAccountBook() *PrivateRestSpotAccountBookAPI {
	return &PrivateRestSpotAccountBookAPI{
		client: client,
		req:    &PrivateRestSpotAccountBookReq{},
	}
}

func (api *PrivateRestSpotAccountBookAPI) Do() (*GateRestRes[PrivateRestSpotAccountBookRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestSpotAccountBook])
	return gateCallApiWithSecret[PrivateRestSpotAccountBookRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestSpotOpenOrders PrivateRest接口 GET 查询所有挂单
func (client *PrivateRestClient) NewPrivateRestSpotOpenOrders() *PrivateRestSpotOpenOrdersAPI {
	return &PrivateRestSpotOpenOrdersAPI{
		client: client,
		req:    &PrivateRestSpotOpenOrdersReq{},
	}
}

func (api *PrivateRestSpotOpenOrdersAPI) Do() (*GateRestRes[PrivateRestSpotOpenOrdersRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestSpotOpenOrders])
	return gateCallApiWithSecret[PrivateRestSpotOpenOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestSpotOrdersGet PrivateRest接口 GET 查询订单列表
func (client *PrivateRestClient) NewPrivateRestSpotOrdersGet() *PrivateRestSpotOrdersGetAPI {
	return &PrivateRestSpotOrdersGetAPI{
		client: client,
		req:    &PrivateRestSpotOrdersGetReq{},
	}
}

func (api *PrivateRestSpotOrdersGetAPI) Do() (*GateRestRes[PrivateRestSpotOrdersGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestSpotOrdersGet])
	return gateCallApiWithSecret[PrivateRestSpotOrdersGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestSpotOrdersOrderIdGet PrivateRest接口 GET 查询单个订单详情
func (client *PrivateRestClient) NewPrivateRestSpotOrdersOrderIdGet() *PrivateRestSpotOrdersOrderIdGetAPI {
	return &PrivateRestSpotOrdersOrderIdGetAPI{
		client: client,
		req:    &PrivateRestSpotOrdersOrderIdGetReq{},
	}
}

func (api *PrivateRestSpotOrdersOrderIdGetAPI) Do() (*GateRestRes[PrivateRestSpotOrdersOrderIdGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestSpotOrdersOrderIdGet])
	return gateCallApiWithSecret[PrivateRestSpotOrdersOrderIdGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestSpotOrdersOrderIdPatch PrivateRest接口 PATCH 修改单个订单
func (client *PrivateRestClient) NewPrivateRestSpotOrdersOrderIdPatch() *PrivateRestSpotOrdersOrderIdPatchAPI {
	return &PrivateRestSpotOrdersOrderIdPatchAPI{
		client: client,
		req:    &PrivateRestSpotOrdersOrderIdPatchReq{},
	}
}

func (api *PrivateRestSpotOrdersOrderIdPatchAPI) Do() (*GateRestRes[PrivateRestSpotOrdersOrderIdPatchRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestSpotOrdersOrderIdPatch])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestSpotOrdersOrderIdPatchRes](api.client.c, url, reqBody, PATCH)
}

// gate PrivateRestSpotOrdersOrderIdDelete PrivateRest接口 DELETE 撤销单个订单
func (client *PrivateRestClient) NewPrivateRestSpotOrdersOrderIdDelete() *PrivateRestSpotOrdersOrderIdDeleteAPI {
	return &PrivateRestSpotOrdersOrderIdDeleteAPI{
		client: client,
		req:    &PrivateRestSpotOrdersOrderIdDeleteReq{},
	}
}

func (api *PrivateRestSpotOrdersOrderIdDeleteAPI) Do() (*GateRestRes[PrivateRestSpotOrdersOrderIdDeleteRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestSpotOrdersOrderIdDelete])
	return gateCallApiWithSecret[PrivateRestSpotOrdersOrderIdDeleteRes](api.client.c, url, NIL_REQBODY, DELETE)
}

// gate PrivateRestSpotMyTrades PrivateRest接口 GET 查询个人成交记录
func (client *PrivateRestClient) NewPrivateRestSpotMyTrades() *PrivateRestSpotMyTradesAPI {
	return &PrivateRestSpotMyTradesAPI{
		client: client,
		req:    &PrivateRestSpotMyTradesReq{},
	}
}

func (api *PrivateRestSpotMyTradesAPI) Do() (*GateRestRes[PrivateRestSpotMyTradesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestSpotMyTrades])
	return gateCallApiWithSecret[PrivateRestSpotMyTradesRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestSpotPriceOrdersPost PrivateRest接口 POST 创建价格触发订单
func (client *PrivateRestClient) NewPrivateRestSpotPriceOrdersPost() *PrivateRestSpotPriceOrdersPostAPI {
	return &PrivateRestSpotPriceOrdersPostAPI{
		client: client,
		req:    &PrivateRestSpotPriceOrdersPostReq{},
	}
}

func (api *PrivateRestSpotPriceOrdersPostAPI) Do() (*GateRestRes[PrivateRestSpotPriceOrdersPostRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestSpotPriceOrdersPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestSpotPriceOrdersPostRes](api.client.c, url, reqBody, POST)
}
