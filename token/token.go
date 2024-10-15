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
	DOT       = "."

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

	FOR = "FOR"
	IN  = "IN"

	PRINT = "PRINT"

	AND = "&&"
	OR  = "||"

	// Postfix operators
	INCREMENT = "++"
	DECREMENT = "--"

	PLUS_EQ  = "+="
	MINUS_EQ = "-="
)

var keywords = map[string]TokenType{
	"fn":            FUNCTION,
	"var":           VAR,
	"daca":          IF,
	"altfel":        ELSE,
	"returneaza":    RETURN,
	"adevarat":      TRUE,
	"fals":          FALSE,
	"si":            AND,
	"sau":           OR,
	"incrementeaza": INCREMENT,
	"decrementeaza": DECREMENT,
	"pentru":        FOR,
	"in":            IN,
	"afiseaza":      PRINT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
