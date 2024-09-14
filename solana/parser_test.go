package solana

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	"github.com/0xjeffro/tx-parser/solana/programs/sysComputeBudget"
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
		assert.Equal(t, sellAction.InstructionName, "sell")
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
	if setAction, ok := action.(*types.SysComputeBudgetSetComputeUnitLimitAction); ok {
		assert.Equal(t, setAction.ProgramID, sysComputeBudget.Program)
		assert.Equal(t, setAction.ProgramName, "ComputeBudget")
		assert.Equal(t, setAction.InstructionName, "setComputeUnitLimit")
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
	if setAction, ok := action.(*types.SysComputeBudgetSetComputeUnitPriceAction); ok {
		assert.Equal(t, setAction.ProgramID, sysComputeBudget.Program)
		assert.Equal(t, setAction.ProgramName, "ComputeBudget")
		assert.Equal(t, setAction.InstructionName, "setComputeUnitPrice")
		assert.Equal(t, setAction.MicroLamports, uint64(315000))
	} else {
		t.Errorf("Error type assertion")
	}
}
