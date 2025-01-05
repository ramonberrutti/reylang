package reylang_test

import (
	"testing"

	reylangpb "github.com/ramonberrutti/reylang/protogen/proto/reylang"
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
												LiteralType: &reylangpb.LiteralNode_String_{
													String_: "123",
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
														NodeType: &reylangpb.Node_LiteralNode{
															LiteralNode: &reylangpb.LiteralNode{
																LiteralType: &reylangpb.LiteralNode_String_{
																	String_: "match.id",
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
		},
	}

	_ = ast

}
