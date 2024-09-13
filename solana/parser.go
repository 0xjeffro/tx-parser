package solana

import (
	"encoding/json"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func Parser(rpcData []byte) ([]types.ParsedResult, error) {
	var txs types.RawTxs
	err := json.Unmarshal(rpcData, &txs)
	if err != nil {
		return nil, err
	}
	var results []types.ParsedResult
	for _, tx := range txs {
		result := parser(tx)
		results = append(results, *result)
	}
	return results, nil
}

func parser(tx types.RawTx) *types.ParsedResult {
	var result types.ParsedResult
	result.RawTx = tx                                                               // set raw tx
	result = *getAccountList(&result)                                               // get account list
	result.Actions = make([]types.Action, len(tx.Transaction.Message.Instructions)) // init actions

	for i, _ := range tx.Transaction.Message.Instructions {
		action, err := router(&result, i)
		if err != nil {
			result.Actions[i] = types.UnknownAction{
				BaseAction: types.BaseAction{
					ProgramID:       result.AccountList[tx.Transaction.Message.Instructions[i].ProgramIDIndex],
					ProgramName:     "Unknown",
					InstructionName: "Unknown",
				},
				Error: err,
			}
		} else {
			result.Actions[i] = action
		}
	}
	return &result
}

func getAccountList(result *types.ParsedResult) *types.ParsedResult {
	length := len(result.RawTx.Transaction.Message.AccountKeys) +
		len(result.RawTx.Meta.LoadedAddresses.Writable) +
		len(result.RawTx.Meta.LoadedAddresses.Readonly)
	result.AccountList = make([]string, length)

	var i = 0
	for _, v := range result.RawTx.Transaction.Message.AccountKeys {
		result.AccountList[i] = v
		i++
	}
	for _, v := range result.RawTx.Meta.LoadedAddresses.Writable {
		result.AccountList[i] = v
		i++
	}
	for _, v := range result.RawTx.Meta.LoadedAddresses.Readonly {
		result.AccountList[i] = v
		i++
	}
	return result
}
