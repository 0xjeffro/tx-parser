package computeBudget

import "github.com/0xjeffro/tx-parser/solana/types"

type SetComputeUnitLimitAction struct {
	types.BaseAction
	ComputeUnitLimit uint32 `json:"computeUnitLimit"`
}

type SetComputeUnitPriceAction struct {
	types.BaseAction
	MicroLamports uint64 `json:"microLamports"`
}
