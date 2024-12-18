package mygateapi

type PublicRestSpotCurrenciesReq struct {
}

type PublicRestSpotCurrenciesAPI struct {
	client *PublicRestClient
	req    *PublicRestSpotCurrenciesReq
}

type PublicRestSpotCurrenciesCurrencyReq struct {
	Currency *string `json:"currency"`
	Test     *string `json:"test"`
}

type PublicRestSpotCurrenciesCurrencyAPI struct {
	client *PublicRestClient
	req    *PublicRestSpotCurrenciesCurrencyReq
}

func (api *PublicRestSpotCurrenciesCurrencyAPI) Currency(currency string) *PublicRestSpotCurrenciesCurrencyAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

func (api *PublicRestSpotCurrenciesCurrencyAPI) Test(test string) *PublicRestSpotCurrenciesCurrencyAPI {
	api.req.Test = GetPointer(test)
	return api
}

type PublicRestSpotCurrencyPairsAllReq struct{}

type PublicRestSpotCurrencyPairsAllAPI struct {
	client *PublicRestClient
	req    *PublicRestSpotCurrencyPairsAllReq
}

type PublicRestSpotCurrencyPairsSingleReq struct {
	CurrencyPair *string `json:"currency_pair"`
}

// currency_pair	请求参数	string	是	交易对名称
func (api *PublicRestSpotCurrencyPairsSingleAPI) CurrencyPair(currencyPair string) *PublicRestSpotCurrencyPairsSingleAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

type PublicRestSpotCurrencyPairsSingleAPI struct {
	client *PublicRestClient
	req    *PublicRestSpotCurrencyPairsSingleReq
}

type PublicRestSpotTickersReq struct {
	CurrencyPair *string `json:"currency_pair"`
	Timezone     *string `json:"timezone"`
}

// currency_pair	请求参数	string	否	交易对
func (api *PublicRestSpotTickersAPI) CurrencyPair(currencyPair string) *PublicRestSpotTickersAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

// timezone	请求参数	string	否	时区
func (api *PublicRestSpotTickersAPI) Timezone(timezone string) *PublicRestSpotTickersAPI {
	api.req.Timezone = GetPointer(timezone)
	return api
}

type PublicRestSpotTickersAPI struct {
	client *PublicRestClient
	req    *PublicRestSpotTickersReq
}

type PublicRestSpotOrderBookReq struct {
	CurrencyPair *string `json:"currency_pair"`
	Interval     *string `json:"interval"`
	Limit        *int    `json:"limit"`
	WithId       *bool   `json:"with_id"`
}

// currency_pair	请求参数	string	是	交易对
func (api *PublicRestSpotOrderBookAPI) CurrencyPair(currencyPair string) *PublicRestSpotOrderBookAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

// interval	请求参数	string	否	深度间隔
func (api *PublicRestSpotOrderBookAPI) Interval(interval string) *PublicRestSpotOrderBookAPI {
	api.req.Interval = GetPointer(interval)
	return api
}

// limit	请求参数	int	否	返回深度数量
func (api *PublicRestSpotOrderBookAPI) Limit(limit int) *PublicRestSpotOrderBookAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// with_id	请求参数	boolean	否	是否返回深度ID
func (api *PublicRestSpotOrderBookAPI) WithId(withId bool) *PublicRestSpotOrderBookAPI {
	api.req.WithId = GetPointer(withId)
	return api
}

type PublicRestSpotOrderBookAPI struct {
	client *PublicRestClient
	req    *PublicRestSpotOrderBookReq
}

type PublicRestSpotTradesReq struct {
	CurrencyPair *string `json:"currency_pair"`
	Limit        *int    `json:"limit"`
	LastId       *string `json:"last_id"`
	Reverse      *bool   `json:"reverse"`
	From         *int64  `json:"from"`
	To           *int64  `json:"to"`
	Page         *int    `json:"page"`
}

// currency_pair	请求参数	string	是	交易对
func (api *PublicRestSpotTradesAPI) CurrencyPair(currencyPair string) *PublicRestSpotTradesAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

// limit	请求参数	integer(int32)	否	列表返回的最大数量。默认为100，最小1，最大1000。
func (api *PublicRestSpotTradesAPI) Limit(limit int) *PublicRestSpotTradesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// last_id	请求参数	string	否	以上个列表的最后一条记录的 ID 作为下个列表的起点
func (api *PublicRestSpotTradesAPI) LastId(lastId string) *PublicRestSpotTradesAPI {
	api.req.LastId = GetPointer(lastId)
	return api
}

// reverse	请求参数	boolean	否	是否获取小于 last_id 的数据，默认返回大于 last_id 的记录。
func (api *PublicRestSpotTradesAPI) Reverse(reverse bool) *PublicRestSpotTradesAPI {
	api.req.Reverse = GetPointer(reverse)
	return api
}

// from	请求参数	integer(int64)	否	查询记录的起始时间
func (api *PublicRestSpotTradesAPI) From(from int64) *PublicRestSpotTradesAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	integer(int64)	否	查询记录的结束时间，不指定则默认为当前时间
func (api *PublicRestSpotTradesAPI) To(to int64) *PublicRestSpotTradesAPI {
	api.req.To = GetPointer(to)
	return api
}

// page	请求参数	integer(int32)	否	列表页数
func (api *PublicRestSpotTradesAPI) Page(page int) *PublicRestSpotTradesAPI {
	api.req.Page = GetPointer(page)
	return api
}

type PublicRestSpotTradesAPI struct {
	client *PublicRestClient
	req    *PublicRestSpotTradesReq
}

type PublicRestSpotCandlesticksReq struct {
	CurrencyPair *string `json:"currency_pair"`
	Limit        *int    `json:"limit"`
	From         *int64  `json:"from"`
	To           *int64  `json:"to"`
	Interval     *string `json:"interval"`
}

// currency_pair	请求参数	string	是	交易对
func (api *PublicRestSpotCandlesticksAPI) CurrencyPair(currencyPair string) *PublicRestSpotCandlesticksAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

// limit	请求参数	integer	否	指定数据点的数量，适用于取最近 limit 数量的数据，该字段与 from, to 互斥，如果指定了 from, to 中的任意字段，该字段会被拒绝
func (api *PublicRestSpotCandlesticksAPI) Limit(limit int) *PublicRestSpotCandlesticksAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// from	请求参数	integer(int64)	否	指定 K 线图的起始时间，注意时间格式为秒(s)精度的 Unix 时间戳，不指定则默认为 to - 100 * interval，即向前最多 100 个点的时间
func (api *PublicRestSpotCandlesticksAPI) From(from int64) *PublicRestSpotCandlesticksAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	integer(int64)	否	指定 K 线图的结束时间，不指定则默认当前时间，注意时间格式为秒(s)精度的 Unix 时间戳
func (api *PublicRestSpotCandlesticksAPI) To(to int64) *PublicRestSpotCandlesticksAPI {
	api.req.To = GetPointer(to)
	return api
}

// interval	请求参数	string	否	数据点的时间间隔， 注意 30d 代表的是自然月，不是按30天对齐
func (api *PublicRestSpotCandlesticksAPI) Interval(interval string) *PublicRestSpotCandlesticksAPI {
	api.req.Interval = GetPointer(interval)
	return api
}

type PublicRestSpotCandlesticksAPI struct {
	client *PublicRestClient
	req    *PublicRestSpotCandlesticksReq
}

type PublicRestSpotFeeReq struct {
	CurrencyPair *string `json:"currency_pair"`
}
