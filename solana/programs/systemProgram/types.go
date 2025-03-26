package systemProgram

import "github.com/0xjeffro/tx-parser/solana/types"

type TransferAction struct {
	types.BaseAction
	From     string `json:"from"`
	To       string `json:"to"`
	Lamports uint64 `json:"lamports"`
}

type CreateAccountWithSeedAction struct {
	types.BaseAction
	Who        string `json:"who"`
	NewAccount string `json:"newAccount"`
	Base       string `json:"base"`
	Seed       string `json:"seed"`
	Lamports   uint64 `json:"lamports"`
	Space      uint64 `json:"space"`
	Owner      string `json:"owner"`
}
