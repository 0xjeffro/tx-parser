package parsers

import (
	"encoding/binary"
	"github.com/thetafunction/tx-parser/solana/globals"
	"github.com/thetafunction/tx-parser/solana/programs/pumpfun"
	"github.com/thetafunction/tx-parser/solana/programs/system_program"
	systemParsers "github.com/thetafunction/tx-parser/solana/programs/system_program/parsers"
	"github.com/thetafunction/tx-parser/solana/types"
	"github.com/mr-tron/base58"
	"github.com/near/borsh-go"
)

type BuyData struct {
	Discriminator uint64
	Amount        uint64
	MaxSolCost    uint64
}

func BuyParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*pumpfun.BuyAction, error) {
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

	var innerInstructions []types.Instruction // get all innerInstructions for this instruction
	for _, innerInstruction := range result.RawTx.Meta.InnerInstructions {
		if innerInstruction.Index == instructionIndex {
			innerInstructions = innerInstruction.Instructions
			break
		}
	}

	buyTokenAmount := uint64(0)
	buySolAmount := uint64(0)
	feeAmount := uint64(0)

	feeRecipient := result.AccountList[instruction.Accounts[1]]

	for _, instr := range innerInstructions {
		programId := result.AccountList[instr.ProgramIDIndex]
		data := instr.Data
		decode, err := base58.Decode(data)
		if err != nil {
			return nil, err
		}

		if programId == pumpfun.Program {
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
		} else if programId == system_program.Program {
			discriminator := binary.LittleEndian.Uint32(decode[0:4])
			if discriminator == system_program.TransferDiscriminator {
				action, err := systemParsers.TransferParser(result, instr, decode)
				if err == nil {
					if action.To == feeRecipient {
						feeAmount = action.Lamports
					}
				}
			}
		}
	}

	action := pumpfun.BuyAction{
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
		FeeAmount:       feeAmount,
	}
	return &action, nil
}
