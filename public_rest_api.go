package mygateapi

// 公共Rest接口
type PublicRestAPI int

const (
	// Spot公共接口
	PublicRestSpotCurrencies                PublicRestAPI = iota //查询所有币种信息
	PublicRestSpotCurrenciesCurrency                             //查询单个币种信息
	PublicRestSpotCurrencyPairs                                  //查询支持的所有交易对
	PublicRestSpotCurrencyPairsCurrencyPair                      //查询单个交易对详情
	PublicRestSpotTickers                                        //获取交易对ticker信息
	PublicRestSpotOrderBook                                      //获取市场深度信息
	PublicRestSpotTrades                                         //查询市场成交记录
	PublicRestSpotCandlesticks                                   //市场K线图
	PublicRestSpotTime                                           //获取服务器当前时间

	// Margin公共接口
	PublicRestMarginCrossCurrencies //全仓杠杆支持的币种列表

	// Futures公共接口
	PublicRestFuturesSettleContracts         //查询所有的合约信息
	PublicRestFuturesSettleContractsContract //查询单个合约信息
	PublicRestFuturesSettleOrderBook         //查询合约市场深度信息
	PublicRestFuturesSettleTrades            //合约市场成交记录
	PublicRestFuturesSettleCandlesticks      //合约市场K线图

	// Delivery公共接口
	PublicRestDeliverySettleContracts         //查询所有的合约信息
	PublicRestDeliverySettleContractsContract //查询单个合约信息
	PublicRestDeliverySettleOrderBook         //查询合约市场深度信息
	PublicRestDeliverySettleTrades            //合约市场成交记录
	PublicRestDeliverySettleCandlesticks      //合约市场K线图
	PublicRestDeliverySettleTickers           //获取所有合约交易中行情统计
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

	// Margin公共接口
	PublicRestMarginCrossCurrencies: "/api/v4/margin/cross/currencies", //GET 全仓杠杆支持的币种列表

	// Futures公共接口
	PublicRestFuturesSettleContracts:         "/api/v4/futures/{settle}/contracts",            //GET 查询所有的合约信息
	PublicRestFuturesSettleContractsContract: "/api/v4/futures/{settle}/contracts/{contract}", //GET 查询单个合约信息
	PublicRestFuturesSettleOrderBook:         "/api/v4/futures/{settle}/order_book",           //GET 查询合约市场深度信息
	PublicRestFuturesSettleTrades:            "/api/v4/futures/{settle}/trades",               //GET 合约市场成交记录
	PublicRestFuturesSettleCandlesticks:      "/api/v4/futures/{settle}/candlesticks",         //GET 合约市场K线图

	// Delivery公共接口
	PublicRestDeliverySettleContracts:         "/api/v4/delivery/{settle}/contracts",            //GET 查询所有的合约信息
	PublicRestDeliverySettleContractsContract: "/api/v4/delivery/{settle}/contracts/{contract}", //GET 查询单个合约信息
	PublicRestDeliverySettleOrderBook:         "/api/v4/delivery/{settle}/order_book",           //GET 查询合约市场深度信息
	PublicRestDeliverySettleTrades:            "/api/v4/delivery/{settle}/trades",               //GET 合约市场成交记录
	PublicRestDeliverySettleCandlesticks:      "/api/v4/delivery/{settle}/candlesticks",         //GET 合约市场K线图
	PublicRestDeliverySettleTickers:           "/api/v4/delivery/{settle}/tickers",              //GET 获取所有合约交易中行情统计
}
