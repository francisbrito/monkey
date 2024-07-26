package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
		let five = 5;
		let ten = 10;
		let add = fn(x, y) {
			x + y;
		};
		let result = add(five, ten);
		!-/*5;
		5 < 10 > 5;
		if 5 < 10 {
			return true;
		} else {
			return false;
		}
		10 == 10;
		8 != 10;
	`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.LET, "let"},
		{"IDENT", "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{"IDENT", "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{"IDENT", "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{"IDENT", "x"},
		{",", ","},
		{"IDENT", "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{"IDENT", "x"},
		{token.PLUS, "+"},
		{"IDENT", "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{"IDENT", "result"},
		{token.ASSIGN, "="},
		{"IDENT", "add"},
		{token.LPAREN, "("},
		{"IDENT", "five"},
		{",", ","},
		{"IDENT", "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.NOT, "!"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.MULTIPLICATION, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQUALS, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "8"},
		{token.NOT_EQUALS, "!="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tc := range tests {
		tok := l.NextToken()
		if tok.Type != tc.expectedType {
			t.Fatalf("test[%d] expected=%q, got=%q, literal=%q", i, tc.expectedType, tok.Type, tok.Literal)
		}
		if tok.Literal != tc.expectedLiteral {
			t.Fatalf("test[%d] expected=%q, got=%q", i, tc.expectedLiteral, tok.Literal)
		}
	}
}
