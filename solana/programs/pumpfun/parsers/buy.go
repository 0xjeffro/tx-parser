package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/systemProgram/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type BuyData struct {
	Discriminator uint64
	Amount        uint64
	MaxSolCost    uint64
}

func BuyParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*types.PumpFunBuyAction, error) {
	var buyData BuyData
	err := borsh.Deserialize(&buyData, decodedData)
	if err != nil {
		return nil, err
	}

	var instructionIndex int
	for idx, instr := range result.RawTx.Transaction.Message.Instructions {
		if result.AccountList[instr.ProgramIDIndex] == pumpfun.Program && instr.Data == instruction.Data {
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

	feeRecipient := result.AccountList[instruction.Accounts[1]]
	bondingCurve := result.AccountList[instruction.Accounts[3]]

	feeAmount := uint64(0)
	buyTokenAmount := uint64(0)

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
				if p.To == feeRecipient {
					feeAmount += p.Lamports
				} else if p.To == bondingCurve {
					buyTokenAmount += p.Lamports
				}
			}
		default:
			continue
		}
	}

	action := types.PumpFunBuyAction{
		BaseAction: types.BaseAction{
			ProgramID:       pumpfun.Program,
			ProgramName:     pumpfun.ProgramName,
			InstructionName: "Buy",
		},
		Who:             result.AccountList[instruction.Accounts[6]],
		ToToken:         result.AccountList[instruction.Accounts[2]],
		FromToken:       globals.WSOL,
		ToTokenAmount:   buyData.Amount,
		FromTokenAmount: buyTokenAmount,
		FeeAmount:       feeAmount,
	}

	return &action, nil
}
