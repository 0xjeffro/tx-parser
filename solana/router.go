package solana

import (
	"github.com/thetafunction/tx-parser/solana/programs/compute_budget"
	ComputeBudgetParsers "github.com/thetafunction/tx-parser/solana/programs/compute_budget/parsers"
	"github.com/thetafunction/tx-parser/solana/programs/jupiter_aggregator_v6"
	JupiterAggregatorV6 "github.com/thetafunction/tx-parser/solana/programs/jupiter_aggregator_v6/parsers"
	"github.com/thetafunction/tx-parser/solana/programs/jupiter_dca"
	JupiterDCA "github.com/thetafunction/tx-parser/solana/programs/jupiter_dca/parsers"
	"github.com/thetafunction/tx-parser/solana/programs/okx_dex_aggregation_router_v2"
	OKXDEXAggregationRouterV2Parsers "github.com/thetafunction/tx-parser/solana/programs/okx_dex_aggregation_router_v2/parsers"
	"github.com/thetafunction/tx-parser/solana/programs/photon_program"
	PhotonProgramParsers "github.com/thetafunction/tx-parser/solana/programs/photon_program/parsers"
	"github.com/thetafunction/tx-parser/solana/programs/pumpfun"
	PumpfunParsers "github.com/thetafunction/tx-parser/solana/programs/pumpfun/parsers"
	"github.com/thetafunction/tx-parser/solana/programs/raydium_liquidity_pool_v4"
	RaydiumLiquidityPoolV4 "github.com/thetafunction/tx-parser/solana/programs/raydium_liquidity_pool_v4/parsers"
	"github.com/thetafunction/tx-parser/solana/programs/system_program"
	SystemProgramParsers "github.com/thetafunction/tx-parser/solana/programs/system_program/parsers"
	"github.com/thetafunction/tx-parser/solana/programs/token_program"
	TokenProgramParsers "github.com/thetafunction/tx-parser/solana/programs/token_program/parsers"
	"github.com/thetafunction/tx-parser/solana/types"
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
	case token_program.Program:
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
	case photon_program.Program:
		return PhotonProgramParsers.InstructionRouter(result, instruction)

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
