package svrdefine

type NetworkInfo struct {
	Network                 string `json:"network"`
	Coin                    string `json:"coin"`
	EntityTag               string `json:"entityTag"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
	IsDefault               bool   `json:"isDefault"`
	DepositEnable           bool   `json:"depositEnable"`
	WithdrawEnable          bool   `json:"withdrawEnable"`
	DepositDesc             string `json:"depositDesc"`
	WithdrawDesc            string `json:"withdrawDesc"`
	SpecialTips             string `json:"specialTips"`
	SpecialWithdrawTips     string `json:"specialWithdrawTips"`
	Name                    string `json:"name"`
	ResetAddressStatus      bool   `json:"resetAddressStatus"`
	AddressRegex            string `json:"addressRegex"`
	AddressRule             string `json:"addressRule"`
	MemoRegex               string `json:"memoRegex"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawMin             string `json:"withdrawMin"`
	WithdrawMax             string `json:"withdrawMax"`
	MinConfirm              int    `json:"minConfirm"`
	UnLockConfirm           int    `json:"unLockConfirm"`
	SameAddress             bool   `json:"sameAddress"`
	EstimatedArrivalTime    int    `json:"estimatedArrivalTime"`
	Busy                    bool   `json:"busy"`
	Country                 string `json:"country"`
	ContractAddressUrl      string `json:"contractAddressUrl"`
	ContractAddress         string `json:"contractAddress"`
}

type CoinInfo struct {
	Coin              string        `json:"coin"`
	DepositAllEnable  bool          `json:"depositAllEnable"`
	WithdrawAllEnable bool          `json:"withdrawAllEnable"`
	Name              string        `json:"name"`
	Free              string        `json:"free"`
	Locked            string        `json:"locked"`
	Freeze            string        `json:"freeze"`
	Withdrawing       string        `json:"withdrawing"`
	Ipoing            string        `json:"ipoing"`
	Ipoable           string        `json:"ipoable"`
	Storage           string        `json:"storage"`
	IsLegalMoney      bool          `json:"isLegalMoney"`
	Trading           bool          `json:"trading"`
	NetworkList       []NetworkInfo `json:"networkList"`
}
