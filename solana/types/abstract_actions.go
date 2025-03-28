package types

type Action interface {
	GetProgramID() string
	GetProgramName() string
	GetInstructionName() string
}

type BaseAction struct {
	ProgramID       string `json:"programId"`
	ProgramName     string `json:"programName"`
	InstructionName string `json:"instructionName"`
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

type UnknownAction struct {
	BaseAction
	Error error `json:"error"`
}
