package mygateapi

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	GATE_API_WS_SPOT    = "api.gateio.ws"
	GATE_API_WS_FUTURES = "fx-ws.gateio.ws"
)
const (
	GATE_API_WS_SPOT_PATH         = "/ws/v4/"     //现货公共私有频道合一
	GATE_API_WS_FUTURES_PATH_USDT = "/v4/ws/usdt" //合约USDT结算路径
	GATE_API_WS_FUTURES_PATH_BTC  = "/v4/ws/btc"  //合约BTC结算路径
)

type ContractType string

const (
	USDT_CONTRACT ContractType = "USDT"
	BTC_CONTRACT  ContractType = "BTC"
)

// event类型
const (
	SUBSCRIBE   = "subscribe"   //订阅
	UNSUBSCRIBE = "unsubscribe" //取消订阅
)

var (
	WebsocketTimeout        = time.Second * 60
	WebsocketPingTicker     = time.Second * 5
	WebsocketKeepalive      = false
	SUBSCRIBE_INTERVAL_TIME = 500 * time.Millisecond //订阅间隔
)

var node *snowflake.Node

func init() {
	node, _ = snowflake.NewNode(33)
}

// 数据流订阅基础客户端
type WsStreamClient struct {
	client       *RestClient
	apiType      APIType
	contractType ContractType
	conn         *websocket.Conn
	connId       string

	waitSubscribeResMap MySyncMap[int64, *Subscription[WsSubscribeResult[WsSubscribeStatus]]] //等待订阅结果

	currentSubMap MySyncMap[int64, *Subscription[WsSubscribeResult[WsSubscribeStatus]]] //当前订阅成功的所有结果
	candleSubMap  MySyncMap[string, *Subscription[WsSubscribeResult[WsCandles]]]        //K线推送订阅频道

	resultChan chan []byte
	errChan    chan error
	isClose    bool

	AutoReConnectTimes int //自动重连次数
	writeMu            sync.Mutex
}

// 授权结构体，所需签名等结构
type WsAuth struct {
	Method string `json:"method"` //String 身份验证方法。目前只接受一个方法 api_key
	KEY    string `json:"KEY"`    //String Gate APIv4 用户 key 字符串
	SIGN   string `json:"SIGN"`   //String 使用 GateAPIv4 secret 和请求信息生成的认证签名
}

func (subReq *WsSubscribeReq) SetAuth(client *RestClient) {
	signData := fmt.Sprintf("channel=%s&event=%s&time=%d", subReq.Channel, subReq.Event, subReq.Time)
	sign := HmacSha512(client.c.ApiSecret, signData)
	subReq.Auth = &WsAuth{
		Method: "api_key",
		KEY:    client.c.ApiKey,
		SIGN:   sign,
	}
}

// 订阅请求结构体，订阅私有频道时wsAuth不可为空
type WsSubscribeReq struct {
	Time    int64    `json:"time"`              // Integer 是 精确到秒的请求时间。请求时间和服务器时间之间的间隔不能超过 60 秒。
	Id      int64    `json:"id"`                // Integer 否 请求 id（可选），服务器将此 id 发送回来，用来帮助您确定服务器响应哪个请求
	Channel string   `json:"channel"`           // String  是 需要订阅的 websocket 频道
	Event   string   `json:"event"`             // String  是 频道的操作事件。例如 subscribe, unsubscribe
	Payload []string `json:"payload,omitempty"` // Any     否 请求详细参数（可选）
	Auth    *WsAuth  `json:"auth,omitempty"`    // Auth    否 私有频道的身份验证凭证。详见 身份验证（Authentication） 细节部分
}

// 订阅返回结果，可返回正常推送数据、订阅失败结果、订阅成功结果
type WsSubscribeResult[T any] struct {
	Time    int64             `json:"time"`    //Integer	精确到秒的响应时间
	TimeMs  int64             `json:"time_ms"` //Integer	精确到毫秒的响应时间
	ConnId  string            `json:"conn_id"` //String	连接 ID
	Id      int64             `json:"id"`      //Integer	从客户端请求 payload 中提取的请求 ID （如果请求参数中有的话）
	Channel string            `json:"channel"` //String	WebSocket 频道名称
	Event   string            `json:"event"`   //String	服务端频道事件（即，update）或 用于从客户端发起的请求的 event
	Error   *WsSubscribeError `json:"error"`   //Error	如果服务端正常接受客户端的请求，则返回为空；否则，返回请求被拒绝的详情。
	Payload []string          `json:"payload"` //Any	返回来自服务端的新数据通知 或 对客户端请求的响应。如果有错误返回则 error 不为空，没有错误则此字段为空。
	Result  *T                `json:"result"`  //Any	返回来自服务端的新数据通知 或 对客户端请求的响应。如果有错误返回则 error 不为空，没有错误则此字段为空。
}

