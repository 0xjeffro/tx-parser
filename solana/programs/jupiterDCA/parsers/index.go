package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/jupiterDCA"
	"github.com/0xjeffro/tx-parser/solana/types"
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
	case jupiterDCA.OpenDcaV2Discriminator:
		return OpenDcaV2Parser(result, instruction, decode)
	case jupiterDCA.EndAndCloseDiscriminator:
		return EndAndCloseParser(result, instruction, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     jupiterDCA.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
