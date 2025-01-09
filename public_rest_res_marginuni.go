package mygateapi

type PublicRestMarginUniCurrencyPairsResRow struct {
	CurrencyPair         string `json:"currency_pair"`           //交易对
	BaseMinBorrowAmount  string `json:"base_min_borrow_amount"`  //交易货币最小借入数量
	QuoteMinBorrowAmount string `json:"quote_min_borrow_amount"` //计价货币最小借入数量
	Leverage             string `json:"leverage"`                //杠杆倍数
}

type PublicRestMarginUniCurrencyPairsRes []PublicRestMarginUniCurrencyPairsResRow

type PublicRestMarginUniCurrencyPairsCurrencyPairRes PublicRestMarginUniCurrencyPairsResRow
