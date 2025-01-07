package reylang_test

import (
	"strings"
	"testing"

	"github.com/ramonberrutti/reylang"
)

func TestLexer(t *testing.T) {
	code := `
	for _, match := range matches {
		if match.id == "123" {
			golangFunction(match.id)
		}
	}
`

	lexer := reylang.NewLexer(strings.NewReader(code))

	for {
		token := lexer.NextToken()
		if token == nil {
			break
		}
		t.Logf("Token: %v", token)
	}
}
