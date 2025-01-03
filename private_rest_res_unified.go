package mygateapi

type PrivateRestUnifiedAccountsRes struct {
	UserId      int64 `json:"user_id"`      // 用户 ID
	RefreshTime int64 `json:"refresh_time"` // 最近一次刷新时间
	Locked      bool  `json:"locked"`       // 账户是否被锁定
	Balances    struct {
		UnifiedBalance struct {
			Available    string `json:"available"`     // 可用额度
			Freeze       string `json:"freeze"`        // 被锁定的额度
			Borrowed     string `json:"borrowed"`      // 借入额度
			NegativeLiab string `json:"negative_liab"` // 负余额借贷
			//FuturesPosLiab          string `json:"futures_pos_liab"`          // 合约开仓借币(已废弃,待下线字段)
			Equity string `json:"equity"` // 权益
			//TotalFreeze             string `json:"total_freeze"`              // 总占用(已废弃,待下线字段)
			TotalLiab      string `json:"total_liab"`      // 总借款
			SpotInUse      string `json:"spot_in_use"`     // 现货对冲占用数量
			Funding        string `json:"funding"`         // 余币宝理财数量
			FundingVersion string `json:"funding_version"` // 余币宝理财版本号
		} `json:"unified_balance"`
		//Total                   string `json:"total"`                     // 折算成 USD 的账户总资产，即所有币种 (available + freeze) * price 之和,(已废弃，待下线字段，用unified_account_total代替)
		Borrowed                   string `json:"borrowed"`                      // 折算成 USD 的账户总借入数量，即所有币种(不包括点卡)的 borrowed * price 之和
		TotalInitialMargin         string `json:"total_initial_margin"`          // 总初始保证金
		TotalMarginBalance         string `json:"total_margin_balance"`          // 总保证金余额
		TotalMaintenanceMargin     string `json:"total_maintenance_margin"`      // 总维持保证金
		TotalInitialMarginRate     string `json:"total_initial_margin_rate"`     // 总初始保证金率
		TotalMaintenanceMarginRate string `json:"total_maintenance_margin_rate"` // 总维持保证金率
		TotalAvailableMargin       string `json:"total_available_margin"`        // 可用的保证金额度
		UnifiedAccountTotal        string `json:"unified_account_total"`         // 统一账户总资产
		UnifiedAccountTotalLiab    string `json:"unified_account_total_liab"`    // 统一账户总借贷
		UnifiedAccountTotalEquity  string `json:"unified_account_total_equity"`  // 统一账户总权益
		Leverage                   string `json:"leverage"`                      // 实际杠杆
		SpotOrderLoss              string `json:"spot_order_loss"`               // 总挂单损失,单位USDT
		SpotHedge                  bool   `json:"spot_hedge"`                    // 现货对冲状态, true - 启用，false - 未启用
		UseFunding                 bool   `json:"use_funding"`                   // 是否将余币宝理财资金作为保证金
	} `json:"balances"` // 账户余额
}

type PrivateRestUnifiedTransferableRes struct {
	Currency string `json:"currency"` // 币种信息
	Amount   string `json:"amount"`   // 最多可转出的额度
}

type PrivateRestUnifiedRiskUnitsRes struct {
	UserId    int64 `json:"user_id"`    // 用户 ID
	SpotHedge bool  `json:"spot_hedge"` // 现货对冲状态, true - 启用，false - 未启用
	RiskUnits []struct {
		Symbol         string `json:"symbol"`          // 风险单元标志
		SpotInUse      string `json:"spot_in_use"`     // 现货对冲占用数量
		MaintainMargin string `json:"maintain_margin"` // 风险单元的维持保证金
		InitialMargin  string `json:"initial_margin"`  // 风险单元的起始保证金
		Delta          string `json:"delta"`           // 风险单元的 总 delta
		Gamma          string `json:"gamma"`           // 风险单元的 总 gamma
		Theta          string `json:"theta"`           // 风险单元的 总 theta
		Vega           string `json:"vega"`            // 风险单元的 总 vega
	} `json:"risk_units"` // 风险单元
}

type PrivateRestUnifiedUnifiedModePutRes struct{}

type PrivateRestUnifiedUnifiedModeGetRes struct {
	Mode     string `json:"mode"` // 统一账户模式
	Settings struct {
		UsdtFutures bool `json:"usdt_futures"` // USDT合约开关
		SpotHedge   bool `json:"spot_hedge"`   // 现货对冲开关
		UseFunding  bool `json:"use_funding"`  // 是否将余币宝理财资金作为保证金
		Options     bool `json:"options"`      // 期权开关
	} `json:"settings"` // 设置
}

type PrivateRestUnifiedLeverageUserCurrencyConfigRes struct {
	CurrentLeverage          string `json:"current_leverage"`           // 当前杠杆倍数
	MinLeverage              string `json:"min_leverage"`               // 最小可调杠杆倍数
	MaxLeverage              string `json:"max_leverage"`               // 最大可调杠杆倍数
	Debit                    string `json:"debit"`                      // 当前负债
	AvailableMargin          string `json:"available_margin"`           // 可用保证金
	Borrowable               string `json:"borrowable"`                 // 当前选择杠杆可借
	ExceptLeverageBorrowable string `json:"except_leverage_borrowable"` // 保证金最大可借、余币宝最大可借，两者取较小的值
}

type PrivateRestUnifiedLeverageUserCurrencySettingGetRes struct {
	Currency string `json:"currency"` // 币种名称
	Leverage string `json:"leverage"` // 倍数
}

type PrivateRestUnifiedLeverageUserCurrencySettingPostRes struct{}
