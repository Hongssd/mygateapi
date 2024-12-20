package mygateapi

type PublicRestFuturesSettleContractsReq struct {
	Settle *string `json:"settle"` // settle	URL	string	是	结算货币
	Limit  *int    `json:"limit"`  // limit	请求参数	integer	否	列表返回的最大数量
	Offset *int    `json:"offset"` // offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PublicRestFuturesSettleContractsAPI) Settle(settle string) *PublicRestFuturesSettleContractsAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PublicRestFuturesSettleContractsAPI) Limit(limit int) *PublicRestFuturesSettleContractsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
func (api *PublicRestFuturesSettleContractsAPI) Offset(offset int) *PublicRestFuturesSettleContractsAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

type PublicRestFuturesSettleContractsAPI struct {
	client *PublicRestClient
	req    *PublicRestFuturesSettleContractsReq
}

type PublicRestFuturesSettleContractsContractReq struct {
	Settle   *string `json:"settle"`   // settle	URL	string	是	结算货币
	Contract *string `json:"contract"` // contract	URL	string	是	合约标识
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PublicRestFuturesSettleContractsContractAPI) Settle(settle string) *PublicRestFuturesSettleContractsContractAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PublicRestFuturesSettleContractsContractAPI) Contract(contract string) *PublicRestFuturesSettleContractsContractAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

type PublicRestFuturesSettleContractsContractAPI struct {
	client *PublicRestClient
	req    *PublicRestFuturesSettleContractsContractReq
}

type PublicRestFuturesSettleOrderBookReq struct {
	Settle   *string `json:"settle"`   // settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
	Contract *string `json:"contract"` // contract	URL	string	是	合约标识
	Interval *string `json:"interval"` // interval	请求参数	string	否	合并深度指定的价格精度，0 为不合并，不指定则默认为 0
	Limit    *int    `json:"limit"`    // limit	请求参数	integer	否	深度档位数量
	WithId   *bool   `json:"with_id"`  // with_id	请求参数	boolean	否	是否返回深度更新 ID。深度每发生一次变化，该 ID 自增 1
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PublicRestFuturesSettleOrderBookAPI) Settle(settle string) *PublicRestFuturesSettleOrderBookAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PublicRestFuturesSettleOrderBookAPI) Contract(contract string) *PublicRestFuturesSettleOrderBookAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// interval	请求参数	string	否	合并深度指定的价格精度，0 为不合并，不指定则默认为 0
func (api *PublicRestFuturesSettleOrderBookAPI) Interval(interval string) *PublicRestFuturesSettleOrderBookAPI {
	api.req.Interval = GetPointer(interval)
	return api
}

// limit	请求参数	integer	否	深度档位数量
func (api *PublicRestFuturesSettleOrderBookAPI) Limit(limit int) *PublicRestFuturesSettleOrderBookAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// with_id	请求参数	boolean	否	是否返回深度更新 ID。深度每发生一次变化，该 ID 自增 1
func (api *PublicRestFuturesSettleOrderBookAPI) WithId(withId bool) *PublicRestFuturesSettleOrderBookAPI {
	api.req.WithId = GetPointer(withId)
	return api
}

type PublicRestFuturesSettleOrderBookAPI struct {
	client *PublicRestClient
	req    *PublicRestFuturesSettleOrderBookReq
}

type PublicRestFuturesSettleTradesReq struct {
	Settle   *string `json:"settle"`   // settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
	Contract *string `json:"contract"` // contract	URL	string	是	合约标识
	Limit    *int    `json:"limit"`    // limit	请求参数	integer	否	列表返回的最大数量
	Offset   *int    `json:"offset"`   // offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
	LastId   *string `json:"last_id"`  // last_id	请求参数	string	否	以上个列表的最后一条记录的 ID 作为下个列表的起点。
	From     *int64  `json:"from"`     // from	请求参数	integer(int64)	否	指定起始时间，时间格式为秒(s)精度的 Unix 时间戳。 不指定则按照 to 和 limit 来限定返回的数量。
	To       *int64  `json:"to"`       // to	请求参数	integer(int64)	否	指定结束时间，不指定则默认当前时间，时间格式为秒(s)精度的 Unix 时间戳
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PublicRestFuturesSettleTradesAPI) Settle(settle string) *PublicRestFuturesSettleTradesAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PublicRestFuturesSettleTradesAPI) Contract(contract string) *PublicRestFuturesSettleTradesAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PublicRestFuturesSettleTradesAPI) Limit(limit int) *PublicRestFuturesSettleTradesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// offset	请求参数	integer	否	列表返回的偏移量，从 0 开始
func (api *PublicRestFuturesSettleTradesAPI) Offset(offset int) *PublicRestFuturesSettleTradesAPI {
	api.req.Offset = GetPointer(offset)
	return api
}

