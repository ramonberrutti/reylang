package reylang_test

import (
	"testing"

	"github.com/ramonberrutti/reylang"

	reylangpb "github.com/ramonberrutti/reylang/protogen/reylang"
)

func TestVM(t *testing.T) {

	vm := reylang.NewVM([]any{
		"123",
	})

	vm.SetSymbol("matches", []any{
		map[string]any{
			"id": "123",
		},
	})
	_ = vm

	vm.Run([]reylang.Instruction{
		{Code: reylangpb.OpCode_ITER_START, Operand: []string{"_", "match", "matches"}},
		{Code: reylangpb.OpCode_ITER_NEXT, Operand: 9},
		{Code: reylangpb.OpCode_LOAD_VAR, Operand: "match.id"},
		{Code: reylangpb.OpCode_LOAD_CONST, Operand: 0},
		{Code: reylangpb.OpCode_COMPARE, Operand: reylangpb.ComparisonNode_OPERATOR_EQ},
		{Code: reylangpb.OpCode_JUMP_IF_FALSE, Operand: 8},
		{Code: reylangpb.OpCode_LOAD_VAR, Operand: "match.id"},
		{Code: reylangpb.OpCode_CALL, Operand: "print"},
		{Code: reylangpb.OpCode_JUMP, Operand: 1},
		{Code: reylangpb.OpCode_ITER_END, Operand: nil},
	})
}
