package mygateapi

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
)

// K线通用订阅
func subscribeCandleCommon(ws *WsStreamClient, isSubscribe bool, symbols, intervals []string) (*MultipleSubscription[WsSubscribeResult[WsCandles]], error) {
	channel := ""
	switch ws.apiType {
	case WS_SPOT:
		channel = "spot.candlesticks"
	case WS_FUTURES, WS_DELIVERY:
		channel = "futures.candlesticks"
	default:
		return nil, fmt.Errorf("invalid apiType: %v", ws.apiType)
	}

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

			if isSubscribe {
				thisSub, err := subscribe[WsSubscribeResult[WsCandles]](ws, channel, SUBSCRIBE, payload, subKeys, false)
				if err != nil {
					return err
				}
				ws.candleSubMap.Store(subKeys[0], thisSub)
				mu.Lock()
				thisSubs = append(thisSubs, thisSub)
				mu.Unlock()
			} else {
				_, err := subscribe[WsSubscribeResult[WsCandles]](ws, channel, SUBSCRIBE, payload, subKeys, false)
				if err != nil {
					return err
				}
				ws.candleSubMap.Delete(subKeys[0])
			}
			return nil
		})
	}

	if err := errG.Wait(); err != nil {
		return nil, err
	}

	if isSubscribe {
		subscription, err := mergeSubscription(thisSubs...)
		if err != nil {
			return nil, err
		}

		return subscription, nil
	} else {
		return nil, nil
	}

}

func (ws *WsStreamClient) SubscribeCandle(symbol string, interval string) (*MultipleSubscription[WsSubscribeResult[WsCandles]], error) {
	return ws.SubscribeCandleMultiple([]string{symbol}, []string{interval})
}
func (ws *WsStreamClient) UnSubscribeCandle(symbol string, interval string) error {
	return ws.UnSubscribeCandleMultiple([]string{symbol}, []string{interval})
}
func (ws *WsStreamClient) SubscribeCandleMultiple(symbols, intervals []string) (*MultipleSubscription[WsSubscribeResult[WsCandles]], error) {
	return subscribeCandleCommon(ws, true, symbols, intervals)
}
func (ws *WsStreamClient) UnSubscribeCandleMultiple(symbols, intervals []string) error {
	_, err := subscribeCandleCommon(ws, false, symbols, intervals)
	return err
}

// 公共成交流通用订阅
func subscribeTradeCommon(ws *WsStreamClient, isSubscribe bool, symbols []string) (*MultipleSubscription[WsSubscribeResult[WsTrade]], error) {
	channel := ""
	switch ws.apiType {
	case WS_SPOT:
		channel = "spot.trades"
	case WS_FUTURES, WS_DELIVERY:
		channel = "futures.trades"
	default:
		return nil, fmt.Errorf("invalid apiType: %v", ws.apiType)
	}

	payload := symbols
	subKeys := symbols

	var thisSubs []*Subscription[WsSubscribeResult[WsTrade]]

	if isSubscribe {
		thisSub, err := subscribe[WsSubscribeResult[WsTrade]](ws, channel, SUBSCRIBE, payload, subKeys, false)
		if err != nil {
			return nil, err
		}
		for _, subKey := range subKeys {
			ws.tradeSubMap.Store(subKey, thisSub)
		}
		thisSubs = append(thisSubs, thisSub)
		subscription, err := mergeSubscription(thisSubs...)
		if err != nil {
			return nil, err
		}
		return subscription, nil
	} else {
		_, err := subscribe[WsSubscribeResult[WsTrade]](ws, channel, UNSUBSCRIBE, payload, subKeys, false)
		if err != nil {
			return nil, err
		}
		for _, subKey := range subKeys {
			ws.tradeSubMap.Delete(subKey)
		}
		return nil, nil
	}
}

func (ws *WsStreamClient) SubscribeTrade(symbol string) (*MultipleSubscription[WsSubscribeResult[WsTrade]], error) {
	return ws.SubscribeTradeMultiple([]string{symbol})
}
func (ws *WsStreamClient) UnSubscribeTrade(symbol string) error {
	return ws.UnSubscribeTradeMultiple([]string{symbol})
}
func (ws *WsStreamClient) SubscribeTradeMultiple(symbols []string) (*MultipleSubscription[WsSubscribeResult[WsTrade]], error) {
	return subscribeTradeCommon(ws, true, symbols)
}
func (ws *WsStreamClient) UnSubscribeTradeMultiple(symbols []string) error {
	_, err := subscribeTradeCommon(ws, false, symbols)
	return err
}

