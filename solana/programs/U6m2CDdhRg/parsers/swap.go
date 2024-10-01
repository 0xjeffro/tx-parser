package parsers

import (
	"encoding/binary"
	"github.com/0xjeffro/tx-parser/solana/programs/U6m2CDdhRg"
	"github.com/0xjeffro/tx-parser/solana/types"
	"strconv"
)

func SwapParser(result *types.ParsedResult, instruction types.Instruction, decodedData []byte) (*types.U6m2CDdhRgSwapAction, error) {
	who := result.AccountList[instruction.Accounts[0]]
	fromToken := result.AccountList[instruction.Accounts[3]]
	toToken := result.AccountList[instruction.Accounts[4]]
	fromTokenAmount := binary.LittleEndian.Uint64(decodedData[8:16])
	toTokenAmount := uint64(0)

	preTokenBalances := result.RawTx.Meta.PreTokenBalances
	postTokenBalances := result.RawTx.Meta.PostTokenBalances

	var preToTokenAmount, postToTokenAmount uint64
	var fromTokenDecimals, toTokenDecimals uint64
	for _, b := range preTokenBalances {
		if b.Mint == toToken && b.Owner == who {
			var err error
			preToTokenAmount, err = strconv.ParseUint(b.UITokenAmount.Amount, 10, 64)
			if err != nil {
				return nil, err
			}
		}
		if b.Mint == fromToken {
			fromTokenDecimals = b.UITokenAmount.Decimals
		}
		if b.Mint == toToken {
			toTokenDecimals = b.UITokenAmount.Decimals
		}
	}
	for _, b := range postTokenBalances {
		if b.Mint == toToken && b.Owner == who {
			var err error
			postToTokenAmount, err = strconv.ParseUint(b.UITokenAmount.Amount, 10, 64)
			if err != nil {
				return nil, err
			}
			break
		}
		if b.Mint == fromToken {
			fromTokenDecimals = b.UITokenAmount.Decimals
		}
		if b.Mint == toToken {
			toTokenDecimals = b.UITokenAmount.Decimals
		}
	}
	toTokenAmount = postToTokenAmount - preToTokenAmount

	action := types.U6m2CDdhRgSwapAction{
		BaseAction: types.BaseAction{
			ProgramID:       U6m2CDdhRg.Program,
			ProgramName:     U6m2CDdhRg.ProgramName,
			InstructionName: "Unknown",
		},
		Who:               who,
		FromToken:         fromToken,
		ToToken:           toToken,
		FromTokenAmount:   fromTokenAmount,
		ToTokenAmount:     toTokenAmount,
		FromTokenDecimals: fromTokenDecimals,
		ToTokenDecimals:   toTokenDecimals,
	}
	return &action, nil
}
