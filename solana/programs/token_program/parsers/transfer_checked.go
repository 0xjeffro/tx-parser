package parsers

import (
	"encoding/binary"
	"github.com/thetafunction/tx-parser/solana/programs/token_program"
	"github.com/thetafunction/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

func TransferCheckedParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*token_program.TransferCheckedAction, error) {
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
	action := token_program.TransferCheckedAction{
		BaseAction: types.BaseAction{
			ProgramID:       token_program.Program,
			ProgramName:     token_program.ProgramName,
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
