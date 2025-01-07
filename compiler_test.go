package reylang_test

import (
	"testing"

	"github.com/ramonberrutti/reylang"
	reylangpb "github.com/ramonberrutti/reylang/protogen/reylang"
)

func TestCompiler(t *testing.T) {
	ast := &reylangpb.Node{
		NodeType: &reylangpb.Node_ForRangeNode{
			ForRangeNode: &reylangpb.ForRangeNode{
				Key: &reylangpb.IdentifierNode{
					Name: "_",
				},
				Value: &reylangpb.IdentifierNode{
					Name: "match",
				},
				Collection: &reylangpb.IdentifierNode{
					Name: "matches",
				},
				Body: []*reylangpb.Node{
					{
						NodeType: &reylangpb.Node_IfNode{
							IfNode: &reylangpb.IfNode{
								Condition: &reylangpb.ComparisonNode{
									Left: &reylangpb.Node{
										NodeType: &reylangpb.Node_IdentifierNode{
											IdentifierNode: &reylangpb.IdentifierNode{
												Name: "match.id",
											},
										},
									},
									Operator: reylangpb.ComparisonNode_OPERATOR_EQ,
									Right: &reylangpb.Node{
										NodeType: &reylangpb.Node_LiteralNode{
											LiteralNode: &reylangpb.LiteralNode{
												Value: &reylangpb.LiteralNode_StringValue{
													StringValue: "123",
												},
											},
										},
									},
								},
								Body: []*reylangpb.Node{
									{
										NodeType: &reylangpb.Node_FunctionCallNode{
											FunctionCallNode: &reylangpb.FunctionCallNode{
												Name: &reylangpb.IdentifierNode{
													Name: "print",
												},
												Arguments: []*reylangpb.Node{
													{
														NodeType: &reylangpb.Node_IdentifierNode{
															IdentifierNode: &reylangpb.IdentifierNode{
																Name: "match.id",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	_ = ast
	compiler := reylang.NewCompiler()
	compiler.Compile(ast)

	bytecode := compiler.Bytecode()
	for i, instr := range bytecode {
		t.Logf("Instruction %d: %s: %v", i, instr.Code, instr.Operand)
	}
}
