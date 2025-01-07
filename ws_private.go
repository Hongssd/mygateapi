package mygateapi

import (
	"golang.org/x/sync/errgroup"
	"sync"
)

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

// 订阅订单频道
func (ws *FuturesWsStreamClient) SubscribeOrder(userId, contract string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesOrder]], error) {
	return ws.SubscribeOrderMultiple(userId, []string{contract})
}
func (ws *FuturesWsStreamClient) SubscribeOrderMultiple(userId string, contracts []string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesOrder]], error) {
	channel := "futures.orders"
	payloads := [][]string{}
	subKeys := contracts
	for _, c := range contracts {
		payload := []string{userId, c}
		payloads = append(payloads, payload)
	}

	var thisSubs []*Subscription[WsSubscribeResult[[]WsFuturesOrder]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesOrder]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, true)
			if err != nil {
				return err
			}
			ws.futuresOrderSubMap.Store(payload[1], thisSub)
			mu.Lock()
			thisSubs = append(thisSubs, thisSub)
			mu.Unlock()
			return nil
		})
	}

	if err := errG.Wait(); err != nil {
		return nil, err
	}

	subscription, err := mergeSubscription(thisSubs...)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
func (ws *FuturesWsStreamClient) UnSubscribeOrder(userId, contract string) error {
	return ws.UnSubscribeOrderMultiple(userId, []string{contract})
}
func (ws *FuturesWsStreamClient) UnSubscribeOrderMultiple(userId string, contracts []string) error {
	channel := "futures.orders"
	payloads := [][]string{}
	subKeys := contracts
	for _, c := range contracts {
		payload := []string{userId, c}
		payloads = append(payloads, payload)
	}

	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			_, err := subscribe[WsSubscribeResult[[]WsFuturesOrder]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, true)
			if err != nil {
				return err
			}
			mu.Lock()
			ws.futuresOrderSubMap.Delete(payload[1])
			mu.Unlock()
			return nil
		})
	}

	if err := errG.Wait(); err != nil {
		return err
	}

	return nil
}

// 订阅现货余额频道
func (ws *FuturesWsStreamClient) SubscribeBalance(userId string) (*Subscription[WsSubscribeResult[[]WsFuturesBalance]], error) {
	channel := "futures.balances"
	payload := []string{userId}
	subKey := []string{userId}

	thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesBalance]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKey, true)
	if err != nil {
		return nil, err
	}

	ws.futuresBalanceSubMap.Store(subKey[0], thisSub)

	return thisSub, nil
}
func (ws *FuturesWsStreamClient) UnSubscribeBalance(userId string) error {
	channel := "futures.balances"
	payload := []string{userId}
	subKey := []string{userId}

	_, err := subscribe[WsSubscribeResult[[]WsFuturesBalance]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKey, true)
	if err != nil {
		return err
	}

	ws.futuresBalanceSubMap.Delete(subKey[0])

	return nil
}

// 订阅仓位频道
func (ws *FuturesWsStreamClient) SubscribePosition(userId, contract string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesPosition]], error) {
	return ws.SubscribePositionMultiple(userId, []string{contract})
}
func (ws *FuturesWsStreamClient) SubscribePositionMultiple(userId string, contracts []string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesPosition]], error) {
	channel := "futures.positions"
	payloads := [][]string{}
	subKeys := contracts
	for _, c := range contracts {
		payload := []string{userId, c}
		payloads = append(payloads, payload)
	}

	var thisSubs []*Subscription[WsSubscribeResult[[]WsFuturesPosition]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesPosition]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, true)
			if err != nil {
				return err
			}
			ws.futuresPositionSubMap.Store(payload[1], thisSub)
			mu.Lock()
			thisSubs = append(thisSubs, thisSub)
			mu.Unlock()
			return nil
		})
	}

	if err := errG.Wait(); err != nil {
		return nil, err
	}

	subscription, err := mergeSubscription(thisSubs...)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
func (ws *FuturesWsStreamClient) UnSubscribePosition(userId, contract string) error {
	return ws.UnSubscribePositionMultiple(userId, []string{contract})
}
func (ws *FuturesWsStreamClient) UnSubscribePositionMultiple(userId string, contracts []string) error {
	channel := "futures.positions"
	payloads := [][]string{}
	subKeys := contracts
	for _, c := range contracts {
		payload := []string{userId, c}
		payloads = append(payloads, payload)
	}

	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			_, err := subscribe[WsSubscribeResult[[]WsFuturesPosition]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, true)
			if err != nil {
				return err
			}
			mu.Lock()
			ws.futuresPositionSubMap.Delete(payload[1])
			mu.Unlock()
			return nil
		})
	}

	if err := errG.Wait(); err != nil {
		return err
	}

	return nil
}

// 订阅余额频道
func (ws *DeliveryWsStreamClient) SubscribeBalance(userId string) (*Subscription[WsSubscribeResult[[]WsFuturesBalance]], error) {
	channel := "futures.balances"
	payload := []string{userId}
	subKey := []string{channel}

	thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesBalance]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKey, true)
	if err != nil {
		return nil, err
	}

	ws.futuresBalanceSubMap.Store(subKey[0], thisSub)

	return thisSub, nil
}
func (ws *DeliveryWsStreamClient) UnSubscribeBalance(userId string) error {
	channel := "futures.balances"
	payload := []string{userId}
	subKey := []string{userId}

	thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesBalance]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKey, false)
	if err != nil {
		return err
	}

	ws.futuresBalanceSubMap.Store(subKey[0], thisSub)

	return nil
}

// 订阅仓位频道
func (ws *DeliveryWsStreamClient) SubscribePosition(userId, contract string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesPosition]], error) {
	return ws.SubscribePositionMultiple([]string{userId}, []string{contract})
}
func (ws *DeliveryWsStreamClient) SubscribePositionMultiple(userIds, contracts []string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesPosition]], error) {
	channel := "futures.positions"
	payloads := [][]string{}
	for _, u := range userIds {
		for _, c := range contracts {
			payload := []string{u, c}
			payloads = append(payloads, payload)
		}
	}
	subKeys := userIds

	var thisSubs []*Subscription[WsSubscribeResult[[]WsFuturesPosition]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesPosition]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, true)
			if err != nil {
				return err
			}
			ws.futuresPositionSubMap.Store(payload[0], thisSub)
			mu.Lock()
			thisSubs = append(thisSubs, thisSub)
			mu.Unlock()
			return nil
		})
	}

	if err := errG.Wait(); err != nil {
		return nil, err
	}

	subscription, err := mergeSubscription(thisSubs...)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
func (ws *DeliveryWsStreamClient) UnSubscribePosition(userId, contract string) error {
	return ws.UnSubscribePositionMultiple([]string{userId}, []string{contract})
}
func (ws *DeliveryWsStreamClient) UnSubscribePositionMultiple(userIds, contracts []string) error {
	channel := "futures.positions"
	payloads := [][]string{}
	for _, u := range userIds {
		for _, c := range contracts {
			payload := []string{u, c}
			payloads = append(payloads, payload)
		}
	}
	subKeys := userIds

	var thisSubs []*Subscription[WsSubscribeResult[[]WsFuturesPosition]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesPosition]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, true)
			if err != nil {
				return err
			}
			ws.futuresPositionSubMap.Store(payload[0], thisSub)
			mu.Lock()
			thisSubs = append(thisSubs, thisSub)
			mu.Unlock()
			return nil
		})
	}

	if err := errG.Wait(); err != nil {
		return err
	}

	return nil
}
