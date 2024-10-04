package parsers

import (
	"encoding/binary"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	"github.com/0xjeffro/tx-parser/solana/types"
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
	case systemProgram.TransferDiscriminator:
		return TransferParser(result, instruction, decode)
	case systemProgram.CreateAccountWithSeedDiscriminator:
		return CreateAccountWithSeedParser(result, instruction, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     systemProgram.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
