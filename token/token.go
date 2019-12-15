package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL" // token or character we don't know about
	EOF     = "EOF" // end of file

	// identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
