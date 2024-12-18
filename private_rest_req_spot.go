package mygateapi

import "github.com/shopspring/decimal"

type PrivateRestSpotAccountsReq struct {
	Currency *string `json:"currency"` //string	否	指定币种名称查询
}

type PrivateRestSpotAccountsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestSpotAccountsReq
}

// currency	请求参数	string	否	指定币种名称查询
func (api *PrivateRestSpotAccountsAPI) Currency(currency string) *PrivateRestSpotAccountsAPI {
	api.req.Currency = &currency
	return api
}

// 参数
// 名称	位置	类型	必选	描述
// x-gate-exptime	请求头部	integer(int64)	否	指定过期时间(毫秒); 如果 Gate 收到请求的时间大于过期时间, 请求将被拒绝
// body	body	Order	是
// » text	body	string	否	订单自定义信息，用户可以用该字段设置自定义 ID，用户自定义字段必须满足以下条件：
// » currency_pair	body	string	是	交易货币对
// » type	body	string	否	订单类型
// » account	body	string	否	账户类型，spot - 现货账户，margin - 杠杆账户，cross_margin - 全仓杠杆账户，unified - 统一账户
// » side	body	string	是	买单或者卖单
// » amount	body	string	是	交易数量
// » price	body	string	否	交易价,type=limit时必填
// » time_in_force	body	string	否	Time in force 策略。
// » iceberg	body	string	否	冰山下单显示的数量，不指定或传 0 都默认为普通下单。目前不支持全部冰山。
// » auto_borrow	body	boolean	否	杠杆(包括逐仓全仓)交易时，如果账户余额不足，是否由系统自动借入不足部分
// » auto_repay	body	boolean	否	全仓杠杆下单是否开启自动还款，默认关闭。需要注意的是:
// » stp_act	body	string	否	Self-Trading Prevention Action,用户可以用该字段设置自定义限制自成交策略。
// » action_mode	body	string	否	处理模式:
// 详细描述
// » text: 订单自定义信息，用户可以用该字段设置自定义 ID，用户自定义字段必须满足以下条件：
//
// 必须以 t- 开头
// 不计算 t- ，长度不能超过 28 字节
// 输入内容只能包含数字、字母、下划线(_)、中划线(-) 或者点(.)
// 除用户自定义信息以外，以下为内部保留字段，标识订单来源: 101 代表 android 下单 102 代表 IOS 下单 103 代表 IPAD 下单 104 代表 webapp 下单 3 代表 web 下单 2 代表 apiv2 下单 apiv4 代表 apiv4 下单
//
// » type: 订单类型
//
// limit : 限价单
// market : 市价单
// » account: 账户类型，spot - 现货账户，margin - 杠杆账户，cross_margin - 全仓杠杆账户，unified - 统一账户 统一账户（旧）只能设置 cross_margin
//
// » amount: 交易数量
// type为limit时，指交易货币，即需要交易的货币，如BTC_USDT中指BTC。
// type为market时，根据买卖不同指代不同
//
// side : buy 指代计价货币，BTC_USDT中指USDT
// side : sell 指代交易货币，BTC_USDT中指BTC
// » time_in_force: Time in force 策略。
//
// gtc: GoodTillCancelled
// ioc: ImmediateOrCancelled，立即成交或者取消，只吃单不挂单
// poc: PendingOrCancelled，被动委托，只挂单不吃单
// fok: FillOrKill，全部成交或者全部取消
// type=market时仅支持ioc和fok
//
// » auto_repay: 全仓杠杆下单是否开启自动还款，默认关闭。需要注意的是:
//
// 此字段仅针对全仓杠杆有效。逐仓杠杆不支持订单级别的自动还款设置，只能通过 POST /margin/auto_repay 修改用户级别的设置
// auto_borrow 与 auto_repay 支持同时开启
// » stp_act: Self-Trading Prevention Action,用户可以用该字段设置自定义限制自成交策略。
//
// 用户在设置加入STP用户组后，可以通过传递 stp_act 来限制用户发生自成交的策略，没有传递 stp_act 默认按照 cn 的策略。
// 用户在没有设置加入STP用户组时，传递 stp_act 参数会报错。
// 用户没有使用 stp_act 发生成交的订单，stp_act 返回 -。
// cn: Cancel newest,取消新订单，保留老订单
// co: Cancel oldest,取消⽼订单，保留新订单
// cb: Cancel both,新旧订单都取消
// » action_mode: 处理模式: 下单时根据action_mode返回不同的字段, 该字段只在请求时有效，响应结果中不包含该字段 ACK: 异步模式，只返回订单关键字段 RESULT: 无清算信息 FULL: 完整模式（默认）
//
// #枚举值列表
// 参数	值
// » type	limit
// » type	market
// » side	buy
// » side	sell
// » time_in_force	gtc
// » time_in_force	ioc
// » time_in_force	poc
// » time_in_force	fok
// » stp_act	cn
// » stp_act	co
// » stp_act	cb
// » stp_act	-
type PrivateRestSpotOrdersPostReq struct {
	Text         *string `json:"text,omitempty"`          //string	否	订单自定义信息，用户可以用该字段设置自定义 ID，用户自定义字段必须满足以下条件：必须以 t- 开头；不计算 t- ，长度不能超过 28 字节；输入内容只能包含数字、字母、下划线(_)、中划线(-) 或者点(.)；除用户自定义信息以外，以下为内部保留字段，标识订单来源: 101 代表 android 下单 102 代表 IOS 下单 103 代表 IPAD 下单 104 代表 webapp 下单 3 代表 web 下单 2 代表 apiv2 下单 apiv4 代表 apiv4 下单
	CurrencyPair *string `json:"currency_pair,omitempty"` //string	是	交易货币对
	Type         *string `json:"type,omitempty"`          //string	否	订单类型 limit : 限价单 market : 市价单
	Account      *string `json:"account,omitempty"`       //string	否	账户类型，spot - 现货账户，margin - 杠杆账户，cross_margin - 全仓杠杆账户，unified - 统一账户
	Side         *string `json:"side,omitempty"`          //string	是	买单或者卖单 buy sell
	Amount       *string `json:"amount,omitempty"`        //string	是	交易数量
	Price        *string `json:"price,omitempty"`         //string	否	交易价,type=limit时必填
	TimeInForce  *string `json:"time_in_force,omitempty"` //string	否	Time in force 策略。 gtc ioc poc fok type=market时仅支持ioc和fok
	Iceberg      *string `json:"iceberg,omitempty"`       //string	否	冰山下单显示的数量，不指定或传 0 都默认为普通下单。目前不支持全部冰山。
	AutoBorrow   *bool   `json:"auto_borrow,omitempty"`   //boolean	否	杠杆(包括逐仓全仓)交易时，如果账户余额不足，是否由系统自动借入不足部分
	AutoRepay    *bool   `json:"auto_repay,omitempty"`    //boolean	否	全仓杠杆下单是否开启自动还款，默认关闭。需要注意的是: auto_borrow 与 auto_repay 支持同时开启
	StpAct       *string `json:"stp_act,omitempty"`       //string	否	Self-Trading Prevention Action,用户可以用该字段设置自定义限制自成交策略。 cn co cb
	ActionMode   *string `json:"action_mode,omitempty"`   //string	否	处理模式: 下单时根据action_mode返回不同的字段, 该字段只在请求时有效，响应结果中不包含该字段 ACK: 异步模式，只返回订单关键字段 RESULT: 无清算信息 FULL: 完整模式（默认）
}

type PrivateRestSpotOrdersPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestSpotOrdersPostReq
}

// string 否 订单自定义信息，用户可以用该字段设置自定义 ID，用户自定义字段必须满足以下条件：必须以 t- 开头；不计算 t- ，长度不能超过 28 字节；输入内容只能包含数字、字母、下划线(_)、中划线(-) 或者点(.)；除用户自定义信息以外，以下为内部保留字段，标识订单来源: 101 代表 android 下单 102 代表 IOS 下单 103 代表 IPAD 下单 104 代表 webapp 下单 3 代表 web 下单 2 代表 apiv2 下单 apiv4 代表 apiv4 下单
func (api *PrivateRestSpotOrdersPostAPI) Text(text string) *PrivateRestSpotOrdersPostAPI {
	api.req.Text = GetPointer(text)
	return api
}

// string 是 交易货币对
func (api *PrivateRestSpotOrdersPostAPI) CurrencyPair(currencyPair string) *PrivateRestSpotOrdersPostAPI {
	api.req.CurrencyPair = GetPointer(currencyPair)
	return api
}

// string 否 订单类型 limit : 限价单 market : 市价单
func (api *PrivateRestSpotOrdersPostAPI) Type(t string) *PrivateRestSpotOrdersPostAPI {
	api.req.Type = GetPointer(t)
	return api
}

// string 否 账户类型，spot - 现货账户，margin - 杠杆账户，cross_margin - 全仓杠杆账户，unified - 统一账户
func (api *PrivateRestSpotOrdersPostAPI) Account(account string) *PrivateRestSpotOrdersPostAPI {
	api.req.Account = GetPointer(account)
	return api
}

// string 是 买单或者卖单 buy sell
func (api *PrivateRestSpotOrdersPostAPI) Side(side string) *PrivateRestSpotOrdersPostAPI {
	api.req.Side = GetPointer(side)
	return api
}

// string 是 交易数量
func (api *PrivateRestSpotOrdersPostAPI) Amount(amount decimal.Decimal) *PrivateRestSpotOrdersPostAPI {
	api.req.Amount = GetPointer(amount.String())
	return api
}

// string 否 交易价,type=limit时必填
func (api *PrivateRestSpotOrdersPostAPI) Price(price decimal.Decimal) *PrivateRestSpotOrdersPostAPI {
	api.req.Price = GetPointer(price.String())
	return api
}

