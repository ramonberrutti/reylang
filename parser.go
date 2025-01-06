package reylang

import (
	"fmt"
	"strconv"

	reylangpb "github.com/ramonberrutti/reylang/protogen/reylang"
)

// Convert tokens into AST
type Parser struct {
	tokens  []*reylangpb.Token
	current int
}

func NewParser(tokens []*reylangpb.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) peek() *reylangpb.Token {
	if p.current >= len(p.tokens) {
		return nil
	}
	return p.tokens[p.current]
}

func (p *Parser) next() *reylangpb.Token {
	p.current++
	return p.peek()
}

func (p *Parser) match(tokenTypes ...reylangpb.TokenType) bool {
	if p.peek() != nil {
		for _, t := range tokenTypes {
			if p.peek().Token == t {
				return true
			}
		}
	}
	return false
}

func (p *Parser) Parse() (*reylangpb.Node, error) {
	return p.parse()
}

func (p *Parser) parse() (*reylangpb.Node, error) {
	switch {
	case p.match(reylangpb.TokenType_KEYWORD) && p.peek().Value == "for":
		return p.parseFor()
	case p.match(reylangpb.TokenType_KEYWORD) && p.peek().Value == "if":
		return p.parseIf()
	case p.match(reylangpb.TokenType_IDENTIFIER):
		return p.parseIdentifier()
	// Add more cases here
	default:
		return nil, fmt.Errorf("unexpected token: %v", p.peek())
	}
}

func (p *Parser) parseFor() (*reylangpb.Node, error) {
	p.next() // Skip "for"

	if !p.match(reylangpb.TokenType_IDENTIFIER) {
		return nil, fmt.Errorf("expected identifier after 'for'")
	}
	key := &reylangpb.IdentifierNode{
		Name: p.peek().Value,
	}

	p.next() // Skip key

	if !p.match(reylangpb.TokenType_PUNCTUATION) || p.peek().Value != "," {
		return nil, fmt.Errorf("expected ',' after key")
	}

	p.next() // Skip ","

	if !p.match(reylangpb.TokenType_IDENTIFIER) {
		return nil, fmt.Errorf("expected identifier after 'for'")
	}
	value := &reylangpb.IdentifierNode{
		Name: p.peek().Value,
	}

	p.next() // Skip value

	if !p.match(reylangpb.TokenType_PUNCTUATION) || p.peek().Value != ":=" {
		return nil, fmt.Errorf("expected ':=' after value")
	}

	p.next() // Skip ":="

	if !p.match(reylangpb.TokenType_KEYWORD) || p.peek().Value != "range" {
		return nil, fmt.Errorf("expected 'range' after ':='")
	}

	p.next() // Skip "range"

	if !p.match(reylangpb.TokenType_IDENTIFIER) {
		return nil, fmt.Errorf("expected identifier after 'range'")
	}
	collection := &reylangpb.IdentifierNode{
		Name: p.peek().Value,
	}

	p.next() // Skip collection

	if !p.match(reylangpb.TokenType_PUNCTUATION) || p.peek().Value != "{" {
		return nil, fmt.Errorf("expected '{' after condition")
	}

	p.next() // Skip "{"

	body := []*reylangpb.Node{}
	for !p.match(reylangpb.TokenType_PUNCTUATION) || p.peek().Value != "}" {
		node, err := p.parse()
		if err != nil {
			return nil, err
		}
		body = append(body, node)
	}

	p.next() // Skip "}"

	return &reylangpb.Node{
		NodeType: &reylangpb.Node_ForRangeNode{
			ForRangeNode: &reylangpb.ForRangeNode{
				Key:        key,
				Value:      value,
				Collection: collection,
				Body:       body,
			},
		},
	}, nil
}

func (p *Parser) parseIf() (*reylangpb.Node, error) {
	p.next() // Skip "if"

	condition, err := p.parseExpression()
	if err != nil {
		return nil, err
	}

	if !p.match(reylangpb.TokenType_PUNCTUATION) || p.peek().Value != "{" {
		return nil, fmt.Errorf("expected '{' after condition")
	}

	p.next() // Skip "{"

	body := []*reylangpb.Node{}
	for !p.match(reylangpb.TokenType_PUNCTUATION) || p.peek().Value != "}" {
		node, err := p.parse()
		if err != nil {
			return nil, err
		}
		body = append(body, node)
	}

	p.next() // Skip "}"

	elseBody := []*reylangpb.Node{}
	if p.match(reylangpb.TokenType_KEYWORD) && p.peek().Value == "else" {
		p.next() // Skip "else"
		if !p.match(reylangpb.TokenType_PUNCTUATION) || p.peek().Value != "{" {
			return nil, fmt.Errorf("expected '{' after 'else'")
		}

		p.next() // Skip "{"

		for !p.match(reylangpb.TokenType_PUNCTUATION) || p.peek().Value != "}" {
			node, err := p.parse()
			if err != nil {
				return nil, err
			}
			elseBody = append(elseBody, node)
		}

		p.next() // Skip "}"
	}

	return &reylangpb.Node{
		NodeType: &reylangpb.Node_IfNode{
			IfNode: &reylangpb.IfNode{
				Condition: condition.GetComparisonNode(),
				Body:      body,
				ElseBody:  elseBody,
			},
		},
	}, nil
}

