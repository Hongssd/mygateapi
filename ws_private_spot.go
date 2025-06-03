package mygateapi

// !all订阅全部订单推送
func (ws *SpotWsStreamClient) SubscribeOrders(symbols ...string) (*MultipleSubscription[WsSubscribeResult[WsSpotOrder]], error) {
	//订阅订单推送频道
	channel := "spot.orders"
	if len(symbols) == 0 {
		symbols = []string{"!all"}
	}

	var thisSubs []*Subscription[WsSubscribeResult[WsSpotOrder]]
	payload := symbols
	subKeys := symbols

	thisSub, err := subscribe[WsSubscribeResult[WsSpotOrder]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, true)
	if err != nil {
		return nil, err
	}
	ws.spotOrderSubMap.Store(subKeys[0], thisSub)
	thisSubs = append(thisSubs, thisSub)

	subscription, err := mergeSubscription(thisSubs...)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}

// 现货余额更新推送
func (ws *SpotWsStreamClient) SubscribeBalance() (*Subscription[WsSubscribeResult[WsSpotBalance]], error) {
	channel := "spot.balances"
	payload := []string{}
	subKeys := []string{channel}

	thisSub, err := subscribe[WsSubscribeResult[WsSpotBalance]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, true)
	if err != nil {
		return nil, err
	}

	ws.spotBalanceSubMap.Store(subKeys[0], thisSub)

	return thisSub, nil
}
func (ws *SpotWsStreamClient) UnSubscribeBalance() error {
	channel := "spot.balances"
	subKeys := []string{channel}

	thisSub, err := subscribe[WsSubscribeResult[WsSpotBalance]](&ws.WsStreamClient, channel, UNSUBSCRIBE, []string{}, subKeys, true)
	if err != nil {
		return err
	}

	log.Infof("UnSubscribeBalance Success: args:%v", thisSub.Res)
	ws.spotBalanceSubMap.Delete(subKeys[0])

	return nil
}

// 订阅全仓杠杆余额频道
func (ws *SpotWsStreamClient) SubscribeCrossBalance() (*Subscription[WsSubscribeResult[[]WsSpotCrossBalance]], error) {
	channel := "futures.cross_balances"
	payload := []string{}
	subKey := []string{channel}

	thisSub, err := subscribe[WsSubscribeResult[[]WsSpotCrossBalance]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKey, true)
	if err != nil {
		return nil, err
	}

	ws.spotCrossBalanceSubMap.Store(subKey[0], thisSub)

	return thisSub, nil
}
