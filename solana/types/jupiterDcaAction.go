package types

type JupiterDcaOpenV2Action struct {
	BaseAction
	InAmount         uint64  `json:"inAmount"`
	InAmountPerCycle uint64  `json:"inAmountPerCycle"`
	CycleFrequency   int64   `json:"cycleFrequency"`
	MinOutAmount     *uint64 `json:"minOutAmount"`
	MaxOutAmount     *uint64 `json:"maxOutAmount"`
	StartAt          *int64  `json:"startAt"`

	Dca        string `json:"dca"`
	User       string `json:"user"`
	Payer      string `json:"payer"`
	InputMint  string `json:"inputMint"`
	OutputMint string `json:"outputMint"`
	UserAta    string `json:"userAta"`
	InAta      string `json:"inAta"`
	OutAta     string `json:"payerAta"`
}

type JupiterDcaCloseAction struct {
	BaseAction
	User       string `json:"user"`
	Dca        string `json:"dca"`
	InputMint  string `json:"inputMint"`
	OutputMint string `json:"outputMint"`
	InAta      string `json:"inAta"`
	OutAta     string `json:"outAta"`
	UserInAta  string `json:"userInAta"`
	UserOutAta string `json:"userOutAta"`
}

type JupiterDcaEndAndCloseAction struct {
	BaseAction
	Keeper     string `json:"keeper"`
	Dca        string `json:"dca"`
	InputMint  string `json:"inputMint"`
	OutputMint string `json:"outputMint"`
	InAta      string `json:"inAta"`
	OutAta     string `json:"outAta"`
	User       string `json:"user"`
	UserOutAta string `json:"userOutAta"`
}

type JupiterWithDrawAction struct {
	BaseAction
	// WithDrawParams
	// ...
	User       string `json:"user"`
	Dca        string `json:"dca"`
	InputMint  string `json:"inputMint"`
	OutputMint string `json:"outputMint"`
	DcaAta     string `json:"dcaAta"`
	UserInAta  string `json:"userInAta"`
	UserOutAta string `json:"userOutAta"`
}

type JupiterDcaDepositAction struct {
	BaseAction
	DepositIn uint64 `json:"depositIn"`
	User      string `json:"user"`
	Dca       string `json:"dca"`
	InAta     string `json:"inAta"`
	UserInAta string `json:"userInAta"`
}
