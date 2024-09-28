package solana

import (
	"github.com/0xjeffro/tx-parser/solana/programs/U6m2CDdhRg"
	U6m2CDdhRgParsers "github.com/0xjeffro/tx-parser/solana/programs/U6m2CDdhRg/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/computeBudget"
	ComputeBudgetParsers "github.com/0xjeffro/tx-parser/solana/programs/computeBudget/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/jupiterDCA"
	JupiterDCA "github.com/0xjeffro/tx-parser/solana/programs/jupiterDCA/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	PumpfunParsers "github.com/0xjeffro/tx-parser/solana/programs/pumpfun/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/systemProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	TokenProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/tokenProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func router(result *types.ParsedResult, i int) (action types.Action, err error) {
	defer func() {
		if err := recover(); err != nil {
			action = types.UnknownAction{
				BaseAction: types.BaseAction{
					ProgramID:       "Unknown",
					ProgramName:     "Unknown",
					InstructionName: "Unknown",
				},
				Error: err.(error),
			}
		}
	}()
	programID := result.AccountList[result.RawTx.Transaction.Message.Instructions[i].ProgramIDIndex]
	instruction := result.RawTx.Transaction.Message.Instructions[i]
	switch programID {
	case systemProgram.Program:
		return SystemProgramParsers.InstructionRouter(result, instruction)
	case tokenProgram.Program:
		return TokenProgramParsers.InstructionRouter(result, instruction)
	case computeBudget.Program:
		return ComputeBudgetParsers.InstructionRouter(result, instruction)
	case pumpfun.Program:
		return PumpfunParsers.InstructionRouter(result, instruction)
	case U6m2CDdhRg.Program:
		return U6m2CDdhRgParsers.InstructionRouter(result, instruction)
	case jupiterDCA.Program:
		return JupiterDCA.InstructionRouter(result, instruction)
	default:
		return types.UnknownAction{
			BaseAction: types.BaseAction{
				ProgramID:       programID,
				ProgramName:     "Unknown",
				InstructionName: "Unknown",
			},
		}, nil
	}
}
