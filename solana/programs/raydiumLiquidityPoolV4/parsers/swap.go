package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/raydiumLiquidityPoolV4"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	TokenProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/tokenProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type SwapData struct {
	Discriminator    uint8
	AmountIn         uint64
	MinimumAmountOut uint64
}

func SwapParser(result *types.ParsedResult, instruction types.Instruction, instructionIdx int, decodedData []byte) (*types.RaydiumLiquidityPoolV4SwapAction, error) {
	var swapData SwapData
	err := borsh.Deserialize(&swapData, decodedData)
	if err != nil {
		return nil, err
	}

	var fromToken, toToken string = globals.WSOL, globals.WSOL
	var fromTokenDecimals, toTokenDecimals uint64 = globals.SOLDecimals, globals.SOLDecimals

	accountsLen := len(instruction.Accounts)
	who := result.AccountList[instruction.Accounts[accountsLen-1]]
	userSourceTokenAccount := result.AccountList[instruction.Accounts[accountsLen-3]]
	userDestinationTokenAccount := result.AccountList[instruction.Accounts[accountsLen-2]]

	tokenBalances := append([]types.TokenBalance{}, result.RawTx.Meta.PreTokenBalances[:]...)
	tokenBalances = append(tokenBalances, result.RawTx.Meta.PostTokenBalances[:]...)

	for _, tb := range tokenBalances {
		tokenAccount := result.AccountList[tb.AccountIndex]
		if tokenAccount == userSourceTokenAccount {
			fromToken = tb.Mint
			fromTokenDecimals = tb.UITokenAmount.Decimals
		} else if tokenAccount == userDestinationTokenAccount {
			toToken = tb.Mint
			toTokenDecimals = tb.UITokenAmount.Decimals
		}
	}

	var toTokenAmount uint64

	var associatedInnerInstructions []types.Instruction
	for _, innerInstruction := range result.RawTx.Meta.InnerInstructions {
		if innerInstruction.Index == instructionIdx {
			associatedInnerInstructions = append(associatedInnerInstructions, innerInstruction.Instructions...)
			break
		}
	}
	for _, innerInstruction := range associatedInnerInstructions {
		if result.AccountList[innerInstruction.ProgramIDIndex] == tokenProgram.Program {
			action, err := TokenProgramParsers.InstructionRouter(result, innerInstruction)
			if err != nil {
				continue
			}
			transferAction, ok := action.(*types.TokenProgramTransferAction)
			if ok {
				if transferAction.To == userDestinationTokenAccount {
					toTokenAmount = transferAction.Amount
				}
			}
		}
	}

	action := types.RaydiumLiquidityPoolV4SwapAction{
		BaseAction: types.BaseAction{
			ProgramID:       raydiumLiquidityPoolV4.Program,
			ProgramName:     raydiumLiquidityPoolV4.ProgramName,
			InstructionName: "Swap",
		},
		Who:               who,
		FromToken:         fromToken,
		FromTokenAmount:   swapData.AmountIn,
		FromTokenDecimals: fromTokenDecimals,
		ToToken:           toToken,
		ToTokenAmount:     toTokenAmount,
		ToTokenDecimals:   toTokenDecimals,

		MinimumAmountOut: swapData.MinimumAmountOut,
	}

	return &action, nil
}
