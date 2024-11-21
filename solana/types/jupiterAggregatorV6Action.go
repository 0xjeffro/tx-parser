package types

type JupiterAggregatorV6SharedAccountRouteAction struct {
	BaseAction
	Who             string `json:"who"`
	FromToken       string `json:"fromToken"`
	FromTokenAmount uint64 `json:"fromTokenAmount"`
	ToToken         string `json:"toToken"`
	ToTokenAmount   uint64 `json:"toTokenAmount"`
}
