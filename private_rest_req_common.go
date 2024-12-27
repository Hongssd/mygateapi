package mygateapi

type GateFuturesOrderReqCommon struct {
	Settle     *string `json:"settle,omitempty"`      //结算货币
	Contract   *string `json:"contract,omitempty"`    //合约标识
	Size       *int64  `json:"size,omitempty"`        //交易数量，正数为买入，负数为卖出。平仓委托则设置为0。
	Iceberg    *int64  `json:"iceberg,omitempty"`     //冰山委托显示数量。0为完全不隐藏。注意，隐藏部分成交按照taker收取手
	Price      *string `json:"price,omitempty"`       //委托价。价格为0并且tif为ioc，代表市价委托。
	Close      *bool   `json:"close,omitempty"`       //设置为 true 的时候执行平仓操作，并且size应设置为0
	ReduceOnly *bool   `json:"reduce_only,omitempty"` //设置为 true 的时候，为只减仓委托
	Tif        *string `json:"tif,omitempty"`         //Time in force 策略，市价单当前只支持 ioc 模式
	Text       *string `json:"text,omitempty"`        //订单自定义信息，用户可以用该字段设置自定义 ID，用户自定义字段必须满足以下条件：
	AutoSize   *string `json:"auto_size,omitempty"`   //双仓模式下用于设置平仓的方向，close_long 平多头， close_short 平空头，需要同时设置 size 为 0
	StpAct     *string `json:"stp_act,omitempty"`     //Self-Trading Prevention Action,用户可以用该字段设置自定义限制自成交策略。
}

func removeSettleFromReqBody(req interface{}) ([]byte, error) {
	// remove settle from body
	bodyMap := map[string]interface{}{}
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(reqBody, &bodyMap)
	if err != nil {
		return nil, err
	}
	delete(bodyMap, "settle")
	reqBody, err = json.Marshal(bodyMap)
	if err != nil {
		return nil, err
	}
	return reqBody, nil
}
