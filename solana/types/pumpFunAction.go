package types

type PumpFunBuyAction struct {
	BaseAction
	Who             string `json:"who"`
	FromToken       string `json:"fromToken"`
	FromTokenAmount uint64 `json:"fromTokenAmount"`
	ToToken         string `json:"toToken"`
	ToTokenAmount   uint64 `json:"toTokenAmount"`
	FeeAmount       uint64 `json:"feeAmount"`
}

type PumpFunSellAction struct {
	BaseAction
	Who             string `json:"who"`
	FromToken       string `json:"fromToken"`
	FromTokenAmount uint64 `json:"fromTokenAmount"`
	ToToken         string `json:"toToken"`
	ToTokenAmount   uint64 `json:"toTokenAmount"`
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
