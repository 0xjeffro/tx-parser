package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/computeBudget"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/mr-tron/base58"
)

func Router(result *types.ParsedResult, i int) (types.Action, error) {
	instruction := result.RawTx.Transaction.Message.Instructions[i]
	data := instruction.Data
	decode, err := base58.Decode(data)
	if err != nil {
		return nil, err
	}
	discriminator := decode[0]

	switch discriminator {
	case computeBudget.SetComputeUnitLimitDiscriminator:
		return SetComputeUnitLimitParser(result, i, decode)
	case computeBudget.SetComputeUnitPriceDiscriminator:
		return SetComputeUnitPriceParser(result, i, decode)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       result.AccountList[instruction.ProgramIDIndex],
				ProgramName:     "computeBudget",
				InstructionName: "Unknown",
			},
		}, nil
	}
}
