package types

type RawTxs []RawTx

type RawTx struct {
	BlockTime        int64       `json:"blockTime"`
	IndexWithinBlock int64       `json:"indexWithinBlock"`
	Meta             Meta        `json:"meta"`
	Slot             int64       `json:"slot"`
	Transaction      Transaction `json:"transaction"`
}

type Instruction struct {
	Accounts       []int  `json:"accounts"`
	Data           string `json:"data"`
	ProgramIDIndex int    `json:"programIdIndex"`
}
type InnerInstructions struct {
	Index        int           `json:"index"`
	Instructions []Instruction `json:"instructions"`
}
type LoadedAddresses struct {
	Readonly []string `json:"readonly"`
	Writable []string `json:"writable"`
}
type UITokenAmount struct {
	Amount         string  `json:"amount"`
	Decimals       int     `json:"decimals"`
	UIAmount       float64 `json:"uiAmount"`
	UIAmountString string  `json:"uiAmountString"`
}
type TokenBalance struct {
	AccountIndex  int           `json:"accountIndex"`
	Mint          string        `json:"mint"`
	Owner         string        `json:"owner"`
	ProgramID     string        `json:"programId"`
	UITokenAmount UITokenAmount `json:"uiTokenAmount"`
}

type Meta struct {
	Err               interface{}         `json:"err"`
	Fee               uint64              `json:"fee"`
	InnerInstructions []InnerInstructions `json:"innerInstructions"`
	LoadedAddresses   LoadedAddresses     `json:"loadedAddresses"`
	LogMessages       []string            `json:"logMessages"`
	PostBalances      []uint64            `json:"postBalances"`
	PostTokenBalances []TokenBalance      `json:"postTokenBalances"`
	PreBalances       []uint64            `json:"preBalances"`
	PreTokenBalances  []TokenBalance      `json:"preTokenBalances"`
	Rewards           []interface{}       `json:"rewards"`
}
type Header struct {
	NumReadonlySignedAccounts   int `json:"numReadonlySignedAccounts"`
	NumReadonlyUnsignedAccounts int `json:"numReadonlyUnsignedAccounts"`
	NumRequiredSignatures       int `json:"numRequiredSignatures"`
}
type Message struct {
	AccountKeys         []string      `json:"accountKeys"`
	AddressTableLookups interface{}   `json:"addressTableLookups"`
	Header              Header        `json:"header"`
	Instructions        []Instruction `json:"instructions"`
	RecentBlockHash     string        `json:"recentBlockhash"`
}

type Transaction struct {
	Message    Message  `json:"message"`
	Signatures []string `json:"signatures"`
}

type ParsedTx struct {
	AccountData []struct {
		Account             string        `json:"account"`
		NativeBalanceChange float64       `json:"nativeBalanceChange"`
		TokenBalanceChanges []interface{} `json:"tokenBalanceChanges"`
	} `json:"accountData"`
	Description  string `json:"description"`
	Fee          int    `json:"fee"`
	FeePayer     string `json:"feePayer"`
	Instructions []struct {
		Accounts          []string `json:"accounts"`
		Data              string   `json:"data"`
		InnerInstructions []struct {
			Accounts  []string `json:"accounts"`
			Data      string   `json:"data"`
			ProgramID string   `json:"programId"`
		} `json:"innerInstructions"`
		ProgramID string `json:"programId"`
	} `json:"instructions"`
	NativeTransfers []struct {
		Amount          float64 `json:"amount"`
		FromUserAccount string  `json:"fromUserAccount"`
		ToUserAccount   string  `json:"toUserAccount"`
	} `json:"nativeTransfers"`
	Signature      string `json:"signature"`
	Slot           int    `json:"slot"`
	Source         string `json:"source"`
	Timestamp      int    `json:"timestamp"`
	TokenTransfers []struct {
		FromTokenAccount string  `json:"fromTokenAccount"`
		FromUserAccount  string  `json:"fromUserAccount"`
		Mint             string  `json:"mint"`
		ToTokenAccount   string  `json:"toTokenAccount"`
		ToUserAccount    string  `json:"toUserAccount"`
		TokenAmount      float64 `json:"tokenAmount"`
		TokenStandard    string  `json:"tokenStandard"`
	} `json:"tokenTransfers"`
	TransactionError interface{} `json:"transactionError"`
	Type             string      `json:"type"`
}
