package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
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
	case tokenProgram.TransferDiscriminator:
		return TransferParser(result, instruction, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     tokenProgram.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
