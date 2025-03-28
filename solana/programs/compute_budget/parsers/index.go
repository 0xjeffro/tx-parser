package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/compute_budget"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/mr-tron/base58"
)

func InstructionRouter(result *types.ParsedResult, instruction types.Instruction) (types.Action, error) {
	data := instruction.Data
	decode, err := base58.Decode(data)
	if err != nil {
		return nil, err
	}
	discriminator := decode[0]

	switch discriminator {
	case compute_budget.SetComputeUnitLimitDiscriminator:
		return SetComputeUnitLimitParser(result, instruction, decode)
	case compute_budget.SetComputeUnitPriceDiscriminator:
		return SetComputeUnitPriceParser(result, instruction, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     compute_budget.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
