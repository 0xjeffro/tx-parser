package types

type U6m2CDdhRgSwapAction struct {
	BaseAction
	Who               string `json:"who"`
	FromToken         string `json:"fromToken"`
	FromTokenAmount   uint64 `json:"fromTokenAmount"`
	FromTokenDecimals uint64 `json:"fromTokenDecimals"`
	ToToken           string `json:"toToken"`
	ToTokenAmount     uint64 `json:"toTokenAmount"`
	ToTokenDecimals   uint64 `json:"toTokenDecimals"`
}
