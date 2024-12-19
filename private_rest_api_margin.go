package mygateapi

// gate PrivateRestMarginAccounts PrivateRest接口 GET 获取杠杆交易账户列表
func (client *PrivateRestClient) NewPrivateRestMarginAccounts() *PrivateRestMarginAccountsAPI {
	return &PrivateRestMarginAccountsAPI{
		client: client,
		req:    &PrivateRestMarginAccountsReq{},
	}
}

func (api *PrivateRestMarginAccountsAPI) Do() (*GateRestRes[PrivateRestMarginAccountsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestMarginAccounts])
	return gateCallApiWithSecret[PrivateRestMarginAccountsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestMarginAccountBook PrivateRest接口 GET 查询杠杆账户变动历史
func (client *PrivateRestClient) NewPrivateRestMarginAccountBook() *PrivateRestMarginAccountBookAPI {
	return &PrivateRestMarginAccountBookAPI{
		client: client,
		req:    &PrivateRestMarginAccountBookReq{},
	}
}

func (api *PrivateRestMarginAccountBookAPI) Do() (*GateRestRes[PrivateRestMarginAccountBookRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestMarginAccountBook])
	return gateCallApiWithSecret[PrivateRestMarginAccountBookRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestMarginTransferable PrivateRest接口 GET 逐仓杠杆允许的最大转出
func (client *PrivateRestClient) NewPrivateRestMarginTransferable() *PrivateRestMarginTransferableAPI {
	return &PrivateRestMarginTransferableAPI{
		client: client,
		req:    &PrivateRestMarginTransferableReq{},
	}
}

func (api *PrivateRestMarginTransferableAPI) Do() (*GateRestRes[PrivateRestMarginTransferableRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestMarginTransferable])
	return gateCallApiWithSecret[PrivateRestMarginTransferableRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestMarginCrossTransferable PrivateRest接口 GET 全仓杠杆允许的最大转出
func (client *PrivateRestClient) NewPrivateRestMarginCrossTransferable() *PrivateRestMarginCrossTransferableAPI {
	return &PrivateRestMarginCrossTransferableAPI{
		client: client,
		req:    &PrivateRestMarginCrossTransferableReq{},
	}
}

func (api *PrivateRestMarginCrossTransferableAPI) Do() (*GateRestRes[PrivateRestMarginCrossTransferableRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestMarginCrossTransferable])
	return gateCallApiWithSecret[PrivateRestMarginCrossTransferableRes](api.client.c, url, NIL_REQBODY, GET)
}
