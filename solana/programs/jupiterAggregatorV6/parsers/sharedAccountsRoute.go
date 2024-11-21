package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/jupiterAggregatorV6"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/systemProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	TokenProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/tokenProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func SharedAccountsRouteParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*types.JupiterAggregatorV6SharedAccountRouteAction, error) {
	user := result.AccountList[instruction.Accounts[2]]
	fromTokenAccount := result.AccountList[instruction.Accounts[3]]
	toTokenAccount := result.AccountList[instruction.Accounts[6]]
	fromToken := result.AccountList[instruction.Accounts[7]]
	toToken := result.AccountList[instruction.Accounts[8]]

	var fromTokenAmount, toTokenAmount uint64

	var instructionIndex int
	for idx, instr := range result.RawTx.Transaction.Message.Instructions {
		if result.AccountList[instr.ProgramIDIndex] == jupiterAggregatorV6.Program && instr.Data == instruction.Data {
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
		case systemProgram.Program:
			parsedData, err := SystemProgramParsers.InstructionRouter(result, instr)
			if err != nil {
				continue
			}
			switch p := parsedData.(type) {
			case *types.SystemProgramTransferAction:
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

	return &types.JupiterAggregatorV6SharedAccountRouteAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     jupiterAggregatorV6.ProgramName,
			InstructionName: "SharedAccountsRoute",
		},
		Who:             user,
		FromToken:       fromToken,
		FromTokenAmount: fromTokenAmount,
		ToToken:         toToken,
		ToTokenAmount:   toTokenAmount,
	}, nil
}