// last_id	请求参数	string	否	以上个列表的最后一条记录的 ID 作为下个列表的起点。
func (api *PublicRestFuturesSettleTradesAPI) LastId(lastId string) *PublicRestFuturesSettleTradesAPI {
	api.req.LastId = GetPointer(lastId)
	return api
}

// from	请求参数	integer(int64)	否	指定起始时间，时间格式为秒(s)精度的 Unix 时间戳。 不指定则按照 to 和 limit 来限定返回的数量。
func (api *PublicRestFuturesSettleTradesAPI) From(from int64) *PublicRestFuturesSettleTradesAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	integer(int64)	否	指定结束时间，不指定则默认当前时间，时间格式为秒(s)精度的 Unix 时间戳
func (api *PublicRestFuturesSettleTradesAPI) To(to int64) *PublicRestFuturesSettleTradesAPI {
	api.req.To = GetPointer(to)
	return api
}

type PublicRestFuturesSettleTradesAPI struct {
	client *PublicRestClient
	req    *PublicRestFuturesSettleTradesReq
}

type PublicRestFuturesSettleCandlesticksReq struct {
	Settle   *string `json:"settle"`   // settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
	Contract *string `json:"contract"` // contract	请求参数	string	是	合约标识
	From     *int64  `json:"from"`     // from	请求参数	integer(int64)	否	指定 K 线图的起始时间，注意时间格式为秒(s)精度的 Unix 时间戳，不指定则默认为 to - 100 * interval，即向前最多 100 个点的时间
	To       *int64  `json:"to"`       // to	请求参数	integer(int64)	否	指定 K 线图的结束时间，不指定则默认当前时间，注意时间格式为秒(s)精度的 Unix 时间戳
	Limit    *int    `json:"limit"`    // limit	请求参数	integer	否	指定数据点的数量，适用于取最近 limit 数量的数据，该字段与 from, to 互斥，如果指定了 from, to 中的任意字段，该字段会被拒绝
	Interval *string `json:"interval"` // interval	请求参数	string	否	数据点的时间间隔，注意 1w 代表一个自然周，7d 的时间是和 Unix 初始时间对齐, 30d 代表一个自然月
}

// settle	URL	string	是	结算货币 例如：usdt（需小写，大写会返回resource not found）
func (api *PublicRestFuturesSettleCandlesticksAPI) Settle(settle string) *PublicRestFuturesSettleCandlesticksAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PublicRestFuturesSettleCandlesticksAPI) Contract(contract string) *PublicRestFuturesSettleCandlesticksAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// from	请求参数	integer(int64)	否	指定 K 线图的起始时间，注意时间格式为秒(s)精度的 Unix 时间戳，不指定则默认为 to - 100 * interval，即向前最多 100 个点的时间
func (api *PublicRestFuturesSettleCandlesticksAPI) From(from int64) *PublicRestFuturesSettleCandlesticksAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	integer(int64)	否	指定 K 线图的结束时间，不指定则默认当前时间，注意时间格式为秒(s)精度的 Unix 时间戳
func (api *PublicRestFuturesSettleCandlesticksAPI) To(to int64) *PublicRestFuturesSettleCandlesticksAPI {
	api.req.To = GetPointer(to)
	return api
}

// limit	请求参数	integer	否	指定数据点的数量，适用于取最近 limit 数量的数据，该字段与 from, to 互斥，如果指定了 from, to 中的任意字段，该字段会被拒绝
func (api *PublicRestFuturesSettleCandlesticksAPI) Limit(limit int) *PublicRestFuturesSettleCandlesticksAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

type PublicRestFuturesSettleCandlesticksAPI struct {
	client *PublicRestClient
	req    *PublicRestFuturesSettleCandlesticksReq
}
