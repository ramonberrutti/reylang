package reylang

import (
	reylangpb "github.com/ramonberrutti/reylang/protogen/reylang"
)

type Instruction struct {
	Code    reylangpb.OpCode
	Operand any
}
type Compiler struct {
	bytecode  []Instruction
	constants []any
}

func NewCompiler() *Compiler {
	return &Compiler{
		bytecode:  make([]Instruction, 0),
		constants: make([]any, 0),
	}
}

func (c *Compiler) Compile(node *reylangpb.Node) {
	switch n := node.NodeType.(type) {
	case *reylangpb.Node_ForRangeNode:
		c.compileForRange(n.ForRangeNode)
	case *reylangpb.Node_IfNode:
		c.compileIf(n.IfNode)
	case *reylangpb.Node_ComparisonNode:
		c.compileComparison(n.ComparisonNode)
	case *reylangpb.Node_FunctionCallNode:
		c.compileFunctionCall(n.FunctionCallNode)
	case *reylangpb.Node_IdentifierNode:
		c.emit(reylangpb.OpCode_LOAD_VAR, n.IdentifierNode.Name)
	case *reylangpb.Node_LiteralNode:
		c.compileLiteral(n.LiteralNode)
	}
}

func (c *Compiler) Bytecode() []Instruction {
	return c.bytecode
}

func (c *Compiler) emit(code reylangpb.OpCode, operand any) int {
	pos := len(c.bytecode)
	c.bytecode = append(c.bytecode, Instruction{code, operand})
	return pos
}

func (c *Compiler) compileLiteral(node *reylangpb.LiteralNode) {
	var value any
	switch v := node.Value.(type) {
	case *reylangpb.LiteralNode_StringValue:
		value = v.StringValue
	case *reylangpb.LiteralNode_IntegerValue:
		value = v.IntegerValue
	case *reylangpb.LiteralNode_FloatValue:
		value = v.FloatValue
	case *reylangpb.LiteralNode_BooleanValue:
		value = v.BooleanValue
	}

	constIndex := len(c.constants)
	c.constants = append(c.constants, value)

	c.emit(reylangpb.OpCode_LOAD_CONST, constIndex)
}

func (c *Compiler) compileFunctionCall(node *reylangpb.FunctionCallNode) {
	for _, arg := range node.Arguments {
		c.Compile(arg)
	}

	c.emit(reylangpb.OpCode_CALL, node.Name.Name)
}

func (c *Compiler) compileForRange(node *reylangpb.ForRangeNode) {
	c.emit(reylangpb.OpCode_ITER_START, []string{
		node.Key.Name,
		node.Value.Name,
		node.Collection.Name,
	})

	// Next
	loopStart := c.emit(reylangpb.OpCode_ITER_NEXT, nil)

	// Compile the body
	for _, child := range node.Body {
		c.Compile(child)
	}

	// Jump back to the start of the loop
	c.emit(reylangpb.OpCode_JUMP, loopStart)

	loopEnd := c.emit(reylangpb.OpCode_ITER_END, nil)
	c.bytecode[loopStart].Operand = loopEnd
}

func (c *Compiler) compileIf(node *reylangpb.IfNode) {
	c.compileComparison(node.Condition)

	jumpIfFalseIndex := c.emit(reylangpb.OpCode_JUMP_IF_FALSE, nil)

	// Compile the body
	for _, child := range node.Body {
		c.Compile(child)
	}

	if node.ElseBody != nil && len(node.ElseBody) > 0 {
		// Jump over
		jumpToEndIndex := c.emit(reylangpb.OpCode_JUMP, nil)

		// Set the jumpIfFalseIndex to the current position
		c.bytecode[jumpIfFalseIndex].Operand = len(c.bytecode)

		// Compile the else body
		for _, child := range node.ElseBody {
			c.Compile(child)
		}

		// Set the jumpToEndIndex to the current position
		c.bytecode[jumpToEndIndex].Operand = len(c.bytecode)
	} else {
		// Set the jumpIfFalseIndex to the current position
		c.bytecode[jumpIfFalseIndex].Operand = len(c.bytecode)
	}

}

func (c *Compiler) compileComparison(node *reylangpb.ComparisonNode) {
	c.Compile(node.Left)
	c.Compile(node.Right)

	c.emit(reylangpb.OpCode_COMPARE, node.Operator)
}
