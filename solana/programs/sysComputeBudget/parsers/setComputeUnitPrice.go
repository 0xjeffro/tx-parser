package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type SetComputeUnitPriceData struct {
	Discriminator uint8
	Units         uint64
}

func SetComputeUnitPriceParser(result *types.ParsedResult, i int, decodedData []byte) (*types.SysComputeBudgetSetComputeUnitPriceAction, error) {
	var data SetComputeUnitPriceData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.SysComputeBudgetSetComputeUnitPriceAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[result.RawTx.Transaction.Message.Instructions[i].ProgramIDIndex],
			ProgramName:     "ComputeBudget",
			InstructionName: "setComputeUnitPrice",
		},
		MicroLamports: data.Units,
	}
	return &action, nil
}
