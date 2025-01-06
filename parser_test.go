package reylang_test

import (
	"testing"

	"github.com/ramonberrutti/reylang"
	"google.golang.org/protobuf/encoding/protojson"

	reylangpb "github.com/ramonberrutti/reylang/protogen/reylang"
)

func TestParser(t *testing.T) {
	tokens := []*reylangpb.Token{
		{Token: reylangpb.TokenType_KEYWORD, Value: "for"},
		{Token: reylangpb.TokenType_IDENTIFIER, Value: "i"},
		{Token: reylangpb.TokenType_PUNCTUATION, Value: ","},
		{Token: reylangpb.TokenType_IDENTIFIER, Value: "match"},
		{Token: reylangpb.TokenType_PUNCTUATION, Value: ":="},
		{Token: reylangpb.TokenType_KEYWORD, Value: "range"},
		{Token: reylangpb.TokenType_IDENTIFIER, Value: "matches"},
		{Token: reylangpb.TokenType_PUNCTUATION, Value: "{"},
		{Token: reylangpb.TokenType_KEYWORD, Value: "if"},
		{Token: reylangpb.TokenType_IDENTIFIER, Value: "match.Id"},
		{Token: reylangpb.TokenType_OPERATOR, Value: "=="},
		{Token: reylangpb.TokenType_STRING, Value: "123"},
		{Token: reylangpb.TokenType_PUNCTUATION, Value: "{"},
		{Token: reylangpb.TokenType_IDENTIFIER, Value: "golangFunction"},
		{Token: reylangpb.TokenType_PUNCTUATION, Value: "("},
		{Token: reylangpb.TokenType_IDENTIFIER, Value: "match.Id"},
		{Token: reylangpb.TokenType_PUNCTUATION, Value: ")"},
		{Token: reylangpb.TokenType_PUNCTUATION, Value: "}"},
		{Token: reylangpb.TokenType_PUNCTUATION, Value: "}"},
	}

	parser := reylang.NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		t.Fatalf("Error parsing: %v", err)
	}
	_ = ast

	j, err := protojson.MarshalOptions{Indent: "  "}.Marshal(ast)
	if err != nil {
		t.Fatalf("Error marshalling: %v", err)
	}

	t.Logf("AST: %s", j)
}
