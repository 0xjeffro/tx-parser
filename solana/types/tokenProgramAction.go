package types

type TokenProgramTransferAction struct {
	BaseAction
	From   string `json:"from"`
	To     string `json:"to"`
	Amount uint32 `json:"amount"`
}
