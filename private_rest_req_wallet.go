package mygateapi

type PrivateRestWalletTotalBalanceReq struct {
	Currency *string `json:"currency"` //用于统计换算的目标货币类型，可接受BTC，CNY，USD，USDT四个值。USDT是默认值
}

// currency	返回参数	string	否	用于统计换算的目标货币类型，可接受BTC，CNY，USD，USDT四个值。USDT是默认值
func (api *PrivateRestWalletTotalBalanceAPI) Currency(currency string) *PrivateRestWalletTotalBalanceAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

type PrivateRestWalletTotalBalanceAPI struct {
	client *PrivateRestClient
	req    *PrivateRestWalletTotalBalanceReq
}

type PrivateRestWalletFeeReq struct {
	CurrencyPair *string `json:"currency_pair"`
	Settle       *string `json:"settle"`
}

// currency_pair	请求参数	string	否	指定交易对获取更准确的费率设置。
func (api *PrivateRestWalletFeeAPI) CurrencyPair(currencyPair string) *PrivateRestWalletFeeAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

// settle	请求参数	string	否	指定合约的结算币种获取更准确的费率设置。
func (api *PrivateRestWalletFeeAPI) Settle(settle string) *PrivateRestWalletFeeAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

type PrivateRestWalletFeeAPI struct {
	client *PrivateRestClient
	req    *PrivateRestWalletFeeReq
}
