package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/raydium_liquidity_pool_v4"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/mr-tron/base58"
)

func InstructionRouter(result *types.ParsedResult, instruction types.Instruction, instructionIdx int) (types.Action, error) {
	data := instruction.Data
	decode, err := base58.Decode(data)
	if err != nil {
		return nil, err
	}
	discriminator := decode[0]
	switch discriminator {
	case raydium_liquidity_pool_v4.SwapDiscriminator:
		return SwapParser(result, instruction, instructionIdx, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     raydium_liquidity_pool_v4.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