// 深度通用订阅
func subscribeOrderBookCommon(ws *WsStreamClient, isSubscribe, isUpdate bool, symbols []string, USpeed string, level string) (*MultipleSubscription[WsSubscribeResult[WsOrderBook]], error) {

	if USpeed == "" {
		USpeed = "100ms"
	}

	if level == "" {
		level = "20"
	}

	payloadLen := 0
	channel := ""
	payloads := [][]string{}
	switch ws.apiType {
	case WS_SPOT:
		if isUpdate {
			channel = "spot.order_book_update"
			for _, s := range symbols {
				payload := []string{s, USpeed}
				payloads = append(payloads, payload)
			}
			payloadLen = 2
		} else {
			channel = "spot.order_book"
			for _, s := range symbols {
				payload := []string{s, level, USpeed}
				payloads = append(payloads, payload)
			}
			payloadLen = 3
		}
	case WS_FUTURES, WS_DELIVERY:
		if isUpdate {
			channel = "futures.order_book_update"
			for _, s := range symbols {
				payload := []string{s, USpeed, level}
				payloads = append(payloads, payload)
			}
		} else {
			channel = "futures.order_book"
			for _, s := range symbols {
				payload := []string{s, level, "0"}
				payloads = append(payloads, payload)
			}
		}
		payloadLen = 3
	default:
		return nil, fmt.Errorf("invalid apiType: %v", ws.apiType)
	}

	var thisSubs []*Subscription[WsSubscribeResult[WsOrderBook]]
	var mu sync.Mutex
	var errG errgroup.Group
	for _, payload := range payloads {
		payload := payload
		errG.Go(func() error {
			if len(payload) != payloadLen {
				return fmt.Errorf("invalid payload: %v", payload)
			}
			subKeys := []string{payload[0]}
			if isSubscribe {
				thisSub, err := subscribe[WsSubscribeResult[WsOrderBook]](ws, channel, SUBSCRIBE, payload, subKeys, false)
				if err != nil {
					return err
				}
				ws.orderBookSubMap.Store(subKeys[0], thisSub)
				mu.Lock()
				thisSubs = append(thisSubs, thisSub)
				mu.Unlock()
			} else {
				_, err := subscribe[WsSubscribeResult[WsOrderBook]](ws, channel, UNSUBSCRIBE, payload, subKeys, false)
				if err != nil {
					return err
				}
				ws.orderBookSubMap.Delete(subKeys[0])
			}

			return nil
		})
	}

	if err := errG.Wait(); err != nil {
		return nil, err
	}

	if isSubscribe {
		subscription, err := mergeSubscription(thisSubs...)
		if err != nil {
			return nil, err
		}

		return subscription, nil
	} else {
		return nil, nil
	}
}

func (ws *WsStreamClient) SubscribeOrderBookUpdate(symbol, USpeed, level string) (*MultipleSubscription[WsSubscribeResult[WsOrderBook]], error) {
	return ws.SubscribeOrderBookUpdateMultiple([]string{symbol}, USpeed, level)
}
func (ws *WsStreamClient) UnSubscribeOrderBookUpdate(symbol, USpeed, level string) error {
	return ws.UnSubscribeOrderBookUpdateMultiple([]string{symbol}, USpeed, level)
}
func (ws *WsStreamClient) SubscribeOrderBookUpdateMultiple(symbols []string, USpeed, level string) (*MultipleSubscription[WsSubscribeResult[WsOrderBook]], error) {
	return subscribeOrderBookCommon(ws, true, true, symbols, USpeed, level)
}
func (ws *WsStreamClient) UnSubscribeOrderBookUpdateMultiple(symbols []string, USpeed, level string) error {
	_, err := subscribeOrderBookCommon(ws, false, true, symbols, USpeed, level)
	return err
}

func (ws *WsStreamClient) SubscribeOrderBook(symbol, USpeed, level string) (*MultipleSubscription[WsSubscribeResult[WsOrderBook]], error) {
	return ws.SubscribeOrderBookMultiple([]string{symbol}, USpeed, level)
}
func (ws *WsStreamClient) UnSubscribeOrderBook(symbol, USpeed, level string) error {
	return ws.UnSubscribeOrderBookMultiple([]string{symbol}, USpeed, level)
}
func (ws *WsStreamClient) SubscribeOrderBookMultiple(symbols []string, USpeed, level string) (*MultipleSubscription[WsSubscribeResult[WsOrderBook]], error) {
	return subscribeOrderBookCommon(ws, true, false, symbols, USpeed, level)
}
func (ws *WsStreamClient) UnSubscribeOrderBookMultiple(symbols []string, USpeed, level string) error {
	_, err := subscribeOrderBookCommon(ws, false, false, symbols, USpeed, level)
	return err
}
