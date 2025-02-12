package mygateapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"regexp"
	"sync"
	"time"

	"net/url"
	"reflect"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

const (
	BIT_BASE_10 = 10
	BIT_SIZE_64 = 64
	//BIT_SIZE_32 = 32
)

type RequestType string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
	PATCH  = "PATCH"
)

var NIL_REQBODY = []byte{}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var log = logrus.New()

func SetLogger(logger *logrus.Logger) {
	log = logger
}

var httpTimeout = 100 * time.Second

func SetHttpTimeout(timeout time.Duration) {
	httpTimeout = timeout
}

func GetPointer[T any](v T) *T {
	return &v
}

func HmacSha512(secret, data string) string {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

type MyGate struct {
}

const (
	GATE_API_HTTP      = "api.gateio.ws"
	TEST_GATE_API_HTTP = "fx-api-testnet.gateio.ws"
	IS_GZIP            = false
)

type NetType int

const (
	MAIN_NET NetType = iota
	TEST_NET
)

var NowNetType = MAIN_NET

func SetNetType(netType NetType) {
	NowNetType = netType
}

type APIType int

const (
	REST APIType = iota
	WS_SPOT
	WS_FUTURES
	WS_DELIVERY
)

type Client struct {
	ApiKey    string
	ApiSecret string
}
type RestClient struct {
	c *Client
}
type PublicRestClient RestClient

type PrivateRestClient RestClient

func NewRestClient(apiKey string, apiSecret string) *RestClient {
	client := &RestClient{
		&Client{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		},
	}
	return client
}

func (c *RestClient) PublicRestClient() *PublicRestClient {
	return &PublicRestClient{
		c: c.c,
	}
}

func (c *RestClient) PrivateRestClient() *PrivateRestClient {
	return &PrivateRestClient{
		c: c.c,
	}
}

var serverTimeDelta int64 = 0

func SetServerTimeDelta(delta int64) {
	serverTimeDelta = delta
}

// 通用接口调用
func gateCallAPI[T any](client *Client, url url.URL, reqBody []byte, method string) (*GateRestRes[T], error) {
	body, respHeaderMap, statusCode, err := Request(url.String(), reqBody, method, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body, respHeaderMap, statusCode)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

// 通用鉴权接口调用
func gateCallApiWithSecret[T any](client *Client, url url.URL, reqBody []byte, method string) (*GateRestRes[T], error) {
	timestamp := time.Now().Unix() //秒级时间戳
	requestPath := url.Path
	query := url.RawQuery
	reqBodySum512 := sha512.Sum512(reqBody)
	reqBodySum512String := hex.EncodeToString(reqBodySum512[:])

	//格式 Request Method + "\n" + Request URL + "\n" + Query String + "\n" + HexEncode(SHA512(Request Payload)) + "\n" + Timestamp
	hmacSha512Data := method + "\n" + requestPath + "\n" + query + "\n" + reqBodySum512String + "\n" + strconv.FormatInt(timestamp, BIT_BASE_10)
	sign := HmacSha512(client.ApiSecret, hmacSha512Data)

	// log.Warn(hmacSha512Data)
	// log.Warn("timestamp: ", timestamp)
	// log.Warn("method: ", method)
	// log.Warn("requestPath: ", requestPath)
	// log.Warn("query: ", query)
	// log.Warn("reqBody: ", string(reqBody))
	// log.Warn("hmacSha512Data: ", hmacSha512Data)
	// log.Warn("sign: ", sign)

	body, respHeaderMap, statusCode, err := RequestWithHeader(url.String(), reqBody, method,
		map[string]string{
			"KEY":       client.ApiKey,
			"SIGN":      sign,
			"Timestamp": strconv.FormatInt(timestamp, BIT_BASE_10),
		}, IS_GZIP)
	if err != nil {
		return nil, err
	}
	res, err := handlerCommonRest[T](body, respHeaderMap, statusCode)
	if err != nil {
		return nil, err
	}
	return res, res.handlerError()
}

// URL标准封装 带路径参数
func gateHandlerRequestAPIWithPathQueryParam[T any](apiType APIType, request *T, name string) url.URL {
	// 正则匹配name字符串中包含{xxx}的参数，其中xxx为路径参数名
	// 定义正则表达式，匹配 `{xxx}` 的部分
	re := regexp.MustCompile(`\{([a-zA-Z0-9_]+)}`)
	// 查找所有匹配的子字符串
	matches := re.FindAllStringSubmatch(name, -1)
	// 提取路径参数名
	var params []string
	for _, match := range matches {
		if len(match) > 1 {
			params = append(params, match[1])
		}
	}
	//fmt.Println("路径参数:", params)

	//将指定参数回填
	// 参数值映射
	query, paramsMap := gateHandlerReq(request, true, params...)

	result := name
	if len(paramsMap) != 0 {
		// 替换占位符
		result = re.ReplaceAllStringFunc(name, func(m string) string {
			// 提取参数名
			key := re.FindStringSubmatch(m)[1]
			// 替换为参数值
			if value, exists := paramsMap[key]; exists {
				return value
			}
			// 如果参数未提供，保留原占位符
			return m
		})
	}

	//fmt.Println("填充后的路径:", result)

	u := url.URL{
		Scheme:   "https",
		Host:     GateGetRestHostByAPIType(apiType),
		Path:     result,
		RawQuery: query,
	}
	return u
}

// URL标准封装 不带路径参数
func gateHandlerRequestAPIWithoutPathQueryParam[T any](apiType APIType, request *T, name string) url.URL {
	// 正则匹配name字符串中包含{xxx}的参数，其中xxx为路径参数名
	// 定义正则表达式，匹配 `{xxx}` 的部分
	re := regexp.MustCompile(`\{([a-zA-Z0-9_]+)}`)
	// 查找所有匹配的子字符串
	matches := re.FindAllStringSubmatch(name, -1)
	// 提取路径参数名
	var params []string
	for _, match := range matches {
		if len(match) > 1 {
			params = append(params, match[1])
		}
	}
	//fmt.Println("路径参数:", params)

	//将指定参数回填
	// 参数值映射
	query, paramsMap := gateHandlerReq(request, false, params...)

	result := name
	log.Debug("paramsMap:", paramsMap)
	if len(paramsMap) != 0 {
		// 替换占位符
		result = re.ReplaceAllStringFunc(name, func(m string) string {
			// 提取参数名
			key := re.FindStringSubmatch(m)[1]
			log.Debug("key:", key)
			// 替换为参数值
			if value, exists := paramsMap[key]; exists {
				return value
			}
			// 如果参数未提供，保留原占位符
			return m
		})
	}

	//fmt.Println("填充后的路径:", result)

	// query := gateHandlerReq(request)
	u := url.URL{
		Scheme:   "https",
		Host:     GateGetRestHostByAPIType(apiType),
		Path:     result,
		RawQuery: query,
	}

	log.Debug(u.String())

	return u
}

func gateHandlerReq[T any](req *T, isPathQuery bool, pathParams ...string) (string, map[string]string) {
	var argBuffer bytes.Buffer
	pathParamsMap := map[string]string{}

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)
	if v.IsNil() {
		return "", pathParamsMap
	}
	t = t.Elem()
	v = v.Elem()
	count := v.NumField()
	for i := 0; i < count; i++ {
		argName := t.Field(i).Tag.Get("json")
		argName = strings.ReplaceAll(argName, ",omitempty", "")
		isPathParam := false
		for _, pathParam := range pathParams {
			if argName == pathParam {
				isPathParam = true
				break
			}
		}
		if isPathParam {
			switch v.Field(i).Elem().Kind() {
			case reflect.String:
				pathParamsMap[argName] = v.Field(i).Elem().String()
			case reflect.Int, reflect.Int64:
				pathParamsMap[argName] = strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10)
			case reflect.Float32, reflect.Float64:
				pathParamsMap[argName] = decimal.NewFromFloat(v.Field(i).Elem().Float()).String()
			case reflect.Bool:
				pathParamsMap[argName] = strconv.FormatBool(v.Field(i).Elem().Bool())
			case reflect.Struct:
				sv := reflect.ValueOf(v.Field(i).Interface())
				ToStringMethod := sv.MethodByName("String")
				args := make([]reflect.Value, 0)
				result := ToStringMethod.Call(args)
				pathParamsMap[argName] = result[0].String()
			case reflect.Slice:
				s := v.Field(i).Interface()
				d, _ := json.Marshal(s)
				pathParamsMap[argName] = url.QueryEscape(string(d))
			case reflect.Invalid:
			default:
				log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
			}
		} else {
			switch v.Field(i).Elem().Kind() {
			case reflect.String:
				argBuffer.WriteString(argName + "=" + v.Field(i).Elem().String() + "&")
			case reflect.Int, reflect.Int64:
				argBuffer.WriteString(argName + "=" + strconv.FormatInt(v.Field(i).Elem().Int(), BIT_BASE_10) + "&")
			case reflect.Float32, reflect.Float64:
				argBuffer.WriteString(argName + "=" + decimal.NewFromFloat(v.Field(i).Elem().Float()).String() + "&")
			case reflect.Bool:
				argBuffer.WriteString(argName + "=" + strconv.FormatBool(v.Field(i).Elem().Bool()) + "&")
			case reflect.Struct:
				sv := reflect.ValueOf(v.Field(i).Interface())
				ToStringMethod := sv.MethodByName("String")
				args := make([]reflect.Value, 0)
				result := ToStringMethod.Call(args)
				argBuffer.WriteString(argName + "=" + result[0].String() + "&")
			case reflect.Slice:
				s := v.Field(i).Interface()
				d, _ := json.Marshal(s)
				argBuffer.WriteString(argName + "=" + url.QueryEscape(string(d)) + "&")
			case reflect.Invalid:
			default:
				log.Errorf("req type error %s:%s", argName, v.Field(i).Elem().Kind())
			}
		}

	}

	if !isPathQuery {
		return "", pathParamsMap
	}

	return strings.TrimRight(argBuffer.String(), "&"), pathParamsMap
}

