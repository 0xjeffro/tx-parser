package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/jupiterDCA"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func CloseDcaParser(result *types.ParsedResult, instruction types.Instruction) (*jupiterDCA.CloseDcaAction, error) {
	return &jupiterDCA.CloseDcaAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     jupiterDCA.ProgramName,
			InstructionName: "CloseDca",
		},
		User:       result.AccountList[instruction.Accounts[0]],
		Dca:        result.AccountList[instruction.Accounts[1]],
		InputMint:  result.AccountList[instruction.Accounts[2]],
		OutputMint: result.AccountList[instruction.Accounts[3]],
		InAta:      result.AccountList[instruction.Accounts[4]],
		OutAta:     result.AccountList[instruction.Accounts[5]],
		UserInAta:  result.AccountList[instruction.Accounts[6]],
		UserOutAta: result.AccountList[instruction.Accounts[7]],
	}, nil
}
