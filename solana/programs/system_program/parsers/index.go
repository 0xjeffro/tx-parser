package parsers

import (
	"encoding/binary"
	"github.com/thetafunction/tx-parser/solana/programs/system_program"
	"github.com/thetafunction/tx-parser/solana/types"
	"github.com/mr-tron/base58"
)

func InstructionRouter(result *types.ParsedResult, instruction types.Instruction) (types.Action, error) {
	data := instruction.Data
	decode, err := base58.Decode(data)
	if err != nil {
		return nil, err
	}
	discriminator := binary.LittleEndian.Uint32(decode[0:4])

	switch discriminator {
	case system_program.TransferDiscriminator:
		return TransferParser(result, instruction, decode)
	case system_program.CreateAccountWithSeedDiscriminator:
		return CreateAccountWithSeedParser(result, instruction, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     system_program.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
