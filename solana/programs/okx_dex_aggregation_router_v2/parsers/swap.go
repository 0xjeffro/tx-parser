package parsers

import (
	"encoding/binary"
	"github.com/thetafunction/tx-parser/solana/programs/okx_dex_aggregation_router_v2"
	"github.com/thetafunction/tx-parser/solana/programs/system_program"
	SystemProgramParsers "github.com/thetafunction/tx-parser/solana/programs/system_program/parsers"
	"github.com/thetafunction/tx-parser/solana/programs/token_program"
	TokenProgramParsers "github.com/thetafunction/tx-parser/solana/programs/token_program/parsers"
	"github.com/thetafunction/tx-parser/solana/types"
)

func SwapParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*okx_dex_aggregation_router_v2.SwapAction, error) {
	who := result.AccountList[instruction.Accounts[0]]
	fromToken := result.AccountList[instruction.Accounts[3]]
	toToken := result.AccountList[instruction.Accounts[4]]
	fromTokenAmount := binary.LittleEndian.Uint64(decodedData[8:16])
	toTokenAmount := uint64(0)

	toTokenAccount := result.AccountList[instruction.Accounts[2]]

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
				if p.To == toTokenAccount {
					toTokenAmount += p.Amount
				}
			case *token_program.TransferCheckedAction:
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

	action := okx_dex_aggregation_router_v2.SwapAction{
		BaseAction: types.BaseAction{
			ProgramID:       okx_dex_aggregation_router_v2.Program,
			ProgramName:     okx_dex_aggregation_router_v2.ProgramName,
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
