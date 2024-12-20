package mygateapi

// gate PrivateRestAccountDetail PrivateRest接口 GET 获取用户账户信息
func (client *PrivateRestClient) NewPrivateRestAccountDetail() *PrivateRestAccountDetailAPI {
	return &PrivateRestAccountDetailAPI{
		client: client,
		req:    &PrivateRestAccountDetailReq{},
	}
}

func (api *PrivateRestAccountDetailAPI) Do() (*GateRestRes[PrivateRestAccountDetailRes], error) {
	url := gateHandlerRequestAPIWithPathQueryParam(REST, api.req, PrivateRestAPIMap[PrivateRestAccountDetail])
	return gateCallApiWithSecret[PrivateRestAccountDetailRes](api.client.c, url, NIL_REQBODY, GET)
}
