package photon_program

import "github.com/0xjeffro/tx-parser/solana/types"

type PumpfunBuyAction struct {
	types.BaseAction
	types.SwapActionMixin
	PumpfunFeeAmount uint64 `json:"pumpfunFee"`
	PhotonFeeAmount  uint64 `json:"photonFee"`
}
