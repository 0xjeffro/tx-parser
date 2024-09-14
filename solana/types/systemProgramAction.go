package types

type SystemProgramTransferAction struct {
	BaseAction
	From     string `json:"from"`
	To       string `json:"to"`
	Lamports uint64 `json:"lamports"`
}
