package mygateapi

// gate PrivateRestSpotAccounts PrivateRest接口 GET 获取现货交易账户列表
func (client *PrivateRestClient) NewPrivateRestPrivateInstruments() *PrivateRestSpotAccountsAPI {
	return &PrivateRestSpotAccountsAPI{
		client: client,
		req:    &PrivateRestSpotAccountsReq{},
	}
}
func (api *PrivateRestSpotAccountsAPI) Do() (*GateRestRes[PrivateRestSpotAccountsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateSpotAccounts])
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
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateSpotOrdersPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestSpotOrdersPostRes](api.client.c, url, reqBody, POST)
}

// gate PrivateSpotAccountBook PrivateRest接口 GET 查询现货账户变动历史
func (client *PrivateRestClient) NewPrivateRestPrivateAccountBook() *PrivateRestSpotAccountBookAPI {
	return &PrivateRestSpotAccountBookAPI{
		client: client,
		req:    &PrivateRestSpotAccountBookReq{},
	}
}

func (api *PrivateRestSpotAccountBookAPI) Do() (*GateRestRes[PrivateRestSpotAccountBookRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateSpotAccountBook])
	return gateCallApiWithSecret[PrivateRestSpotAccountBookRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateSpotOpenOrders PrivateRest接口 GET 查询所有挂单
func (client *PrivateRestClient) NewPrivateRestPrivateOpenOrders() *PrivateRestSpotOpenOrdersAPI {
	return &PrivateRestSpotOpenOrdersAPI{
		client: client,
		req:    &PrivateRestSpotOpenOrdersReq{},
	}
}

func (api *PrivateRestSpotOpenOrdersAPI) Do() (*GateRestRes[PrivateRestSpotOpenOrdersRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateSpotOpenOrders])
	return gateCallApiWithSecret[PrivateRestSpotOpenOrdersRes](api.client.c, url, NIL_REQBODY, GET)
}