func GateGetRestHostByAPIType(apiType APIType) string {
	switch apiType {
	case REST:
		if NowNetType == MAIN_NET {
			return GATE_API_HTTP
		} else {
			return TEST_GATE_API_HTTP
		}
	}

	return ""

}

func interfaceStringToFloat64(inter interface{}) float64 {
	return stringToFloat64(inter.(string))
}

func interfaceStringToInt64(inter interface{}) int64 {
	return int64(inter.(float64))
}

func stringToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, BIT_SIZE_64)
	return f
}

type MySyncMap[K any, V any] struct {
	smap sync.Map
}

func NewMySyncMap[K any, V any]() MySyncMap[K, V] {
	return MySyncMap[K, V]{
		smap: sync.Map{},
	}
}
func (m *MySyncMap[K, V]) Load(k K) (V, bool) {
	v, ok := m.smap.Load(k)

	if ok {
		return v.(V), true
	}
	var resv V
	return resv, false
}
func (m *MySyncMap[K, V]) Store(k K, v V) {
	m.smap.Store(k, v)
}

func (m *MySyncMap[K, V]) Delete(k K) {
	m.smap.Delete(k)
}
func (m *MySyncMap[K, V]) Range(f func(k K, v V) bool) {
	m.smap.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}

func (m *MySyncMap[K, V]) Length() int {
	length := 0
	m.Range(func(k K, v V) bool {
		length += 1
		return true
	})
	return length
}

func (m *MySyncMap[K, V]) MapValues(f func(k K, v V) V) *MySyncMap[K, V] {
	var res = NewMySyncMap[K, V]()
	m.Range(func(k K, v V) bool {
		res.Store(k, f(k, v))
		return true
	})
	return &res
}