// For now, only support simple expressions with literals
func (p *Parser) parseExpression() (*reylangpb.Node, error) {
	// parse left operand
	var left *reylangpb.Node
	var err error
	switch p.peek().Token {
	case reylangpb.TokenType_IDENTIFIER:
		left, err = p.parseIdentifier()
	default:
		left, err = p.parseLiteral()
	}
	if err != nil {
		return nil, err
	}

	// parse operator
	if !p.match(reylangpb.TokenType_OPERATOR) {
		return nil, fmt.Errorf("expected operator")
	}
	var operator reylangpb.ComparisonNode_Operator
	switch p.peek().Value {
	case "==":
		operator = reylangpb.ComparisonNode_OPERATOR_EQ
	case "!=":
		operator = reylangpb.ComparisonNode_OPERATOR_NE
	case "<":
		operator = reylangpb.ComparisonNode_OPERATOR_LT
	case "<=":
		operator = reylangpb.ComparisonNode_OPERATOR_LE
	case ">":
		operator = reylangpb.ComparisonNode_OPERATOR_GT
	case ">=":
		operator = reylangpb.ComparisonNode_OPERATOR_GE
	default:
		return nil, fmt.Errorf("unexpected operator: %v", p.peek())
	}

	p.next() // Skip operator

	// parse right operand
	var right *reylangpb.Node
	switch p.peek().Token {
	case reylangpb.TokenType_IDENTIFIER:
		right, err = p.parseIdentifier()
	default:
		right, err = p.parseLiteral()
	}
	if err != nil {
		return nil, err
	}

	return &reylangpb.Node{
		NodeType: &reylangpb.Node_ComparisonNode{
			ComparisonNode: &reylangpb.ComparisonNode{
				Left:     left,
				Operator: operator,
				Right:    right,
			},
		},
	}, nil
}

func (p *Parser) parseIdentifier() (*reylangpb.Node, error) {
	if !p.match(reylangpb.TokenType_IDENTIFIER) {
		return nil, fmt.Errorf("expected identifier")
	}
	identifierNode := &reylangpb.IdentifierNode{
		Name: p.peek().Value,
	}

	p.next() // Skip identifier

	// check if it's a function call
	if p.peek().Token == reylangpb.TokenType_PUNCTUATION && p.peek().Value == "(" {
		p.next() // Skip "("
		var args []*reylangpb.Node
		for !p.match(reylangpb.TokenType_PUNCTUATION) || p.peek().Value != ")" {
			arg, err := p.parse()
			if err != nil {
				return nil, err
			}
			args = append(args, arg)

			if p.peek().Token == reylangpb.TokenType_PUNCTUATION && p.peek().Value == "," {
				p.next() // Skip ","
			}
		}
		p.next() // Skip ")"

		return &reylangpb.Node{
			NodeType: &reylangpb.Node_FunctionCallNode{
				FunctionCallNode: &reylangpb.FunctionCallNode{
					Name:      identifierNode,
					Arguments: args,
				},
			},
		}, nil
	}

	return &reylangpb.Node{
		NodeType: &reylangpb.Node_IdentifierNode{
			IdentifierNode: identifierNode,
		},
	}, nil
}

func (p *Parser) parseLiteral() (*reylangpb.Node, error) {
	defer p.next()

	switch p.peek().Token {
	case reylangpb.TokenType_STRING:
		return &reylangpb.Node{
			NodeType: &reylangpb.Node_LiteralNode{
				LiteralNode: &reylangpb.LiteralNode{
					LiteralType: &reylangpb.LiteralNode_String_{
						String_: p.peek().Value,
					},
				},
			},
		}, nil

	case reylangpb.TokenType_INTEGER:
		val, err := strconv.ParseInt(p.peek().Value, 10, 32)
		if err != nil {
			return nil, err
		}

		return &reylangpb.Node{
			NodeType: &reylangpb.Node_LiteralNode{
				LiteralNode: &reylangpb.LiteralNode{
					LiteralType: &reylangpb.LiteralNode_Integer{
						Integer: int32(val),
					},
				},
			},
		}, nil

	case reylangpb.TokenType_FLOAT:
		val, err := strconv.ParseFloat(p.peek().Value, 32)
		if err != nil {
			return nil, err
		}

		return &reylangpb.Node{
			NodeType: &reylangpb.Node_LiteralNode{
				LiteralNode: &reylangpb.LiteralNode{
					LiteralType: &reylangpb.LiteralNode_Float{
						Float: float32(val),
					},
				},
			},
		}, nil

	case reylangpb.TokenType_BOOLEAN:
		val, err := strconv.ParseBool(p.peek().Value)
		if err != nil {
			return nil, err
		}

		return &reylangpb.Node{
			NodeType: &reylangpb.Node_LiteralNode{
				LiteralNode: &reylangpb.LiteralNode{
					LiteralType: &reylangpb.LiteralNode_Boolean{
						Boolean: val,
					},
				},
			},
		}, nil
	}

	return nil, fmt.Errorf("unexpected token: %v", p.peek())
}
