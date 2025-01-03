package mygateapi

// gate PublicRestDeliverySettleContracts PrivateRest接口 GET 获取所有的合约信息
func (client *PublicRestClient) NewPublicRestDeliverySettleContracts() *PublicRestDeliverySettleContractsAPI {
	return &PublicRestDeliverySettleContractsAPI{
		client: client,
		req:    &PublicRestDeliverySettleContractsReq{},
	}
}

func (api *PublicRestDeliverySettleContractsAPI) Do() (*GateRestRes[PublicRestDeliverySettleContractsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestDeliverySettleContracts])
	return gateCallAPI[PublicRestDeliverySettleContractsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestDeliverySettleContractsContract PrivateRest接口 GET 获取合约信息
func (client *PublicRestClient) NewPublicRestDeliverySettleContractsContract() *PublicRestDeliverySettleContractsContractAPI {
	return &PublicRestDeliverySettleContractsContractAPI{
		client: client,
		req:    &PublicRestDeliverySettleContractsContractReq{},
	}
}

func (api *PublicRestDeliverySettleContractsContractAPI) Do() (*GateRestRes[PublicRestDeliverySettleContractsContractRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestDeliverySettleContractsContract])
	return gateCallAPI[PublicRestDeliverySettleContractsContractRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestDeliverySettleOrderBook PrivateRest接口 GET 查询合约市场深度信息
func (client *PublicRestClient) NewPublicRestDeliverySettleOrderBook() *PublicRestDeliverySettleOrderBookAPI {
	return &PublicRestDeliverySettleOrderBookAPI{
		client: client,
		req:    &PublicRestDeliverySettleOrderBookReq{},
	}
}

func (api *PublicRestDeliverySettleOrderBookAPI) Do() (*GateRestRes[PublicRestDeliverySettleOrderBookRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestDeliverySettleOrderBook])
	return gateCallAPI[PublicRestDeliverySettleOrderBookRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestDeliverySettleTrades PrivateRest接口 GET 合约市场成交记录
func (client *PublicRestClient) NewPublicRestDeliverySettleTrades() *PublicRestDeliverySettleTradesAPI {
	return &PublicRestDeliverySettleTradesAPI{
		client: client,
		req:    &PublicRestDeliverySettleTradesReq{},
	}
}

func (api *PublicRestDeliverySettleTradesAPI) Do() (*GateRestRes[PublicRestDeliverySettleTradesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestDeliverySettleTrades])
	return gateCallAPI[PublicRestDeliverySettleTradesRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestDeliverySettleCandlesticks PrivateRest接口 GET 合约市场K线数据
func (client *PublicRestClient) NewPublicRestDeliverySettleCandlesticks() *PublicRestDeliverySettleCandlesticksAPI {
	return &PublicRestDeliverySettleCandlesticksAPI{
		client: client,
		req:    &PublicRestDeliverySettleCandlesticksReq{},
	}
}

func (api *PublicRestDeliverySettleCandlesticksAPI) Do() (*GateRestRes[PublicRestDeliverySettleCandlesticksRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestDeliverySettleCandlesticks])
	return gateCallAPI[PublicRestDeliverySettleCandlesticksRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestDeliverySettleTickers PrivateRest接口 GET 获取所有合约交易行情统计
func (client *PublicRestClient) NewPublicRestDeliverySettleTickers() *PublicRestDeliverySettleTickersAPI {
	return &PublicRestDeliverySettleTickersAPI{
		client: client,
		req:    &PublicRestDeliverySettleTickersReq{},
	}
}

func (api *PublicRestDeliverySettleTickersAPI) Do() (*GateRestRes[PublicRestDeliverySettleTickersRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestDeliverySettleTickers])
	return gateCallAPI[PublicRestDeliverySettleTickersRes](api.client.c, url, NIL_REQBODY, GET)
}
