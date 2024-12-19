package mygateapi

// gate PrivateRestWalletTotalBalance PrivateRest接口 GET 查询个人账户总额
func (client *PrivateRestClient) NewPrivateRestWalletTotalBalance() *PrivateRestWalletTotalBalanceAPI {
	return &PrivateRestWalletTotalBalanceAPI{
		client: client,
		req:    &PrivateRestWalletTotalBalanceReq{},
	}
}

func (api *PrivateRestWalletTotalBalanceAPI) Do() (*GateRestRes[PrivateRestWalletTotalBalanceRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestWalletTotalBalance])
	return gateCallApiWithSecret[PrivateRestWalletTotalBalanceRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PrivateRestWalletFee PrivateRest接口 GET 查询个人交易费率
func (client *PrivateRestClient) NewPrivateRestWalletFee() *PrivateRestWalletFeeAPI {
	return &PrivateRestWalletFeeAPI{
		client: client,
		req:    &PrivateRestWalletFeeReq{},
	}
}

func (api *PrivateRestWalletFeeAPI) Do() (*GateRestRes[PrivateRestWalletFeeRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestWalletFee])
	return gateCallApiWithSecret[PrivateRestWalletFeeRes](api.client.c, url, NIL_REQBODY, GET)
}