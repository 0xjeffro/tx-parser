package raydiumLiquidityPoolV4

import "github.com/0xjeffro/tx-parser/solana/types"

type SwapAction struct {
	types.BaseAction
	Who               string `json:"who"`
	FromToken         string `json:"fromToken"`
	FromTokenAmount   uint64 `json:"fromTokenAmount"`
	FromTokenDecimals uint64 `json:"fromTokenDecimals"`
	ToToken           string `json:"toToken"`
	ToTokenAmount     uint64 `json:"toTokenAmount"`
	ToTokenDecimals   uint64 `json:"toTokenDecimals"`

	MinimumAmountOut uint64 `json:"minimumAmountOut"`
}
