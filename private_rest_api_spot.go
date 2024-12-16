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
