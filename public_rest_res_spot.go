package mygateapi

//返回格式
//状态码 200
//
//名称	类型	描述
//None	array
//» currency	string	币种名称
//» delisted	boolean	是否下架
//» withdraw_disabled	boolean	是否暂停提现
//» withdraw_delayed	boolean	提现是否存在延迟
//» deposit_disabled	boolean	是否暂停充值
//» trade_disabled	boolean	是否暂停交易
//» fixed_rate	string	固定交易手续费率。仅限固定交易费率的币种，普通币种该字段无效
//» chain	string	币对应的链

type PublicRestSpotCurrenciesResRow struct {
	Currency         string `json:"currency"`          // string 币种名称
	Delisted         bool   `json:"delisted"`          // boolean 是否下架
	WithdrawDisabled bool   `json:"withdraw_disabled"` // boolean 是否暂停提现
	WithdrawDelayed  bool   `json:"withdraw_delayed"`  // boolean 提现是否存在延迟
	DepositDisabled  bool   `json:"deposit_disabled"`  // boolean 是否暂停充值
	TradeDisabled    bool   `json:"trade_disabled"`    // boolean 是否暂停交易
	FixedRate        string `json:"fixed_rate"`        // string 固定交易手续费率。仅限固定交易费率的币种，普通币种该字段无效
	Chain            string `json:"chain"`             // string 币对应的链
}
type PublicRestSpotCurrenciesRes []PublicRestSpotCurrenciesResRow
