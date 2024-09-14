package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type TransferData struct {
	Discriminator uint32
	Lamports      uint64
}

func TransferParser(result *types.ParsedResult, i int, decodedData []byte) (*types.SystemProgramTransferAction, error) {
	var data TransferData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.SystemProgramTransferAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[result.RawTx.Transaction.Message.Instructions[i].ProgramIDIndex],
			ProgramName:     systemProgram.ProgramName,
			InstructionName: "Transfer",
		},
		From:     result.AccountList[result.RawTx.Transaction.Message.Instructions[i].Accounts[0]],
		To:       result.AccountList[result.RawTx.Transaction.Message.Instructions[i].Accounts[1]],
		Lamports: data.Lamports,
	}

	return &action, nil
}
