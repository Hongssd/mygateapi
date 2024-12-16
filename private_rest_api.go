package mygateapi

// 私有Rest接口
type PrivateRestAPI int

const (

	//Spot私有接口
	PrivateSpotAccounts PrivateRestAPI = iota //获取现货交易账户列表

	PrivateSpotOrdersPost
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	//Spot公共接口
	PrivateSpotAccounts: "/api/v4/spot/accounts", //GET 获取现货交易账户列表

	PrivateSpotOrdersPost: "/api/v4/spot/orders", //POST 下单
}
