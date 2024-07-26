package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	INT = "INT"

	// Operators
	ASSIGN         = "="
	PLUS           = "+"
	MINUS          = "-"
	MULTIPLICATION = "*"
	DIVIDE         = "/"
	EQUALS         = "=="
	NOT_EQUALS     = "!="
	NOT            = "!"
	GT             = ">"
	LT             = "<"

	SEMICOLON = ";"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func ResolveIdent(ident string) Type {
	if t, ok := keywords[ident]; ok {
		return t
	}
	return "IDENT"
}