type WsSubscribeError struct {
	Code int    `json:"code"`    //String	错误代码
	Msg  string `json:"message"` //String	错误信息
}

type WsSubscribeStatus struct {
	Status string `json:"status"` //String	请求状态。如果请求成功，将返回 "success"；如果请求失败，将返回 "error"
}

// 数据流订阅标准结构体
type Subscription[T any] struct {
	SubId      int64                                 //订阅ID
	Ws         *WsStreamClient                       //订阅的连接
	Req        *WsSubscribeReq                       //订阅参数
	Res        *WsSubscribeResult[WsSubscribeStatus] //订阅返回结果
	Error      *WsSubscribeError                     //订阅错误的结果
	SubKeys    []string                              //订阅的自定义唯一key,如交易对 或 K线周期_交易对等格式的唯一字符串
	resultChan chan T                                //接收订阅结果的通道
	errChan    chan error                            //接收订阅错误的通道
	closeChan  chan struct{}                         //接收订阅关闭的通道
}

// 多订阅通道合一
type MultipleSubscription[T any] struct {
	Subs       []*Subscription[T]
	Ws         *WsStreamClient //订阅的连接
	SubKeys    []string        //订阅的自定义唯一key,如交易对 或 K线周期_交易对等格式的唯一字符串
	resultChan chan T          //接收订阅结果的通道
	errChan    chan error      //接收订阅错误的通道
	closeChan  chan struct{}   //接收订阅关闭的通道
}

// 获取订阅结果
func (sub *Subscription[T]) ResultChan() chan T {
	return sub.resultChan
}

// 获取错误订阅
func (sub *Subscription[T]) ErrChan() chan error {
	return sub.errChan
}

// 获取关闭订阅信号
func (sub *Subscription[T]) CloseChan() chan struct{} {
	return sub.closeChan
}

// 取消订阅
func (sub *Subscription[T]) Unsubscribe() error {
	_, err := subscribe[WsSubscribeResult[WsSubscribeStatus]](sub.Ws, sub.Req.Channel, UNSUBSCRIBE, sub.Req.Payload, sub.SubKeys, sub.Req.Auth == nil)
	if err != nil {
		return err
	}
	log.Debugf("Unsubscribe Success args:%v ", sub.Req)

	//取消订阅成功，给所有订阅消息的通道发送关闭信号
	sub.Ws.sendUnSubscribeSuccessToCloseChan(sub.SubId, sub.SubKeys)
	return nil
}

// 获取订阅结果
func (sub *MultipleSubscription[T]) ResultChan() chan T {
	return sub.resultChan
}

// 获取错误订阅
func (sub *MultipleSubscription[T]) ErrChan() chan error {
	return sub.errChan
}

// 获取关闭订阅信号
func (sub *MultipleSubscription[T]) CloseChan() chan struct{} {
	return sub.closeChan
}

func (sub *MultipleSubscription[T]) Unsubscribe() error {
	for _, s := range sub.Subs {
		err := s.Unsubscribe()
		if err != nil {
			return err
		}
	}
	return nil
}

func mergeSubscription[T any](subs ...*Subscription[T]) (*MultipleSubscription[T], error) {
	if len(subs) == 0 {
		return nil, errors.New("no subscription")
	}
	subKeys := []string{}
	for _, sub := range subs {
		subKeys = append(subKeys, sub.SubKeys...)
	}
	mSub := &MultipleSubscription[T]{
		Subs:       subs,
		Ws:         subs[0].Ws,
		SubKeys:    subKeys,
		resultChan: make(chan T, 100),
		errChan:    make(chan error, 50),
		closeChan:  make(chan struct{}, 50),
	}

	for _, s := range subs {
		sub := s
		go func() {
			for {
				select {
				case err := <-sub.ErrChan():
					mSub.errChan <- err
				case result := <-sub.ResultChan():
					mSub.resultChan <- result
				case <-sub.CloseChan():
					mSub.closeChan <- struct{}{}
				}
			}
		}()
	}

	return mSub, nil
}

func (ws *WsStreamClient) GetConn() *websocket.Conn {
	return ws.conn
}

