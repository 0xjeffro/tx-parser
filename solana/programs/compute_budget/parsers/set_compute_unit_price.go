package parsers

import (
	"github.com/thetafunction/tx-parser/solana/programs/compute_budget"
	"github.com/thetafunction/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type SetComputeUnitPriceData struct {
	Discriminator uint8
	Units         uint64
}

func SetComputeUnitPriceParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*compute_budget.SetComputeUnitPriceAction, error) {
	var data SetComputeUnitPriceData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := compute_budget.SetComputeUnitPriceAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     compute_budget.ProgramName,
			InstructionName: "SetComputeUnitPrice",
		},
		MicroLamports: data.Units,
	}
	return &action, nil
}
