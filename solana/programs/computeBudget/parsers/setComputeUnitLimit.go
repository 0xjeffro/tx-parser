package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/computeBudget"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type SetComputeUnitLimitData struct {
	Discriminator uint8
	Unit          uint32
}

func SetComputeUnitLimitParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*computeBudget.SetComputeUnitLimitAction, error) {
	var data SetComputeUnitLimitData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := computeBudget.SetComputeUnitLimitAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     computeBudget.ProgramName,
			InstructionName: "SetComputeUnitLimit",
		},
		ComputeUnitLimit: data.Unit,
	}

	return &action, nil
}