func subscribe[T any](ws *WsStreamClient, channel, event string, payload []string, subKeys []string, needAuth bool) (*Subscription[T], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is close")
	}

	reqId := node.Generate().Int64()
	now := time.Now().Unix()

	subscribeReq := WsSubscribeReq{
		Time:    now,
		Id:      reqId,
		Channel: channel,
		Event:   event,
		Payload: payload,
	}

	if needAuth {
		subscribeReq.SetAuth(ws.client)
	}

	data, err := json.Marshal(subscribeReq)
	if err != nil {
		return nil, err
	}
	//发送请求前，构造一个接收订阅结果的订阅者
	waitSubResult := &Subscription[WsSubscribeResult[WsSubscribeStatus]]{
		SubId:      reqId,
		Ws:         ws,
		Req:        &subscribeReq,
		Res:        nil,
		Error:      nil,
		SubKeys:    subKeys,
		resultChan: make(chan WsSubscribeResult[WsSubscribeStatus], 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
	}

	ws.waitSubscribeResMap.Store(reqId, waitSubResult)

	log.Debugf("send msg: %s", string(data))
	ws.writeMu.Lock()
	err = ws.conn.WriteMessage(websocket.TextMessage, data)
	ws.writeMu.Unlock()
	if err != nil {
		return nil, err
	}

	//同步捕获订阅结果
	err = ws.catchSubscribeResult(waitSubResult)
	if err != nil {
		return nil, err
	}

	dataSubResult := &Subscription[T]{
		SubId:      reqId,
		Ws:         ws,
		Req:        &subscribeReq,
		Res:        waitSubResult.Res,
		Error:      waitSubResult.Error,
		SubKeys:    subKeys,
		resultChan: make(chan T, 50),
		errChan:    make(chan error),
		closeChan:  make(chan struct{}),
	}

	return dataSubResult, nil
}

func (ws *WsStreamClient) Close() error {
	ws.isClose = true

	err := ws.conn.Close()
	if err != nil {
		return err
	}
	//手动关闭成功，给所有订阅发送关闭信号
	ws.sendWsCloseToAllSub()

	//初始化连接状态
	ws.conn = nil
	ws.connId = ""
	close(ws.resultChan)
	close(ws.errChan)
	ws.resultChan = nil
	ws.errChan = nil
	ws.waitSubscribeResMap = NewMySyncMap[int64, *Subscription[WsSubscribeResult[WsSubscribeStatus]]]()

	ws.currentSubMap = NewMySyncMap[int64, *Subscription[WsSubscribeResult[WsSubscribeStatus]]]()
	ws.candleSubMap = NewMySyncMap[string, *Subscription[WsSubscribeResult[WsCandles]]]()

	return nil
}

func (ws *WsStreamClient) OpenConn() error {
	if ws.resultChan == nil {
		ws.resultChan = make(chan []byte)
	}
	if ws.errChan == nil {
		ws.errChan = make(chan error)
	}
	apiUrl := handlerWsStreamRequestApi(ws.apiType, ws.contractType)
	if ws.conn == nil {
		conn, err := wsStreamServe(apiUrl, ws.resultChan, ws.errChan)
		ws.conn = conn
		ws.connId = ""
		ws.isClose = false
		log.Info("OpenConn success to ", apiUrl)
		ws.handleResult(ws.resultChan, ws.errChan)
		return err
	} else {
		conn, err := wsStreamServe(apiUrl, ws.resultChan, ws.errChan)
		ws.conn = conn
		ws.connId = ""
		log.Info("Auto ReOpenConn success to ", apiUrl)
		return err
	}

}

type SpotWsStreamClient struct {
	WsStreamClient
}
type FuturesWsStreamClient struct {
	WsStreamClient
}

func NewSpotWsStreamClient(client *RestClient) *SpotWsStreamClient {
	return &SpotWsStreamClient{
		WsStreamClient{
			apiType:       WS_SPOT,
			client:        client,
			currentSubMap: NewMySyncMap[int64, *Subscription[WsSubscribeResult[WsSubscribeStatus]]](),
			candleSubMap:  NewMySyncMap[string, *Subscription[WsSubscribeResult[WsCandles]]](),
		},
	}
}
func NewFuturesWsStreamClient(client *RestClient, contractType ContractType) *FuturesWsStreamClient {
	return &FuturesWsStreamClient{
		WsStreamClient{
			apiType:       WS_FUTURES,
			contractType:  contractType,
			client:        client,
			currentSubMap: NewMySyncMap[int64, *Subscription[WsSubscribeResult[WsSubscribeStatus]]](),
			candleSubMap:  NewMySyncMap[string, *Subscription[WsSubscribeResult[WsCandles]]](),
		},
	}
}

