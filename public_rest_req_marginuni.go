package mygateapi

type PublicRestMarginUniCurrencyPairsReq struct{}

type PublicRestMarginUniCurrencyPairsAPI struct {
	client *PublicRestClient
	req    *PublicRestMarginUniCurrencyPairsReq
}

type PublicRestMarginUniCurrencyPairsCurrencyPairReq struct {
	CurrencyPair *string `json:"currency_pair"` //交易对
}

// 是 string 交易对
func (api *PublicRestMarginUniCurrencyPairsCurrencyPairAPI) CurrencyPair(currencyPair string) *PublicRestMarginUniCurrencyPairsCurrencyPairAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

type PublicRestMarginUniCurrencyPairsCurrencyPairAPI struct {
	client *PublicRestClient
	req    *PublicRestMarginUniCurrencyPairsCurrencyPairReq
}
