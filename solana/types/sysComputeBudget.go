package types

type SysComputeBudgetSetComputeUnitLimitAction struct {
	BaseAction
	ComputeUnitLimit uint32 `json:"computeUnitLimit"`
}
