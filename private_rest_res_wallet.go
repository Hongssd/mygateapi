package mygateapi

type PrivateRestWalletTotalBalanceResAdditionalProperties struct {
	Amount        string `json:"amount"`                   //账户总额数字
	Currency      string `json:"currency"`                 //币种
	UnrealisedPnl string `json:"unrealised_pnl,omitempty"` //未实现盈亏总和,这个字段只会在futures,options,delivery,total 账户中出现
	Borrowed      string `json:"borrowed,omitempty"`       //杠杆借贷总和,这个字段只会在margin,cross_margin账户中出现
}

type PrivateRestWalletTotalBalanceRes struct {
	Total   PrivateRestWalletTotalBalanceResAdditionalProperties `json:"total"` //换算成目标币种的账户总额
	Details struct {
		CrossMargin PrivateRestWalletTotalBalanceResAdditionalProperties `json:"cross_margin"` //全仓杠杆账户
		Spot        PrivateRestWalletTotalBalanceResAdditionalProperties `json:"spot"`         //现货账户
		Finance     PrivateRestWalletTotalBalanceResAdditionalProperties `json:"finance"`      //金融账户
		Margin      PrivateRestWalletTotalBalanceResAdditionalProperties `json:"margin"`       //杠杆账户
		Quant       PrivateRestWalletTotalBalanceResAdditionalProperties `json:"quant"`        //量化账户
		Futures     PrivateRestWalletTotalBalanceResAdditionalProperties `json:"futures"`      //永续合约账户
		Delivery    PrivateRestWalletTotalBalanceResAdditionalProperties `json:"delivery"`     //交割合约账户
		Warrant     PrivateRestWalletTotalBalanceResAdditionalProperties `json:"warrant"`      //warrant 账户
		Cbbc        PrivateRestWalletTotalBalanceResAdditionalProperties `json:"cbbc"`         //牛熊证账户
	} `json:"details"` //各账户总额
}

type PrivateRestWalletFeeRes struct {
	UserID           int    `json:"user_id"`            //用户 ID
	TakerFee         string `json:"taker_fee"`          //taker 费率
	MakerFee         string `json:"maker_fee"`          //maker 费率
	GtDiscount       bool   `json:"gt_discount"`        //是否开启 GT 抵扣折扣
	GtTakerFee       string `json:"gt_taker_fee"`       //GT 抵扣 taker 费率，未开启 GT 抵扣则为 0
	GtMakerFee       string `json:"gt_maker_fee"`       //GT 抵扣 maker 费率，未开启 GT 抵扣则为 0
	LoanFee          string `json:"loan_fee"`           //杠杆理财的费率
	PointType        string `json:"point_type"`         //点卡类型，0 - 初版点卡，1 - 202009 启用的新点卡
	FuturesTakerFee  string `json:"futures_taker_fee"`  //合约 taker 费率
	FuturesMakerFee  string `json:"futures_maker_fee"`  //合约 maker 费率
	DeliveryTakerFee string `json:"delivery_taker_fee"` //交割合约 taker 费率
	DeliveryMakerFee string `json:"delivery_maker_fee"` //交割合约 maker 费率
	DebitFee         int    `json:"debit_fee"`          //费率抵扣类型 , 1 - GT抵扣 , 2 - 点卡抵扣 , 3 - VIP费率
}

type PrivateRestSpotMyTradesResRow struct {
	ID           string `json:"id"`             //成交记录 ID
	CreateTime   string `json:"create_time"`    //成交时间
	CreateTimeMs string `json:"create_time_ms"` //成交时间，毫秒精度
	CurrencyPair string `json:"currency_pair"`  //交易货币对
	Side         string `json:"side"`           //买单或者卖单
	Role         string `json:"role"`           //交易角色，公共接口无此字段返回
	Amount       string `json:"amount"`         //交易数量
	Price        string `json:"price"`          //交易价
	OrderID      string `json:"order_id"`       //关联的订单 ID，公共接口无此字段返回
	Fee          string `json:"fee"`            //成交扣除的手续费，公共接口无此字段返回
	FeeCurrency  string `json:"fee_currency"`   //手续费计价单位，公共接口无此字段返回
	PointFee     string `json:"point_fee"`      //手续费抵扣使用的点卡数量，公共接口无此字段返回
	GtFee        string `json:"gt_fee"`         //手续费抵扣使用的 GT 数量，公共接口无此字段返回
	AmendText    string `json:"amend_text"`     //用户修改订单时备注的信息
	SequenceID   string `json:"sequence_id"`    //单市场连续成交ID
	Text         string `json:"text"`           //订单的自定义信息，公共接口无此字段返回
}

type PrivateRestSpotMyTradesRes []PrivateRestSpotMyTradesResRow

type PrivateRestWalletTransfersRes struct {
	TxId string `json:"tx_id"` //操作单号
}
