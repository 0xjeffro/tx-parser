package parsers

import (
	"github.com/thetafunction/tx-parser/solana/programs/jupiter_dca"
	"github.com/thetafunction/tx-parser/solana/types"
)

func EndAndCloseParser(result *types.ParsedResult, instruction types.Instruction) (*jupiter_dca.EndAndCloseAction, error) {
	return &jupiter_dca.EndAndCloseAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     jupiter_dca.ProgramName,
			InstructionName: "EndAndClose",
		},
		Keeper:     result.AccountList[instruction.Accounts[0]],
		Dca:        result.AccountList[instruction.Accounts[1]],
		InputMint:  result.AccountList[instruction.Accounts[2]],
		OutputMint: result.AccountList[instruction.Accounts[3]],
		InAta:      result.AccountList[instruction.Accounts[4]],
		OutAta:     result.AccountList[instruction.Accounts[5]],
		User:       result.AccountList[instruction.Accounts[6]],
		UserOutAta: result.AccountList[instruction.Accounts[7]],
	}, nil
}
