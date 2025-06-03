package mygateapi

import (
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"
)

// 订阅订单频道
func (ws *FutureAndDeliveryWsStreamClient) SubscribeOrders(symbols ...string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesOrder]], error) {
	return ws.SubscribeOrderMultiple(symbols...)
}
func (ws *FutureAndDeliveryWsStreamClient) SubscribeOrderMultiple(symbols ...string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesOrder]], error) {
	channel := "futures.orders"
	payloads := [][]string{}

	err := ws.CheckUserId()
	if err != nil {
		return nil, err
	}

	if len(symbols) == 0 {
		symbols = []string{"!all"}
	}

	for _, c := range symbols {
		payload := []string{strconv.FormatInt(ws.userId, 10), c}
		payloads = append(payloads, payload)
	}

	subKeys := symbols

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
func (ws *FutureAndDeliveryWsStreamClient) UnSubscribeOrder(symbols ...string) error {
	return ws.UnSubscribeOrderMultiple(symbols...)
}
func (ws *FutureAndDeliveryWsStreamClient) UnSubscribeOrderMultiple(symbols ...string) error {
	channel := "futures.orders"
	payloads := [][]string{}

	if len(symbols) == 0 {
		symbols = []string{"!all"}
	}
	for _, c := range symbols {
		payload := []string{strconv.FormatInt(ws.userId, 10), c}
		payloads = append(payloads, payload)
	}
	subKeys := symbols

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
func (ws *FutureAndDeliveryWsStreamClient) SubscribeBalance(userId string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesBalance]], error) {
	channel := "futures.balances"
	payload := []string{userId}
	subKey := []string{userId}

	thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesBalance]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKey, true)
	if err != nil {
		return nil, err
	}

	ws.futuresBalanceSubMap.Store(subKey[0], thisSub)
	subscription, err := mergeSubscription([]*Subscription[WsSubscribeResult[[]WsFuturesBalance]]{thisSub}...)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}
func (ws *FutureAndDeliveryWsStreamClient) UnSubscribeBalance(userId string) error {
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
func (ws *FutureAndDeliveryWsStreamClient) SubscribePosition(userId, contract string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesPosition]], error) {
	return ws.SubscribePositionMultiple(userId, []string{contract})
}
func (ws *FutureAndDeliveryWsStreamClient) SubscribePositionMultiple(userId string, contracts []string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesPosition]], error) {
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
func (ws *FutureAndDeliveryWsStreamClient) UnSubscribePosition(userId, contract string) error {
	return ws.UnSubscribePositionMultiple(userId, []string{contract})
}
func (ws *FutureAndDeliveryWsStreamClient) UnSubscribePositionMultiple(userId string, contracts []string) error {
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
