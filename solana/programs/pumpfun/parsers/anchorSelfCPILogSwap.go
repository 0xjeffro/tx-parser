package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/mr-tron/base58"
	"github.com/near/borsh-go"
)

type AnchorSelfCPILogSwapData struct {
	Discriminator        [16]byte
	Mint                 [32]byte
	SolAmount            uint64
	TokenAmount          uint64
	IsBuy                bool
	User                 [32]byte
	Timestamp            int64
	VirtualSolReserves   uint64
	VirtualTokenReserves uint64
}

func AnchorSelfCPILogSwapParser(decodedData []byte) (*types.PumpFunAnchorSelfCPILogSwapAction, error) {
	var data AnchorSelfCPILogSwapData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.PumpFunAnchorSelfCPILogSwapAction{
		BaseAction: types.BaseAction{
			ProgramID:       pumpfun.Program,
			ProgramName:     "pumpfun",
			InstructionName: "AnchorSelfCPILog Swap",
		},
		Mint:                 base58.Encode(data.Mint[:]),
		SolAmount:            data.SolAmount,
		TokenAmount:          data.TokenAmount,
		IsBuy:                data.IsBuy,
		User:                 base58.Encode(data.User[:]),
		Timestamp:            data.Timestamp,
		VirtualTokenReserves: data.VirtualTokenReserves,
		VirtualSolReserves:   data.VirtualSolReserves,
	}

	return &action, nil
}
