package mygateapi

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
)

const (
	WS_FUTURES_ORDERBOOK_DEPTH_LIMIT_1   = "1"
	WS_FUTURES_ORDERBOOK_DEPTH_LIMIT_5   = "5"
	WS_FUTURES_ORDERBOOK_DEPTH_LIMIT_10  = "10"
	WS_FUTURES_ORDERBOOK_DEPTH_LIMIT_20  = "20"
	WS_FUTURES_ORDERBOOK_DEPTH_LIMIT_50  = "50"
	WS_FUTURES_ORDERBOOK_DEPTH_LIMIT_100 = "100"

	WS_FUTURES_ORDERBOOK_FREQUENCY_100MS  = "100ms"
	WS_FUTURES_ORDERBOOK_FREQUENCY_1000MS = "1000ms"

	WS_FUTURES_ORDERBOOK_LEVEL_20  = "20"
	WS_FUTURES_ORDERBOOK_LEVEL_50  = "50"
	WS_FUTURES_ORDERBOOK_LEVEL_100 = "100"

	WS_FUTURES_ORDERBOOK_INTERVAL_0 = "0"
)

// 订阅单个K线 如: "BTCUSDT","1m"
func (ws *FuturesWsStreamClient) SubscribeCandle(symbol string, interval string) (*MultipleSubscription[WsSubscribeResult[WsCandles]], error) {
	return ws.SubscribeCandleMultiple([]string{symbol}, []string{interval})
}
func (ws *FuturesWsStreamClient) SubscribeCandleMultiple(symbols, intervals []string) (*MultipleSubscription[WsSubscribeResult[WsCandles]], error) {
	//订阅现货K线频道
	channel := "futures.candlesticks"
	payloads := [][]string{}
	for _, s := range symbols {
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

// 订阅ticker频道
func (ws *FuturesWsStreamClient) SubscribeTicker(symbol string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesTicker]], error) {
	return ws.SubscribeTickerMultiple([]string{symbol})
}
func (ws *FuturesWsStreamClient) SubscribeTickerMultiple(symbols []string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesTicker]], error) {
	channel := "futures.tickers"
	payload := symbols
	subKeys := symbols

	var thisSubs []*Subscription[WsSubscribeResult[[]WsFuturesTicker]]
	thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesTicker]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, false)
	if err != nil {
		return nil, err
	}
	thisSubs = append(thisSubs, thisSub)
	for _, subKey := range subKeys {
		ws.futuresTickerSubMap.Store(subKey, thisSub)
	}

	subscription, err := mergeSubscription(thisSubs...)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
func (ws *FuturesWsStreamClient) UnSubscribeTicker(symbol string) error {
	return ws.UnSubscribeTickerMultiple([]string{symbol})
}
func (ws *FuturesWsStreamClient) UnSubscribeTickerMultiple(symbols []string) error {
	channel := "futures.tickers"
	payload := symbols
	subKeys := symbols

	_, err := subscribe[WsSubscribeResult[WsFuturesTicker]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, false)
	if err != nil {
		return err
	}
	for _, subKey := range subKeys {
		ws.futuresTickerSubMap.Delete(subKey)
	}
	return nil
}

// 订阅深度全量更新频道
func (ws *FuturesWsStreamClient) SubscribeOrderBook(contract, limit, interval string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBook]], error) {
	return ws.SubscribeOrderBookMultiple([]string{contract}, []string{limit}, []string{interval})
}
func (ws *FuturesWsStreamClient) SubscribeOrderBookMultiple(contracts, limit, interval []string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBook]], error) {
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
func (ws *FuturesWsStreamClient) UnSubscribeOrderBook(contract, limit, interval string) error {
	return ws.UnSubscribeOrderBookMultiple([]string{contract}, []string{limit}, []string{interval})
}
func (ws *FuturesWsStreamClient) UnSubscribeOrderBookMultiple(contracts, limit, interval []string) error {
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

	for _, payload := range payloads {
		_, err := subscribe[WsSubscribeResult[WsFuturesOrderBook]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, false)
		if err != nil {
			return err
		}

		ws.futuresOrderBookSubMap.Delete(payload[0])
	}

	return nil
}

// 订阅深度更新推送频道
func (ws *FuturesWsStreamClient) SubscribeOrderBookUpdate(contract, frequency, interval string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBookUpdate]], error) {
	return ws.SubscribeOrderBookUpdateMultiple([]string{contract}, []string{frequency}, []string{interval})
}
func (ws *FuturesWsStreamClient) SubscribeOrderBookUpdateMultiple(contracts, frequency, interval []string) (*MultipleSubscription[WsSubscribeResult[WsFuturesOrderBookUpdate]], error) {
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
func (ws *FuturesWsStreamClient) UnSubscribeOrderBookUpdate(contract, frequency, interval string) error {
	return ws.UnSubscribeOrderBookUpdateMultiple([]string{contract}, []string{frequency}, []string{interval})
}
func (ws *FuturesWsStreamClient) UnSubscribeOrderBookUpdateMultiple(contracts, frequency, interval []string) error {
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

	for _, payload := range payloads {
		_, err := subscribe[WsSubscribeResult[WsFuturesOrderBookUpdate]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, false)
		if err != nil {
			return err
		}

		ws.futuresOrderBookUpdateSubMap.Delete(payload[0])
	}

	return nil
}

// 订阅共有成交频道
func (ws *FuturesWsStreamClient) SubscribeTrade(contract string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesTrade]], error) {
	return ws.SubscribeTradeMultiple([]string{contract})
}
func (ws *FuturesWsStreamClient) SubscribeTradeMultiple(contracts []string) (*MultipleSubscription[WsSubscribeResult[[]WsFuturesTrade]], error) {
	channel := "futures.trades"
	payload := contracts
	subKeys := contracts

	var thisSubs []*Subscription[WsSubscribeResult[[]WsFuturesTrade]]
	thisSub, err := subscribe[WsSubscribeResult[[]WsFuturesTrade]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, false)
	if err != nil {
		return nil, err
	}
	thisSubs = append(thisSubs, thisSub)
	for _, subKey := range subKeys {
		ws.futuresTradeSubMap.Store(subKey, thisSub)
	}

	subscription, err := mergeSubscription(thisSubs...)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
func (ws *FuturesWsStreamClient) UnSubscribeTrade(contract string) error {
	return ws.UnSubscribeTradeMultiple([]string{contract})
}
func (ws *FuturesWsStreamClient) UnSubscribeTradeMultiple(contracts []string) error {
	channel := "futures.trades"
	payload := contracts
	subKeys := contracts

	_, err := subscribe[WsSubscribeResult[[]WsFuturesTrade]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, true)
	if err != nil {
		return err
	}
	for _, subKey := range subKeys {
		ws.futuresTradeSubMap.Delete(subKey)
	}
	return nil
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
