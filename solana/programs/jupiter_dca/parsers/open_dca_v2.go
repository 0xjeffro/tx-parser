package parsers

import (
	"encoding/binary"
	"github.com/thetafunction/tx-parser/solana/programs/jupiter_dca"
	"github.com/thetafunction/tx-parser/solana/types"
)

func OpenDcaV2Parser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*jupiter_dca.OpenV2Action, error) {
	var inAmount uint64
	var inAmountPerCycle uint64
	var cycleFrequency int64
	var minOutAmount *uint64
	var maxOutAmount *uint64
	var startAt *int64
	dca := result.AccountList[instruction.Accounts[0]]
	user := result.AccountList[instruction.Accounts[1]]
	payer := result.AccountList[instruction.Accounts[2]]
	inputMint := result.AccountList[instruction.Accounts[3]]
	outputMint := result.AccountList[instruction.Accounts[4]]
	userAta := result.AccountList[instruction.Accounts[5]]
	inAta := result.AccountList[instruction.Accounts[6]]
	outAta := result.AccountList[instruction.Accounts[7]]

	inAmount = binary.LittleEndian.Uint64(decodedData[16:24])
	inAmountPerCycle = binary.LittleEndian.Uint64(decodedData[24:32])
	cycleFrequency = int64(binary.LittleEndian.Uint64(decodedData[32:40]))

	if decodedData[40] == 1 {
		minOutAmount = new(uint64)
		*minOutAmount = binary.LittleEndian.Uint64(decodedData[41:49])
	}
	if decodedData[49] == 1 {
		maxOutAmount = new(uint64)
		*maxOutAmount = binary.LittleEndian.Uint64(decodedData[50:58])
	}
	if decodedData[58] == 1 {
		startAt = new(int64)
		*startAt = int64(binary.LittleEndian.Uint64(decodedData[59:67]))
	}

	return &jupiter_dca.OpenV2Action{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     jupiter_dca.ProgramName,
			InstructionName: "OpenDcaV2",
		},
		InAmount:         inAmount,
		InAmountPerCycle: inAmountPerCycle,
		CycleFrequency:   cycleFrequency,
		MinOutAmount:     minOutAmount,
		MaxOutAmount:     maxOutAmount,
		StartAt:          startAt,

		Dca:        dca,
		User:       user,
		Payer:      payer,
		InputMint:  inputMint,
		OutputMint: outputMint,
		UserAta:    userAta,
		InAta:      inAta,
		OutAta:     outAta,
	}, nil
}
