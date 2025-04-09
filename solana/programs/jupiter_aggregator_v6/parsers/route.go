package parsers

import (
	"github.com/thetafunction/tx-parser/solana/globals"
	"github.com/thetafunction/tx-parser/solana/programs/jupiter_aggregator_v6"
	"github.com/thetafunction/tx-parser/solana/programs/system_program"
	SystemProgramParsers "github.com/thetafunction/tx-parser/solana/programs/system_program/parsers"
	"github.com/thetafunction/tx-parser/solana/programs/token_program"
	TokenProgramParsers "github.com/thetafunction/tx-parser/solana/programs/token_program/parsers"
	"github.com/thetafunction/tx-parser/solana/types"
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
		case system_program.Program:
			parsedData, err := SystemProgramParsers.InstructionRouter(result, instr)
			if err != nil {
				continue
			}
			switch p := parsedData.(type) {
			case *system_program.TransferAction:
				if p.From == fromTokenAccount {
					fromTokenAmount += p.Lamports
				}
				if p.To == toTokenAccount {
					toTokenAmount += p.Lamports
				}
			}
		case token_program.Program:
			parsedData, err := TokenProgramParsers.InstructionRouter(result, instr)
			if err != nil {
				continue
			}
			switch p := parsedData.(type) {
			case *token_program.TransferAction:
				if p.From == fromTokenAccount {
					fromTokenAmount += p.Amount
				}
				if p.To == toTokenAccount {
					toTokenAmount += p.Amount
				}
			case *token_program.TransferCheckedAction:
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
