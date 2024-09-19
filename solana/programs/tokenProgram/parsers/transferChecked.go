package parsers

import (
	"encoding/binary"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

func TransferCheckedParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*types.TokenProgramTransferCheckedAction, error) {
	var data TransferData
	err := borsh.Deserialize(&data, decodedData)
	if err != nil {
		return nil, err
	}
	amount := binary.LittleEndian.Uint64(decodedData[1:9])
	decimals := binary.LittleEndian.Uint64(decodedData[9:17])
	fromTokenAccount := result.AccountList[instruction.Accounts[0]]
	mint := result.AccountList[instruction.Accounts[1]]
	toTokenAccount := result.AccountList[instruction.Accounts[2]]
	action := types.TokenProgramTransferCheckedAction{
		BaseAction: types.BaseAction{
			ProgramID:       tokenProgram.Program,
			ProgramName:     tokenProgram.ProgramName,
			InstructionName: "TransferChecked",
		},
		From:     fromTokenAccount,
		To:       toTokenAccount,
		Mint:     mint,
		Amount:   amount,
		Decimals: decimals,
	}
	return &action, nil
}
