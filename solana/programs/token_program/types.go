package token_program

import "github.com/thetafunction/tx-parser/solana/types"

type TransferAction struct {
	types.BaseAction
	From   string `json:"from"`
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
}

type TransferCheckedAction struct {
	types.BaseAction
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   uint64 `json:"amount"`
	Mint     string `json:"mint"`
	Decimals uint64 `json:"decimals"`
}

type InitializeAccountAction struct {
	types.BaseAction
	Account    string `json:"account"`
	Mint       string `json:"mint"`
	Owner      string `json:"owner"`
	RentSysvar string `json:"rentSysvar"`
}
