package mygateapi

type PublicRestDeliverySettleContractsReq struct {
	Settle *string `json:"settle"` //结算货币
}

// settle	URL	string	是	结算货币
func (api *PublicRestDeliverySettleContractsAPI) Settle(settle string) *PublicRestDeliverySettleContractsAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

type PublicRestDeliverySettleContractsAPI struct {
	client *PublicRestClient
	req    *PublicRestDeliverySettleContractsReq
}

type PublicRestDeliverySettleContractsContractReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
}

// settle	URL	string	是	结算货币
func (api *PublicRestDeliverySettleContractsContractAPI) Settle(settle string) *PublicRestDeliverySettleContractsContractAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	URL	string	是	合约标识
func (api *PublicRestDeliverySettleContractsContractAPI) Contract(contract string) *PublicRestDeliverySettleContractsContractAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

type PublicRestDeliverySettleContractsContractAPI struct {
	client *PublicRestClient
	req    *PublicRestDeliverySettleContractsContractReq
}

// settle	URL	string	是	结算货币
// contract	请求参数	string	是	合约标识
// interval	请求参数	string	否	合并深度指定的价格精度，0 为不合并，不指定则默认为 0
// limit	请求参数	integer	否	深度档位数量
// with_id	请求参数	boolean	否	是否返回深度更新 ID。深度每发生一次变化，该 ID 自增 1
type PublicRestDeliverySettleOrderBookReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
	Interval *string `json:"interval"` //合并深度指定的价格精度，0 为不合并，不指定则默认为 0
	Limit    *int    `json:"limit"`    //深度档位数量
	WithId   *bool   `json:"with_id"`  //是否返回深度更新 ID。深度每发生一次变化，该 ID 自增 1
}

// settle	URL	string	是	结算货币
func (api *PublicRestDeliverySettleOrderBookAPI) Settle(settle string) *PublicRestDeliverySettleOrderBookAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PublicRestDeliverySettleOrderBookAPI) Contract(contract string) *PublicRestDeliverySettleOrderBookAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// interval	请求参数	string	否	合并深度指定的价格精度，0 为不合并，不指定则默认为 0
func (api *PublicRestDeliverySettleOrderBookAPI) Interval(interval string) *PublicRestDeliverySettleOrderBookAPI {
	api.req.Interval = GetPointer(interval)
	return api
}

// limit	请求参数	integer	否	深度档位数量
func (api *PublicRestDeliverySettleOrderBookAPI) Limit(limit int) *PublicRestDeliverySettleOrderBookAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// with_id	请求参数	boolean	否	是否返回深度更新 ID。深度每发生一次变化，该 ID 自增 1
func (api *PublicRestDeliverySettleOrderBookAPI) WithId(withId bool) *PublicRestDeliverySettleOrderBookAPI {
	api.req.WithId = GetPointer(withId)
	return api
}

type PublicRestDeliverySettleOrderBookAPI struct {
	client *PublicRestClient
	req    *PublicRestDeliverySettleOrderBookReq
}

type PublicRestDeliverySettleTradesReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
	Limit    *int    `json:"limit"`    //列表返回的最大数量
	LastId   *string `json:"last_id"`  //以上个列表的最后一条记录的 ID 作为下个列表的起点。
	From     *int64  `json:"from"`     //指定起始时间，时间格式为秒(s)精度的 Unix 时间戳。 不指定则按照 to 和 limit 来限定返回的数量。
	To       *int64  `json:"to"`       //指定结束时间，不指定则默认当前时间，时间格式为秒(s)精度的 Unix 时间戳
}

// settle	URL	string	是	结算货币
func (api *PublicRestDeliverySettleTradesAPI) Settle(settle string) *PublicRestDeliverySettleTradesAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PublicRestDeliverySettleTradesAPI) Contract(contract string) *PublicRestDeliverySettleTradesAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PublicRestDeliverySettleTradesAPI) Limit(limit int) *PublicRestDeliverySettleTradesAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// last_id	请求参数	string	否	以上个列表的最后一条记录的 ID 作为下个列表的起点。
func (api *PublicRestDeliverySettleTradesAPI) LastId(lastId string) *PublicRestDeliverySettleTradesAPI {
	api.req.LastId = GetPointer(lastId)
	return api
}

