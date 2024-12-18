package mygateapi

// 私有Rest接口
type PrivateRestAPI int

const (
	//Spot私有接口

	// 账户相关
	PrivateRestSpotAccounts    PrivateRestAPI = iota //获取现货交易账户列表
	PrivateRestSpotAccountBook                       //查询现货账户变动历史

	// 钱包相关
	PrivateRestWalletFee          //查询个人交易费率
	PrivateRestWalletTotalBalance //查询个人账户总额

	// 订单相关
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
	//Spot私有接口

	// 钱包相关
	PrivateRestWalletFee:          "/api/v4/wallet/fee",           //GET 查询个人交易费率
	PrivateRestWalletTotalBalance: "/api/v4/wallet/total_balance", //GET 查询个人账户总额

	// 账户相关
	PrivateRestSpotAccounts:    "/api/v4/spot/accounts",     //GET 获取现货交易账户列表
	PrivateRestSpotAccountBook: "/api/v4/spot/account_book", //GET 查询现货账户变动历史

	// 订单相关
	PrivateRestSpotOrdersPost:          "/api/v4/spot/orders",            //POST 下单
	PrivateRestSpotPriceOrdersPost:     "/api/v4/spot/price_orders",      //POST 创建价格触发订单
	PrivateRestSpotOrdersOrderIdGet:    "/api/v4/spot/orders/{order_id}", //GET 查询单个订单详情
	PrivateRestSpotOrdersOrderIdPatch:  "/api/v4/spot/orders/{order_id}", //PATCH 修改单个订单
	PrivateRestSpotOrdersOrderIdDelete: "/api/v4/spot/orders/{order_id}", //DELETE 撤销单个订单
	PrivateRestSpotOrdersGet:           "/api/v4/spot/orders",            //GET 查询订单列表
	PrivateRestSpotOpenOrders:          "/api/v4/spot/open_orders",       //GET 查询所有挂单
	PrivateRestSpotMyTrades:            "/api/v4/spot/my_trades",         //GET 查询个人成交记录
}