func (ws *WsStreamClient) sendSubscribeResultToChan(result WsSubscribeResult[WsSubscribeStatus]) {
	if waitSub, ok := ws.waitSubscribeResMap.Load(result.Id); ok {
		if result.Error == nil && result.Result != nil && result.Result.Status == "success" {
			//订阅成功
			waitSub.resultChan <- result
		} else {
			//订阅失败
			waitSub.errChan <- fmt.Errorf("errHandler: %+v", result.Error)
		}
	}
}

func (ws *WsStreamClient) sendUnSubscribeSuccessToCloseChan(reqId int64, subKeys []string) {
	//解除订阅，删除不需要的订阅者
	//清除当前该订阅
	if _, ok := ws.currentSubMap.Load(reqId); ok {
		ws.currentSubMap.Delete(reqId)
	}

	for _, key := range subKeys {
		//删除K线订阅者
		if sub, ok := ws.candleSubMap.Load(key); ok {
			ws.candleSubMap.Delete(key)
			if sub.closeChan != nil {
				sub.closeChan <- struct{}{}
				sub.closeChan = nil
			}
		}
	}

}

// 解除所有订阅，清除所有订阅者
func (ws *WsStreamClient) sendWsCloseToAllSub() {
	ws.currentSubMap.Range(func(reqId int64, sub *Subscription[WsSubscribeResult[WsSubscribeStatus]]) bool {
		ws.sendUnSubscribeSuccessToCloseChan(reqId, sub.SubKeys)
		return true
	})
}

// 重订阅
func (ws *WsStreamClient) reSubscribeForReconnect() error {

	var errG errgroup.Group
	ws.currentSubMap.Range(func(reqId int64, sub *Subscription[WsSubscribeResult[WsSubscribeStatus]]) bool {
		errG.Go(func() error {
			newSub, err := subscribe[WsSubscribeResult[WsSubscribeStatus]](ws, sub.Req.Channel, SUBSCRIBE, sub.Req.Payload, sub.SubKeys, sub.Req.Auth == nil)
			if err != nil {
				log.Error(err)
				return err
			}
			log.Infof("reSubscribe Success: args:%v", sub.Req)
			for _, subKey := range sub.SubKeys {

				//重新订阅成功，更新订阅者
				dataSub, ok := ws.candleSubMap.Load(subKey)
				if ok {
					dataSub.SubId = newSub.SubId
				}
			}
			sub.SubId = newSub.SubId
			return nil
		})
		return true
	})
	if err := errG.Wait(); err != nil {
		return err
	}
	return nil
}

func (ws *WsStreamClient) handleResult(resultChan chan []byte, errChan chan error) {
	go func() {
		for {
			select {
			case err, ok := <-errChan:
				if !ok {
					log.Error("errChan is closed")
					return
				}
				log.Error(err)
				//错误处理 重连等
				//ws标记为非关闭 且返回错误包含EOF、close、reset时自动重连
				if !ws.isClose && (strings.Contains(err.Error(), "EOF") ||
					strings.Contains(err.Error(), "close") ||
					strings.Contains(err.Error(), "reset")) {
					//重连
					err := ws.OpenConn()
					for err != nil {
						time.Sleep(1500 * time.Millisecond)
						err = ws.OpenConn()
					}
					ws.AutoReConnectTimes += 1
					go func() {
						//重新订阅
						err = ws.reSubscribeForReconnect()
						if err != nil {
							log.Error(err)
						}
					}()
				} else {
					continue
				}
			case data, ok := <-resultChan:
				if !ok {
					log.Error("resultChan is closed")
					return
				}
				// log.Debug("receive result: ", string(data))
				//处理订阅或查询订阅列表请求返回结果
				if strings.Contains(string(data), "event\":\"subscribe") || strings.Contains(string(data), "event\":\"unsubscribe") {
					result := WsSubscribeResult[WsSubscribeStatus]{}
					err := json.Unmarshal(data, &result)
					if err != nil {
						log.Error(err)
						continue
					}
					ws.sendSubscribeResultToChan(result)
					continue
				}

				//处理正常数据的返回结果
				//K线处理
				if strings.Contains(string(data), "spot.candlesticks") && strings.Contains(string(data), "event\":\"update") {
					c, err := handleWsData[WsCandles](data)
					//使用交易所规则，interval_symbol作为唯一key
					key := c.Result.Interval
					if sub, ok := ws.candleSubMap.Load(key); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *c
					}
					continue
				}
				//K线处理
				if strings.Contains(string(data), "futures.candlesticks") && strings.Contains(string(data), "event\":\"update") {
					c, err := handleWsData[[]WsFutureCandles](data)
					if err != nil {
						log.Error(err)
						continue
					}
					//使用交易所规则，interval_symbol作为唯一key
					cs := splitSlice[WsFutureCandles, WsCandles](c, func(o WsFutureCandles) *WsCandles {
						return &WsCandles{
							Timestamp:   strconv.FormatInt(o.Timestamp, 10),
							Interval:    o.Interval,
							Open:        o.Open,
							High:        o.High,
							Low:         o.Low,
							Close:       o.Close,
							Volume:      strconv.FormatInt(o.Volume, 10),
							Amount:      o.Amount,
							WindowClose: o.WindowClose,
						}
					})
					for _, candle := range cs {
						key := candle.Result.Interval
						if sub, ok := ws.candleSubMap.Load(key); ok {
							if err != nil {
								sub.errChan <- err
								continue
							}
							sub.resultChan <- *candle
						}
					}
					continue
				}
			}
		}
	}()
}

