package parsers

import (
	"github.com/thetafunction/tx-parser/solana/programs/token_program"
	"github.com/thetafunction/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type TransferData struct {
	Discriminator uint8
	Amount        uint64
}

func TransferParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*token_program.TransferAction, error) {
	var data TransferData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := token_program.TransferAction{
		BaseAction: types.BaseAction{
			ProgramID:       token_program.Program,
			ProgramName:     token_program.ProgramName,
			InstructionName: "Transfer",
		},
		From:   result.AccountList[instruction.Accounts[0]],
		To:     result.AccountList[instruction.Accounts[1]],
		Amount: data.Amount,
	}
	return &action, nil
}
