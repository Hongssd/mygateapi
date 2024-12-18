package mygateapi

// 私有Rest接口
type PrivateRestAPI int

const (
	//Spot私有接口

	// 账户相关
	PrivateSpotAccounts    PrivateRestAPI = iota //获取现货交易账户列表
	PrivateSpotAccountBook                       //查询现货账户变动历史

	// 订单相关
	PrivateSpotOrdersPost //下单
	PrivateSpotOpenOrders //查询所有挂单
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	//Spot私有接口

	// 账户相关
	PrivateSpotAccounts:    "/api/v4/spot/accounts",     //GET 获取现货交易账户列表
	PrivateSpotAccountBook: "/api/v4/spot/account_book", //GET 查询现货账户变动历史

	// 订单相关
	PrivateSpotOrdersPost: "/api/v4/spot/orders",      //POST 下单
	PrivateSpotOpenOrders: "/api/v4/spot/open_orders", //GET 查询所有挂单
}
