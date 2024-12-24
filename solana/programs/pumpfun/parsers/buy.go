package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/mr-tron/base58"
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

	buyTokenAmount := uint64(0)
	buySolAmount := uint64(0)

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
					buyTokenAmount = action.TokenAmount
					buySolAmount = action.SolAmount
				}
			}
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
		ToTokenAmount:   buyTokenAmount,
		FromTokenAmount: buySolAmount,
		MaxSolCost:      buyData.MaxSolCost,
	}

	return &action, nil
}
