package mygateapi

// 公共Rest接口
type PublicRestAPI int

const (
	//Spot公共接口
	PublicRestSpotCurrencies                PublicRestAPI = iota //查询所有币种信息
	PublicRestSpotCurrenciesCurrency                             //查询单个币种信息
	PublicRestSpotCurrencyPairs                                  //查询支持的所有交易对
	PublicRestSpotCurrencyPairsCurrencyPair                      //查询单个交易对详情
	PublicRestSpotTickers                                        //获取交易对ticker信息
	PublicRestSpotOrderBook                                      //获取市场深度信息
	PublicRestSpotTrades                                         //查询市场成交记录
	PublicRestSpotCandlesticks                                   //市场K线图
	PublicRestSpotTime                                           //获取服务器当前时间
)

var PublicRestAPIMap = map[PublicRestAPI]string{
	//Spot公共接口
	PublicRestSpotCurrencies:                "/api/v4/spot/currencies",                     //GET 查询所有币种信息
	PublicRestSpotCurrenciesCurrency:        "/api/v4/spot/currencies/{currency}",          //GET 查询单个币种信息
	PublicRestSpotCurrencyPairs:             "/api/v4/spot/currency_pairs",                 //GET 查询支持的所有交易对
	PublicRestSpotCurrencyPairsCurrencyPair: "/api/v4/spot/currency_pairs/{currency_pair}", //GET 查询单个交易对详情
	PublicRestSpotTickers:                   "/api/v4/spot/tickers",                        //GET 获取交易对tickers信息
	PublicRestSpotOrderBook:                 "/api/v4/spot/order_book",                     //GET 获取市场深度信息
	PublicRestSpotTrades:                    "/api/v4/spot/trades",                         //GET 查询市场成交记录
	PublicRestSpotCandlesticks:              "/api/v4/spot/candlesticks",                   //GET 获取K线图
	PublicRestSpotTime:                      "/api/v4/spot/time",                           //GET 获取服务器当前时间
}
