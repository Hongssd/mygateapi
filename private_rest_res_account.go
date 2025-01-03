package mygateapi

type PrivateRestAccountDetailRes struct {
	IpWhitelist   []string `json:"ip_whitelist"`
	CurrencyPairs []string `json:"currency_pairs"`
	UserId        int64    `json:"user_id"`
	Tier          int64    `json:"tier"`
	Key           struct {
		Mode int `json:"mode"` // 模式： 1 - 经典模式 2 - 旧版统一模式
	} `json:"key"`
	CopyTradingRole int `json:"copy_trading_role"`
}
