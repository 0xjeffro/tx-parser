package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/photon_program"
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
	case photon_program.PumpfunBuyDiscriminator:
		return PumpfunBuyParser(result, instruction)
	case photon_program.PumpfunSellDiscriminator:
		return PumpfunSellParser(result, instruction)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     photon_program.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
