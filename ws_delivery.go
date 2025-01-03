package mygateapi

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
)

const (
	WS_DELIVERY_ORDERBOOK_LIMIT_1   = "1"
	WS_DELIVERY_ORDERBOOK_LIMIT_5   = "5"
	WS_DELIVERY_ORDERBOOK_LIMIT_10  = "10"
	WS_DELIVERY_ORDERBOOK_LIMIT_20  = "20"
	WS_DELIVERY_ORDERBOOK_LIMIT_50  = "50"
	WS_DELIVERY_ORDERBOOK_LIMIT_100 = "100"

	WS_DELIVERY_ORDERBOOK_FREQUENCY_1000 = "1000"
	WS_DELIVERY_ORDERBOOK_FREQUENCY_100  = "100"

	WS_DELIVERY_ORDERBOOK_LEVEL_5   = "5"
	WS_DELIVERY_ORDERBOOK_LEVEL_10  = "10"
	WS_DELIVERY_ORDERBOOK_LEVEL_20  = "20"
	WS_DELIVERY_ORDERBOOK_LEVEL_50  = "50"
	WS_DELIVERY_ORDERBOOK_LEVEL_100 = "100"

	WS_DELIVERY_ORDERBOOK_INTERVAL_0 = "0"
)

// 订阅单个K线 如: "BTCUSDT","1m"
// 如果合约字段 contract 包含 mark_ 前缀的话, 将订阅合同的标记价格 K 线数据。
func (ws *DeliveryWsStreamClient) SubscribeCandle(contract, interval string) (*MultipleSubscription[WsSubscribeResult[WsCandles]], error) {
	return ws.SubscribeCandleMultiple([]string{contract}, []string{interval})
}
func (ws *DeliveryWsStreamClient) SubscribeCandleMultiple(contracts, intervals []string) (*MultipleSubscription[WsSubscribeResult[WsCandles]], error) {
	channel := "futures.candlesticks"
	payloads := [][]string{}
	for _, s := range contracts {
		for _, i := range intervals {
			payload := []string{i, s}
			payloads = append(payloads, payload)
		}
	}

	var thisSubs []*Subscription[WsSubscribeResult[WsCandles]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			if len(payload) != 2 {
				return fmt.Errorf("invalid payload: %v", payload)
			}
			subKeys := []string{fmt.Sprintf("%s_%s", payload[0], payload[1])}
			thisSub, err := subscribe[WsSubscribeResult[WsCandles]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, false)
			if err != nil {
				return err
			}
			ws.candleSubMap.Store(subKeys[0], thisSub)
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

// 订阅深度全量更新频道
func (ws *DeliveryWsStreamClient) SubscribeOrderBook(contract, limit, interval string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBook]], error) {
	return ws.SubscribeOrderBookMultiple([]string{contract}, []string{limit}, []string{interval})
}
func (ws *DeliveryWsStreamClient) SubscribeOrderBookMultiple(contracts, limit, interval []string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBook]], error) {
	channel := "futures.order_book"
	payloads := [][]string{}
	for _, c := range contracts {
		for _, l := range limit {
			for _, i := range interval {
				payload := []string{c, l, i}
				payloads = append(payloads, payload)
			}
		}
	}
	subKeys := contracts

	var thisSubs []*Subscription[WsSubscribeResult[WsFuturesOrderBook]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			thisSub, err := subscribe[WsSubscribeResult[WsFuturesOrderBook]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, false)
			if err != nil {
				return err
			}
			ws.futuresOrderBookSubMap.Store(payload[0], thisSub)
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
func (ws *DeliveryWsStreamClient) UnSubscribeOrderBook(contract, limit, interval string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBook]], error) {
	return ws.UnSubscribeOrderBookMultiple([]string{contract}, []string{limit}, []string{interval})
}
func (ws *DeliveryWsStreamClient) UnSubscribeOrderBookMultiple(contracts, limit, interval []string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBook]], error) {
	channel := "futures.order_book"
	payloads := [][]string{}
	for _, c := range contracts {
		for _, l := range limit {
			for _, i := range interval {
				payload := []string{c, l, i}
				payloads = append(payloads, payload)
			}
		}
	}
	subKeys := contracts

	var thisSubs []*Subscription[WsSubscribeResult[WsFuturesOrderBook]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			thisSub, err := subscribe[WsSubscribeResult[WsFuturesOrderBook]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, false)
			if err != nil {
				return err
			}
			ws.futuresOrderBookSubMap.Store(payload[0], thisSub)
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

// 订阅深度增量更新频道
func (ws *DeliveryWsStreamClient) SubscribeOrderBookUpdate(contract, frequency, interval string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBookUpdate]], error) {
	return ws.SubscribeOrderBookUpdateMultiple([]string{contract}, []string{frequency}, []string{interval})
}
func (ws *DeliveryWsStreamClient) SubscribeOrderBookUpdateMultiple(contracts, frequency, interval []string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBookUpdate]], error) {
	channel := "futures.order_book_update"
	payloads := [][]string{}
	for _, c := range contracts {
		for _, f := range frequency {
			for _, i := range interval {
				payload := []string{c, f, i}
				payloads = append(payloads, payload)
			}
		}
	}
	subKeys := contracts

	var thisSubs []*Subscription[WsSubscribeResult[WsFuturesOrderBookUpdate]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			thisSub, err := subscribe[WsSubscribeResult[WsFuturesOrderBookUpdate]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, false)
			if err != nil {
				return err
			}
			ws.futuresOrderBookUpdateSubMap.Store(payload[0], thisSub)
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
func (ws *DeliveryWsStreamClient) UnSubscribeOrderBookUpdate(contract, frequency, interval string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBookUpdate]], error) {
	return ws.UnSubscribeOrderBookUpdateMultiple([]string{contract}, []string{frequency}, []string{interval})
}
func (ws *DeliveryWsStreamClient) UnSubscribeOrderBookUpdateMultiple(contracts, frequency, interval []string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBookUpdate]], error) {
	channel := "futures.order_book_update"
	payloads := [][]string{}
	for _, c := range contracts {
		for _, f := range frequency {
			for _, i := range interval {
				payload := []string{c, f, i}
				payloads = append(payloads, payload)
			}
		}
	}
	subKeys := contracts

	var thisSubs []*Subscription[WsSubscribeResult[WsFuturesOrderBookUpdate]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			thisSub, err := subscribe[WsSubscribeResult[WsFuturesOrderBookUpdate]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, false)
			if err != nil {
				return err
			}
			ws.futuresOrderBookUpdateSubMap.Store(payload[0], thisSub)
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

// 订阅交易频道
func (ws *DeliveryWsStreamClient) SubscribeTrade(contract string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesTrade]], error) {
	return ws.SubscribeTradeMultiple([]string{contract})
}
func (ws *DeliveryWsStreamClient) SubscribeTradeMultiple(contracts []string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesTrade]], error) {
	channel := "futures.trades"
	subKeys := contracts

	var thisSubs []*Subscription[WsSubscribeResult[[]WsFuturesTrade]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, c := range contracts {
		c := c
		errG.Go(func() error {
			thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesTrade]](&ws.WsStreamClient, channel, SUBSCRIBE, []string{c}, subKeys, true)
			if err != nil {
				return err
			}
			ws.futuresTradeSubMap.Store(c, thisSub)
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
func (ws *DeliveryWsStreamClient) UnSubscribeTrade(contract string) error {
	return ws.UnSubscribeTradeMultiple([]string{contract})
}
func (ws *DeliveryWsStreamClient) UnSubscribeTradeMultiple(contracts []string) error {
	channel := "futures.trades"
	subKeys := contracts

	var thisSubs []*Subscription[WsSubscribeResult[[]WsFuturesTrade]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, c := range contracts {
		c := c
		errG.Go(func() error {
			thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesTrade]](&ws.WsStreamClient, channel, UNSUBSCRIBE, []string{c}, subKeys, false)
			if err != nil {
				return err
			}
			ws.futuresTradeSubMap.Store(c, thisSub)
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
