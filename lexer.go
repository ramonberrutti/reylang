package reylang

import (
	"io"
	"strings"
	"unicode"

	reylangpb "github.com/ramonberrutti/reylang/protogen/reylang"
)

type Lexer struct {
	reader io.Reader
	buf    [1]byte
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{reader: reader}
}

func (l *Lexer) next() rune {
	r := l.peek()
	n, err := l.reader.Read(l.buf[:])
	if err != nil || n == 0 {
		return 0
	}
	if r == 0 { // What is this?
		return rune(l.buf[0])
	}

	return r
}

func (l *Lexer) peek() rune {
	return rune(l.buf[0])
}

func (l *Lexer) NextToken() *reylangpb.Token {
	// Read from reader
	// Tokenize
	// Return token
	for {
		ch := l.next()
		if ch == 0 {
			return nil
		}

		if unicode.IsSpace(ch) {
			continue // skip whitespace
		}

		if unicode.IsLetter(ch) {
			identifier := l.readIdentifier(ch)
			if isKeyword(identifier) {
				return &reylangpb.Token{
					Token: reylangpb.TokenType_KEYWORD,
					Value: identifier,
				}
			} else {
				return &reylangpb.Token{
					Token: reylangpb.TokenType_IDENTIFIER,
					Value: identifier,
				}
			}
		} else if unicode.IsDigit(ch) {
			number := l.readNumber(ch)
			return &reylangpb.Token{
				Token: reylangpb.TokenType_INTEGER,
				Value: number,
			}
		} else if ch == '"' {
			str := l.readString()
			return &reylangpb.Token{
				Token: reylangpb.TokenType_STRING,
				Value: str,
			}
		} else if strings.ContainsRune("=<>!+-*/", ch) {
			operator := l.readOperator(ch)
			return &reylangpb.Token{
				Token: reylangpb.TokenType_OPERATOR,
				Value: operator,
			}
		} else if strings.ContainsRune(",:{}()[]", ch) {
			return &reylangpb.Token{
				Token: reylangpb.TokenType_PUNCTUATION,
				Value: string(ch),
			}
		} else {
			return &reylangpb.Token{
				Token: reylangpb.TokenType_UNKNOWN,
				Value: string(ch),
			}
		}
	}
}

func (l *Lexer) readIdentifier(start rune) string {
	var builder strings.Builder
	builder.WriteRune(start)
	for unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) || l.peek() == '_' || l.peek() == '.' {
		builder.WriteRune(l.next())
	}

	return builder.String()
}

func (l *Lexer) readNumber(start rune) string {
	var builder strings.Builder
	builder.WriteRune(start)
	for unicode.IsDigit(l.peek()) {
		builder.WriteRune(l.next())
	}

	return builder.String()
}

func (l *Lexer) readString() string {
	var builder strings.Builder
	for {
		ch := l.next()
		if ch == '"' {
			break
		}

		builder.WriteRune(ch)
	}

	return builder.String()
}

func (l *Lexer) readOperator(start rune) string {
	var builder strings.Builder
	builder.WriteRune(start)
	for strings.ContainsRune("=<>!", l.peek()) {
		builder.WriteRune(l.next())
	}

	return builder.String()
}

func isKeyword(s string) bool {
	switch s {
	case "if", "else", "for", "range", "break", "continue", "return":
		return true
	}
	return false
}
