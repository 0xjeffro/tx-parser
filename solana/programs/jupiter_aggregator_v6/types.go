package jupiter_aggregator_v6

import "github.com/thetafunction/tx-parser/solana/types"

type SharedAccountRouteAction struct {
	types.BaseAction
	Who               string `json:"who"`
	FromToken         string `json:"fromToken"`
	FromTokenAmount   uint64 `json:"fromTokenAmount"`
	FromTokenDecimals uint64 `json:"fromTokenDecimals"`
	ToToken           string `json:"toToken"`
	ToTokenAmount     uint64 `json:"toTokenAmount"`
	ToTokenDecimals   uint64 `json:"toTokenDecimals"`
}

type RouteAction struct {
	types.BaseAction
	types.SwapActionMixin
}