// 捕获订阅结果
func (ws *WsStreamClient) catchSubscribeResult(sub *Subscription[WsSubscribeResult[WsSubscribeStatus]]) error {
	isBreak := false
	for {
		select {
		case err := <-sub.ErrChan():
			log.Error(err)
			return fmt.Errorf("SubAction Error: %v", err)
		case subResult := <-sub.ResultChan():
			if subResult.Error != nil || subResult.Result == nil || subResult.Result.Status != "success" {
				log.Errorf("%d:%v", subResult.Error.Code, subResult.Error.Msg)
				return fmt.Errorf("%d:%v", subResult.Error.Code, subResult.Error.Msg)
			}
			if ws.connId == "" {
				ws.connId = subResult.ConnId
			}
			sub.Res = &subResult
			ws.waitSubscribeResMap.Delete(sub.SubId)
			ws.currentSubMap.Store(sub.SubId, sub)

			isBreak = true
		case <-sub.CloseChan():
			return errors.New("subActionError: use of subscription closeChan")
		}
		if isBreak {
			break
		}
	}
	d, _ := json.Marshal(sub.Res)
	log.Debugf("catchResults: %s", string(d))

	//log.Debugf("ws: %+v", ws)
	return nil
}

// 标准订阅方法
func wsStreamServe(api string, resultChan chan []byte, errChan chan error) (*websocket.Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(api, nil)
	if err != nil {
		return nil, err
	}
	c.SetReadLimit(655350)
	go func() {
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout, WebsocketPingTicker)
		}
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- message
		}
	}()
	return c, err
}

// 获取数据流请求URL
func handlerWsStreamRequestApi(apiType APIType, contractType ContractType) string {
	u := url.URL{
		Scheme:   "wss",
		Host:     getWsApiHost(apiType),
		Path:     getWsApiPath(apiType, contractType),
		RawQuery: "",
	}
	return u.String()
}

// 获取数据流请求Host
func getWsApiHost(apiType APIType) string {
	switch apiType {
	case WS_SPOT:
		return GATE_API_WS_SPOT
	case WS_FUTURES:
		return GATE_API_WS_FUTURES
	default:
		return ""
	}
}

// 获取数据流请求Path
func getWsApiPath(apiType APIType, contractType ContractType) string {
	switch apiType {
	case WS_SPOT:
		return GATE_API_WS_SPOT_PATH
	case WS_FUTURES:
		switch contractType {
		case BTC_CONTRACT:
			return GATE_API_WS_FUTURES_PATH_BTC
		case USDT_CONTRACT:
			return GATE_API_WS_FUTURES_PATH_USDT
		default:
			return ""
		}
	default:
		return ""
	}
}

// 发送ping/pong消息以检查连接稳定性
func keepAlive(c *websocket.Conn, timeout, tickerTimeout time.Duration) {
	ticker := time.NewTicker(tickerTimeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		log.Warn(msg)
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				log.Error(err)
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				err := c.Close()
				if err != nil {
					log.Error(err)
					return
				}
				return
			}
		}
	}()
}
