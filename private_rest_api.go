package mygateapi

// 私有Rest接口
type PrivateRestAPI int

const (
	// 账户相关
	PrivateRestAccountDetail PrivateRestAPI = iota // 获取用户账户信息

	// 钱包相关
	PrivateRestWalletTransfers    //交易账户互转
	PrivateRestWalletFee          //查询个人交易费率
	PrivateRestWalletTotalBalance //查询个人账户总额

	// Spot & Margin 私有接口
	// Spot 账户相关
	PrivateRestSpotAccounts    //获取现货交易账户列表
	PrivateRestSpotAccountBook //查询现货账户变动历史

	// Margin 账户相关
	PrivateRestMarginAccounts          //获取杠杆交易账户列表
	PrivateRestMarginAccountBook       //查询杠杆账户变动历史
	PrivateRestMarginTransferable      //逐仓杠杆允许的最大转出
	PrivateRestMarginCrossTransferable //全仓杠杆允许的最大转出

	// Spot & Margin 订单相关
	PrivateRestSpotOrdersPost               //下单
	PrivateRestSpotOrdersOrderIdGet         //查询单个订单详情
	PrivateRestSpotOrdersOrderIdPatch       //修改单个订单
	PrivateRestSpotOrdersOrderIdDelete      //撤销单个订单
	PrivateRestSpotOpenOrders               //查询所有挂单
	PrivateRestSpotOrdersGet                //查询订单列表
	PrivateRestSpotMyTrades                 //查询个人成交记录
	PrivateRestSpotPriceOrdersPost          //创建价格触发订单
	PrivateRestSpotPriceOrdersGet           //查询进行中自动订单列表
	PrivateRestSpotPriceOrdersDelete        //批量取消自动订单
	PrivateRestSpotPriceOrdersOrderIdGet    //查询单个自动订单详情
	PrivateRestSpotPriceOrdersOrderIdDelete //撤销单个自动订单

	// Futures 账户相关
	PrivateRestFuturesSettleAccounts                           //获取合约账户
	PrivateRestFuturesSettleAccountBook                        //查询合约账户变更历史
	PrivateRestFuturesSettlePositions                          //获取用户仓位列表
	PrivateRestFuturesSettlePositionsContract                  //获取单个仓位信息
	PrivateRestFuturesSettlePositionsContractMargin            //更新仓位保证金
	PrivateRestFuturesSettlePositionsContractLeverage          //更新仓位杠杆
	PrivateRestFuturesSettlePositionsContractRiskLimit         //更新仓位风险限额
	PrivateRestFuturesSettleDualMode                           //设置持仓模式
	PrivateRestFuturesSettleDualCompPositionsContract          //获取双仓模式下的持仓信息
	PrivateRestFuturesSettleDualCompPositionsContractMargin    //更新双仓模式下的仓位保证金
	PrivateRestFuturesSettleDualCompPositionsContractLeverage  //更新双仓模式下的仓位杠杆
	PrivateRestFuturesSettleDualCompPositionsContractRiskLimit //更新双仓模式下的仓位风险限额

	// Futures 订单相关
	PrivateRestFuturesSettleOrdersPost               //合约交易下单
	PrivateRestFuturesSettleOrdersGet                //查询合约订单列表
	PrivateRestFuturesSettleOrdersTimeRange          //查询合约订单列表(时间区间)
	PrivateRestFuturesSettleOrdersOrderIdGet         //查询单个订单详情
	PrivateRestFuturesSettleOrdersOrderIdDelete      //撤销单个订单
	PrivateRestFuturesSettleOrdersOrderIdPut         //修改单个订单
	PrivateRestFuturesSettleMyTrades                 //查询个人成交记录
	PrivateRestFuturesSettlePriceOrdersPost          //创建价格触发订单
	PrivateRestFuturesSettlePriceOrdersGet           //查询进行中自动订单列表
	PrivateRestFuturesSettlePriceOrdersDelete        //批量取消自动订单
	PrivateRestFuturesSettlePriceOrdersOrderIdGet    //查询单个自动订单详情
	PrivateRestFuturesSettlePriceOrdersOrderIdDelete //撤销单个自动订单

	// Delivery 账户相关
	PrivateRestDeliverySettleAccounts                   //获取合约账户
	PrivateRestDeliverySettleAccountBook                //查询合约账户变动历史
	PrivateRestDeliverySettlePositions                  //获取用户仓位列表
	PrivateRestDeliverySettlePositionsContract          //获取单个仓位信息
	PrivateRestDeliverySettlePositionsContractMargin    //更新仓位保证金
	PrivateRestDeliverySettlePositionsContractLeverage  //更新仓位杠杆
	PrivateRestDeliverySettlePositionsContractRiskLimit //更新仓位风险限额

	// Delivery 订单相关
	PrivateRestDeliverySettleOrdersPost               //合约交易下单
	PrivateRestDeliverySettleOrdersGet                //查询合约订单列表
	PrivateRestDeliverySettleOrdersDelete             //批量取消状态为open的订单
	PrivateRestDeliverySettleOrdersOrderIdGet         //查询单个订单详情
	PrivateRestDeliverySettleOrdersOrderIdDelete      //撤销单个订单
	PrivateRestDeliverySettleMyTrades                 //查询个人成交记录
	PrivateRestDeliverySettlePositionClose            //查询平仓历史
	PrivateRestDeliverySettleLiquidates               //查询强平历史
	PrivateRestDeliverySettleSettlements              //查询结算记录
	PrivateRestDeliverySettlePriceOrdersPost          //创建价格触发订单
	PrivateRestDeliverySettlePriceOrdersGet           //查询自动订单列表
	PrivateRestDeliverySettlePriceOrdersDelete        //批量取消自动订单
	PrivateRestDeliverySettlePriceOrdersOrderIdGet    //查询单个自动订单详情
	PrivateRestDeliverySettlePriceOrdersOrderIdDelete //撤销单个自动订单

	// Unified 账户相关
	PrivateRestUnifiedAccounts                        //获取统一账户信息
	PrivateRestUnifiedTransferable                    //获取统一账户最多可转出
	PrivateRestUnifiedRiskUnits                       //获取用户风险单元详情（仅在组合保证金模式有效）
	PrivateRestUnifiedUnifiedModePut                  //设置统一账户模式
	PrivateRestUnifiedUnifiedModeGet                  //查询统一账户模式
	PrivateRestUnifiedLeverageUserCurrencyConfig      // 用户最大、最小可设置币种杠杆倍数
	PrivateRestUnifiedLeverageUserCurrencySettingGet  // 获取用户币种杠杆倍数，currency不传则查询全部币种
	PrivateRestUnifiedLeverageUserCurrencySettingPost // 设置币种杠杆倍数
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	// 账户相关
	PrivateRestAccountDetail: "/api/v4/account/detail", //GET 获取用户账户信息

	// 钱包相关
	PrivateRestWalletTransfers:    "/api/v4/wallet/transfers",     //POST 交易账户互转
	PrivateRestWalletFee:          "/api/v4/wallet/fee",           //GET 查询个人交易费率
	PrivateRestWalletTotalBalance: "/api/v4/wallet/total_balance", //GET 查询个人账户总额

	// Spot & Margin 私有接口
	// Spot 账户相关
	PrivateRestSpotAccounts:    "/api/v4/spot/accounts",     //GET 获取现货交易账户列表
	PrivateRestSpotAccountBook: "/api/v4/spot/account_book", //GET 查询现货账户变动历史

	// Margin 账户相关
	PrivateRestMarginAccounts:          "/api/v4/margin/accounts",           //GET 获取杠杆交易账户列表
	PrivateRestMarginAccountBook:       "/api/v4/margin/account_book",       //GET 查询杠杆账户变动历史
	PrivateRestMarginTransferable:      "/api/v4/margin/transferable",       //GET 逐仓杠杆允许的最大转出
	PrivateRestMarginCrossTransferable: "/api/v4/margin/cross/transferable", //GET 全仓杠杆允许的最大转出

	// Spot & Margin 订单相关
	PrivateRestSpotOrdersPost:               "/api/v4/spot/orders",                  //POST 下单
	PrivateRestSpotOrdersOrderIdGet:         "/api/v4/spot/orders/{order_id}",       //GET 查询单个订单详情
	PrivateRestSpotOrdersOrderIdPatch:       "/api/v4/spot/orders/{order_id}",       //PATCH 修改单个订单
	PrivateRestSpotOrdersOrderIdDelete:      "/api/v4/spot/orders/{order_id}",       //DELETE 撤销单个订单
	PrivateRestSpotOrdersGet:                "/api/v4/spot/orders",                  //GET 查询订单列表
	PrivateRestSpotOpenOrders:               "/api/v4/spot/open_orders",             //GET 查询所有挂单
	PrivateRestSpotMyTrades:                 "/api/v4/spot/my_trades",               //GET 查询个人成交记录
	PrivateRestSpotPriceOrdersPost:          "/api/v4/spot/price_orders",            //POST 创建价格触发订单
	PrivateRestSpotPriceOrdersGet:           "/api/v4/spot/price_orders",            //GET 查询价格触发订单列表
	PrivateRestSpotPriceOrdersDelete:        "/api/v4/spot/price_orders",            //DELETE 批量取消自动订单
	PrivateRestSpotPriceOrdersOrderIdGet:    "/api/v4/spot/price_orders/{order_id}", //GET 查询单个价格触发订单详情
	PrivateRestSpotPriceOrdersOrderIdDelete: "/api/v4/spot/price_orders/{order_id}", //DELETE 撤销单个价格触发订单

	// Futures 账户相关
	PrivateRestFuturesSettleAccounts:                           "/api/v4/futures/{settle}/accounts",                                  //GET 获取合约账户
	PrivateRestFuturesSettleAccountBook:                        "/api/v4/futures/{settle}/account_book",                              //GET 查询合约账户变更历史
	PrivateRestFuturesSettlePositions:                          "/api/v4/futures/{settle}/positions",                                 //GET 获取用户仓位列表
	PrivateRestFuturesSettlePositionsContract:                  "/api/v4/futures/{settle}/positions/{contract}",                      //GET 获取单个仓位信息
	PrivateRestFuturesSettlePositionsContractMargin:            "/api/v4/futures/{settle}/positions/{contract}/margin",               //POST 更新仓位保证金
	PrivateRestFuturesSettlePositionsContractLeverage:          "/api/v4/futures/{settle}/positions/{contract}/leverage",             //POST 更新仓位杠杆
	PrivateRestFuturesSettlePositionsContractRiskLimit:         "/api/v4/futures/{settle}/positions/{contract}/risk_limit",           //POST 更新仓位风险限额
	PrivateRestFuturesSettleDualMode:                           "/api/v4/futures/{settle}/dual_mode",                                 //POST 设置持仓模式
	PrivateRestFuturesSettleDualCompPositionsContract:          "/api/v4/futures/{settle}/dual_comp/positions/{contract}",            //GET 获取双仓模式下的持仓信息
	PrivateRestFuturesSettleDualCompPositionsContractMargin:    "/api/v4/futures/{settle}/dual_comp/positions/{contract}/margin",     //POST 更新双仓模式下的仓位保证金
	PrivateRestFuturesSettleDualCompPositionsContractLeverage:  "/api/v4/futures/{settle}/dual_comp/positions/{contract}/leverage",   //POST 更新双仓模式下的仓位杠杆
	PrivateRestFuturesSettleDualCompPositionsContractRiskLimit: "/api/v4/futures/{settle}/dual_comp/positions/{contract}/risk_limit", //POST 更新双仓模式下的仓位风险限额

	// Futures 订单相关
	PrivateRestFuturesSettleOrdersPost:               "/api/v4/futures/{settle}/orders",                  //POST 合约交易下单
	PrivateRestFuturesSettleOrdersGet:                "/api/v4/futures/{settle}/orders",                  //GET 查询合约订单列表
	PrivateRestFuturesSettleOrdersTimeRange:          "/api/v4/futures/{settle}/orders_timerange",        //GET 查询合约订单列表(时间区间)
	PrivateRestFuturesSettleOrdersOrderIdGet:         "/api/v4/futures/{settle}/orders/{order_id}",       //GET 查询单个订单详情
	PrivateRestFuturesSettleOrdersOrderIdDelete:      "/api/v4/futures/{settle}/orders/{order_id}",       //DELETE 撤销单个订单
	PrivateRestFuturesSettleOrdersOrderIdPut:         "/api/v4/futures/{settle}/orders/{order_id}",       //PUT 修改单个订单
	PrivateRestFuturesSettleMyTrades:                 "/api/v4/futures/{settle}/my_trades",               //GET 查询个人成交记录
	PrivateRestFuturesSettlePriceOrdersPost:          "/api/v4/futures/{settle}/price_orders",            //POST 创建价格触发订单
	PrivateRestFuturesSettlePriceOrdersGet:           "/api/v4/futures/{settle}/price_orders",            //GET 查询进行中自动订单列表
	PrivateRestFuturesSettlePriceOrdersDelete:        "/api/v4/futures/{settle}/price_orders",            //DELETE 批量取消自动订单
	PrivateRestFuturesSettlePriceOrdersOrderIdGet:    "/api/v4/futures/{settle}/price_orders/{order_id}", //GET 查询单个自动订单详情
	PrivateRestFuturesSettlePriceOrdersOrderIdDelete: "/api/v4/futures/{settle}/price_orders/{order_id}", //DELETE 撤销单个自动订单

	// Delivery 账户相关
	PrivateRestDeliverySettleAccounts:                   "/api/v4/delivery/{settle}/accounts",                        //GET 获取合约账户
	PrivateRestDeliverySettleAccountBook:                "/api/v4/delivery/{settle}/account_book",                    //GET 查询合约账户变动历史
	PrivateRestDeliverySettlePositions:                  "/api/v4/delivery/{settle}/positions",                       //GET 获取用户仓位列表
	PrivateRestDeliverySettlePositionsContract:          "/api/v4/delivery/{settle}/positions/{contract}",            //GET 获取单个仓位信息
	PrivateRestDeliverySettlePositionsContractMargin:    "/api/v4/delivery/{settle}/positions/{contract}/margin",     //POST 更新仓位保证金
	PrivateRestDeliverySettlePositionsContractLeverage:  "/api/v4/delivery/{settle}/positions/{contract}/leverage",   //POST 更新仓位杠杆
	PrivateRestDeliverySettlePositionsContractRiskLimit: "/api/v4/delivery/{settle}/positions/{contract}/risk_limit", //POST 更新仓位风险限额

	// Delivery 订单相关
	PrivateRestDeliverySettleOrdersPost:               "/api/v4/delivery/{settle}/orders",                  //POST 合约交易下单
	PrivateRestDeliverySettleOrdersGet:                "/api/v4/delivery/{settle}/orders",                  //GET 查询合约订单列表
	PrivateRestDeliverySettleOrdersDelete:             "/api/v4/delivery/{settle}/orders",                  //DELETE 批量取消状态为open的订单
	PrivateRestDeliverySettleOrdersOrderIdGet:         "/api/v4/delivery/{settle}/orders/{order_id}",       //GET 查询单个订单详情
	PrivateRestDeliverySettleOrdersOrderIdDelete:      "/api/v4/delivery/{settle}/orders/{order_id}",       //DELETE 撤销单个订单
	PrivateRestDeliverySettleMyTrades:                 "/api/v4/delivery/{settle}/my_trades",               //GET 查询个人成交记录
	PrivateRestDeliverySettlePositionClose:            "/api/v4/delivery/{settle}/position_close",          //GET 查询平仓历史
	PrivateRestDeliverySettleLiquidates:               "/api/v4/delivery/{settle}/liquidates",              //GET 查询强平历史
	PrivateRestDeliverySettleSettlements:              "/api/v4/delivery/{settle}/settlements",             //GET 查询结算记录
	PrivateRestDeliverySettlePriceOrdersPost:          "/api/v4/delivery/{settle}/price_orders",            //POST 创建价格触发订单
	PrivateRestDeliverySettlePriceOrdersGet:           "/api/v4/delivery/{settle}/price_orders",            //GET 查询自动订单列表
	PrivateRestDeliverySettlePriceOrdersDelete:        "/api/v4/delivery/{settle}/price_orders",            //DELETE 批量取消自动订单
	PrivateRestDeliverySettlePriceOrdersOrderIdGet:    "/api/v4/delivery/{settle}/price_orders/{order_id}", //GET 查询单个自动订单详情
	PrivateRestDeliverySettlePriceOrdersOrderIdDelete: "/api/v4/delivery/{settle}/price_orders/{order_id}", //DELETE 撤销单个自动订单

	// Unified 账户相关
	PrivateRestUnifiedAccounts:                        "/api/v4/unified/accounts",                       //GET 获取统一账户信息
	PrivateRestUnifiedTransferable:                    "/api/v4/unified/transferable",                   //GET 获取统一账户最多可转出
	PrivateRestUnifiedRiskUnits:                       "/api/v4/unified/risk_units",                     //GET 获取用户风险单元详情（仅在组合保证金模式有效）
	PrivateRestUnifiedUnifiedModePut:                  "/api/v4/unified/unified_mode",                   //PUT 设置统一账户模式
	PrivateRestUnifiedUnifiedModeGet:                  "/api/v4/unified/unified_mode",                   //GET 查询统一账户模式
	PrivateRestUnifiedLeverageUserCurrencyConfig:      "/api/v4/unified/leverage/user_currency_config",  //GET 用户最大、最小可设置币种杠杆倍数
	PrivateRestUnifiedLeverageUserCurrencySettingGet:  "/api/v4/unified/leverage/user_currency_setting", //GET 获取用户币种杠杆倍数，currency不传则查询全部币种
	PrivateRestUnifiedLeverageUserCurrencySettingPost: "/api/v4/unified/leverage/user_currency_setting", //POST 设置币种杠杆倍数
}
