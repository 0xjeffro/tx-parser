package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/computeBudget"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type SetComputeUnitPriceData struct {
	Discriminator uint8
	Units         uint64
}

func SetComputeUnitPriceParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*computeBudget.SetComputeUnitPriceAction, error) {
	var data SetComputeUnitPriceData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := computeBudget.SetComputeUnitPriceAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     computeBudget.ProgramName,
			InstructionName: "SetComputeUnitPrice",
		},
		MicroLamports: data.Units,
	}
	return &action, nil
}
