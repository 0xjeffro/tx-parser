package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type TransferData struct {
	Discriminator uint8
	Amount        uint32
}

func TransferParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*types.TokenProgramTransferAction, error) {
	var data TransferData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.TokenProgramTransferAction{
		BaseAction: types.BaseAction{
			ProgramID:       tokenProgram.Program,
			ProgramName:     tokenProgram.ProgramName,
			InstructionName: "Transfer",
		},
		From:   result.AccountList[instruction.Accounts[0]],
		To:     result.AccountList[instruction.Accounts[1]],
		Amount: data.Amount,
	}
	return &action, nil
}
