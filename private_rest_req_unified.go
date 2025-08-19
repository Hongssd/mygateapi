package mygateapi

type PrivateRestUnifiedAccountsReq struct {
	Currency *string `json:"currency"` //String	否	指定币种名称查询
}

// currency	String	否	指定币种名称查询
func (api *PrivateRestUnifiedAccountsAPI) Currency(currency string) *PrivateRestUnifiedAccountsAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

type PrivateRestUnifiedAccountsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUnifiedAccountsReq
}

type PrivateRestUnifiedTransferableReq struct {
	Currency *string `json:"currency"` //String	否	指定币种名称查询
}

// currency	String	否	指定币种名称查询
func (api *PrivateRestUnifiedTransferableAPI) Currency(currency string) *PrivateRestUnifiedTransferableAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

type PrivateRestUnifiedTransferableAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUnifiedTransferableReq
}

type PrivateRestUnifiedRiskUnitsReq struct{}

type PrivateRestUnifiedRiskUnitsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUnifiedRiskUnitsReq
}

type PrivateRestUnifiedUnifiedModePutReq struct {
	Mode     *string                                `json:"mode"`     //String	是	统一账户模式
	Settings *PrivateRestUnifiedUnifiedModeSettings `json:"settings"` //Object	否
}
type PrivateRestUnifiedUnifiedModeSettings struct {
	UsdtFutures *bool `json:"usdt_futures"` //Boolean	否	USDT合约开关。不传时,取当前开关值,首次开通不传时默认为关
	SpotHedge   *bool `json:"spot_hedge"`   //Boolean	否	现货对冲开关。不传时,取当前开关值,首次开通不传时默认为关
	UseFunding  *bool `json:"use_funding"`  //Boolean	否	当mode为组合保证金模式时,是否将余币宝理财资金作为保证金
	Options     *bool `json:"options"`      //Boolean	否	期权开关。不传时,取当前开关值,首次开通不传时默认为关
}

// mode	String	是	统一账户模式
func (api *PrivateRestUnifiedUnifiedModePutAPI) Mode(mode string) *PrivateRestUnifiedUnifiedModePutAPI {
	api.req.Mode = GetPointer(mode)
	return api
}

// usdtFutures	Boolean	否	USDT合约开关。不传时,取当前开关值,首次开通不传时默认为关
func (api *PrivateRestUnifiedUnifiedModePutAPI) UsdtFutures(usdtFutures bool) *PrivateRestUnifiedUnifiedModePutAPI {
	if api.req.Settings == nil {
		api.req.Settings = &PrivateRestUnifiedUnifiedModeSettings{}
	}
	api.req.Settings.UsdtFutures = GetPointer(usdtFutures)
	return api
}

// spotHedge	Boolean	否	现货对冲开关。不传时,取当前开关值,首次开通不传时默认为关
func (api *PrivateRestUnifiedUnifiedModePutAPI) SpotHedge(spotHedge bool) *PrivateRestUnifiedUnifiedModePutAPI {
	if api.req.Settings == nil {
		api.req.Settings = &PrivateRestUnifiedUnifiedModeSettings{}
	}
	api.req.Settings.SpotHedge = GetPointer(spotHedge)
	return api
}

// useFunding	Boolean	否	当mode为组合保证金模式时,是否将余币宝理财资金作为保证金
func (api *PrivateRestUnifiedUnifiedModePutAPI) UseFunding(useFunding bool) *PrivateRestUnifiedUnifiedModePutAPI {
	if api.req.Settings == nil {
		api.req.Settings = &PrivateRestUnifiedUnifiedModeSettings{}
	}
	api.req.Settings.UseFunding = GetPointer(useFunding)
	return api
}

// options	Boolean	否	期权开关。不传时,取当前开关值,首次开通不传时默认为关
func (api *PrivateRestUnifiedUnifiedModePutAPI) Options(options bool) *PrivateRestUnifiedUnifiedModePutAPI {
	if api.req.Settings == nil {
		api.req.Settings = &PrivateRestUnifiedUnifiedModeSettings{}
	}
	api.req.Settings.Options = GetPointer(options)
	return api
}

type PrivateRestUnifiedUnifiedModePutAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUnifiedUnifiedModePutReq
}

type PrivateRestUnifiedUnifiedModeGetReq struct{}

type PrivateRestUnifiedUnifiedModeGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUnifiedUnifiedModeGetReq
}

type PrivateRestUnifiedLeverageUserCurrencyConfigReq struct {
	Currency *string `json:"currency"` //String	是	币种
}

// currency	String	是	币种
func (api *PrivateRestUnifiedLeverageUserCurrencyConfigAPI) Currency(currency string) *PrivateRestUnifiedLeverageUserCurrencyConfigAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

type PrivateRestUnifiedLeverageUserCurrencyConfigAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUnifiedLeverageUserCurrencyConfigReq
}

type PrivateRestUnifiedLeverageUserCurrencySettingGetReq struct {
	Currency *string `json:"currency"` //String	否	币种
}

// currency	String	否	币种
func (api *PrivateRestUnifiedLeverageUserCurrencySettingGetAPI) Currency(currency string) *PrivateRestUnifiedLeverageUserCurrencySettingGetAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

type PrivateRestUnifiedLeverageUserCurrencySettingGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUnifiedLeverageUserCurrencySettingGetReq
}

type PrivateRestUnifiedLeverageUserCurrencySettingPostReq struct {
	Currency *string `json:"currency"` //String	否	币种
	Leverage *string `json:"leverage"` //String	否	杠杆倍数
}

// currency	String	否	币种
func (api *PrivateRestUnifiedLeverageUserCurrencySettingPostAPI) Currency(currency string) *PrivateRestUnifiedLeverageUserCurrencySettingPostAPI {
	api.req.Currency = GetPointer(currency)
	return api
}

// leverage	String	否	杠杆倍数
func (api *PrivateRestUnifiedLeverageUserCurrencySettingPostAPI) Leverage(leverage string) *PrivateRestUnifiedLeverageUserCurrencySettingPostAPI {
	api.req.Leverage = GetPointer(leverage)
	return api
}

type PrivateRestUnifiedLeverageUserCurrencySettingPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestUnifiedLeverageUserCurrencySettingPostReq
}
