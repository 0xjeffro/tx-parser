package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
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
	case pumpfun.BuyDiscriminator:
		return BuyParser(result, instruction, decode)
	case pumpfun.SellDiscriminator:
		return SellParser(result, instruction, decode)
	case pumpfun.CreateDiscriminator:
		return CreateParser(result, instruction, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     pumpfun.ProgramName,
				InstructionName: "Unknown",
			},
		}, nil
	}
}
