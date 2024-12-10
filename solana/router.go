package solana

import (
	"github.com/0xjeffro/tx-parser/solana/programs/OKXDEXAggregationRouterV2"
	OKXDEXAggregationRouterV2Parsers "github.com/0xjeffro/tx-parser/solana/programs/OKXDEXAggregationRouterV2/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/computeBudget"
	ComputeBudgetParsers "github.com/0xjeffro/tx-parser/solana/programs/computeBudget/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/jupiterAggregatorV6"
	JupiterAggregatorV6 "github.com/0xjeffro/tx-parser/solana/programs/jupiterAggregatorV6/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/jupiterDCA"
	JupiterDCA "github.com/0xjeffro/tx-parser/solana/programs/jupiterDCA/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	PumpfunParsers "github.com/0xjeffro/tx-parser/solana/programs/pumpfun/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/raydiumLiquidityPoolV4"
	RaydiumLiquidityPoolV4 "github.com/0xjeffro/tx-parser/solana/programs/raydiumLiquidityPoolV4/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/systemProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	TokenProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/tokenProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func router(result *types.ParsedResult, instructionIdx int) (action types.Action, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			action = nil
		}
	}()
	programID := result.AccountList[result.RawTx.Transaction.Message.Instructions[instructionIdx].ProgramIDIndex]
	instruction := result.RawTx.Transaction.Message.Instructions[instructionIdx]
	switch programID {
	case systemProgram.Program:
		return SystemProgramParsers.InstructionRouter(result, instruction)
	case tokenProgram.Program:
		return TokenProgramParsers.InstructionRouter(result, instruction)
	case computeBudget.Program:
		return ComputeBudgetParsers.InstructionRouter(result, instruction)
	case pumpfun.Program:
		return PumpfunParsers.InstructionRouter(result, instruction)
	case jupiterDCA.Program:
		return JupiterDCA.InstructionRouter(result, instruction)
	case raydiumLiquidityPoolV4.Program:
		return RaydiumLiquidityPoolV4.InstructionRouter(result, instruction, instructionIdx)
	case jupiterAggregatorV6.Program:
		return JupiterAggregatorV6.InstructionRouter(result, instruction)
	case OKXDEXAggregationRouterV2.Program:
		return OKXDEXAggregationRouterV2Parsers.InstructionRouter(result, instruction)

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
