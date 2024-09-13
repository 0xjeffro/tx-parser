package solana

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestPumpFunSell(t *testing.T) {
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
	if buyAction, ok := action.(*types.PumpFunSellAction); ok {
		assert.Equal(t, buyAction.ProgramID, pumpfun.Program)
		assert.Equal(t, buyAction.ProgramName, "PumpFun")
		assert.Equal(t, buyAction.InstructionName, "sell")
		assert.Equal(t, buyAction.Who, "4DdrfiDHpmx55i4SPssxVzS9ZaKLb8qr45NKY9Er9nNh")
		assert.Equal(t, buyAction.FromToken, "CnNVDyM7GXBBcH8giuRYm17YCn6kpFTTbnd6Tx4hpump")
		assert.Equal(t, buyAction.ToToken, globals.WSOL)
		assert.Equal(t, buyAction.FromTokenAmount, uint64(592443959000000))
		assert.Equal(t, buyAction.ToTokenAmount, uint64(35951023733))
	} else {
		t.Errorf("Error type assertion")
	}
}
