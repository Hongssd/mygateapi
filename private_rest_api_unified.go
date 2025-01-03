package mygateapi

// gate PrivateRestUnifiedAccounts PrivateRest接口 GET 获取统一账户信息
func (client *PrivateRestClient) NewPrivateRestUnifiedAccounts() *PrivateRestUnifiedAccountsAPI {
	return &PrivateRestUnifiedAccountsAPI{
		client: client,
		req:    &PrivateRestUnifiedAccountsReq{},
	}
}

func (api *PrivateRestUnifiedAccountsAPI) Do() (*GateRestRes[PrivateRestUnifiedAccountsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUnifiedAccounts])
	return gateCallApiWithSecret[PrivateRestUnifiedAccountsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestUnifiedTransferable PrivateRest接口 GET 查询统一账户最多可转出
func (client *PrivateRestClient) NewPrivateRestUnifiedTransferable() *PrivateRestUnifiedTransferableAPI {
	return &PrivateRestUnifiedTransferableAPI{
		client: client,
		req:    &PrivateRestUnifiedTransferableReq{},
	}
}

func (api *PrivateRestUnifiedTransferableAPI) Do() (*GateRestRes[PrivateRestUnifiedTransferableRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUnifiedTransferable])
	return gateCallApiWithSecret[PrivateRestUnifiedTransferableRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestUnifiedRiskUnits PrivateRest接口 GET 获取用户风险单元详情（仅在组合保证金模式有效）
func (client *PrivateRestClient) NewPrivateRestUnifiedRiskUnits() *PrivateRestUnifiedRiskUnitsAPI {
	return &PrivateRestUnifiedRiskUnitsAPI{
		client: client,
		req:    &PrivateRestUnifiedRiskUnitsReq{},
	}
}

func (api *PrivateRestUnifiedRiskUnitsAPI) Do() (*GateRestRes[PrivateRestUnifiedRiskUnitsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUnifiedRiskUnits])
	return gateCallApiWithSecret[PrivateRestUnifiedRiskUnitsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestUnifiedUnifiedModePut PrivateRest接口 PUT 设置统一账户模式
func (client *PrivateRestClient) NewPrivateRestUnifiedUnifiedModePut() *PrivateRestUnifiedUnifiedModePutAPI {
	return &PrivateRestUnifiedUnifiedModePutAPI{
		client: client,
		req:    &PrivateRestUnifiedUnifiedModePutReq{},
	}
}

func (api *PrivateRestUnifiedUnifiedModePutAPI) Do() (*GateRestRes[PrivateRestUnifiedUnifiedModePutRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUnifiedUnifiedModePut])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestUnifiedUnifiedModePutRes](api.client.c, url, reqBody, PUT)
}

// gate PrivateRestUnifiedUnifiedModeGet PrivateRest接口 GET 获取统一账户模式
func (client *PrivateRestClient) NewPrivateRestUnifiedUnifiedModeGet() *PrivateRestUnifiedUnifiedModeGetAPI {
	return &PrivateRestUnifiedUnifiedModeGetAPI{
		client: client,
		req:    &PrivateRestUnifiedUnifiedModeGetReq{},
	}
}

func (api *PrivateRestUnifiedUnifiedModeGetAPI) Do() (*GateRestRes[PrivateRestUnifiedUnifiedModeGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUnifiedUnifiedModeGet])
	return gateCallApiWithSecret[PrivateRestUnifiedUnifiedModeGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestUnifiedLeverageUserCurrencyConfig PrivateRest接口 GET 用户最大、最小可设置币种杠杆倍数
func (client *PrivateRestClient) NewPrivateRestUnifiedLeverageUserCurrencyConfig() *PrivateRestUnifiedLeverageUserCurrencyConfigAPI {
	return &PrivateRestUnifiedLeverageUserCurrencyConfigAPI{
		client: client,
		req:    &PrivateRestUnifiedLeverageUserCurrencyConfigReq{},
	}
}

func (api *PrivateRestUnifiedLeverageUserCurrencyConfigAPI) Do() (*GateRestRes[PrivateRestUnifiedLeverageUserCurrencyConfigRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUnifiedLeverageUserCurrencyConfig])
	return gateCallApiWithSecret[PrivateRestUnifiedLeverageUserCurrencyConfigRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestUnifiedLeverageUserCurrencySettingGet PrivateRest接口 GET 获取用户币种杠杆倍数，currency不传则查询全部币种
func (client *PrivateRestClient) NewPrivateRestUnifiedLeverageUserCurrencySettingGet() *PrivateRestUnifiedLeverageUserCurrencySettingGetAPI {
	return &PrivateRestUnifiedLeverageUserCurrencySettingGetAPI{
		client: client,
		req:    &PrivateRestUnifiedLeverageUserCurrencySettingGetReq{},
	}
}

func (api *PrivateRestUnifiedLeverageUserCurrencySettingGetAPI) Do() (*GateRestRes[PrivateRestUnifiedLeverageUserCurrencySettingGetRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUnifiedLeverageUserCurrencySettingGet])
	return gateCallApiWithSecret[PrivateRestUnifiedLeverageUserCurrencySettingGetRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestUnifiedLeverageUserCurrencySettingPost PrivateRest接口 POST 设置币种杠杆倍数
func (client *PrivateRestClient) NewPrivateRestUnifiedLeverageUserCurrencySettingPost() *PrivateRestUnifiedLeverageUserCurrencySettingPostAPI {
	return &PrivateRestUnifiedLeverageUserCurrencySettingPostAPI{
		client: client,
		req:    &PrivateRestUnifiedLeverageUserCurrencySettingPostReq{},
	}
}

func (api *PrivateRestUnifiedLeverageUserCurrencySettingPostAPI) Do() (*GateRestRes[PrivateRestUnifiedLeverageUserCurrencySettingPostRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestUnifiedLeverageUserCurrencySettingPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestUnifiedLeverageUserCurrencySettingPostRes](api.client.c, url, reqBody, POST)
}