// string 否 Time in force 策略。 gtc ioc poc fok type=market时仅支持ioc和fok
func (api *PrivateRestSpotOrdersPostAPI) TimeInForce(timeInForce string) *PrivateRestSpotOrdersPostAPI {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}

// string 否 冰山下单显示的数量，不指定或传 0 都默认为普通下单。目前不支持全部冰山。
func (api *PrivateRestSpotOrdersPostAPI) Iceberg(iceberg string) *PrivateRestSpotOrdersPostAPI {
	api.req.Iceberg = GetPointer(iceberg)
	return api
}

// boolean 否 杠杆(包括逐仓全仓)交易时，如果账户余额不足，是否由系统自动借入不足部分
func (api *PrivateRestSpotOrdersPostAPI) AutoBorrow(autoBorrow bool) *PrivateRestSpotOrdersPostAPI {
	api.req.AutoBorrow = GetPointer(autoBorrow)
	return api
}

// boolean 否 全仓杠杆下单是否开启自动还款，默认关闭。需要注意的是: auto_borrow 与 auto_repay 支持同时开启
func (api *PrivateRestSpotOrdersPostAPI) AutoRepay(autoRepay bool) *PrivateRestSpotOrdersPostAPI {
	api.req.AutoRepay = GetPointer(autoRepay)
	return api
}

// string 否 Self-Trading Prevention Action,用户可以用该字段设置自定义限制自成交策略。 cn co cb
func (api *PrivateRestSpotOrdersPostAPI) StpAct(stpAct string) *PrivateRestSpotOrdersPostAPI {
	api.req.StpAct = GetPointer(stpAct)
	return api
}

type PrivateRestSpotAccountBookReq struct {
	Currency *string `json:"currency"` //string	否	指定币种名称查询
	From     *int64  `json:"from"`     //integer(int64)	否	查询记录的起始时间
	To       *int64  `json:"to"`       //integer(int64)	否	查询记录的结束时间，不指定则默认为当前时间
	Page     *int    `json:"page"`     //integer(int32)	否	列表页数
	Limit    *int    `json:"limit"`    //integer	否	列表返回的最大数量
	Type     *string `json:"type"`     //string	否	指定账户变动类型查询，不指定则包含全部变动类型
}

// currency	请求参数	string	否	指定币种名称查询
func (api *PrivateRestSpotAccountBookAPI) Currency(currency string) *PrivateRestSpotAccountBookAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

// from	请求参数	integer(int64)	否	查询记录的起始时间
func (api *PrivateRestSpotAccountBookAPI) From(from int64) *PrivateRestSpotAccountBookAPI {
	api.req.From = GetPointer(from)
	return api
}

// to	请求参数	integer(int64)	否	查询记录的结束时间，不指定则默认为当前时间
func (api *PrivateRestSpotAccountBookAPI) To(to int64) *PrivateRestSpotAccountBookAPI {
	api.req.To = GetPointer(to)
	return api
}

// page	请求参数	integer(int32)	否	列表页数
func (api *PrivateRestSpotAccountBookAPI) Page(page int) *PrivateRestSpotAccountBookAPI {
	api.req.Page = GetPointer(page)
	return api
}

// limit	请求参数	integer	否	列表返回的最大数量
func (api *PrivateRestSpotAccountBookAPI) Limit(limit int) *PrivateRestSpotAccountBookAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// type	请求参数	string	否	指定账户变动类型查询，不指定则包含全部变动类型
func (api *PrivateRestSpotAccountBookAPI) Type(t string) *PrivateRestSpotAccountBookAPI {
	api.req.Type = GetPointer(t)
	return api
}

type PrivateRestSpotAccountBookAPI struct {
	client *PrivateRestClient
	req    *PrivateRestSpotAccountBookReq
}

type PrivateRestSpotOpenOrdersReq struct {
	Page    *int    `json:"page"`    //integer(int32)	否	列表页数
	Limit   *int    `json:"limit"`   //integer	否	每个交易对每页最多返回的数量
	Account *string `json:"account"` //string	否	指定查询账户。不指定默认现货，保证金和逐仓杠杆账户。指定 cross_margin 则查询全仓杠杆账户。
}

// page	请求参数	integer(int32)	否	列表页数
func (api *PrivateRestSpotOpenOrdersAPI) Page(page int) *PrivateRestSpotOpenOrdersAPI {
	api.req.Page = GetPointer(page)
	return api
}

// limit	请求参数	integer	否	每个交易对每页最多返回的数量
func (api *PrivateRestSpotOpenOrdersAPI) Limit(limit int) *PrivateRestSpotOpenOrdersAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// account	请求参数	string	否	指定查询账户。不指定默认现货，保证金和逐仓杠杆账户。指定 cross_margin 则查询全仓杠杆账户。
func (api *PrivateRestSpotOpenOrdersAPI) Account(account string) *PrivateRestSpotOpenOrdersAPI {
	api.req.Account = GetPointer(account)
	return api
}

type PrivateRestSpotOpenOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestSpotOpenOrdersReq
}
