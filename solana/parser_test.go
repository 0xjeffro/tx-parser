package solana

import (
	"github.com/0xjeffro/tx-parser/solana/globals"
	"github.com/0xjeffro/tx-parser/solana/programs/U6m2CDdhRg"
	"github.com/0xjeffro/tx-parser/solana/programs/computeBudget"
	"github.com/0xjeffro/tx-parser/solana/programs/jupiterDCA"
	"github.com/0xjeffro/tx-parser/solana/programs/pumpfun"
	"github.com/0xjeffro/tx-parser/solana/programs/systemProgram"
	"github.com/0xjeffro/tx-parser/solana/programs/tokenProgram"
	"github.com/0xjeffro/tx-parser/solana/types"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestBrokenData_0(t *testing.T) {
	jsonFile, err := os.Open("data/broken_data_0.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	_, err = Parser(byteValue)
	assert.NotEqual(t, err, nil)
}

func TestBrokenData_1(t *testing.T) {
	jsonFile, err := os.Open("data/broken_data_1.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	results, err := Parser(byteValue)
	assert.Equal(t, results[0].Actions[0].GetProgramID(), "Unknown")
	assert.Equal(t, results[0].Actions[0].GetProgramName(), "Unknown")
	assert.Equal(t, results[0].Actions[0].GetInstructionName(), "Unknown")

}

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

func TestU6m2CDdhRgSwap(t *testing.T) {
	jsonFile, err := os.Open("data/U6m2CDdhRg_swap_0.json")
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

	if transferAction, ok := action.(*types.U6m2CDdhRgSwapAction); ok {
		assert.Equal(t, transferAction.ProgramID, U6m2CDdhRg.Program)
		assert.Equal(t, transferAction.ProgramName, U6m2CDdhRg.ProgramName)
		assert.Equal(t, transferAction.InstructionName, "Unknown")
		assert.Equal(t, transferAction.FromToken, "7atgF8KQo4wJrD5ATGX7t1V2zVvykPJbFfNeVf1icFv1")
		assert.Equal(t, transferAction.FromTokenAmount, uint64(358800))
		assert.Equal(t, transferAction.FromTokenDecimals, uint64(2))
		assert.Equal(t, transferAction.ToToken, "ED5nyyWEzpPPiWimP8vYm7sD7TD3LAt3Q3gRTWHzPJBY")
		assert.Equal(t, transferAction.ToTokenAmount, uint64(74619))
		assert.Equal(t, transferAction.ToTokenDecimals, uint64(6))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestU6m2CDdhRgSwap1(t *testing.T) {
	jsonFile, err := os.Open("data/U6m2CDdhRg_swap_1.json")
	if err != nil {
		t.Errorf("Error opening JSON file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		t.Errorf("Error reading JSON file: %v", err)
	}
	results, _ := Parser(byteValue)
	action := results[0].Actions[7]

	if transferAction, ok := action.(*types.U6m2CDdhRgSwapAction); ok {
		assert.Equal(t, transferAction.ProgramID, U6m2CDdhRg.Program)
		assert.Equal(t, transferAction.ProgramName, U6m2CDdhRg.ProgramName)
		assert.Equal(t, transferAction.InstructionName, "Unknown")
		assert.Equal(t, transferAction.FromToken, "So11111111111111111111111111111111111111112")
		assert.Equal(t, transferAction.FromTokenAmount, uint64(10000000000))
		assert.Equal(t, transferAction.FromTokenDecimals, uint64(9))
		assert.Equal(t, transferAction.ToToken, "KMnDBXcPXoz6oMJW5XG4tXdwSWpmWEP2RQM1Uujpump")
		assert.Equal(t, transferAction.ToTokenAmount, uint64(998528432013))
		assert.Equal(t, transferAction.ToTokenDecimals, uint64(6))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestJupiterDcaOpenDcaV2_0(t *testing.T) {
	jsonFile, err := os.Open("data/jupiterDca_openDcaV2_0.json")
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

	if DcaAction, ok := action.(*types.JupiterDcaOpenV2Action); ok {
		assert.Equal(t, DcaAction.ProgramID, jupiterDCA.Program)
		assert.Equal(t, DcaAction.ProgramName, jupiterDCA.ProgramName)
		assert.Equal(t, DcaAction.Dca, "7F3Wg8gzekghzqPLGubCSSjZGj6ihVb14A6QmGKjNL92")
		assert.Equal(t, DcaAction.User, "BxDjGoj9y33tmkKMK5dRHeUGkSGWEs4H51uHoQaBv9Wz")
		assert.Equal(t, DcaAction.Payer, "BxDjGoj9y33tmkKMK5dRHeUGkSGWEs4H51uHoQaBv9Wz")
		assert.Equal(t, DcaAction.InputMint, "ED5nyyWEzpPPiWimP8vYm7sD7TD3LAt3Q3gRTWHzPJBY")
		assert.Equal(t, DcaAction.OutputMint, "So11111111111111111111111111111111111111112")
		assert.Equal(t, DcaAction.UserAta, "8f53JuPmhMhYDhJwGawysBvAWmWyAA1oqXmUT475QjDr")
		assert.Equal(t, DcaAction.InAta, "A9Go8ThBWWue7Jcpnrzd8RegjJ1weJQMjBkJWiEuao6f")
		assert.Equal(t, DcaAction.OutAta, "3YYfQQkd4c97KR3ieHgMJvs52VaydLyWhCi5wPXmGgx1")
		assert.Equal(t, DcaAction.InAmount, uint64(217363000000))
		assert.Equal(t, DcaAction.InAmountPerCycle, uint64(3952054545))
		assert.Equal(t, DcaAction.CycleFrequency, int64(60))
		assert.Equal(t, *DcaAction.MinOutAmount, uint64(0))
		assert.Equal(t, *DcaAction.MaxOutAmount, uint64(0))
		assert.Equal(t, *DcaAction.StartAt, int64(0))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestJupiterDcaOpenDcaV2_1(t *testing.T) {
	jsonFile, err := os.Open("data/jupiterDca_openDcaV2_1.json")
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

	if DcaAction, ok := action.(*types.JupiterDcaOpenV2Action); ok {
		assert.Equal(t, DcaAction.ProgramID, jupiterDCA.Program)
		assert.Equal(t, DcaAction.ProgramName, jupiterDCA.ProgramName)
		assert.Equal(t, DcaAction.Dca, "GxMofmfnZw6ia9DE2KoqjgEd1vp7VuSrhqYqLWLW6bXf")
		assert.Equal(t, DcaAction.User, "4vYWDeg6NHUt1VoUZoV8RqZA51AUEd4VCCghA4arfuH7")
		assert.Equal(t, DcaAction.Payer, "4vYWDeg6NHUt1VoUZoV8RqZA51AUEd4VCCghA4arfuH7")
		assert.Equal(t, DcaAction.InputMint, "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")
		assert.Equal(t, DcaAction.OutputMint, "63LfDmNb3MQ8mw9MtZ2To9bEA2M71kZUUGq5tiJxcqj9")
		assert.Equal(t, DcaAction.UserAta, "92Q2QXq4oJBsi1a5KWPLVPaFqbrKYKX6USArFeyyrBma")
		assert.Equal(t, DcaAction.InAta, "69Z3oyToqfVhw4b5qzW4zdu4QV5UMfhN5wG9J2qxoXQp")
		assert.Equal(t, DcaAction.OutAta, "3Gfqf1vRXqd7V6VTn9sgdLzJEqfRUueDMKHv8CUGL1ud")
		assert.Equal(t, DcaAction.InAmount, uint64(500000000))
		assert.Equal(t, DcaAction.InAmountPerCycle, uint64(25000000))
		assert.Equal(t, DcaAction.CycleFrequency, int64(300))
		assert.Equal(t, *DcaAction.MinOutAmount, uint64(119047619))
		assert.Equal(t, *DcaAction.MaxOutAmount, uint64(138888888))
		assert.Equal(t, *DcaAction.StartAt, int64(0))
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestJupiterDcaEndAndClose_0(t *testing.T) {
	jsonFile, err := os.Open("data/jupiterDca_endAndClose_0.json")
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

	if DcaAction, ok := action.(*types.JupiterDcaEndAndCloseAction); ok {
		assert.Equal(t, DcaAction.ProgramID, jupiterDCA.Program)
		assert.Equal(t, DcaAction.ProgramName, jupiterDCA.ProgramName)
		assert.Equal(t, DcaAction.Keeper, "JD25qVdtd65FoiXNmR89JjmoJdYk9sjYQeSTZAALFiMy")
		assert.Equal(t, DcaAction.Dca, "Cf8vzN89jMzfgg3XpNGMTkzwihzbm1AHnQy1bVyeLVcy")
		assert.Equal(t, DcaAction.InputMint, "E1kvzJNxShvvWTrudokpzuc789vRiDXfXG3duCuY6ooE")
		assert.Equal(t, DcaAction.OutputMint, "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")
		assert.Equal(t, DcaAction.InAta, "DZAo37cFYiGJYhN9MjJ7EBz5CHvCJsax6yhZwgJmTxBT")
		assert.Equal(t, DcaAction.OutAta, "9EjKBrGmodH9vVWgiMu4oK1LnKxMQK9rFMeLpTzVRsDE")
		assert.Equal(t, DcaAction.User, "E4R1iXBJh8dN869akjGskM8DYkawncMRmtUsuYcTnM5S")
		assert.Equal(t, DcaAction.UserOutAta, "2HXai2nbYof8Mc5LEXSxLKHHtjJBnFqFFoSjtSj2LuFS")
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestJupiterDcaEndAndClose_1(t *testing.T) {
	jsonFile, err := os.Open("data/jupiterDca_endAndClose_1.json")
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

	if DcaAction, ok := action.(*types.JupiterDcaEndAndCloseAction); ok {
		assert.Equal(t, DcaAction.ProgramID, jupiterDCA.Program)
		assert.Equal(t, DcaAction.ProgramName, jupiterDCA.ProgramName)
		assert.Equal(t, DcaAction.Keeper, "JD38n7ynKYcgPpF7k1BhXEeREu1KqptU93fVGy3S624k")
		assert.Equal(t, DcaAction.Dca, "DdtdfnA7mPfzbxDfU2hD7DWNe64fWL9eMk3Tb9jw9ncr")
		assert.Equal(t, DcaAction.InputMint, "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")
		assert.Equal(t, DcaAction.OutputMint, "So11111111111111111111111111111111111111112")
		assert.Equal(t, DcaAction.InAta, "DtzRqfDTAjNV2YmSyMgbPTw1kKpkRqUwtsHoVuAL829T")
		assert.Equal(t, DcaAction.OutAta, "4WFAeSZU7AWyw8ZCyo4xFvJRevaasXVHxQcPJ4TtKe1y")
		assert.Equal(t, DcaAction.User, "5aT271zahvBR27MUUkLj8B7KtDGF8r2rZpFHerZi1yEn")
		assert.Equal(t, DcaAction.UserOutAta, "DCA265Vj8a9CEuX1eb1LWRnDT7uK6q1xMipnNyatn23M")
	} else {
		t.Errorf("Error type assertion")
	}
}

func TestJupiterDcaCloseDca_0(t *testing.T) {
	jsonFile, err := os.Open("data/jupiterDca_CloseDca_0.json")
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

	if DcaAction, ok := action.(*types.JupiterDcaCloseDcaAction); ok {
		assert.Equal(t, DcaAction.ProgramID, jupiterDCA.Program)
		assert.Equal(t, DcaAction.ProgramName, jupiterDCA.ProgramName)
		assert.Equal(t, DcaAction.User, "3k2pJD3FtFT8zZLoYRgHEQgays1jYs6LYCKi5MWNPLKJ")
		assert.Equal(t, DcaAction.Dca, "565uLdnjfP69rUpGuE7d5rB65V6BoqYknnWrFdt5ebFk")
		assert.Equal(t, DcaAction.InputMint, "6T44rfi9BDUdZbEvVddZWVfsGrpC6N1sSSKYnCsLpump")
		assert.Equal(t, DcaAction.OutputMint, "AiYhnwWiqbdSiEHgAzqrurcdoZx4V21mnuMt5ps2pump")
		assert.Equal(t, DcaAction.InAta, "HQpRWqAzEUsnNZR3tt8tpHHk7yyw4Q9sJQ27ZQLQ6hoz")
		assert.Equal(t, DcaAction.OutAta, "4zVid4YvihF4zeQd6RvVNMui7tesa3Jsjmuqvfovx898")
		assert.Equal(t, DcaAction.UserInAta, "4TQuSH6cQUHFZUxYPX4neGLXZXfR4gqWB9442aoPqDoJ")
		assert.Equal(t, DcaAction.UserOutAta, "GMjhyph1BXpt6qQV1MmeD1gzdR7GASJFjNo2CBgZj2wa")
	} else {
		t.Errorf("Error type assertion")
	}
}
