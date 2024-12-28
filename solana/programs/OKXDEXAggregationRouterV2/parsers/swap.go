package parsers

import (
	"encoding/binary"
	"github.com/0xjeffro/tx-parser/solana/programs/OKXDEXAggregationRouterV2"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/systemProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	TokenProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/tokenProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func SwapParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*OKXDEXAggregationRouterV2.SwapAction, error) {
	who := result.AccountList[instruction.Accounts[0]]
	fromToken := result.AccountList[instruction.Accounts[3]]
	toToken := result.AccountList[instruction.Accounts[4]]
	fromTokenAmount := binary.LittleEndian.Uint64(decodedData[8:16])
	toTokenAmount := uint64(0)

	toTokenAccount := result.AccountList[instruction.Accounts[2]]

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
			case *types.SystemProgramTransferAction:
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
				if p.To == toTokenAccount {
					toTokenAmount += p.Amount
				}
			case *types.TokenProgramTransferCheckedAction:
				if p.To == toTokenAccount {
					toTokenAmount += p.Amount
				}
			}
		default:
			continue
		}
	}

	preTokenBalances := result.RawTx.Meta.PreTokenBalances
	postTokenBalances := result.RawTx.Meta.PostTokenBalances

	var fromTokenDecimals, toTokenDecimals uint64
	for _, b := range preTokenBalances {
		if b.Mint == fromToken {
			fromTokenDecimals = b.UITokenAmount.Decimals
		}
		if b.Mint == toToken {
			toTokenDecimals = b.UITokenAmount.Decimals
		}
	}
	for _, b := range postTokenBalances {
		if b.Mint == fromToken {
			fromTokenDecimals = b.UITokenAmount.Decimals
		}
		if b.Mint == toToken {
			toTokenDecimals = b.UITokenAmount.Decimals
		}
	}

	action := OKXDEXAggregationRouterV2.SwapAction{
		BaseAction: types.BaseAction{
			ProgramID:       OKXDEXAggregationRouterV2.Program,
			ProgramName:     OKXDEXAggregationRouterV2.ProgramName,
			InstructionName: "Swap",
		},
		Who:               who,
		FromToken:         fromToken,
		ToToken:           toToken,
		FromTokenAmount:   fromTokenAmount,
		ToTokenAmount:     toTokenAmount,
		FromTokenDecimals: fromTokenDecimals,
		ToTokenDecimals:   toTokenDecimals,
	}
	return &action, nil
}
