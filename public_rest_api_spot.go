package mygateapi

// gate PublicRestSpotCurrencies PublicRest接口 GET 查询所有币种信息
func (client *PublicRestClient) NewPublicRestSpotCurrencies() *PublicRestSpotCurrenciesAPI {
	return &PublicRestSpotCurrenciesAPI{
		client: client,
		req:    &PublicRestSpotCurrenciesReq{},
	}
}
func (api *PublicRestSpotCurrenciesAPI) Do() (*GateRestRes[PublicRestSpotCurrenciesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicSpotCurrencies])
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
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicSpotCurrenciesCurrency])
	return gateCallAPI[PublicRestSpotCurrenciesCurrencyRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestSpotCurrencyPairs PublicRest接口 GET 查询支持的所有交易对
func (client *PublicRestClient) NewPublicRestPublicCurrencyPairsAll() *PublicRestSpotCurrencyPairsAllAPI {
	return &PublicRestSpotCurrencyPairsAllAPI{
		client: client,
		req:    &PublicRestSpotCurrencyPairsAllReq{},
	}
}

func (api *PublicRestSpotCurrencyPairsAllAPI) Do() (*GateRestRes[PublicRestSpotCurrencyPairsAllRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicSpotCurrencyPairsAll])
	return gateCallAPI[PublicRestSpotCurrencyPairsAllRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicSpotCurrencyPairsSingle PublicRest接口 GET 查询单个交易对详情
func (client *PublicRestClient) NewPublicRestPublicCurrencyPairsSingle() *PublicRestSpotCurrencyPairsSingleAPI {
	return &PublicRestSpotCurrencyPairsSingleAPI{
		client: client,
		req:    &PublicRestSpotCurrencyPairsSingleReq{},
	}
}

func (api *PublicRestSpotCurrencyPairsSingleAPI) Do() (*GateRestRes[PublicRestSpotCurrencyPairsSingleRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicSpotCurrencyPairsSingle])
	return gateCallAPI[PublicRestSpotCurrencyPairsSingleRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicSpotTickers PublicRest接口 GET 获取交易对tickers信息
// 如果指定 currency_pair 则只查询该交易对，否则返回全部信息
func (client *PublicRestClient) NewPublicRestPublicTickers() *PublicRestSpotTickersAPI {
	return &PublicRestSpotTickersAPI{
		client: client,
		req:    &PublicRestSpotTickersReq{},
	}
}

func (api *PublicRestSpotTickersAPI) Do() (*GateRestRes[PublicRestSpotTickersRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicSpotTickers])
	return gateCallAPI[PublicRestSpotTickersRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicSpotOrderBook PublicRest接口 GET 获取市场深度信息
func (client *PublicRestClient) NewPublicRestPublicOrderBook() *PublicRestSpotOrderBookAPI {
	return &PublicRestSpotOrderBookAPI{
		client: client,
		req:    &PublicRestSpotOrderBookReq{},
	}
}

func (api *PublicRestSpotOrderBookAPI) Do() (*GateRestRes[PublicRestSpotOrderBookRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicSpotOrderBook])
	return gateCallAPI[PublicRestSpotOrderBookRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicSpotTrades PublicRest接口 GET 查询市场成交记录
func (client *PublicRestClient) NewPublicRestPublicTrades() *PublicRestSpotTradesAPI {
	return &PublicRestSpotTradesAPI{
		client: client,
		req:    &PublicRestSpotTradesReq{},
	}
}

func (api *PublicRestSpotTradesAPI) Do() (*GateRestRes[PublicRestSpotTradesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicSpotTrades])
	return gateCallAPI[PublicRestSpotTradesRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicSpotCandlesticks PublicRest接口 GET 市场K线图
func (client *PublicRestClient) NewPublicRestPublicCandlesticks() *PublicRestSpotCandlesticksAPI {
	return &PublicRestSpotCandlesticksAPI{
		client: client,
		req:    &PublicRestSpotCandlesticksReq{},
	}
}

func (api *PublicRestSpotCandlesticksAPI) Do() (*GateRestRes[PublicRestSpotCandlesticksRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicSpotCandlesticks])
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
