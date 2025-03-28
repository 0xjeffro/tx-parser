package parsers

import (
	"encoding/binary"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	"github.com/0xjeffro/tx-parser/solana/types"
	solanago "github.com/gagliardetto/solana-go"
)

func CreateAccountWithSeedParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*systemProgram.CreateAccountWithSeedAction, error) {
	basePubKey := solanago.PublicKeyFromBytes(decodedData[4:36])
	seedLength := binary.LittleEndian.Uint64(decodedData[36:44])
	seed := string(decodedData[44 : 44+seedLength])
	lamports := binary.LittleEndian.Uint64(decodedData[44+seedLength : 44+seedLength+8])
	space := binary.LittleEndian.Uint64(decodedData[44+seedLength+8 : 44+seedLength+16])
	ownerPubKey := solanago.PublicKeyFromBytes(decodedData[44+seedLength+16 : 44+seedLength+16+32])

	action := systemProgram.CreateAccountWithSeedAction{
		BaseAction: types.BaseAction{
			ProgramID:       result.AccountList[instruction.ProgramIDIndex],
			ProgramName:     systemProgram.ProgramName,
			InstructionName: "CreateAccountWithSeed",
		},
		Who:        result.AccountList[instruction.Accounts[0]],
		NewAccount: result.AccountList[instruction.Accounts[1]],
		Base:       basePubKey.String(),
		Seed:       seed,
		Lamports:   lamports,
		Space:      space,
		Owner:      ownerPubKey.String(),
	}

	return &action, nil
}
