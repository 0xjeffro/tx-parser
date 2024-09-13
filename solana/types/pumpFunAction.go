package types

type PumpFunBuyAction struct {
	BaseAction
	Who             string `json:"who"`
	FromToken       string `json:"fromToken"`
	FromTokenAmount uint64 `json:"fromTokenAmount"`
	ToToken         string `json:"toToken"`
	ToTokenAmount   uint64 `json:"toTokenAmount"`
}

type PumpFunSellAction struct {
	BaseAction
	Who             string `json:"who"`
	FromToken       string `json:"fromToken"`
	FromTokenAmount uint64 `json:"fromTokenAmount"`
	ToToken         string `json:"toToken"`
	ToTokenAmount   uint64 `json:"toTokenAmount"`
}
