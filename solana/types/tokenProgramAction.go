package types

type TokenProgramTransferAction struct {
	BaseAction
	From   string `json:"from"`
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
}

type TokenProgramTransferCheckedAction struct {
	BaseAction
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   uint64 `json:"amount"`
	Mint     string `json:"mint"`
	Decimals uint64 `json:"decimals"`
}

type TokenProgramInitializeAccountAction struct {
	BaseAction
	Account    string `json:"account"`
	Mint       string `json:"mint"`
	Owner      string `json:"owner"`
	RentSysvar string `json:"rentSysvar"`
}
