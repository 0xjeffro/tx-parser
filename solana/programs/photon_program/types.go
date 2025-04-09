package photon_program

import "github.com/thetafunction/tx-parser/solana/types"

type PumpfunBuyAction struct {
	types.BaseAction
	types.SwapActionMixin
	PumpfunFeeAmount uint64 `json:"pumpfunFeeAmount"`
	PhotonFeeAmount  uint64 `json:"photonFeeAmount"`
}

type PumpfunSellAction struct {
	types.BaseAction
	types.SwapActionMixin
	PhotonFeeAmount uint64 `json:"photonFeeAmount"`
}
