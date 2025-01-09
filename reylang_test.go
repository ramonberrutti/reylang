package reylang_test

import (
	"strings"
	"testing"

	"github.com/ramonberrutti/reylang"
	reylangpb "github.com/ramonberrutti/reylang/protogen/reylang"
)

func TestReyLang(t *testing.T) {
	code := `
	sayHello("Ramon")
	for _, match := range matches {
		if match.id == "123" {
			golangFunction(match.name)
			sayHello(match.name)
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
		"Ramon",
		"123",
	})

	vm.SetSymbol("matches", []any{
		map[string]any{
			"id":   "123",
			"name": "Ramon",
		},
		map[string]any{
			"id":   "123",
			"name": "Ramon2",
		},
	})
	vm.SetSymbol("golangFunction", func(args ...any) any {
		t.Logf("Calling golangFunction with args: %v", args)
		return nil
	})
	vm.SetSymbol("sayHello", func(args ...any) any {
		t.Logf("Calling sayHello with args: %v", args)
		return nil
	})

	vm.Run(bytecode)
}
