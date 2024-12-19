package mygateapi

// 私有Rest接口
type PrivateRestAPI int

const (
	// 钱包相关
	PrivateRestWalletFee          PrivateRestAPI = iota //查询个人交易费率
	PrivateRestWalletTotalBalance                       //查询个人账户总额

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
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	// 钱包相关
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

}
