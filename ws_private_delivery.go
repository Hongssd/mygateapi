package mygateapi

// // 订阅余额频道
// func (ws *DeliveryWsStreamClient) SubscribeBalance(userId string) (*Subscription[WsSubscribeResult[[]WsFuturesBalance]], error) {
// 	channel := "futures.balances"
// 	payload := []string{userId}
// 	subKey := []string{channel}

// 	thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesBalance]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKey, true)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ws.futuresBalanceSubMap.Store(subKey[0], thisSub)

// 	return thisSub, nil
// }
// func (ws *DeliveryWsStreamClient) UnSubscribeBalance(userId string) error {
// 	channel := "futures.balances"
// 	payload := []string{userId}
// 	subKey := []string{userId}

// 	thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesBalance]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKey, false)
// 	if err != nil {
// 		return err
// 	}

// 	ws.futuresBalanceSubMap.Store(subKey[0], thisSub)

// 	return nil
// }

// // 订阅仓位频道
// func (ws *DeliveryWsStreamClient) SubscribePosition(userId, contract string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesPosition]], error) {
// 	return ws.SubscribePositionMultiple([]string{userId}, []string{contract})
// }
// func (ws *DeliveryWsStreamClient) SubscribePositionMultiple(userIds, contracts []string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesPosition]], error) {
// 	channel := "futures.positions"
// 	payloads := [][]string{}
// 	for _, u := range userIds {
// 		for _, c := range contracts {
// 			payload := []string{u, c}
// 			payloads = append(payloads, payload)
// 		}
// 	}
// 	subKeys := userIds

// 	var thisSubs []*Subscription[WsSubscribeResult[[]WsFuturesPosition]]
// 	var mu sync.Mutex
// 	var errG errgroup.Group
// 	for _, payload := range payloads {
// 		payload := payload
// 		errG.Go(func() error {
// 			thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesPosition]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, true)
// 			if err != nil {
// 				return err
// 			}
// 			ws.futuresPositionSubMap.Store(payload[0], thisSub)
// 			mu.Lock()
// 			thisSubs = append(thisSubs, thisSub)
// 			mu.Unlock()
// 			return nil
// 		})
// 	}

// 	if err := errG.Wait(); err != nil {
// 		return nil, err
// 	}

// 	subscription, err := mergeSubscription(thisSubs...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return subscription, nil
// }
// func (ws *DeliveryWsStreamClient) UnSubscribePosition(userId, contract string) error {
// 	return ws.UnSubscribePositionMultiple([]string{userId}, []string{contract})
// }
// func (ws *DeliveryWsStreamClient) UnSubscribePositionMultiple(userIds, contracts []string) error {
// 	channel := "futures.positions"
// 	payloads := [][]string{}
// 	for _, u := range userIds {
// 		for _, c := range contracts {
// 			payload := []string{u, c}
// 			payloads = append(payloads, payload)
// 		}
// 	}
// 	subKeys := userIds

// 	var thisSubs []*Subscription[WsSubscribeResult[[]WsFuturesPosition]]
// 	var mu sync.Mutex
// 	var errG errgroup.Group
// 	for _, payload := range payloads {
// 		payload := payload
// 		errG.Go(func() error {
// 			thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesPosition]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, true)
// 			if err != nil {
// 				return err
// 			}
// 			ws.futuresPositionSubMap.Store(payload[0], thisSub)
// 			mu.Lock()
// 			thisSubs = append(thisSubs, thisSub)
// 			mu.Unlock()
// 			return nil
// 		})
// 	}

// 	if err := errG.Wait(); err != nil {
// 		return err
// 	}

// 	return nil
// }
