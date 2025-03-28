package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func InitializeAccountParser(result *types.ParsedResult, instruction types.Instruction) (*tokenProgram.InitializeAccountAction, error) {

	action := tokenProgram.InitializeAccountAction{
		BaseAction: types.BaseAction{
			ProgramID:       tokenProgram.Program,
			ProgramName:     tokenProgram.ProgramName,
			InstructionName: "InitializeAccount",
		},
		Account:    result.AccountList[instruction.Accounts[0]],
		Mint:       result.AccountList[instruction.Accounts[1]],
		Owner:      result.AccountList[instruction.Accounts[2]],
		RentSysvar: result.AccountList[instruction.Accounts[3]],
	}
	return &action, nil
}
