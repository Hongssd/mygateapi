package mygateapi

// gate PublicRestSpotCurrencies PublicRest接口 GET 查询所有币种信息
func (client *PublicRestClient) NewPublicRestSpotCurrencies() *PublicRestSpotCurrenciesAPI {
	return &PublicRestSpotCurrenciesAPI{
		client: client,
		req:    &PublicRestSpotCurrenciesReq{},
	}
}
func (api *PublicRestSpotCurrenciesAPI) Do() (*GateRestRes[PublicRestSpotCurrenciesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestSpotCurrencies])
	return gateCallAPI[PublicRestSpotCurrenciesRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestSpotCurrenciesCurrency PublicRest接口 GET 查询单个币种信息
func (client *PublicRestClient) NewPublicRestSpotCurrenciesCurrency() *PublicRestSpotCurrenciesCurrencyAPI {
	return &PublicRestSpotCurrenciesCurrencyAPI{
		client: client,
		req:    &PublicRestSpotCurrenciesCurrencyReq{},
	}
}
func (api *PublicRestSpotCurrenciesCurrencyAPI) Do() (*GateRestRes[PublicRestSpotCurrenciesCurrencyRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestSpotCurrenciesCurrency])
	return gateCallAPI[PublicRestSpotCurrenciesCurrencyRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestSpotCurrencyPairs PublicRest接口 GET 查询支持的所有交易对
func (client *PublicRestClient) NewPublicRestCurrencyPairsAll() *PublicRestSpotCurrencyPairsAllAPI {
	return &PublicRestSpotCurrencyPairsAllAPI{
		client: client,
		req:    &PublicRestSpotCurrencyPairsAllReq{},
	}
}

func (api *PublicRestSpotCurrencyPairsAllAPI) Do() (*GateRestRes[PublicRestSpotCurrencyPairsAllRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestSpotCurrencyPairs])
	return gateCallAPI[PublicRestSpotCurrencyPairsAllRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestSpotCurrencyPairsCurrencyPair PublicRest接口 GET 查询单个交易对详情
func (client *PublicRestClient) NewPublicRestCurrencyPairsCurrencyPair() *PublicRestSpotCurrencyPairsCurrencyPairAPI {
	return &PublicRestSpotCurrencyPairsCurrencyPairAPI{
		client: client,
		req:    &PublicRestSpotCurrencyPairsCurrencyPairReq{},
	}
}

func (api *PublicRestSpotCurrencyPairsCurrencyPairAPI) Do() (*GateRestRes[PublicRestSpotCurrencyPairsCurrencyPairRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestSpotCurrencyPairsCurrencyPair])
	return gateCallAPI[PublicRestSpotCurrencyPairsCurrencyPairRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestSpotTickers PublicRest接口 GET 获取交易对tickers信息
// 如果指定 currency_pair 则只查询该交易对，否则返回全部信息
func (client *PublicRestClient) NewPublicRestTickers() *PublicRestSpotTickersAPI {
	return &PublicRestSpotTickersAPI{
		client: client,
		req:    &PublicRestSpotTickersReq{},
	}
}

func (api *PublicRestSpotTickersAPI) Do() (*GateRestRes[PublicRestSpotTickersRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestSpotTickers])
	return gateCallAPI[PublicRestSpotTickersRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestSpotOrderBook PublicRest接口 GET 获取市场深度信息
func (client *PublicRestClient) NewPublicRestSpotOrderBook() *PublicRestSpotOrderBookAPI {
	return &PublicRestSpotOrderBookAPI{
		client: client,
		req:    &PublicRestSpotOrderBookReq{},
	}
}

func (api *PublicRestSpotOrderBookAPI) Do() (*GateRestRes[PublicRestSpotOrderBookRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestSpotOrderBook])
	return gateCallAPI[PublicRestSpotOrderBookRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestSpotTrades PublicRest接口 GET 查询市场成交记录
func (client *PublicRestClient) NewPublicRestSpotTrades() *PublicRestSpotTradesAPI {
	return &PublicRestSpotTradesAPI{
		client: client,
		req:    &PublicRestSpotTradesReq{},
	}
}

func (api *PublicRestSpotTradesAPI) Do() (*GateRestRes[PublicRestSpotTradesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestSpotTrades])
	return gateCallAPI[PublicRestSpotTradesRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestSpotCandlesticks PublicRest接口 GET 市场K线图
func (client *PublicRestClient) NewPublicRestSpotCandlesticks() *PublicRestSpotCandlesticksAPI {
	return &PublicRestSpotCandlesticksAPI{
		client: client,
		req:    &PublicRestSpotCandlesticksReq{},
	}
}

func (api *PublicRestSpotCandlesticksAPI) Do() (*GateRestRes[PublicRestSpotCandlesticksRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestSpotCandlesticks])
	middleRes, err := gateCallAPI[PublicRestSpotCandlesticksMiddle](api.client.c, url, NIL_REQBODY, GET)
	if err != nil {
		return nil, err
	}
	res := &GateRestRes[PublicRestSpotCandlesticksRes]{
		GateTimeRes:  middleRes.GateTimeRes,
		GateErrorRes: middleRes.GateErrorRes,
		Data:         *middleRes.Data.ConvertToRes(),
	}
	return res, nil
}

// gate PublicRestSpotTime PublicRest接口 GET 获取服务器时间
func (client *PublicRestClient) NewPublicRestSpotTime() *PublicRestSpotTimeAPI {
	return &PublicRestSpotTimeAPI{
		client: client,
		req:    &PublicRestSpotTimeReq{},
	}
}

func (api *PublicRestSpotTimeAPI) Do() (*GateRestRes[PublicRestSpotTimeRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestSpotTime])
	return gateCallAPI[PublicRestSpotTimeRes](api.client.c, url, NIL_REQBODY, GET)
}
