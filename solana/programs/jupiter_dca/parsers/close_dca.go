package parsers

import (
	"github.com/thetafunction/tx-parser/solana/programs/jupiter_dca"
	"github.com/thetafunction/tx-parser/solana/types"
)

func CloseDcaParser(result *types.ParsedResult, instruction types.Instruction) (*jupiter_dca.CloseDcaAction, error) {
	return &jupiter_dca.CloseDcaAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     jupiter_dca.ProgramName,
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
