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

func SharedAccountsRouteParser(result *types.ParsedResult, instruction types.Instruction) (*jupiter_aggregator_v6.SharedAccountRouteAction, error) {
	user := result.AccountList[instruction.Accounts[2]]
	fromTokenAccount := result.AccountList[instruction.Accounts[3]]
	toTokenAccount := result.AccountList[instruction.Accounts[6]]
	fromToken := result.AccountList[instruction.Accounts[7]]
	toToken := result.AccountList[instruction.Accounts[8]]

	var fromTokenAmount, toTokenAmount uint64

	var instructionIndex int
	for idx, instr := range result.RawTx.Transaction.Message.Instructions {
		if result.AccountList[instr.ProgramIDIndex] == jupiter_aggregator_v6.Program && instr.Data == instruction.Data {
			instructionIndex = idx
			break
		}
	}

	var instructions []types.Instruction
	for _, innerInstruction := range result.RawTx.Meta.InnerInstructions {
		if innerInstruction.Index == instructionIndex {
			instructions = innerInstruction.Instructions
			break
		}
	}

	for _, instr := range instructions {
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

	var fromTokenDecimals, toTokenDecimals uint64
	if fromToken == globals.WSOL {
		fromTokenDecimals = globals.SOLDecimals
	}
	if toToken == globals.WSOL {
		toTokenDecimals = globals.SOLDecimals
	}

	var tokenBalances []types.TokenBalance
	tokenBalances = append(tokenBalances, result.RawTx.Meta.PreTokenBalances...)
	tokenBalances = append(tokenBalances, result.RawTx.Meta.PostTokenBalances...)

	for _, tokenBalance := range tokenBalances {
		if fromTokenDecimals != 0 && toTokenDecimals != 0 {
			break
		}
		if fromToken == tokenBalance.Mint {
			fromTokenDecimals = tokenBalance.UITokenAmount.Decimals
		}
		if toToken == tokenBalance.Mint {
			toTokenDecimals = tokenBalance.UITokenAmount.Decimals
		}
	}

	return &jupiter_aggregator_v6.SharedAccountRouteAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     jupiter_aggregator_v6.ProgramName,
			InstructionName: "SharedAccountsRoute",
		},
		Who:               user,
		FromToken:         fromToken,
		FromTokenAmount:   fromTokenAmount,
		FromTokenDecimals: fromTokenDecimals,
		ToToken:           toToken,
		ToTokenAmount:     toTokenAmount,
		ToTokenDecimals:   toTokenDecimals,
	}, nil
}
