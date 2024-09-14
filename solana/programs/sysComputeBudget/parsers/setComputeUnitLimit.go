package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type Data struct {
	Discriminator    uint8
	ComputeUnitLimit uint32
}

func SetComputeUnitLimitParser(result *types.ParsedResult, i int, decodedData []byte) (*types.SysComputeBudgetSetComputeUnitLimitAction, error) {
	var data Data
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.SysComputeBudgetSetComputeUnitLimitAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[result.RawTx.Transaction.Message.Instructions[i].ProgramIDIndex],
			ProgramName:     "ComputeBudget",
			InstructionName: "setComputeUnitLimit",
		},
		ComputeUnitLimit: data.ComputeUnitLimit,
	}

	return &action, nil
}
