package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type CreateData struct {
	Discriminator uint64
	Name          string
	Symbol        string
	Uri           string
}

func CreateParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*types.PumpFunCreateAction, error) {
	var createData CreateData
	err := borsh.Deserialize(&createData, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.PumpFunCreateAction{
		BaseAction: types.BaseAction{
			ProgramID:       pumpfun.Program,
			ProgramName:     pumpfun.ProgramName,
			InstructionName: "Create",
		},
		Who:                    result.AccountList[instruction.Accounts[7]],
		Mint:                   result.AccountList[instruction.Accounts[0]],
		MintAuthority:          result.AccountList[instruction.Accounts[1]],
		BondingCurve:           result.AccountList[instruction.Accounts[2]],
		AssociatedBondingCurve: result.AccountList[instruction.Accounts[3]],
		MplTokenMetadata:       result.AccountList[instruction.Accounts[5]],
		MetaData:               result.AccountList[instruction.Accounts[6]],
		Name:                   createData.Name,
		Symbol:                 createData.Symbol,
		Uri:                    createData.Uri,
	}

	return &action, nil
}
