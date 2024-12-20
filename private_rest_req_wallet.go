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

type PrivateRestWalletTransfersReq struct {
	Currency     *string `json:"currency"`      //转账货币名称。关联合约账户时，currency 可以设置的值为POINT(即点卡) 和支持的结算货币(如 BTC, USDT)
	From         *string `json:"from"`          //转出账户
	To           *string `json:"to"`            //转入账户
	Amount       *string `json:"amount"`        //转账额度
	CurrencyPair *string `json:"currency_pair"` //杠杆交易对。转入或转出杠杆账户时必填
	Settle       *string `json:"settle"`        //合约结算币种。 转入转出合约账户时必填
}

// currency	请求参数	string	否	转账货币名称。关联合约账户时，currency 可以设置的值为POINT(即点卡) 和支持的结算货币(如 BTC, USDT)
func (api *PrivateRestWalletTransfersAPI) Currency(currency string) *PrivateRestWalletTransfersAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

// from	请求参数	string	否	转出账户
func (api *PrivateRestWalletTransfersAPI) From(from string) *PrivateRestWalletTransfersAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	string	否	转入账户
func (api *PrivateRestWalletTransfersAPI) To(to string) *PrivateRestWalletTransfersAPI {
	api.req.To = GetPointer(to)
	return api
}

// amount	请求参数	string	否	转账额度
func (api *PrivateRestWalletTransfersAPI) Amount(amount string) *PrivateRestWalletTransfersAPI {
	api.req.Amount = GetPointer(amount)
	return api
}

// currency_pair	请求参数	string	否	杠杆交易对。转入或转出杠杆账户时必填
func (api *PrivateRestWalletTransfersAPI) CurrencyPair(currencyPair string) *PrivateRestWalletTransfersAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

// settle	请求参数	string	否	合约结算币种。 转入转出合约账户时必填
func (api *PrivateRestWalletTransfersAPI) Settle(settle string) *PrivateRestWalletTransfersAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

type PrivateRestWalletTransfersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestWalletTransfersReq
}
