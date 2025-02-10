package mygateapi

// gate PublicRestMarginUniCurrencyPairs PublicRest接口 GET 查询借贷市场列表
func (client *PublicRestClient) NewPublicRestMarginUniCurrencyPairs() *PublicRestMarginUniCurrencyPairsAPI {
	return &PublicRestMarginUniCurrencyPairsAPI{
		client: client,
		req:    &PublicRestMarginUniCurrencyPairsReq{},
	}
}

func (api *PublicRestMarginUniCurrencyPairsAPI) Do() (*GateRestRes[PublicRestMarginUniCurrencyPairsRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarginUniCurrencyPairs])
	return gateCallAPI[PublicRestMarginUniCurrencyPairsRes](api.client.c, url, NIL_REQBODY, GET)
}

// gate PublicRestMarginUniCurrencyPairsCurrencyPair PublicRest接口 GET 查询借贷市场详情
func (client *PublicRestClient) NewPublicRestMarginUniCurrencyPairsCurrencyPair() *PublicRestMarginUniCurrencyPairsCurrencyPairAPI {
	return &PublicRestMarginUniCurrencyPairsCurrencyPairAPI{
		client: client,
		req:    &PublicRestMarginUniCurrencyPairsCurrencyPairReq{},
	}
}

func (api *PublicRestMarginUniCurrencyPairsCurrencyPairAPI) Do() (*GateRestRes[PublicRestMarginUniCurrencyPairsCurrencyPairRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarginUniCurrencyPairsCurrencyPair])
	return gateCallAPI[PublicRestMarginUniCurrencyPairsCurrencyPairRes](api.client.c, url, NIL_REQBODY, GET)
}
