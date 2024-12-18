package mygateapi

// 公共Rest接口
type PublicRestAPI int

const (
	//Spot公共接口
	PublicSpotCurrencies          PublicRestAPI = iota //查询所有币种信息
	PublicSpotCurrenciesCurrency                       //查询单个币种信息
	PublicSpotCurrencyPairsAll                         //查询支持的所有交易对
	PublicSpotCurrencyPairsSingle                      //查询单个交易对详情
	PublicSpotTickers                                  //获取交易对ticker信息
	PublicSpotOrderBook                                //获取市场深度信息
	PublicSpotTrades                                   //查询市场成交记录
	PublicSpotCandlesticks                             //市场K线图
)

var PublicRestAPIMap = map[PublicRestAPI]string{
	//Spot公共接口
	PublicSpotCurrencies:         "/api/v4/spot/currencies",            //GET 查询所有币种信息
	PublicSpotCurrenciesCurrency: "/api/v4/spot/currencies/{currency}", //GET 查询单个币种信息

	PublicSpotCurrencyPairsAll:    "/api/v4/spot/currency_pairs", //GET 查询支持的所有交易对
	PublicSpotCurrencyPairsSingle: "/api/v4/spot/currency_pairs", //GET 查询单个交易对详情
	PublicSpotTickers:             "/api/v4/spot/tickers",        //GET 获取交易对tickers信息
	PublicSpotOrderBook:           "/api/v4/spot/order_book",     //GET 获取市场深度信息
	PublicSpotTrades:              "/api/v4/spot/trades",         //GET 查询市场成交记录
	PublicSpotCandlesticks:        "/api/v4/spot/candlesticks",   //GET 获取K线图

}
