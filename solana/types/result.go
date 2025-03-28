package types

type ParsedResult struct {
	RawTx       RawTx    `json:"rawTx"`
	AccountList []string `json:"accountList"`
	Actions     []Action `json:"actions"`
}
