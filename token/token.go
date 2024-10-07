package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int16
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	LTE = "<="
	GTE = ">="

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACKET = "["
	RBRACKET = "]"
	LBRACE   = "{"
	RBRACE   = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	AND = "&&"
	OR  = "||"

	INCREMENT = "++"
	DECREMENT = "--"
)

var keywords = map[string]TokenType{
	"fn":        FUNCTION,
	"var":       VAR,
	"daca":      IF,
	"altfel":    ELSE,
	"ret":       RETURN,
	"adev":      TRUE,
	"fals":      FALSE,
	"si":        AND,
	"sau":       OR,
	"increment": INCREMENT,
	"decrement": DECREMENT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
