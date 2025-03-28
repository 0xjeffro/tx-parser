package types

type Action interface {
	GetProgramID() string
	GetProgramName() string
	GetInstructionName() string
	GetActionType() string
}

type BaseAction struct {
	ProgramID       string `json:"programId"`
	ProgramName     string `json:"programName"`
	InstructionName string `json:"instructionName"`
	ActionType      string `json:"actionType"`
}

func (a BaseAction) GetProgramID() string {
	return a.ProgramID
}

func (a BaseAction) GetProgramName() string {
	return a.ProgramName
}

func (a BaseAction) GetInstructionName() string {
	return a.InstructionName
}

func (a BaseAction) GetActionType() string {
	return a.ActionType
}

type UnknownAction struct {
	BaseAction
	Error error `json:"error"`
}

type SwapActionMixin struct {
	Who               string `json:"who"`
	FromToken         string `json:"fromToken"`
	FromTokenAmount   uint64 `json:"fromTokenAmount"`
	FromTokenDecimals uint64 `json:"fromTokenDecimals"`
	ToToken           string `json:"toToken"`
	ToTokenAmount     uint64 `json:"toTokenAmount"`
	ToTokenDecimals   uint64 `json:"toTokenDecimals"`
}
