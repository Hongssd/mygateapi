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

// gate PrivateRestWalletTransfers PrivateRest接口 POST 交易账户互转
func (client *PrivateRestClient) NewPrivateRestWalletTransfers() *PrivateRestWalletTransfersAPI {
	return &PrivateRestWalletTransfersAPI{
		client: client,
		req:    &PrivateRestWalletTransfersReq{},
	}
}

func (api *PrivateRestWalletTransfersAPI) Do() (*GateRestRes[PrivateRestWalletTransfersRes], error) {
	url := gateHandlerRequestAPIWithoutPathQueryParam(REST, PrivateRestAPIMap[PrivateRestWalletTransfers])
	log.Info(url)
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return gateCallApiWithSecret[PrivateRestWalletTransfersRes](api.client.c, url, reqBody, POST)
}
