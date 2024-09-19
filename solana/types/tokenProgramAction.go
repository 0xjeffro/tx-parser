package types

type TokenProgramTransferAction struct {
	BaseAction
	From   string `json:"from"`
	To     string `json:"to"`
	Amount uint32 `json:"amount"`
}

type TokenProgramTransferCheckedAction struct {
	BaseAction
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   uint64 `json:"amount"`
	Mint     string `json:"mint"`
	Decimals uint64 `json:"decimals"`
}
