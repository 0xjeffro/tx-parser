package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/near/borsh-go"
)

type SellData struct {
	Discriminator uint64
	Amount        uint64
	MinSolOutput  uint64
}

func SellParser(result *types.ParsedResult, i int, decodedData []byte) (*types.PumpFunSellAction, error) {
	var sellData SellData
	err := borsh.Deserialize(&sellData, decodedData)
	if err != nil {
		return nil, err
	}

	action := types.PumpFunSellAction{
		BaseAction: types.BaseAction{
			ProgramID:       pumpfun.Program,
			ProgramName:     pumpfun.ProgramName,
			InstructionName: "Sell",
		},
		Who:             result.AccountList[result.RawTx.Transaction.Message.Instructions[i].Accounts[6]],
		FromToken:       result.AccountList[result.RawTx.Transaction.Message.Instructions[i].Accounts[2]],
		ToToken:         globals.WSOL,
		FromTokenAmount: sellData.Amount,
		ToTokenAmount:   sellData.MinSolOutput,
	}
	return &action, nil
}
