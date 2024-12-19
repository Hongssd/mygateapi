package mygateapi

type PrivateRestMarginAccountsReq struct {
	CurrencyPair *string `json:"currency_pair"` // currency_pair	请求参数	string	否	交易对
}

// currency_pair	请求参数	string	否	交易对
func (api *PrivateRestMarginAccountsAPI) SetCurrencyPair(currencyPair string) *PrivateRestMarginAccountsAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

type PrivateRestMarginAccountsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestMarginAccountsReq
}

type PrivateRestMarginAccountBookReq struct {
	Currency     *string `json:"currency"`      // currency	请求参数	string	否	指定币种查询历史，如果指定 currency ，必须同时指定 currency_pair
	CurrencyPair *string `json:"currency_pair"` // currency_pair	请求参数	string	否	指定杠杆账户交易对，该字段与 currency 配合使用，如果不指定 currency，该字段忽略
	Type         *string `json:"type"`          // type	请求参数	string	否	指定账户变动类型查询，不指定则包含全部变动类型
	From         *int64  `json:"from"`          // from	请求参数	integer(int64)	否	查询记录的起始时间
	To           *int64  `json:"to"`            // to	请求参数	integer(int64)	否	查询记录的结束时间，不指定则默认为当前时间
	Page         *int    `json:"page"`          // page	请求参数	integer(int32)	否	列表页数
	Limit        *int    `json:"limit"`         // limit	请求参数	integer	否	列表返回的最大数量
}

// currency	请求参数	string	否	指定币种查询历史，如果指定 currency ，必须同时指定 currency_pair
func (api *PrivateRestMarginAccountBookAPI) Currency(currency string) *PrivateRestMarginAccountBookAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

// currency_pair	请求参数	string	否	指定杠杆账户交易对，该字段与 currency 配合使用，如果不指定 currency，该字段忽略
func (api *PrivateRestMarginAccountBookAPI) CurrencyPair(currencyPair string) *PrivateRestMarginAccountBookAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

// type	请求参数	string	否	指定账户变动类型查询，不指定则包含全部变动类型
func (api *PrivateRestMarginAccountBookAPI) Type(type_ string) *PrivateRestMarginAccountBookAPI {
	api.req.Type = GetPointer(type_)
	return api
}

// from	请求参数	integer(int64)	否	查询记录的起始时间
func (api *PrivateRestMarginAccountBookAPI) From(from int64) *PrivateRestMarginAccountBookAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	integer(int64)	否	查询记录的结束时间，不指定则默认为当前时间
func (api *PrivateRestMarginAccountBookAPI) To(to int64) *PrivateRestMarginAccountBookAPI {
	api.req.To = GetPointer(to)
	return api
}

// page	请求参数	integer(int32)	否	列表页数
func (api *PrivateRestMarginAccountBookAPI) Page(page int) *PrivateRestMarginAccountBookAPI {
	api.req.Page = GetPointer(page)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestMarginAccountBookAPI) Limit(limit int) *PrivateRestMarginAccountBookAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PrivateRestMarginAccountBookAPI struct {
	client *PrivateRestClient
	req    *PrivateRestMarginAccountBookReq
}

type PrivateRestMarginTransferableReq struct {
	Currency     *string `json:"currency"`      // currency	请求参数	string	是	指定币种名称查询
	CurrencyPair *string `json:"currency_pair"` // currency_pair	请求参数	string	否	交易对
}

// currency	请求参数	string	是	指定币种名称查询
func (api *PrivateRestMarginTransferableAPI) Currency(currency string) *PrivateRestMarginTransferableAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

// currency_pair	请求参数	string	否	交易对
func (api *PrivateRestMarginTransferableAPI) CurrencyPair(currencyPair string) *PrivateRestMarginTransferableAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

type PrivateRestMarginTransferableAPI struct {
	client *PrivateRestClient
	req    *PrivateRestMarginTransferableReq
}

type PrivateRestMarginCrossTransferableReq struct {
	Currency *string `json:"currency"` // currency	请求参数	string	是	指定币种名称查询
}

// currency	请求参数	string	是	指定币种名称查询
func (api *PrivateRestMarginCrossTransferableAPI) Currency(currency string) *PrivateRestMarginCrossTransferableAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

type PrivateRestMarginCrossTransferableAPI struct {
	client *PrivateRestClient
	req    *PrivateRestMarginCrossTransferableReq
}
