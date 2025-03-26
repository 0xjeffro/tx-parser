package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/OKXDEXAggregationRouterV2"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/systemProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	TokenProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/tokenProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func CommissionSolSwap2Parser(result *types.ParsedResult, instruction types.Instruction) (*OKXDEXAggregationRouterV2.CommissionSolSwap2Action, error) {

	var who string
	var fromToken, toToken string = globals.WSOL, globals.WSOL
	var fromTokenDecimals, toTokenDecimals uint64 = globals.SOLDecimals, globals.SOLDecimals
	var fromTokenAmount, toTokenAmount uint64

	who = result.AccountList[instruction.Accounts[0]]
	fromToken = result.AccountList[instruction.Accounts[3]]
	toToken = result.AccountList[instruction.Accounts[4]]

	var fromTokenAccount, toTokenAccount string
	fromTokenAccount = result.AccountList[instruction.Accounts[1]]
	toTokenAccount = result.AccountList[instruction.Accounts[2]]

	// get index of this instruction
	var instructionIndex int
	for idx, instr := range result.RawTx.Transaction.Message.Instructions {
		if result.AccountList[instr.ProgramIDIndex] == OKXDEXAggregationRouterV2.Program && instr.Data == instruction.Data {
			instructionIndex = idx
			break
		}
	}

	// get all innerInstructions for this instruction
	var innerInstructions []types.Instruction
	for _, innerInstruction := range result.RawTx.Meta.InnerInstructions {
		if innerInstruction.Index == instructionIndex {
			innerInstructions = innerInstruction.Instructions
			break
		}
	}

	for _, instr := range innerInstructions {
		programId := result.AccountList[instr.ProgramIDIndex]
		switch programId {
		case systemProgram.Program:
			parsedData, err := SystemProgramParsers.InstructionRouter(result, instr)
			if err != nil {
				continue
			}
			switch p := parsedData.(type) {
			case *systemProgram.TransferAction:
				if p.From == fromTokenAccount {
					fromTokenAmount += p.Lamports
				}
				if p.To == toTokenAccount {
					toTokenAmount += p.Lamports
				}
			}
		case tokenProgram.Program:
			parsedData, err := TokenProgramParsers.InstructionRouter(result, instr)
			if err != nil {
				continue
			}
			switch p := parsedData.(type) {
			case *types.TokenProgramTransferAction:
				if p.From == fromTokenAccount {
					fromTokenAmount += p.Amount
				}
				if p.To == toTokenAccount {
					toTokenAmount += p.Amount
				}
			case *types.TokenProgramTransferCheckedAction:
				if p.From == fromTokenAccount {
					fromTokenAmount += p.Amount
				}
				if p.To == toTokenAccount {
					toTokenAmount += p.Amount
				}
			}
		default:
			continue
		}
	}

	var tokenBalances []types.TokenBalance
	tokenBalances = append(tokenBalances, result.RawTx.Meta.PreTokenBalances...)
	tokenBalances = append(tokenBalances, result.RawTx.Meta.PostTokenBalances...)

	for _, tokenBalance := range tokenBalances {
		account := result.AccountList[tokenBalance.AccountIndex]
		if account == fromTokenAccount {
			fromToken = tokenBalance.Mint
			fromTokenDecimals = tokenBalance.UITokenAmount.Decimals
		} else if account == toTokenAccount {
			toToken = tokenBalance.Mint
			toTokenDecimals = tokenBalance.UITokenAmount.Decimals
		}
	}

	action := OKXDEXAggregationRouterV2.CommissionSolSwap2Action{
		BaseAction: types.BaseAction{
			ProgramID:       OKXDEXAggregationRouterV2.Program,
			ProgramName:     OKXDEXAggregationRouterV2.ProgramName,
			InstructionName: "CommissionSolSwap2",
		},
		Who:               who,
		ToToken:           toToken,
		FromToken:         fromToken,
		ToTokenAmount:     toTokenAmount,
		FromTokenAmount:   fromTokenAmount,
		FromTokenDecimals: fromTokenDecimals,
		ToTokenDecimals:   toTokenDecimals,
	}

	return &action, nil
}
