package solana

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/computeBudget"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestPumpFunSell_0(t *testing.T) {
	jsonFile, err := os.Open("data/pumpfun_sell_0.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Errorf("Error reading JSON file: %v", err)
	}
	results, _ := Parser(byteValue)
	action := results[0].Actions[3]
	if sellAction, ok := action.(*types.PumpFunSellAction); ok {
		assert.Equal(t, sellAction.ProgramID, pumpfun.Program)
		assert.Equal(t, sellAction.ProgramName, "PumpFun")
		assert.Equal(t, sellAction.InstructionName, "Sell")
		assert.Equal(t, sellAction.Who, "4DdrfiDHpmx55i4SPssxVzS9ZaKLb8qr45NKY9Er9nNh")
		assert.Equal(t, sellAction.FromToken, "CnNVDyM7GXBBcH8giuRYm17YCn6kpFTTbnd6Tx4hpump")
		assert.Equal(t, sellAction.ToToken, globals.WSOL)
		assert.Equal(t, sellAction.FromTokenAmount, uint64(592443959000000))
		assert.Equal(t, sellAction.ToTokenAmount, uint64(35951023733))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestPumpFunBuy_0(t *testing.T) {
	jsonFile, err := os.Open("data/pumpfun_buy_0.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Errorf("Error reading JSON file: %v", err)
	}
	results, _ := Parser(byteValue)
	action := results[0].Actions[2]
	if buyAction, ok := action.(*types.PumpFunBuyAction); ok {
		assert.Equal(t, buyAction.ProgramID, pumpfun.Program)
		assert.Equal(t, buyAction.ProgramName, "PumpFun")
		assert.Equal(t, buyAction.InstructionName, "Buy")
		assert.Equal(t, buyAction.Who, "EfbbhahGNuhqEraRZXrwETfsaKxScngEttdQixWAW4WE")
		assert.Equal(t, buyAction.ToToken, "D66sbPtYqLnwPYcUqjDktrQtb52CnfF77C3sdpNpR8Si")
		assert.Equal(t, buyAction.FromToken, globals.WSOL)
		assert.Equal(t, buyAction.ToTokenAmount, uint64(56716561396327))
		assert.Equal(t, buyAction.FromTokenAmount, uint64(3000000000))
		assert.Equal(t, buyAction.FeeAmount, uint64(30000000))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestPumpFunBuy_1(t *testing.T) {
	jsonFile, err := os.Open("data/pumpfun_buy_1.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Errorf("Error reading JSON file: %v", err)
	}
	results, _ := Parser(byteValue)
	action := results[0].Actions[4]
	if buyAction, ok := action.(*types.PumpFunBuyAction); ok {
		assert.Equal(t, buyAction.ProgramID, pumpfun.Program)
		assert.Equal(t, buyAction.ProgramName, "PumpFun")
		assert.Equal(t, buyAction.InstructionName, "Buy")
		assert.Equal(t, buyAction.Who, "4DdrfiDHpmx55i4SPssxVzS9ZaKLb8qr45NKY9Er9nNh")
		assert.Equal(t, buyAction.ToToken, "7nYhDYAxQdFn2DRPcjBFPQQz5qb8HeFiJv9o9w6fpump")
		assert.Equal(t, buyAction.FromToken, globals.WSOL)
		assert.Equal(t, buyAction.ToTokenAmount, uint64(34612901212710))
		assert.Equal(t, buyAction.FromTokenAmount, uint64(1000000000))
		assert.Equal(t, buyAction.FeeAmount, uint64(10000000))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestComputeBudgetSetComputeUnitLimit(t *testing.T) {
	jsonFile, err := os.Open("data/pumpfun_sell_0.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Errorf("Error reading JSON file: %v", err)
	}
	results, _ := Parser(byteValue)
	action := results[0].Actions[0]
	if setAction, ok := action.(*types.ComputeBudgetSetComputeUnitLimitAction); ok {
		assert.Equal(t, setAction.ProgramID, computeBudget.Program)
		assert.Equal(t, setAction.ProgramName, "ComputeBudget")
		assert.Equal(t, setAction.InstructionName, "SetComputeUnitLimit")
		assert.Equal(t, setAction.ComputeUnitLimit, uint32(100000))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestComputeBudgetSetComputeUnitPrice(t *testing.T) {
	jsonFile, err := os.Open("data/pumpfun_sell_0.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Errorf("Error reading JSON file: %v", err)
	}
	results, _ := Parser(byteValue)
	action := results[0].Actions[2]
	if setAction, ok := action.(*types.ComputeBudgetSetComputeUnitPriceAction); ok {
		assert.Equal(t, setAction.ProgramID, computeBudget.Program)
		assert.Equal(t, setAction.ProgramName, "ComputeBudget")
		assert.Equal(t, setAction.InstructionName, "SetComputeUnitPrice")
		assert.Equal(t, setAction.MicroLamports, uint64(315000))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestSystemProgramTransfer(t *testing.T) {
	jsonFile, err := os.Open("data/pumpfun_sell_0.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Errorf("Error reading JSON file: %v", err)
	}
	results, _ := Parser(byteValue)
	action := results[0].Actions[1]
	if transferAction, ok := action.(*types.SystemProgramTransferAction); ok {
		assert.Equal(t, transferAction.ProgramID, systemProgram.Program)
		assert.Equal(t, transferAction.ProgramName, systemProgram.ProgramName)
		assert.Equal(t, transferAction.InstructionName, "Transfer")
		assert.Equal(t, transferAction.From, "4DdrfiDHpmx55i4SPssxVzS9ZaKLb8qr45NKY9Er9nNh")
		assert.Equal(t, transferAction.To, "HWEoBxYs7ssKuudEjzjmpfJVX7Dvi7wescFsVx2L5yoY")
		assert.Equal(t, transferAction.Lamports, uint64(10000000))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestTokenProgramTransfer(t *testing.T) {
	jsonFile, err := os.Open("data/token_transfer_0.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Errorf("Error reading JSON file: %v", err)
	}
	results, _ := Parser(byteValue)
	action := results[0].Actions[3]

	if transferAction, ok := action.(*types.TokenProgramTransferAction); ok {
		assert.Equal(t, transferAction.ProgramID, tokenProgram.Program)
		assert.Equal(t, transferAction.ProgramName, tokenProgram.ProgramName)
		assert.Equal(t, transferAction.InstructionName, "Transfer")
		assert.Equal(t, transferAction.From, "6MkxxzHTzR9KJfd7PSr1c47ufbY3mfPgyqfgJQso7vtL")
		assert.Equal(t, transferAction.To, "3ZiTyuCBHqFocKpFvU8wPwd6YM284aswxMZfad4w2ode")
		assert.Equal(t, transferAction.Amount, uint32(3000000000))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestTokenProgramTransferChecked(t *testing.T) {
	jsonFile, err := os.Open("data/transferChecked_0.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Errorf("Error reading JSON file: %v", err)
	}
	results, _ := Parser(byteValue)
	action := results[0].Actions[3]

	if transferAction, ok := action.(*types.TokenProgramTransferCheckedAction); ok {
		assert.Equal(t, transferAction.ProgramID, tokenProgram.Program)
		assert.Equal(t, transferAction.ProgramName, tokenProgram.ProgramName)
		assert.Equal(t, transferAction.InstructionName, "TransferChecked")
		assert.Equal(t, transferAction.From, "263W2H6WRhAiXv9r7fpwLya2AweHLGn6GGXH32gLEL7c")
		assert.Equal(t, transferAction.To, "cE8P4G5bRQt4LqR2Moa9bo4hChQGXLXMhNvvLkQt5Tj")
		assert.Equal(t, transferAction.Mint, "AZaxNof3Dy57yXKM99BidjB9vfnzJ3EBuEqiiVnQP71F")
		assert.Equal(t, transferAction.Amount, uint64(222000000000))
		assert.Equal(t, transferAction.Decimals, uint64(9))
	} else {
		t.Errorf("Error type assertion")
	}
}
