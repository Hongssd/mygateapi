package mygateapi

type PublicRestMarginCrossCurrenciesResRow struct {
	Name                 string `json:"name"`                    //币种名称
	Rate                 string `json:"rate"`                    //最小借贷利率（小时利率）
	Prec                 string `json:"prec"`                    //币种精度
	Discount             string `json:"discount"`                //币种价值折扣，计算账户总资产时使用的折扣系数
	MinBorrowAmount      string `json:"min_borrow_amount"`       //该币种最小的借入数量，单位是币
	UserMaxBorrowAmount  string `json:"user_max_borrow_amount"`  //用户最大允许的借入数额，单位是 USDT
	TotalMaxBorrowAmount string `json:"total_max_borrow_amount"` //该币种允许借入的最大数额，单位是 USDT
	Price                string `json:"price"`                   //该币种兑换 USDT 的价格
	Loanable             bool   `json:"loanable"`                //该币种是否可借贷
	Status               int    `json:"status"`                  //币种状态 0 : 禁用 1 : 启用
}

type PublicRestMarginCrossCurrenciesRes []PublicRestMarginCrossCurrenciesResRow
