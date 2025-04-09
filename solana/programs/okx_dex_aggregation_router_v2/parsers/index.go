package parsers

import (
	"github.com/thetafunction/tx-parser/solana/programs/okx_dex_aggregation_router_v2"
	"github.com/thetafunction/tx-parser/solana/types"
	"github.com/mr-tron/base58"
)

func InstructionRouter(result *types.ParsedResult, instruction types.Instruction) (types.Action, error) {
	data := instruction.Data
	decode, err := base58.Decode(data)
	if err != nil {
		return nil, err
	}
	discriminator := *(*[8]byte)(decode[:8])

	switch discriminator {
	case okx_dex_aggregation_router_v2.CommissionSplProxySwapDiscriminator:
		return CommissionSplProxySwapParser(result, instruction)
	case okx_dex_aggregation_router_v2.SwapDiscriminator:
		return SwapParser(result, instruction, decode)
	case okx_dex_aggregation_router_v2.CommissionSolSwap2Discriminator:
		return CommissionSolSwap2Parser(result, instruction)

	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     okx_dex_aggregation_router_v2.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
