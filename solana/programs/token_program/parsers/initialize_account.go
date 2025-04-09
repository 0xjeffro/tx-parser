package parsers

import (
	"github.com/thetafunction/tx-parser/solana/programs/token_program"
	"github.com/thetafunction/tx-parser/solana/types"
)

func InitializeAccountParser(result *types.ParsedResult, instruction types.Instruction) (*token_program.InitializeAccountAction, error) {

	action := token_program.InitializeAccountAction{
		BaseAction: types.BaseAction{
			ProgramID:       token_program.Program,
			ProgramName:     token_program.ProgramName,
			InstructionName: "InitializeAccount",
		},
		Account:    result.AccountList[instruction.Accounts[0]],
		Mint:       result.AccountList[instruction.Accounts[1]],
		Owner:      result.AccountList[instruction.Accounts[2]],
		RentSysvar: result.AccountList[instruction.Accounts[3]],
	}
	return &action, nil
}
