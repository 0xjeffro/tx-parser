package parsers

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/photon_program"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	PumpfunParsers "github.com/0xjeffro/tx-parser/solana/programs/pumpfun/parsers"
	"github.com/0xjeffro/tx-parser/solana/programs/system_program"
	SystemProgramParsers "github.com/0xjeffro/tx-parser/solana/programs/system_program/parsers"
	"github.com/0xjeffro/tx-parser/solana/types"
)

func PumpfunSellParser(result *types.ParsedResult, instruction types.Instruction) (*photon_program.PumpfunSellAction, error) {

	var who string
	var fromToken, toToken string = "", globals.WSOL
	var fromTokenDecimals, toTokenDecimals uint64 = 0, globals.SOLDecimals
	var fromTokenAmount, toTokenAmount uint64
	var photonFeeAmount uint64

	fromTokenAccount := result.AccountList[instruction.Accounts[6]]
	photonFeeAccount := result.AccountList[instruction.Accounts[8]]

	who = result.AccountList[instruction.Accounts[7]]
	fromToken = result.AccountList[instruction.Accounts[3]]

	// get index of this instruction
	var instructionIndex int
	for idx, instr := range result.RawTx.Transaction.Message.Instructions {
		if result.AccountList[instr.ProgramIDIndex] == photon_program.Program && instr.Data == instruction.Data {
			instructionIndex = idx
			break
		}
	}

	// get all innerInstructions for this instruction
	var innerInstructions []types.Instruction
	for _, innerInstruction := range result.RawTx.Meta.InnerInstructions {
		if innerInstruction.Index == instructionIndex {
			innerInstructions = innerInstruction.Instructions
			break
		}
	}

	for _, instr := range innerInstructions {
		programId := result.AccountList[instr.ProgramIDIndex]
		switch programId {
		case system_program.Program:
			parsedData, err := SystemProgramParsers.InstructionRouter(result, instr)
			if err != nil {
				continue
			}
			switch p := parsedData.(type) {
			case *system_program.TransferAction:
				if p.To == photonFeeAccount {
					photonFeeAmount = p.Lamports
				}
			}
		case pumpfun.Program:
			parsedData, err := PumpfunParsers.InstructionRouter(result, instr)
			if err != nil {
				continue
			}
			switch p := parsedData.(type) {
			case *pumpfun.AnchorSelfCPILogSwapAction:
				fromToken = p.Mint
				fromTokenAmount = p.TokenAmount
				toTokenAmount = p.SolAmount
			default:
				continue
			}
		}
	}

	var tokenBalances []types.TokenBalance
	tokenBalances = append(tokenBalances, result.RawTx.Meta.PreTokenBalances...)
	tokenBalances = append(tokenBalances, result.RawTx.Meta.PostTokenBalances...)

	for _, tokenBalance := range tokenBalances {
		account := result.AccountList[tokenBalance.AccountIndex]
		if account == fromTokenAccount {
			fromTokenDecimals = tokenBalance.UITokenAmount.Decimals
		}
	}

	action := photon_program.PumpfunSellAction{
		BaseAction: types.BaseAction{
			ProgramID:       photon_program.Program,
			ProgramName:     photon_program.ProgramName,
			InstructionName: "PumpfunSell",
			ActionLabel:     "SWAP",
		},
		SwapActionMixin: types.SwapActionMixin{
			Who:               who,
			ToToken:           toToken,
			FromToken:         fromToken,
			ToTokenAmount:     toTokenAmount,
			FromTokenAmount:   fromTokenAmount,
			FromTokenDecimals: fromTokenDecimals,
			ToTokenDecimals:   toTokenDecimals,
		},
		PhotonFeeAmount: photonFeeAmount,
	}

	return &action, nil
}
