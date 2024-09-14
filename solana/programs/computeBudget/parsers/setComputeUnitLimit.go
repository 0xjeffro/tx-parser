package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type SetComputeUnitLimitData struct {
	Discriminator uint8
	Unit          uint32
}

func SetComputeUnitLimitParser(result *types.ParsedResult, i int, decodedData []byte) (*types.ComputeBudgetSetComputeUnitLimitAction, error) {
	var data SetComputeUnitLimitData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.ComputeBudgetSetComputeUnitLimitAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[result.RawTx.Transaction.Message.Instructions[i].ProgramIDIndex],
			ProgramName:     "computeBudget",
			InstructionName: "setComputeUnitLimit",
		},
		ComputeUnitLimit: data.Unit,
	}

	return &action, nil
}
