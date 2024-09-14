package types

type SysComputeBudgetSetComputeUnitLimitAction struct {
	BaseAction
	ComputeUnitLimit uint32 `json:"computeUnitLimit"`
}

type SysComputeBudgetSetComputeUnitPriceAction struct {
	BaseAction
	MicroLamports uint64 `json:"microLamports"`
}