// from	请求参数	int64	否	指定起始时间，时间格式为秒(s)精度的 Unix 时间戳。 不指定则按照 to 和 limit 来限定返回的数量。
func (api *PublicRestDeliverySettleTradesAPI) From(from int64) *PublicRestDeliverySettleTradesAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	int64	否	指定结束时间，不指定则默认当前时间，时间格式为秒(s)精度的 Unix 时间戳
func (api *PublicRestDeliverySettleTradesAPI) To(to int64) *PublicRestDeliverySettleTradesAPI {
	api.req.To = GetPointer(to)
	return api
}

type PublicRestDeliverySettleTradesAPI struct {
	client *PublicRestClient
	req    *PublicRestDeliverySettleTradesReq
}

type PublicRestDeliverySettleCandlesticksReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
	From     *int64  `json:"from"`     //指定 K 线图的起始时间，注意时间格式为秒(s)精度的 Unix 时间戳，不指定则默认为 to - 100 * interval，即向前最多 100 个点的时间
	To       *int64  `json:"to"`       //指定 K 线图的结束时间，不指定则默认当前时间，注意时间格式为秒(s)精度的 Unix 时间戳
	Limit    *int    `json:"limit"`    //指定数据点的数量，适用于取最近 limit 数量的数据，该字段与 from, to 互斥，如果指定了 from, to 中的任意字段，该字段会被拒绝
	Interval *string `json:"interval"` //数据点的时间间隔，注意 1w 代表一个自然周，7d 的时间是和 Unix 初始时间对齐
}

// settle	URL	string	是	结算货币
func (api *PublicRestDeliverySettleCandlesticksAPI) Settle(settle string) *PublicRestDeliverySettleCandlesticksAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PublicRestDeliverySettleCandlesticksAPI) Contract(contract string) *PublicRestDeliverySettleCandlesticksAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// from	请求参数	int64	否	指定 K 线图的起始时间，注意时间格式为秒(s)精度的 Unix 时间戳，不指定则默认为 to - 100 * interval，即向前最多 100 个点的时间
func (api *PublicRestDeliverySettleCandlesticksAPI) From(from int64) *PublicRestDeliverySettleCandlesticksAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	int64	否	指定 K 线图的结束时间，不指定则默认当前时间，注意时间格式为秒(s)精度的 Unix 时间戳
func (api *PublicRestDeliverySettleCandlesticksAPI) To(to int64) *PublicRestDeliverySettleCandlesticksAPI {
	api.req.To = GetPointer(to)
	return api
}

// limit	请求参数	int	否	指定数据点的数量，适用于取最近 limit 数量的数据，该字段与 from, to 互斥，如果指定了 from, to 中的任意字段，该字段会被拒绝
func (api *PublicRestDeliverySettleCandlesticksAPI) Limit(limit int) *PublicRestDeliverySettleCandlesticksAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// interval	请求参数	string	否	数据点的时间间隔，注意 1w 代表一个自然周，7d 的时间是和 Unix 初始时间对齐
func (api *PublicRestDeliverySettleCandlesticksAPI) Interval(interval string) *PublicRestDeliverySettleCandlesticksAPI {
	api.req.Interval = GetPointer(interval)
	return api
}

type PublicRestDeliverySettleCandlesticksAPI struct {
	client *PublicRestClient
	req    *PublicRestDeliverySettleCandlesticksReq
}

type PublicRestDeliverySettleTickersReq struct {
	Settle   *string `json:"settle"`   //结算货币
	Contract *string `json:"contract"` //合约标识
}

// settle	URL	string	是	结算货币
func (api *PublicRestDeliverySettleTickersAPI) Settle(settle string) *PublicRestDeliverySettleTickersAPI {
	api.req.Settle = GetPointer(settle)
	return api
}

// contract	请求参数	string	是	合约标识
func (api *PublicRestDeliverySettleTickersAPI) Contract(contract string) *PublicRestDeliverySettleTickersAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

type PublicRestDeliverySettleTickersAPI struct {
	client *PublicRestClient
	req    *PublicRestDeliverySettleTickersReq
}
