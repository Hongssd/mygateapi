package mygateapi

type PrivateRestMarginAccountsResRow struct {
	CurrencyPair string `json:"currency_pair"` //交易对
	Locked       bool   `json:"locked"`        //账户是否被锁定
	Risk         string `json:"risk"`          //杠杆账户当前风险率
	Base         struct {
		Currency  string `json:"currency"`  //货币名称
		Available string `json:"available"` //可用于杠杆交易的额度，available = 保证金 + borrowed
		Locked    string `json:"locked"`    //冻结资金，如已经放在杠杆市场里挂单交易的数额
		Borrowed  string `json:"borrowed"`  //借入资金
		Interest  string `json:"interest"`  //未还利息
	} `json:"base"` //货币账户信息
	Quote struct {
		Currency  string `json:"currency"`  //货币名称
		Available string `json:"available"` //可用于杠杆交易的额度，available = 保证金 + borrowed
		Locked    string `json:"locked"`    //冻结资金，如已经放在杠杆市场里挂单交易的数额
		Borrowed  string `json:"borrowed"`  //借入资金
		Interest  string `json:"interest"`  //未还利息
	} `json:"quote"` //货币账户信息
}

type PrivateRestMarginAccountsRes []PrivateRestMarginAccountsResRow

type PrivateRestMarginAccountBookResRow struct {
	ID           string `json:"id"`             //账户变更记录 ID
	Time         string `json:"time"`           //账户变更时间戳
	TimeMs       int64  `json:"time_ms"`        //账户变更时间戳，毫秒单位
	Currency     string `json:"currency"`       //变更币种
	CurrencyPair string `json:"currency_pair"`  //账户交易对
	Change       string `json:"change"`         //变更金额，正数表示转入，负数表示转出
	Balance      string `json:"balance"`        //变更后账户余额
	Type         string `json:"type,omitempty"` //账户变更类型 , 详见资产流水类型
}

type PrivateRestMarginAccountBookRes []PrivateRestMarginAccountBookResRow

type PrivateRestMarginTransferableRes struct {
	Currency     string `json:"currency"`      //币种信息
	CurrencyPair string `json:"currency_pair"` //交易对
	Amount       string `json:"amount"`        //最大可转出的额度
}

type PrivateRestMarginCrossTransferableRes struct {
	Currency string `json:"currency"` //币种信息
	Amount   string `json:"amount"`   //最大可转出的额度
}
