package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}
	l := New(input)
	for i, tc := range tests {
		tok := l.NextToken()
		if tok == nil {
			t.Fatalf("test[%d] token is nil", i)
		}
		if tok.Type != tc.expectedType {
			t.Fatalf("test[%d] expected=%q, got=%q", i, tc.expectedType, tok.Type)
		}
		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("test[%d] expected=%q, got=%q", i, tc.expectedLiteral, tok.Literal)
		}
	}
}
