package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
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
	discriminator := *(*[8]byte)(decode[:8])

	switch discriminator {
	case pumpfun.BuyDiscriminator:
		// do something
		return nil, nil
	case pumpfun.SellDiscriminator:
		return SellParser(result, i, decode)
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
