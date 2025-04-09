package parsers

import (
	"github.com/thetafunction/tx-parser/solana/programs/token_program"
	"github.com/thetafunction/tx-parser/solana/types"
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
	case token_program.TransferDiscriminator:
		return TransferParser(result, instruction, decode)
	case token_program.TransferCheckedDiscriminator:
		return TransferCheckedParser(result, instruction, decode)
	case token_program.InitializeAccountDiscriminator:
		return InitializeAccountParser(result, instruction)

	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     token_program.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
