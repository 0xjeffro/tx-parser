package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/jupiter_aggregator_v6"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/systemProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	TokenProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/tokenProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func RouteParser(result *types.ParsedResult, instruction types.Instruction) (*jupiter_aggregator_v6.RouteAction, error) {
	user := result.AccountList[instruction.Accounts[1]]
	fromTokenAccount := result.AccountList[instruction.Accounts[2]]
	toTokenAccount := result.AccountList[instruction.Accounts[3]]
	toToken := result.AccountList[instruction.Accounts[5]]

	var fromToken string
	var fromTokenAmount, toTokenAmount uint64
	var fromTokenDecimals, toTokenDecimals uint64

	if toToken == globals.WSOL {
		toTokenDecimals = globals.SOLDecimals
	}

	// get index of this instruction
	var instructionIndex int
	for idx, instr := range result.RawTx.Transaction.Message.Instructions {
		if result.AccountList[instr.ProgramIDIndex] == jupiter_aggregator_v6.Program && instr.Data == instruction.Data {
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
			case *tokenProgram.TransferAction:
				if p.From == fromTokenAccount {
					fromTokenAmount += p.Amount
				}
				if p.To == toTokenAccount {
					toTokenAmount += p.Amount
				}
			case *tokenProgram.TransferCheckedAction:
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

	if fromToken == "" {
		fromToken = globals.WSOL
		fromTokenDecimals = globals.SOLDecimals
	}
	if toToken == "" {
		toToken = globals.WSOL
		toTokenDecimals = globals.SOLDecimals
	}

	return &jupiter_aggregator_v6.RouteAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     jupiter_aggregator_v6.ProgramName,
			InstructionName: "Route",
			ActionLabel:     "SWAP",
		},
		SwapActionMixin: types.SwapActionMixin{
			Who:               user,
			FromToken:         fromToken,
			FromTokenAmount:   fromTokenAmount,
			FromTokenDecimals: fromTokenDecimals,
			ToToken:           toToken,
			ToTokenAmount:     toTokenAmount,
			ToTokenDecimals:   toTokenDecimals,
		},
	}, nil
}
