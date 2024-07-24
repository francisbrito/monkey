package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
)
