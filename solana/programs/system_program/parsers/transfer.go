package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/system_program"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type TransferData struct {
	Discriminator uint32
	Lamports      uint64
}

func TransferParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*system_program.TransferAction, error) {
	var data TransferData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}

	action := system_program.TransferAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     system_program.ProgramName,
			InstructionName: "Transfer",
		},
		From:     result.AccountList[instruction.Accounts[0]],
		To:       result.AccountList[instruction.Accounts[1]],
		Lamports: data.Lamports,
	}

	return &action, nil
}
