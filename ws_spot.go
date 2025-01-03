package mygateapi

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
)

const (
	// Kline intervals
	WS_CANDLE_10S = "10s"
	WS_CANDLE_1M  = "1m"
	WS_CANDLE_5M  = "5m"
	WS_CANDLE_15M = "15m"
	WS_CANDLE_30M = "30m"
	WS_CANDLE_1H  = "1h"
	WS_CANDLE_4H  = "4h"
	WS_CANDLE_8H  = "8h"
	WS_CANDLE_1D  = "1d"
	WS_CANDLE_7D  = "7d"

	// Orderbook levels
	WS_SPOT_ORDERBOOK_LEVEL_5   = "5"
	WS_SPOT_ORDERBOOK_LEVEL_10  = "10"
	WS_SPOT_ORDERBOOK_LEVEL_20  = "20"
	WS_SPOT_ORDERBOOK_LEVEL_30  = "30"
	WS_SPOT_ORDERBOOK_LEVEL_50  = "50"
	WS_SPOT_ORDERBOOK_LEVEL_100 = "100"

	// Orderbook interval
	WS_SPOT_ORDERBOOK_INTERVAL_100MS  = "100ms"
	WS_SPOT_ORDERBOOK_INTERVAL_1000MS = "1000ms"
)

// 订阅单个K线 如: "BTCUSDT","1m"
func (ws *SpotWsStreamClient) SubscribeCandle(symbol string, interval string) (*MultipleSubscription[WsSubscribeResult[WsCandles]], error) {
	return ws.SubscribeCandleMultiple([]string{symbol}, []string{interval})
}
func (ws *SpotWsStreamClient) SubscribeCandleMultiple(symbols, intervals []string) (*MultipleSubscription[WsSubscribeResult[WsCandles]], error) {
	//订阅现货K线频道
	channel := "spot.candlesticks"
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

// 深度最优买卖价订阅
func (ws *SpotWsStreamClient) SubscribeBookTicker(symbol string) (*MultipleSubscription[WsSubscribeResult[WsSpotBookTicker]], error) {
	return ws.SubscribeBookTickerMultiple(symbol)
}
func (ws *SpotWsStreamClient) SubscribeBookTickerMultiple(symbols ...string) (*MultipleSubscription[WsSubscribeResult[WsSpotBookTicker]], error) {
	channel := "spot.book_ticker"
	payload := symbols
	subKeys := symbols

	var thisSubs []*Subscription[WsSubscribeResult[WsSpotBookTicker]]
	thisSub, err := subscribe[WsSubscribeResult[WsSpotBookTicker]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, false)
	if err != nil {
		return nil, err
	}
	for _, subKey := range subKeys {
		ws.spotBookTickerSubMap.Store(subKey, thisSub)
	}
	thisSubs = append(thisSubs, thisSub)

	subscription, err := mergeSubscription(thisSubs...)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
func (ws *SpotWsStreamClient) UnSubscribeBookTicker(symbol string) error {
	return ws.UnSubscribeBookTickerMultiple(symbol)
}
func (ws *SpotWsStreamClient) UnSubscribeBookTickerMultiple(symbols ...string) error {
	channel := "spot.book_ticker"
	subKeys := symbols
	thisSub, err := subscribe[WsSubscribeResult[WsSpotBookTicker]](&ws.WsStreamClient, channel, UNSUBSCRIBE, symbols, subKeys, false)
	if err != nil {
		return err
	}
	log.Infof("UnSubscribeBooks Success: args:%v", thisSub.Res)
	for _, subKey := range subKeys {
		ws.spotBookTickerSubMap.Delete(subKey)
	}

	return nil
}

// 深度增量更新频道订阅
func (ws *SpotWsStreamClient) SubscribeOrderBookUpdate(symbol, interval string) (*MultipleSubscription[WsSubscribeResult[WsSpotOrderBookUdpate]], error) {
	return ws.SubscribeOrderBookUpdateMultiple([]string{symbol}, []string{interval})
}
func (ws *SpotWsStreamClient) SubscribeOrderBookUpdateMultiple(symbols, intervals []string) (*MultipleSubscription[WsSubscribeResult[WsSpotOrderBookUdpate]], error) {
	channel := "spot.order_book_update"
	payloads := [][]string{}
	for _, s := range symbols {
		for _, i := range intervals {
			payload := []string{s, i}
			payloads = append(payloads, payload)
		}
	}

	var thisSubs []*Subscription[WsSubscribeResult[WsSpotOrderBookUdpate]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			if len(payload) != 2 {
				return fmt.Errorf("invalid payload: %v", payload)
			}
			subKeys := []string{payload[0]}
			thisSub, err := subscribe[WsSubscribeResult[WsSpotOrderBookUdpate]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, false)
			if err != nil {
				return err
			}
			ws.spotOrderBookUpdateSubMap.Store(subKeys[0], thisSub)
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
func (ws *SpotWsStreamClient) UnSubscribeOrderBookUpdate(symbol, interval string) error {
	return ws.UnSubscribeOrderBookUpdateMultiple([]string{symbol}, []string{interval})
}
func (ws *SpotWsStreamClient) UnSubscribeOrderBookUpdateMultiple(symbols, intervals []string) error {
	channel := "spot.order_book_update"
	payloads := [][]string{}
	for _, s := range symbols {
		for _, i := range intervals {
			payload := []string{s, i}
			payloads = append(payloads, payload)
		}
	}

	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			if len(payload) != 2 {
				return fmt.Errorf("invalid payload: %v", payload)
			}
			subKeys := []string{payload[0]}
			thisSub, err := subscribe[WsSubscribeResult[WsSpotOrderBookUdpate]](&ws.WsStreamClient, channel, UNSUBSCRIBE, payload, subKeys, false)
			if err != nil {
				return err
			}
			log.Infof("UnSubscribeOrderBookUpdate Success: args:%v", thisSub.Res)
			ws.spotOrderBookUpdateSubMap.Delete(subKeys[0])
			return nil
		})
	}

	if err := errG.Wait(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// 深度全量更新频道订阅
func (ws *SpotWsStreamClient) SubscribeOrderBook(symbol, level, interval string) (*MultipleSubscription[WsSubscribeResult[WsSpotOrderBook]], error) {
	return ws.SubscribeOrderBookMultiple([]string{symbol}, []string{level}, []string{interval})
}
func (ws *SpotWsStreamClient) SubscribeOrderBookMultiple(symbols, level, intervals []string) (*MultipleSubscription[WsSubscribeResult[WsSpotOrderBook]], error) {
	channel := "spot.order_book"
	payloads := [][]string{}
	for _, s := range symbols {
		for _, l := range level {
			for _, i := range intervals {
				payload := []string{s, l, i}
				payloads = append(payloads, payload)
			}
		}
	}

	var thisSubs []*Subscription[WsSubscribeResult[WsSpotOrderBook]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			if len(payload) != 3 {
				return fmt.Errorf("invalid payload: %v", payload)
			}
			subKey := []string{payload[0]}
			thisSub, err := subscribe[WsSubscribeResult[WsSpotOrderBook]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKey, false)
			if err != nil {
				return err
			}
			ws.spotOrderBookSubMap.Store(subKey[0], thisSub)
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
func (ws *SpotWsStreamClient) UnSubscribeOrderBook(symbol, level, interval string) error {
	return ws.UnSubscribeOrderBookMultiple([]string{symbol}, []string{level}, []string{interval})
}
func (ws *SpotWsStreamClient) UnSubscribeOrderBookMultiple(symbols, level, intervals []string) error {
	channel := "spot.order_book"
	payloads := [][]string{}
	for _, s := range symbols {
		for _, l := range level {
			for _, i := range intervals {
				payload := []string{s, l, i}
				payloads = append(payloads, payload)
			}
		}
	}

	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			if len(payload) != 3 {
				return fmt.Errorf("invalid payload: %v", payload)
			}
			subKey := []string{payload[0]}
			thisSub, err := subscribe[WsSubscribeResult[WsSpotOrderBook]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKey, false)
			if err != nil {
				return err
			}
			log.Infof("UnSubscribeOrderBook Success: args:%v", thisSub.Res)
			ws.spotOrderBookSubMap.Delete(subKey[0])
			return nil
		})
	}

	if err := errG.Wait(); err != nil {
		return err
	}

	return nil
}

// ticker订阅
func (ws *SpotWsStreamClient) SubscribeTicker(symbol string) (*MultipleSubscription[WsSubscribeResult[WsSpotTicker]], error) {
	return ws.SubscribeTickerMultiple(symbol)
}
func (ws *SpotWsStreamClient) SubscribeTickerMultiple(symbols ...string) (*MultipleSubscription[WsSubscribeResult[WsSpotTicker]], error) {
	channel := "spot.tickers"
	payload := symbols
	subKeys := symbols

	var thisSubs []*Subscription[WsSubscribeResult[WsSpotTicker]]
	thisSub, err := subscribe[WsSubscribeResult[WsSpotTicker]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, false)
	if err != nil {
		return nil, err
	}
	for _, subKey := range subKeys {
		ws.spotTickerSubMap.Store(subKey, thisSub)
	}
	thisSubs = append(thisSubs, thisSub)

	subscription, err := mergeSubscription(thisSubs...)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
func (ws *SpotWsStreamClient) UnSubscribeTicker(symbol string) error {
	return ws.UnSubscribeTickerMultiple(symbol)
}
func (ws *SpotWsStreamClient) UnSubscribeTickerMultiple(symbols ...string) error {
	channel := "spot.tickers"
	subKeys := symbols
	thisSub, err := subscribe[WsSubscribeResult[WsSpotTicker]](&ws.WsStreamClient, channel, UNSUBSCRIBE, symbols, subKeys, false)
	if err != nil {
		return err
	}
	log.Infof("UnSubscribeTicker Success: args:%v", thisSub.Res)
	for _, subKey := range subKeys {
		ws.spotTickerSubMap.Delete(subKey)
	}

	return nil
}

// 公共成交频道订阅
func (ws *SpotWsStreamClient) SubscribeTrade(symbol string) (*MultipleSubscription[WsSubscribeResult[WsSpotTrade]], error) {
	return ws.SubscribeTradeMultiple([]string{symbol})
}
func (ws *SpotWsStreamClient) SubscribeTradeMultiple(symbols []string) (*MultipleSubscription[WsSubscribeResult[WsSpotTrade]], error) {
	channel := "spot.trades"
	payload := symbols
	subKeys := symbols

	var thisSubs []*Subscription[WsSubscribeResult[WsSpotTrade]]
	thisSub, err := subscribe[WsSubscribeResult[WsSpotTrade]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, false)
	if err != nil {
		return nil, err
	}
	for _, subKey := range subKeys {
		ws.spotTradeSubMap.Store(subKey, thisSub)
	}
	thisSubs = append(thisSubs, thisSub)

	subscription, err := mergeSubscription(thisSubs...)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
func (ws *SpotWsStreamClient) UnSubscribeTrade(symbol string) error {
	return ws.UnSubscribeTradeMultiple([]string{symbol})
}
func (ws *SpotWsStreamClient) UnSubscribeTradeMultiple(symbols []string) error {
	channel := "spot.trades"
	subKeys := symbols
	thisSub, err := subscribe[WsSubscribeResult[WsSpotTrade]](&ws.WsStreamClient, channel, UNSUBSCRIBE, symbols, subKeys, false)
	if err != nil {
		return err
	}
	log.Infof("UnSubscribeTrade Success: args:%v", thisSub.Res)
	for _, subKey := range subKeys {
		ws.spotTradeSubMap.Delete(subKey)
	}

	return nil
}

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
