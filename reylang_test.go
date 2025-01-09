package reylang_test

import (
	"strings"
	"testing"

	"github.com/ramonberrutti/reylang"
	reylangpb "github.com/ramonberrutti/reylang/protogen/reylang"
)

func TestReyLang(t *testing.T) {
	code := `
	for _, match := range matches {
		if match.id == "123" {
			golangFunction(match.id)
		}
	}
`

	lexer := reylang.NewLexer(strings.NewReader(code))

	tokens := []*reylangpb.Token{}
	for {
		token := lexer.NextToken()
		if token == nil {
			break
		}
		t.Logf("Token: %v", token)
		tokens = append(tokens, token)
	}

	parser := reylang.NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Error parsing: %v", err)
	}

	compiler := reylang.NewCompiler()
	compiler.Compile(ast)

	bytecode := compiler.Bytecode()
	for i, instr := range bytecode {
		t.Logf("Instruction %d: %s: %v", i, instr.Code, instr.Operand)
	}

	vm := reylang.NewVM([]any{
		"123",
	})

	vm.SetSymbol("matches", []any{
		map[string]any{
			"id": "123",
		},
	})

	vm.Run(bytecode)
}
