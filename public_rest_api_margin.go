package mygateapi

// gate PublicRestMarginCrossCurrencies PublicRest接口 GET 全仓杠杆支持的币种列表
func (client *PublicRestClient) NewPublicRestMarginCrossCurrencies() *PublicRestMarginCrossCurrenciesAPI {
	return &PublicRestMarginCrossCurrenciesAPI{
		client: client,
		req:    &PublicRestMarginCrossCurrenciesReq{},
	}
}

func (api *PublicRestMarginCrossCurrenciesAPI) Do() (*GateRestRes[PublicRestMarginCrossCurrenciesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestMarginCrossCurrencies])
	return gateCallAPI[PublicRestMarginCrossCurrenciesRes](api.client.c, url, NIL_REQBODY, GET)
}
