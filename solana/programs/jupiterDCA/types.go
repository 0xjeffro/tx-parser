package jupiterDCA

import "github.com/0xjeffro/tx-parser/solana/types"

type OpenV2Action struct {
	types.BaseAction
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

type CloseDcaAction struct {
	types.BaseAction
	User       string `json:"user"`
	Dca        string `json:"dca"`
	InputMint  string `json:"inputMint"`
	OutputMint string `json:"outputMint"`
	InAta      string `json:"inAta"`
	OutAta     string `json:"outAta"`
	UserInAta  string `json:"userInAta"`
	UserOutAta string `json:"userOutAta"`
}

type EndAndCloseAction struct {
	types.BaseAction
	Keeper     string `json:"keeper"`
	Dca        string `json:"dca"`
	InputMint  string `json:"inputMint"`
	OutputMint string `json:"outputMint"`
	InAta      string `json:"inAta"`
	OutAta     string `json:"outAta"`
	User       string `json:"user"`
	UserOutAta string `json:"userOutAta"`
}

type WithdrawAction struct {
	types.BaseAction
	// WithDrawParams
	// ...
	// 58eYP8XhJKfVNm189hQinBcLacDqiYXnTadwzxXbdh9ssDNnmBi9FqaMKMkzgeiykfCDGowfREfh8p7Gziggikm9
	User       string `json:"user"`
	Dca        string `json:"dca"`
	InputMint  string `json:"inputMint"`
	OutputMint string `json:"outputMint"`
	DcaAta     string `json:"dcaAta"`
	UserInAta  string `json:"userInAta"`
	UserOutAta string `json:"userOutAta"`
}

type DepositAction struct {
	types.BaseAction
	DepositIn uint64 `json:"depositIn"`
	User      string `json:"user"`
	Dca       string `json:"dca"`
	InAta     string `json:"inAta"`
	UserInAta string `json:"userInAta"`
}
