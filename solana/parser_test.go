package solana

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/computeBudget"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
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
