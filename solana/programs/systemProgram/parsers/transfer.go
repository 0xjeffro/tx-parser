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

func TransferParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*systemProgram.TransferAction, error) {
	var data TransferData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := systemProgram.TransferAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     systemProgram.ProgramName,
			InstructionName: "Transfer",
		},
		From:     result.AccountList[instruction.Accounts[0]],
		To:       result.AccountList[instruction.Accounts[1]],
		Lamports: data.Lamports,
	}

	return &action, nil
}
