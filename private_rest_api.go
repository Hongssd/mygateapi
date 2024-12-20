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
	PrivateRestSpotOrdersPost          //下单
	PrivateRestSpotPriceOrdersPost     //创建价格触发订单
	PrivateRestSpotOrdersOrderIdGet    //查询单个订单详情
	PrivateRestSpotOrdersOrderIdPatch  //修改单个订单
	PrivateRestSpotOrdersOrderIdDelete //撤销单个订单
	PrivateRestSpotOpenOrders          //查询所有挂单
	PrivateRestSpotOrdersGet           //查询订单列表
	PrivateRestSpotMyTrades            //查询个人成交记录

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
	PrivateRestFuturesSettleOrdersPost          //合约交易下单
	PrivateRestFuturesSettleOrdersGet           //查询合约订单列表
	PrivateRestFuturesSettleOrdersOrderIdGet    //查询单个订单详情
	PrivateRestFuturesSettleOrdersOrderIdDelete //撤销单个订单
	PrivateRestFuturesSettleOrdersOrderIdPut    //修改单个订单
	PrivateRestFuturesSettleMyTrades            //查询个人成交记录
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
	PrivateRestSpotOrdersPost:          "/api/v4/spot/orders",            //POST 下单
	PrivateRestSpotPriceOrdersPost:     "/api/v4/spot/price_orders",      //POST 创建价格触发订单
	PrivateRestSpotOrdersOrderIdGet:    "/api/v4/spot/orders/{order_id}", //GET 查询单个订单详情
	PrivateRestSpotOrdersOrderIdPatch:  "/api/v4/spot/orders/{order_id}", //PATCH 修改单个订单
	PrivateRestSpotOrdersOrderIdDelete: "/api/v4/spot/orders/{order_id}", //DELETE 撤销单个订单
	PrivateRestSpotOrdersGet:           "/api/v4/spot/orders",            //GET 查询订单列表
	PrivateRestSpotOpenOrders:          "/api/v4/spot/open_orders",       //GET 查询所有挂单
	PrivateRestSpotMyTrades:            "/api/v4/spot/my_trades",         //GET 查询个人成交记录

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
	PrivateRestFuturesSettleOrdersPost:          "/api/v4/futures/{settle}/orders",            //POST 合约交易下单
	PrivateRestFuturesSettleOrdersGet:           "/api/v4/futures/{settle}/orders",            //GET 查询合约订单列表
	PrivateRestFuturesSettleOrdersOrderIdGet:    "/api/v4/futures/{settle}/orders/{order_id}", //GET 查询单个订单详情
	PrivateRestFuturesSettleOrdersOrderIdDelete: "/api/v4/futures/{settle}/orders/{order_id}", //DELETE 撤销单个订单
	PrivateRestFuturesSettleOrdersOrderIdPut:    "/api/v4/futures/{settle}/orders/{order_id}", //PUT 修改单个订单
	PrivateRestFuturesSettleMyTrades:            "/api/v4/futures/{settle}/my_trades",         //GET 查询个人成交记录
}
