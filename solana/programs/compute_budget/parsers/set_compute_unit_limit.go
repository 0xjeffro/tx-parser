package parsers

import (
	"github.com/thetafunction/tx-parser/solana/programs/compute_budget"
	"github.com/thetafunction/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type SetComputeUnitLimitData struct {
	Discriminator uint8
	Unit          uint32
}

func SetComputeUnitLimitParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*compute_budget.SetComputeUnitLimitAction, error) {
	var data SetComputeUnitLimitData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := compute_budget.SetComputeUnitLimitAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     compute_budget.ProgramName,
			InstructionName: "SetComputeUnitLimit",
		},
		ComputeUnitLimit: data.Unit,
	}

	return &action, nil
}
