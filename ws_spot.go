package mygateapi

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
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

// !all订阅全部订单推送
func (ws *SpotWsStreamClient) SubscribeOrders(symbols ...string) (*MultipleSubscription[WsSubscribeResult[WsOrder]], error) {
	//订阅订单推送频道
	channel := "spot.orders"
	if len(symbols) == 0 {
		symbols = []string{"!all"}
	}

	var thisSubs []*Subscription[WsSubscribeResult[WsOrder]]
	payload := symbols
	subKeys := symbols

	thisSub, err := subscribe[WsSubscribeResult[WsOrder]](&ws.WsStreamClient, channel, SUBSCRIBE, payload, subKeys, true)
	if err != nil {
		return nil, err
	}
	ws.orderSubMap.Store(subKeys[0], thisSub)
	thisSubs = append(thisSubs, thisSub)

	subscription, err := mergeSubscription(thisSubs...)
	if err != nil {
		return nil, err
	}

	return subscription, nil
}
