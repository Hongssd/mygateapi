package mygateapi

// gate PublicRestSpotCurrencies PublicRest接口 GET 查询所有币种信息
func (client *PublicRestClient) NewPublicRestPublicInstruments() *PublicRestSpotCurrenciesAPI {
	return &PublicRestSpotCurrenciesAPI{
		client: client,
		req:    &PublicRestSpotCurrenciesReq{},
	}
}
func (api *PublicRestSpotCurrenciesAPI) Do() (*GateRestRes[PublicRestSpotCurrenciesRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicSpotCurrencies])
	return gateCallAPI[PublicRestSpotCurrenciesRes](api.client.c, url, NIL_REQBODY, GET)
}
