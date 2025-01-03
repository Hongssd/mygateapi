package mygateapi

// gate PublicRestFuturesSettleContracts PublicRest接口 GET 查询所有的合约信息
func (client *PublicRestClient) NewPublicRestFuturesSettleContracts() *PublicRestFuturesSettleContractsAPI {
	return &PublicRestFuturesSettleContractsAPI{
		client: client,
		req:    &PublicRestFuturesSettleContractsReq{},
	}
}

func (api *PublicRestFuturesSettleContractsAPI) Do() (*GateRestRes[PublicRestFuturesSettleContractsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestFuturesSettleContracts])
	return gateCallAPI[PublicRestFuturesSettleContractsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestFuturesSettleContractsContract PublicRest接口 GET 查询单个合约信息
func (client *PublicRestClient) NewPublicRestFuturesSettleContractsContract() *PublicRestFuturesSettleContractsContractAPI {
	return &PublicRestFuturesSettleContractsContractAPI{
		client: client,
		req:    &PublicRestFuturesSettleContractsContractReq{},
	}
}

func (api *PublicRestFuturesSettleContractsContractAPI) Do() (*GateRestRes[PublicRestFuturesSettleContractsContractRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestFuturesSettleContractsContract])
	return gateCallAPI[PublicRestFuturesSettleContractsContractRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestFuturesSettleOrderBook PublicRest接口 GET 查询合约市场深度信息
func (client *PublicRestClient) NewPublicRestFuturesSettleOrderBook() *PublicRestFuturesSettleOrderBookAPI {
	return &PublicRestFuturesSettleOrderBookAPI{
		client: client,
		req:    &PublicRestFuturesSettleOrderBookReq{},
	}
}

func (api *PublicRestFuturesSettleOrderBookAPI) Do() (*GateRestRes[PublicRestFuturesSettleOrderBookRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestFuturesSettleOrderBook])
	return gateCallAPI[PublicRestFuturesSettleOrderBookRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestFuturesSettleTrades PublicRest接口 GET 合约市场成交记录
func (client *PublicRestClient) NewPublicRestFuturesSettleTrades() *PublicRestFuturesSettleTradesAPI {
	return &PublicRestFuturesSettleTradesAPI{
		client: client,
		req:    &PublicRestFuturesSettleTradesReq{},
	}
}

func (api *PublicRestFuturesSettleTradesAPI) Do() (*GateRestRes[PublicRestFuturesSettleTradesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestFuturesSettleTrades])
	return gateCallAPI[PublicRestFuturesSettleTradesRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestFuturesSettleCandlesticks PublicRest接口 GET 合约市场K线图
func (client *PublicRestClient) NewPublicRestFuturesSettleCandlesticks() *PublicRestFuturesSettleCandlesticksAPI {
	return &PublicRestFuturesSettleCandlesticksAPI{
		client: client,
		req:    &PublicRestFuturesSettleCandlesticksReq{},
	}
}

func (api *PublicRestFuturesSettleCandlesticksAPI) Do() (*GateRestRes[PublicRestFuturesSettleCandlesticksRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestFuturesSettleCandlesticks])
	return gateCallAPI[PublicRestFuturesSettleCandlesticksRes](api.client.c, url, NIL_REQBODY, GET)
}
