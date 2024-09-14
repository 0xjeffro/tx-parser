package solana

import (
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	PumpfunParsers "github.com/0xjeffro/tx-parser/solana/programs/pumpfun/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/sysComputeBudget"
	SysComputeBudgetParsers "github.com/0xjeffro/tx-parser/solana/programs/sysComputeBudget/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func router(result *types.ParsedResult, i int) (types.Action, error) {
	programID := result.AccountList[result.RawTx.Transaction.Message.Instructions[i].ProgramIDIndex]
	switch programID {
	case sysComputeBudget.Program:
		return SysComputeBudgetParsers.Router(result, i)
	case pumpfun.Program:
		return PumpfunParsers.Router(result, i)
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
