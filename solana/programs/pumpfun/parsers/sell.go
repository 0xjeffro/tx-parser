package parsers

import (
	"github.com/thetafunction/tx-parser/solana/globals"
	"github.com/thetafunction/tx-parser/solana/programs/pumpfun"
	"github.com/thetafunction/tx-parser/solana/types"
	"github.com/mr-tron/base58"
	"github.com/near/borsh-go"
)

type SellData struct {
	Discriminator uint64
	Amount        uint64
	MinSolOutput  uint64
}

func SellParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*pumpfun.SellAction, error) {
	var sellData SellData
	err := borsh.Deserialize(&sellData, decodedData)
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

	sellTokenAmount := uint64(0)
	sellSolAmount := uint64(0)

	for _, instr := range instructions {
		programId := result.AccountList[instr.ProgramIDIndex]
		if programId == pumpfun.Program {
			data := instr.Data
			decode, err := base58.Decode(data)
			if err != nil {
				return nil, err
			}
			discriminator := *(*[16]byte)(decode[:16])
			mergedDiscriminator := make([]byte, 0, 16)
			mergedDiscriminator = append(mergedDiscriminator[:], pumpfun.AnchorSelfCPILogDiscriminator[:]...)
			mergedDiscriminator = append(mergedDiscriminator[:], pumpfun.AnchorSelfCPILogSwapDiscriminator[:]...)
			if discriminator == *(*[16]byte)(mergedDiscriminator[:]) {
				action, err := AnchorSelfCPILogSwapParser(decode)
				if err == nil {
					sellTokenAmount = action.TokenAmount
					sellSolAmount = action.SolAmount
				}
			}
		}
	}

	action := pumpfun.SellAction{
		BaseAction: types.BaseAction{
			ProgramID:       pumpfun.Program,
			ProgramName:     pumpfun.ProgramName,
			InstructionName: "Sell",
		},
		Who:             result.AccountList[instruction.Accounts[6]],
		FromToken:       result.AccountList[instruction.Accounts[2]],
		ToToken:         globals.WSOL,
		FromTokenAmount: sellTokenAmount,
		ToTokenAmount:   sellSolAmount,
		MinSolOutput:    sellData.MinSolOutput,
	}
	return &action, nil
}
