package mygateapi

// 公共Rest接口
type PublicRestAPI int

const (

	//Spot公共接口
	PublicSpotCurrencies PublicRestAPI = iota //查询所有币种信息

)

var PublicRestAPIMap = map[PublicRestAPI]string{
	//Spot公共接口
	PublicSpotCurrencies: "/api/v4/spot/currencies", //GET 查询所有币种信息
}
