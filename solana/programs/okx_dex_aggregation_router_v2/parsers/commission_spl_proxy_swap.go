package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/okx_dex_aggregation_router_v2"
	"github.com/0xjeffro/tx-parser/solana/programs/system_program"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/system_program/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/token_program"
	TokenProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/token_program/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func CommissionSplProxySwapParser(result *types.ParsedResult, instruction types.Instruction) (*okx_dex_aggregation_router_v2.CommissionSplProxySwapAction, error) {

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
		if result.AccountList[instr.ProgramIDIndex] == okx_dex_aggregation_router_v2.Program && instr.Data == instruction.Data {
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

	action := okx_dex_aggregation_router_v2.CommissionSplProxySwapAction{
		BaseAction: types.BaseAction{
			ProgramID:       okx_dex_aggregation_router_v2.Program,
			ProgramName:     okx_dex_aggregation_router_v2.ProgramName,
			InstructionName: "CommissionSplProxySwap",
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
