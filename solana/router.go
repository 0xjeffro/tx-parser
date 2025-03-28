package solana

import (
	"github.com/0xjeffro/tx-parser/solana/programs/compute_budget"
	ComputeBudgetParsers "github.com/0xjeffro/tx-parser/solana/programs/compute_budget/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/jupiter_aggregator_v6"
	JupiterAggregatorV6 "github.com/0xjeffro/tx-parser/solana/programs/jupiter_aggregator_v6/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/jupiter_dca"
	JupiterDCA "github.com/0xjeffro/tx-parser/solana/programs/jupiter_dca/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/okx_dex_aggregation_router_v2"
	OKXDEXAggregationRouterV2Parsers "github.com/0xjeffro/tx-parser/solana/programs/okx_dex_aggregation_router_v2/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	PumpfunParsers "github.com/0xjeffro/tx-parser/solana/programs/pumpfun/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/raydium_liquidity_pool_v4"
	RaydiumLiquidityPoolV4 "github.com/0xjeffro/tx-parser/solana/programs/raydium_liquidity_pool_v4/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/system_program"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/system_program/parsers"
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
	case system_program.Program:
		return SystemProgramParsers.InstructionRouter(result, instruction)
	case tokenProgram.Program:
		return TokenProgramParsers.InstructionRouter(result, instruction)
	case compute_budget.Program:
		return ComputeBudgetParsers.InstructionRouter(result, instruction)
	case pumpfun.Program:
		return PumpfunParsers.InstructionRouter(result, instruction)
	case jupiter_dca.Program:
		return JupiterDCA.InstructionRouter(result, instruction)
	case raydium_liquidity_pool_v4.Program:
		return RaydiumLiquidityPoolV4.InstructionRouter(result, instruction, instructionIdx)
	case jupiter_aggregator_v6.Program:
		return JupiterAggregatorV6.InstructionRouter(result, instruction)
	case okx_dex_aggregation_router_v2.Program:
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
