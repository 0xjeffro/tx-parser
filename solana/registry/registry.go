package registry

import "github.com/0xjeffro/tx-parser/solana/types"

type ProgramParser struct {
	ProgramName       string
	ProgramID         string
	ProgramParserFunc func(*types.Instruction, types.Instruction) (types.Action, error)
}

var programParsers = make(map[string]ProgramParser)

func Register(parser ProgramParser) {
	programParsers[parser.ProgramID] = parser
}

func GetAllProgramParsers() []ProgramParser {
	result := make([]ProgramParser, 0, len(programParsers))
	for _, p := range programParsers {
		result = append(result, p)
	}
	return result
}

func GetParserByID(programID string) (ProgramParser, bool) {
	p, ok := programParsers[programID]
	return p, ok
}
