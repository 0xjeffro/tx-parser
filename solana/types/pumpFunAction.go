package types

type PumpFunBuyAction struct {
	BaseAction
	Who             string `json:"who"`
	FromToken       string `json:"fromToken"`
	FromTokenAmount uint64 `json:"fromTokenAmount"`
	ToToken         string `json:"toToken"`
	ToTokenAmount   uint64 `json:"toTokenAmount"`
	MaxSolCost      uint64 `json:"maxSolCost"`
	FeeAmount       uint64 `json:"feeAmount"`
}

type PumpFunSellAction struct {
	BaseAction
	Who             string `json:"who"`
	FromToken       string `json:"fromToken"`
	FromTokenAmount uint64 `json:"fromTokenAmount"`
	ToToken         string `json:"toToken"`
	ToTokenAmount   uint64 `json:"toTokenAmount"`
	MinSolOutput    uint64 `json:"minSolOutput"`
}

type PumpFunCreateAction struct {
	BaseAction
	Who                    string `json:"who"`
	Mint                   string `json:"mint"`
	MintAuthority          string `json:"mintAuthority"`
	BondingCurve           string `json:"bondingCurve"`
	AssociatedBondingCurve string `json:"associatedBondingCurve"`
	MplTokenMetadata       string `json:"mplTokenMetadata"`
	MetaData               string `json:"metaData"`

	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Uri    string `json:"uri"`
}

type PumpFunAnchorSelfCPILogSwapAction struct {
	BaseAction
	Mint                 string `json:"mint"`
	SolAmount            uint64 `json:"solAmount"`
	TokenAmount          uint64 `json:"tokenAmount"`
	IsBuy                bool   `json:"isBuy"`
	User                 string `json:"user"`
	Timestamp            int64  `json:"timestamp"`
	VirtualSolReserves   uint64 `json:"virtualSolReserves"`
	VirtualTokenReserves uint64 `json:"virtualTokenReserves"`
}
